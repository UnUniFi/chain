package types

import (
	fmt "fmt"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultMaxNFTSupplyCap   = 100000
	DefaultMinClassNameLen   = 3
	DefaultMaxClassNameLen   = 128
	DefaultMinUriLen         = 8
	DefaultMaxUriLen         = 512
	DefaultMaxSymbolLen      = 16
	DefaultMaxDescriptionLen = 1024
)

// NewParams returns a new params object
func NewParams(
	maxNFTMintSupplyCap,
	minClassNameLen, maxClassNameLen,
	minUriLen, maxUriLen,
	maxSymbolLen,
	maxDescriptionLen uint64,
) Params {
	return Params{
		MaxNFTSupplyCap:   maxNFTMintSupplyCap,
		MinClassNameLen:   minClassNameLen,
		MaxClassNameLen:   maxClassNameLen,
		MinUriLen:         minUriLen,
		MaxUriLen:         maxClassNameLen,
		MaxSymbolLen:      maxSymbolLen,
		MaxDescriptionLen: maxDescriptionLen,
	}
}

// DefaultParams returns default params for nftmint module
func DefaultParams() Params {
	return Params{
		MaxNFTSupplyCap:   DefaultMaxNFTSupplyCap,
		MinClassNameLen:   DefaultMinClassNameLen,
		MaxClassNameLen:   DefaultMaxClassNameLen,
		MinUriLen:         DefaultMinUriLen,
		MaxUriLen:         DefaultMaxUriLen,
		MaxSymbolLen:      DefaultMaxSymbolLen,
		MaxDescriptionLen: DefaultMaxDescriptionLen,
	}
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Parameter keys
var (
	KeyMaxNFTSupplyCap   = []byte("MaxNFTSupplyCap")
	KeyMinClassNameLen   = []byte("MinClassNameLen")
	KeyMaxClassNameLen   = []byte("MaxClassNameLen")
	KeyMinUriLen         = []byte("MinUriLen")
	KeyMaxUriLen         = []byte("MaxUriLen")
	KeyMaxSymbolLen      = []byte("MaxSymbolLen")
	KeyMaxDescriptionLen = []byte("MaxDescriptionLen")
)

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyMaxNFTSupplyCap, &p.MaxNFTSupplyCap, validateMaxNFTSupplyCap),
		paramstypes.NewParamSetPair(KeyMinClassNameLen, &p.MinClassNameLen, validateMinClassNameLen),
		paramstypes.NewParamSetPair(KeyMaxClassNameLen, &p.MaxClassNameLen, validateMaxClassNameLen),
		paramstypes.NewParamSetPair(KeyMinUriLen, &p.MinUriLen, validateMinUriLen),
		paramstypes.NewParamSetPair(KeyMaxUriLen, &p.MaxUriLen, validateMaxUriLen),
		paramstypes.NewParamSetPair(KeyMaxSymbolLen, &p.MaxSymbolLen, validateMaxSymbolLen),
		paramstypes.NewParamSetPair(KeyMaxDescriptionLen, &p.MaxDescriptionLen, validateMaxDescriptionLen),
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	if err := validateMaxNFTSupplyCap(p.MaxNFTSupplyCap); err != nil {
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

func validateMaxNFTSupplyCap(i interface{}) error {
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
