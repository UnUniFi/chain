package types

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"
)

var emptyDec = sdk.Dec{}

// Defaults for auction params
const (
	// DefaultMaxAuctionDuration max length of auction
	DefaultMaxAuctionDuration time.Duration = 2 * 24 * time.Hour
	// DefaultBidDuration how long an auction gets extended when someone bids
	DefaultBidDuration time.Duration = 1 * time.Hour
)

var (
	// DefaultIncrement is the smallest percent change a new bid must have from the old one
	DefaultIncrement sdk.Dec = sdk.MustNewDecFromStr("0.05")
	// ParamStoreKeyParams Param store key for auction params
	KeyBidDuration         = []byte("BidDuration")
	KeyMaxAuctionDuration  = []byte("MaxAuctionDuration")
	KeyIncrementSurplus    = []byte("IncrementSurplus")
	KeyIncrementDebt       = []byte("IncrementDebt")
	KeyIncrementCollateral = []byte("IncrementCollateral")
)

var _ paramstype.ParamSet = &Params{}

// NewParams returns a new Params object.
func NewParams(maxAuctionDuration, bidDuration time.Duration, incrementSurplus, incrementDebt, incrementCollateral sdk.Dec) Params {
	return Params{
		MaxAuctionDuration:  maxAuctionDuration,
		BidDuration:         bidDuration,
		IncrementSurplus:    incrementSurplus,
		IncrementDebt:       incrementDebt,
		IncrementCollateral: incrementCollateral,
	}
}

// DefaultParams returns the default parameters for auctions.
func DefaultParams() Params {
	return NewParams(
		DefaultMaxAuctionDuration,
		DefaultBidDuration,
		DefaultIncrement,
		DefaultIncrement,
		DefaultIncrement,
	)
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs.
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{
		paramstype.NewParamSetPair(KeyBidDuration, &p.BidDuration, validateBidDurationParam),
		paramstype.NewParamSetPair(KeyMaxAuctionDuration, &p.MaxAuctionDuration, validateMaxAuctionDurationParam),
		paramstype.NewParamSetPair(KeyIncrementSurplus, &p.IncrementSurplus, validateIncrementSurplusParam),
		paramstype.NewParamSetPair(KeyIncrementDebt, &p.IncrementDebt, validateIncrementDebtParam),
		paramstype.NewParamSetPair(KeyIncrementCollateral, &p.IncrementCollateral, validateIncrementCollateralParam),
	}
}

// Equal returns a boolean determining if two Params types are identical.
func (p Params) Equal(p2 Params) bool {
	bz1 := ModuleCdc.MustMarshalLengthPrefixed(&p)
	bz2 := ModuleCdc.MustMarshalLengthPrefixed(&p2)
	return bytes.Equal(bz1, bz2)
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	if err := validateBidDurationParam(p.BidDuration); err != nil {
		return err
	}

	if err := validateMaxAuctionDurationParam(p.MaxAuctionDuration); err != nil {
		return err
	}

	if p.BidDuration > p.MaxAuctionDuration {
		return errors.New("bid duration param cannot be larger than max auction duration")
	}

	if err := validateIncrementSurplusParam(p.IncrementSurplus); err != nil {
		return err
	}

	if err := validateIncrementDebtParam(p.IncrementDebt); err != nil {
		return err
	}

	return validateIncrementCollateralParam(p.IncrementCollateral)
}

func validateBidDurationParam(i interface{}) error {
	bidDuration, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if bidDuration < 0 {
		return fmt.Errorf("bid duration cannot be negative %d", bidDuration)
	}

	return nil
}

func validateMaxAuctionDurationParam(i interface{}) error {
	maxAuctionDuration, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if maxAuctionDuration < 0 {
		return fmt.Errorf("max auction duration cannot be negative %d", maxAuctionDuration)
	}

	return nil
}

func validateIncrementSurplusParam(i interface{}) error {
	incrementSurplus, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if incrementSurplus == emptyDec || incrementSurplus.IsNil() {
		return errors.New("surplus auction increment cannot be nil or empty")
	}

	if incrementSurplus.IsNegative() {
		return fmt.Errorf("surplus auction increment cannot be less than zero %s", incrementSurplus)
	}

	return nil
}

func validateIncrementDebtParam(i interface{}) error {
	incrementDebt, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if incrementDebt == emptyDec || incrementDebt.IsNil() {
		return errors.New("debt auction increment cannot be nil or empty")
	}

	if incrementDebt.IsNegative() {
		return fmt.Errorf("debt auction increment cannot be less than zero %s", incrementDebt)
	}

	return nil
}

func validateIncrementCollateralParam(i interface{}) error {
	incrementCollateral, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if incrementCollateral == emptyDec || incrementCollateral.IsNil() {
		return errors.New("collateral auction increment cannot be nil or empty")
	}

	if incrementCollateral.IsNegative() {
		return fmt.Errorf("collateral auction increment cannot be less than zero %s", incrementCollateral)
	}

	return nil
}
