package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyCommissionRate       = []byte("CommissionRate")
	KeyVaultCreationFee     = []byte("VaultCreationFee")
	KeyVaultCreationDeposit = []byte("VaultCreationDeposit")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	commissionRate sdk.Dec,
	vaultCreationFee sdk.Coin,
	vaultCreationDeposit sdk.Coin,
) Params {
	return Params{
		CommissionRate:       commissionRate,
		VaultCreationFee:     vaultCreationFee,
		VaultCreationDeposit: vaultCreationDeposit,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(sdk.ZeroDec(), sdk.NewInt64Coin("stake", 1000), sdk.NewInt64Coin("stake", 1000))
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyCommissionRate, &p.CommissionRate, validateCommissionRate),
		paramstypes.NewParamSetPair(KeyVaultCreationFee, &p.VaultCreationFee, validateVaultCreationFee),
		paramstypes.NewParamSetPair(KeyVaultCreationDeposit, &p.VaultCreationDeposit, validateVaultCreationDeposit),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateCommissionRate(p.CommissionRate); err != nil {
		return err
	}
	if err := validateVaultCreationFee(p.VaultCreationFee); err != nil {
		return err
	}
	if err := validateVaultCreationDeposit(p.VaultCreationDeposit); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateCommissionRate(i interface{}) error {
	rate, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if rate.IsNil() || rate.IsNegative() || rate.GT(sdk.OneDec()) {
		return fmt.Errorf("invalid rate: %s", rate.String())
	}

	return nil
}

func validateVaultCreationFee(i interface{}) error {
	_, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateVaultCreationDeposit(i interface{}) error {
	_, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
