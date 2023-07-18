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
	return DefaultParams()
}

// DefaultParams returns default params for incentive module
func DefaultParams() Params {
	return Params{
		MinStakingForListing:            sdk.ZeroInt(),
		BidTokens:                       []string{DefaultBidToken},
		NftListingCancelRequiredSeconds: 20,
		BidCancelRequiredSeconds:        20,
		NftListingFullPaymentPeriod:     30,
		NftListingNftDeliveryPeriod:     30,
		NftListingCommissionFee:         5,
	}
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// Parameter keys
var (
	KeyMinStakingForListing            = []byte("MinStakingForListing")
	KeyBidTokens                       = []byte("BidTokens")
	KeyNftListingCancelRequiredSeconds = []byte("NftListingCancelRequiredSeconds")
	KeyBidCancelRequiredSeconds        = []byte("BidCancelRequiredSeconds")
	KeyNftListingFullPaymentPeriod     = []byte("NftListingFullPaymentPeriod")
	KeyNftListingNftDeliveryPeriod     = []byte("NftListingNftDeliveryPeriod")
	KeyNftListingCommissionFee         = []byte("NftListingCommissionFee")
)

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{
		paramstype.NewParamSetPair(KeyMinStakingForListing, &p.MinStakingForListing, validateMinStakingForListing),
		paramstype.NewParamSetPair(KeyBidTokens, &p.BidTokens, validateBidTokens),
		paramstype.NewParamSetPair(KeyNftListingCancelRequiredSeconds, &p.NftListingCancelRequiredSeconds, validateNftListingCancelRequiredSeconds),
		paramstype.NewParamSetPair(KeyBidCancelRequiredSeconds, &p.BidCancelRequiredSeconds, validateBidCancelRequiredSeconds),
		paramstype.NewParamSetPair(KeyNftListingFullPaymentPeriod, &p.NftListingFullPaymentPeriod, validateNftListingFullPaymentPeriod),
		paramstype.NewParamSetPair(KeyNftListingNftDeliveryPeriod, &p.NftListingNftDeliveryPeriod, validateNftListingNftDeliveryPeriod),
		paramstype.NewParamSetPair(KeyNftListingCommissionFee, &p.NftListingCommissionFee, validateNftListingCommissionFee),
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {

	if err := validateMinStakingForListing(p.MinStakingForListing); err != nil {
		return err
	}

	if err := validateBidTokens(p.BidTokens); err != nil {
		return err
	}

	if err := validateNftListingCancelRequiredSeconds(p.BidCancelRequiredSeconds); err != nil {
		return err
	}

	if err := validateBidCancelRequiredSeconds(p.BidCancelRequiredSeconds); err != nil {
		return err
	}

	if err := validateNftListingFullPaymentPeriod(p.NftListingFullPaymentPeriod); err != nil {
		return err
	}

	if err := validateNftListingNftDeliveryPeriod(p.NftListingNftDeliveryPeriod); err != nil {
		return err
	}

	if err := validateNftListingCommissionFee(p.NftListingCommissionFee); err != nil {
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

func validateBidTokens(i interface{}) error {
	_, ok := i.([]string)
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

func validateBidCancelRequiredSeconds(i interface{}) error {
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

func validateNftListingCommissionFee(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
