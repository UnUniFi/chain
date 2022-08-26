package types

import (
	"errors"
	fmt "fmt"
	"strings"

	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
)

var _ paramstype.ParamSet = (*Params)(nil)

// Parameter keys
var (
	KeyNetworks = []byte("Networks")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return DefaultParams()
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	beAddr, _ := sdk.AccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")
	addr := ununifitypes.StringAccAddress(beAddr)
	return Params{
		Networks: []Network{
			{
				NetworkId: "Ethereum",
				Asset:     "ETH",
				Oracles:   []ununifitypes.StringAccAddress{addr},
				Active:    true,
			},
		},
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{
		paramstype.NewParamSetPair(KeyNetworks, &p.Networks, validateNetworksParams),
	}
}

// Validate ensure that params have valid values
func (p Params) Validate() error {
	return validateNetworksParams(p.Networks)
}

// todo: split source
type Networks []Network

func (n Network) Validate() error {
	if strings.TrimSpace(n.NetworkId) == "" {
		return errors.New("network id cannot be blank")
	}
	if strings.TrimSpace(n.Asset) == "" {
		return errors.New("Asset cannot be blank")
	}
	seenOracles := make(map[string]bool)
	for i, oracle := range n.Oracles {
		if oracle.AccAddress().Empty() {
			return fmt.Errorf("oracle %d is empty", i)
		}
		if seenOracles[oracle.AccAddress().String()] {
			return fmt.Errorf("duplicated oracle %s", oracle)
		}
		seenOracles[oracle.AccAddress().String()] = true
	}
	return nil
}

func (ns Networks) Validate() error {
	seenNetworks := make(map[string]bool)
	for _, n := range ns {
		if seenNetworks[n.NetworkId] {
			return fmt.Errorf("duplicated network id %s", n.NetworkId)
		}
		if err := n.Validate(); err != nil {
			return err
		}
		seenNetworks[n.NetworkId] = true
	}
	return nil
}

// String implements fmt.Stringer
func (ns Networks) String() string {
	out := "Networks:\n"
	for _, n := range ns {
		out += fmt.Sprintf("%s\n", n.String())
	}
	return strings.TrimSpace(out)
}

func validateNetworksParams(i interface{}) error {
	networks, ok := i.(Networks)
	if !ok {
		networks, ok = i.([]Network)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
	}

	return networks.Validate()
}
