package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys and default values
var (
	DefaultBidToken = "uguu"
)

// NewParams returns a new params object
func NewParams() Params {
	return Params{
		MinStakingForListing:               sdk.ZeroInt(),
		DefaultBidActiveRank:               1,
		BidTokens:                          []string{DefaultBidToken},
		AutoRelistingCountIfNoBid:          2,
		NftListingDelaySeconds:             10,
		NftListingPeriodInitial:            200,
		NftListingCancelRequiredSeconds:    20,
		NftListingCancelFeePercentage:      5,
		NftListingGapTime:                  20,
		BidCancelRequiredSeconds:           20,
		BidTokenDisburseSecondsAfterCancel: 20,
		NftListingFullPaymentPeriod:        30,
		NftListingNftDeliveryPeriod:        30,
		NftCreatorSharePercentage:          5,
		MarketAdministrator:                "",
		NftListingCommissionFee:            5,
		NftListingExtendSeconds:            30,
		NftListingPeriodExtendFeePerHour:   sdk.NewInt64Coin(DefaultBidToken, 1000000),
	}
}

// DefaultParams returns default params for incentive module
func DefaultParams() Params {
	return Params{
		MinStakingForListing:               sdk.ZeroInt(),
		DefaultBidActiveRank:               1,
		BidTokens:                          []string{DefaultBidToken},
		AutoRelistingCountIfNoBid:          2,
		NftListingDelaySeconds:             10,
		NftListingPeriodInitial:            200,
		NftListingCancelRequiredSeconds:    20,
		NftListingCancelFeePercentage:      5,
		NftListingGapTime:                  20,
		BidCancelRequiredSeconds:           20,
		BidTokenDisburseSecondsAfterCancel: 20,
		NftListingFullPaymentPeriod:        30,
		NftListingNftDeliveryPeriod:        30,
		NftCreatorSharePercentage:          5,
		MarketAdministrator:                "",
		NftListingCommissionFee:            5,
		NftListingExtendSeconds:            30,
		NftListingPeriodExtendFeePerHour:   sdk.NewInt64Coin(DefaultBidToken, 1000000),
	}
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// Parameter keys
var (
	KeyMinStakingForListing               = []byte("MinStakingForListing")
	KeyDefaultBidActiveRank               = []byte("DefaultBidActiveRank")
	KeyBidTokens                          = []byte("BidTokens")
	KeyAutoRelistingCountIfNoBid          = []byte("AutoRelistingCountIfNoBid")
	KeyNftListingDelaySeconds             = []byte("NftListingDelaySeconds")
	KeyNftListingPeriodInitial            = []byte("NftListingPeriodInitial")
	KeyNftListingCancelRequiredSeconds    = []byte("NftListingCancelRequiredSeconds")
	KeyNftListingCancelFeePercentage      = []byte("NftListingCancelFeePercentage")
	KeyNftListingGapTime                  = []byte("NftListingGapTime")
	KeyBidCancelRequiredSeconds           = []byte("BidCancelRequiredSeconds")
	KeyBidTokenDisburseSecondsAfterCancel = []byte("BidTokenDisburseSecondsAfterCancel")
	KeyNftListingFullPaymentPeriod        = []byte("NftListingFullPaymentPeriod")
	KeyNftListingNftDeliveryPeriod        = []byte("NftListingNftDeliveryPeriod")
	KeyNftCreatorSharePercentage          = []byte("NftCreatorSharePercentage")
	KeyMarketAdministrator                = []byte("MarketAdministrator")
	KeyNftListingCommissionFee            = []byte("NftListingCommissionFee")
	KeyNftListingExtendSeconds            = []byte("NftListingExtendSeconds")
	KeyNftListingPeriodExtendFeePerHour   = []byte("NftListingPeriodExtendFeePerHour")
)

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{
		paramstype.NewParamSetPair(KeyMinStakingForListing, &p.MinStakingForListing, validateMinStakingForListing),
		paramstype.NewParamSetPair(KeyDefaultBidActiveRank, &p.DefaultBidActiveRank, validateDefaultBidActiveRank),
		paramstype.NewParamSetPair(KeyBidTokens, &p.BidTokens, validateBidTokens),
		paramstype.NewParamSetPair(KeyAutoRelistingCountIfNoBid, &p.AutoRelistingCountIfNoBid, validateAutoRelistingCountIfNoBid),
		paramstype.NewParamSetPair(KeyNftListingDelaySeconds, &p.NftListingDelaySeconds, validateNftListingDelaySeconds),
		paramstype.NewParamSetPair(KeyNftListingPeriodInitial, &p.NftListingPeriodInitial, validateNftListingPeriodInitial),
		paramstype.NewParamSetPair(KeyNftListingCancelRequiredSeconds, &p.NftListingCancelRequiredSeconds, validateNftListingCancelRequiredSeconds),
		paramstype.NewParamSetPair(KeyNftListingCancelFeePercentage, &p.NftListingCancelFeePercentage, validateNftListingCancelFeePercentage),
		paramstype.NewParamSetPair(KeyNftListingGapTime, &p.NftListingGapTime, validateNftListingGapTime),
		paramstype.NewParamSetPair(KeyBidCancelRequiredSeconds, &p.BidCancelRequiredSeconds, validateBidCancelRequiredSeconds),
		paramstype.NewParamSetPair(KeyBidTokenDisburseSecondsAfterCancel, &p.BidTokenDisburseSecondsAfterCancel, validateBidTokenDisburseSecondsAfterCancel),
		paramstype.NewParamSetPair(KeyNftListingFullPaymentPeriod, &p.NftListingFullPaymentPeriod, validateNftListingFullPaymentPeriod),
		paramstype.NewParamSetPair(KeyNftListingNftDeliveryPeriod, &p.NftListingNftDeliveryPeriod, validateNftListingNftDeliveryPeriod),
		paramstype.NewParamSetPair(KeyNftCreatorSharePercentage, &p.NftCreatorSharePercentage, validateNftCreatorSharePercentage),
		paramstype.NewParamSetPair(KeyMarketAdministrator, &p.MarketAdministrator, validateMarketAdministrator),
		paramstype.NewParamSetPair(KeyNftListingCommissionFee, &p.NftListingCommissionFee, validateNftListingCommissionFee),
		paramstype.NewParamSetPair(KeyNftListingExtendSeconds, &p.NftListingExtendSeconds, validateNftListingExtendSeconds),
		paramstype.NewParamSetPair(KeyNftListingPeriodExtendFeePerHour, &p.NftListingPeriodExtendFeePerHour, validateNftListingPeriodExtendFeePerHour),
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {

	if err := validateMinStakingForListing(p.MinStakingForListing); err != nil {
		return err
	}

	if err := validateDefaultBidActiveRank(p.DefaultBidActiveRank); err != nil {
		return err
	}

	if err := validateBidTokens(p.BidTokens); err != nil {
		return err
	}

	if err := validateAutoRelistingCountIfNoBid(p.AutoRelistingCountIfNoBid); err != nil {
		return err
	}

	if err := validateNftListingDelaySeconds(p.NftListingDelaySeconds); err != nil {
		return err
	}

	if err := validateNftListingPeriodInitial(p.NftListingPeriodInitial); err != nil {
		return err
	}

	if err := validateNftListingCancelRequiredSeconds(p.BidCancelRequiredSeconds); err != nil {
		return err
	}

	if err := validateNftListingCancelFeePercentage(p.NftListingCancelFeePercentage); err != nil {
		return err
	}

	if err := validateNftListingGapTime(p.NftListingGapTime); err != nil {
		return err
	}

	if err := validateBidCancelRequiredSeconds(p.BidCancelRequiredSeconds); err != nil {
		return err
	}

	if err := validateBidTokenDisburseSecondsAfterCancel(p.BidTokenDisburseSecondsAfterCancel); err != nil {
		return err
	}

	if err := validateNftListingFullPaymentPeriod(p.NftListingFullPaymentPeriod); err != nil {
		return err
	}

	if err := validateNftListingNftDeliveryPeriod(p.NftListingNftDeliveryPeriod); err != nil {
		return err
	}

	if err := validateNftCreatorSharePercentage(p.NftCreatorSharePercentage); err != nil {
		return err
	}

	if err := validateMarketAdministrator(p.MarketAdministrator); err != nil {
		return err
	}

	if err := validateNftListingCommissionFee(p.NftListingCommissionFee); err != nil {
		return err
	}

	if err := validateNftListingExtendSeconds(p.NftListingExtendSeconds); err != nil {
		return err
	}

	if err := validateNftListingPeriodExtendFeePerHour(p.NftListingPeriodExtendFeePerHour); err != nil {
		return err
	}

	return nil
}

func validateMinStakingForListing(i interface{}) error {
	_, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateDefaultBidActiveRank(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateBidTokens(i interface{}) error {
	_, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateAutoRelistingCountIfNoBid(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingDelaySeconds(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingPeriodInitial(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingCancelRequiredSeconds(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingCancelFeePercentage(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingGapTime(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateBidCancelRequiredSeconds(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateBidTokenDisburseSecondsAfterCancel(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingFullPaymentPeriod(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingNftDeliveryPeriod(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftCreatorSharePercentage(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateMarketAdministrator(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingCommissionFee(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingExtendSeconds(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNftListingPeriodExtendFeePerHour(i interface{}) error {
	_, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
