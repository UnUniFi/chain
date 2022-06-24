package types

import (
	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys and default values
var ()

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
	return nil
}
