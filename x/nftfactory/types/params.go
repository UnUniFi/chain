package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultFeeDenom                  = "uguu"
	DefaultFeeAmount           int64 = 1_000_000
	DefaultFeeCollectorAddress       = ""
)

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
