package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyPool             = []byte("Pool")
	keyPerpetualFutures = []byte("PerpetualFutures")
	KeyPerpetualOptions = []byte("PerpetualOptions")
)

func DefaultPool() Pool {
	return Pool{
		QuoteTicker:                       "",
		BaseLptMintFee:                    sdk.MustNewDecFromStr("0.001"),
		BaseLptRedeemFee:                  sdk.MustNewDecFromStr("0.001"),
		BorrowingFeeRatePerHour:           sdk.ZeroDec(),
		LiquidationNeededReportRewardRate: sdk.ZeroDec(),
		AcceptedAssets:                    []*Pool_Asset{},
	}
}

func DefaultPerpetualFuturesParams() PerpetualFuturesParams {
	return PerpetualFuturesParams{
		CommissionRate:        sdk.MustNewDecFromStr("0.001"),
		MarginMaintenanceRate: sdk.MustNewDecFromStr("0.5"),
		ImaginaryFundingRateProportionalCoefficient: sdk.MustNewDecFromStr("0.0005"),
		Markets: []*Market{},
	}
}

func DefaultPerpetualOptionsParams() PerpetualOptionsParams {
	return PerpetualOptionsParams{
		PremiumCommissionRate:                       sdk.ZeroDec(),
		StrikeCommissionRate:                        sdk.MustNewDecFromStr("0.001"),
		MarginMaintenanceRate:                       sdk.ZeroDec(),
		ImaginaryFundingRateProportionalCoefficient: sdk.ZeroDec(),
		Markets: []*Market{},
	}
}

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return DefaultParams()
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		Pool:             DefaultPool(),
		PerpetualFutures: DefaultPerpetualFuturesParams(),
		PerpetualOptions: DefaultPerpetualOptionsParams(),
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPool, &p.Pool, validatePool),
		paramtypes.NewParamSetPair(keyPerpetualFutures, &p.PerpetualFutures, validatePerpetualFutures),
		paramtypes.NewParamSetPair(KeyPerpetualOptions, &p.PerpetualOptions, validatePerpetualOptions),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validatePool(p.Pool); err != nil {
		return err
	}

	if err := validatePerpetualFutures(p.PerpetualFutures); err != nil {
		return err
	}

	if err := validatePerpetualOptions(p.PerpetualOptions); err != nil {
		return err
	}

	return nil
}

func validatePool(i interface{}) error {
	// check type
	pool, ok := i.(Pool)
	if !ok {
		return fmt.Errorf("invalid paramter type: %T", i)
	}

	if !pool.BaseLptMintFee.LTE(sdk.OneDec()) {
		return fmt.Errorf("invalid base lpt mint fee: %s", pool.BaseLptMintFee)
	}

	if !pool.BaseLptRedeemFee.LTE(sdk.OneDec()) {
		return fmt.Errorf("invalid base lpt redeem fee: %s", pool.BaseLptRedeemFee)
	}

	if !pool.BorrowingFeeRatePerHour.LTE(sdk.OneDec()) {
		return fmt.Errorf("invalid borrowing fee rate per hour: %s", pool.BorrowingFeeRatePerHour)
	}

	if !pool.LiquidationNeededReportRewardRate.LTE(sdk.OneDec()) {
		return fmt.Errorf("invalid liquidation needed report reward rate: %s", pool.LiquidationNeededReportRewardRate)
	}

	return nil
}

func validatePerpetualFutures(i interface{}) error {
	// check type
	perpetualFuturesParams, ok := i.(PerpetualFuturesParams)
	if !ok {
		return fmt.Errorf("invalid paramter type: %T", i)
	}

	if !perpetualFuturesParams.CommissionRate.LTE(sdk.OneDec()) {
		return fmt.Errorf("invalid commission rate: %s", perpetualFuturesParams.CommissionRate)
	}

	if !perpetualFuturesParams.MarginMaintenanceRate.LTE(sdk.OneDec()) {
		return fmt.Errorf("invalid margin maintenance rate: %s", perpetualFuturesParams.MarginMaintenanceRate)
	}

	return nil
}

func validatePerpetualOptions(i interface{}) error {
	// check type
	_, ok := i.(PerpetualOptionsParams)
	if !ok {
		return fmt.Errorf("invalid paramter type: %T", i)
	}

	return nil
}
