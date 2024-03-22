package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultFeeCollectorAddress = ""
)

// NewParams creates a new Params instance
func NewParams(
	authority string,
	tradeFeeRate sdk.Dec,
) Params {
	return Params{
		Authority:    authority,
		TradeFeeRate: tradeFeeRate,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		"ununifi10d07y265gmmuvt4z0w9aw880jnsr700jvqqhva", // gov module
		sdk.ZeroDec(),
	)
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateTradeFeeRate(p.TradeFeeRate); err != nil {
		return err
	}
	if err := validateAuthority(p.Authority); err != nil {
		return err
	}
	return nil
}

func validateTradeFeeRate(i interface{}) error {
	rate, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if rate.IsNil() || rate.IsNegative() || rate.GT(sdk.OneDec()) {
		return fmt.Errorf("invalid rate: %s", rate.String())
	}

	return nil
}

func validateAuthority(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	_, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err
	}

	return nil
}
