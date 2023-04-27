package types

import (
	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"
)

// NewParams returns a new params object
func NewParams() Params {
	return DefaultParams()
}

// DefaultParams returns default params for incentive module
func DefaultParams() Params {
	return Params{}
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {

	return nil
}

func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{}
}
