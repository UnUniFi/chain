package types

import (
	"fmt"
	"time"

	math "cosmossdk.io/math"
)

// NewParams returns a new params object
func NewParams() Params {
	return DefaultParams()
}

// DefaultParams returns default params for incentive module
func DefaultParams() Params {
	return Params{
		CommissionRate:    math.LegacyMustNewDecFromStr("0.02"),
		FullPaymentPeriod: 30,
		NftDeliveryPeriod: 30,
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {

	if err := validateCommissionRate(p.CommissionRate); err != nil {
		return err
	}

	if err := validateFullPaymentPeriod(p.FullPaymentPeriod); err != nil {
		return err
	}

	if err := validateNftDeliveryPeriod(p.NftDeliveryPeriod); err != nil {
		return err
	}

	return nil
}

func validateCommissionRate(i interface{}) error {
	_, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateFullPaymentPeriod(i interface{}) error {
	_, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftDeliveryPeriod(i interface{}) error {
	_, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if rate.IsNil() || rate.IsNegative() || rate.GT(sdk.OneDec()) {
		return fmt.Errorf("invalid rate: %s", rate.String())
	}

	return nil
}
