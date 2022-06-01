package types

import (
	"fmt"

	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys and default values
var (
	DefaultBidToken = "uguu"
)

// NewParams returns a new params object
func NewParams() Params {
	return Params{}
}

// DefaultParams returns default params for incentive module
func DefaultParams() Params {
	return NewParams()
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {

	if err := validateBidTokens(p.BidTokens); err != nil {
		return err
	}

	return nil
}

func validateBidTokens(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
