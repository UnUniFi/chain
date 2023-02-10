package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter keys
var (
	KeyPool     = []byte("Pool")
	DefaultPool = Pool{
		QuoteTicker: "USD",
	}
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPool, &p.Pool, validatePoolParams),
	}
}

func validatePoolParams(i interface{}) error {
	pool, ok := i.(Pool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return pool.Validate()
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validatePoolParams(p.Pool)
}
