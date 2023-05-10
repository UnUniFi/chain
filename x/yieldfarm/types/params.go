package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	KeyDailyReward = []byte("DailyReward")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(dailyReward uint64) Params {
	return Params{
		DailyReward: dailyReward,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(1)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDailyReward, &p.DailyReward, validateDailyReward),
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {

	if err := validateDailyReward(p.DailyReward); err != nil {
		return err
	}

	return nil
}

func validateDailyReward(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
