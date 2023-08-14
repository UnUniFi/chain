package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyDenomCreationFee              = []byte("ClassCreationFee")
	DefaultFeeDenom                  = "uguu"
	DefaultFeeAmount           int64 = 1_000_000
	KeyFeeCollectorAddress           = []byte("FeeCollectorAddress")
	DefaultFeeCollectorAddress       = "ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6"
)

// ParamTable for tokenfactory module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(denomCreationFee sdk.Coins, feeCollectorAddress string) Params {
	return Params{
		ClassCreationFee:    denomCreationFee,
		FeeCollectorAddress: feeCollectorAddress,
	}
}

// default tokenfactory module parameters.
func DefaultParams() Params {
	return Params{
		ClassCreationFee:    sdk.NewCoins(sdk.NewInt64Coin(DefaultFeeDenom, DefaultFeeAmount)),
		FeeCollectorAddress: DefaultFeeCollectorAddress,
	}
}

// validate params.
func (p Params) Validate() error {
	if err := validateDenomCreationFee(p.ClassCreationFee); err != nil {
		return err
	}

	return validateFeeCollectorAddress(p.FeeCollectorAddress)
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyDenomCreationFee, &p.ClassCreationFee, validateDenomCreationFee),
		paramstypes.NewParamSetPair(KeyFeeCollectorAddress, &p.FeeCollectorAddress, validateFeeCollectorAddress),
	}
}

func validateDenomCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return fmt.Errorf("invalid denom creation fee: %+v, %w", i, err)
	}

	return nil
}

func validateFeeCollectorAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// Fee collector address might be explicitly empty in test environments
	if len(v) == 0 {
		return nil
	}

	_, err := sdk.AccAddressFromBech32(v)
	if err != nil {
		return fmt.Errorf("invalid fee collector address: %w", err)
	}

	return nil
}
