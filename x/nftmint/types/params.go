package types

import (
	fmt "fmt"

	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"
)

// NewParams returns a new params object
func NewParams() Params {
	return DefaultParams()
}

// DefaultParams returns default params for nftmint module
func DefaultParams() Params {
	return Params{
		MaxTokenSupplyCap: 100000,
		MinClassNameLen:   3,
		MaxClassNameLen:   128,
		MinUriLen:         8,
		MaxUriLen:         512,
		MaxSymbolLen:      16,
		MaxDescriptionLen: 1024,
	}
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// Parameter keys
var (
	KeyMaxTokenSupplyCap = []byte("MaxTokenSupplyCap")
	KeyMinClassNameLen   = []byte("MinClassNameLen")
	KeyMaxClassNameLen   = []byte("MaxClassNameLen")
	KeyMinUriLen         = []byte("MinUriLen")
	KeyMaxClassUriLen    = []byte("MaxUriLen")
	KeyMaxSymbolLen      = []byte("MaxSymbolLen")
	KeyMaxDescriptionLen = []byte("MaxDescriptionLen")
)

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{
		paramstype.NewParamSetPair(KeyMaxTokenSupplyCap, &p.MaxTokenSupplyCap, validateMaxTokenSupplyCap),
		paramstype.NewParamSetPair(KeyMinClassNameLen, &p.MinClassNameLen, validateMinClassNameLen),
		paramstype.NewParamSetPair(KeyMaxClassNameLen, &p.MinClassNameLen, validateMaxClassNameLen),
		paramstype.NewParamSetPair(KeyMinUriLen, &p.MinUriLen, validateMinUriLen),
		paramstype.NewParamSetPair(KeyMaxClassUriLen, &p.MaxUriLen, validateMaxUriLen),
		paramstype.NewParamSetPair(KeyMaxSymbolLen, &p.MaxSymbolLen, validateMaxSymbolLen),
		paramstype.NewParamSetPair(KeyMaxDescriptionLen, &p.MaxDescriptionLen, validateMaxDescriptionLen),
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	if err := validateMaxTokenSupplyCap(p.MaxTokenSupplyCap); err != nil {
		return err
	}

	if err := validateMinClassNameLen(p.MinClassNameLen); err != nil {
		return err
	}

	if err := validateMaxClassNameLen(p.MaxClassNameLen); err != nil {
		return err
	}

	if err := validateMinUriLen(p.MinUriLen); err != nil {
		return err
	}

	if err := validateMaxUriLen(p.MaxUriLen); err != nil {
		return err
	}

	if err := validateMaxSymbolLen(p.MaxSymbolLen); err != nil {
		return err
	}

	if err := validateMaxDescriptionLen(p.MaxDescriptionLen); err != nil {
		return err
	}

	return nil
}

func validateMaxTokenSupplyCap(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	return nil
}

func validateMinClassNameLen(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	return nil
}

func validateMaxClassNameLen(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	return nil
}

func validateMinUriLen(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	return nil
}

func validateMaxUriLen(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	return nil
}

func validateMaxSymbolLen(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	return nil
}

func validateMaxDescriptionLen(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	return nil
}
