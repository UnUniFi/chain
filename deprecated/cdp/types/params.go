package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	KeyGlobalDebtLimit      = []byte("GlobalDebtLimit")
	KeyCollateralParams     = []byte("CollateralParams")
	KeyDebtParam            = []byte("DebtParam")
	KeyCircuitBreaker       = []byte("CircuitBreaker")
	KeyDebtThreshold        = []byte("DebtThreshold")
	KeyDebtLot              = []byte("DebtLot")
	KeySurplusThreshold     = []byte("SurplusThreshold")
	KeySurplusLot           = []byte("SurplusLot")
	DefaultCollateralParams = CollateralParams{}
	DefaultDebtParams       = DebtParams{
		DebtParam{
			Denom:                   "jpu",
			ReferenceAsset:          "jpy",
			ConversionFactor:        sdk.NewInt(6),
			DebtFloor:               sdk.NewInt(1),
			GlobalDebtLimit:         sdk.NewCoin(DefaultStableDenom, sdk.ZeroInt()),
			DebtDenom:               "debtjpu",
			SurplusAuctionThreshold: sdk.NewInt(500000000000),
			SurplusAuctionLot:       sdk.NewInt(10000000000),
			DebtAuctionThreshold:    sdk.NewInt(100000000000),
			DebtAuctionLot:          sdk.NewInt(10000000000),
			CircuitBreaker:          false,
		},
		DebtParam{
			Denom:                   "euu",
			ReferenceAsset:          "eur",
			ConversionFactor:        sdk.NewInt(6),
			DebtFloor:               sdk.NewInt(1),
			GlobalDebtLimit:         sdk.NewCoin("euu", sdk.ZeroInt()),
			DebtDenom:               "debteuu",
			SurplusAuctionThreshold: sdk.NewInt(500000000000),
			SurplusAuctionLot:       sdk.NewInt(10000000000),
			DebtAuctionThreshold:    sdk.NewInt(100000000000),
			DebtAuctionLot:          sdk.NewInt(10000000000),
			CircuitBreaker:          false,
		},
	}
	DefaultCdpStartingID = uint64(1)
	DefaultGovDenom      = "uguu"
	DefaultStableDenom   = "jpu"
	minCollateralPrefix  = 0
	maxCollateralPrefix  = 255
	stabilityFeeMax      = sdk.MustNewDecFromStr("1.000000051034942716") // 500% APR
)

// NewParams returns a new params object
func NewParams(
	collateralParams CollateralParams, debtParams DebtParams,
) Params {
	return Params{
		CollateralParams: collateralParams,
		DebtParams:       debtParams,
	}
}

// DefaultParams returns default params for cdp module
func DefaultParams() Params {
	return NewParams(
		DefaultCollateralParams, DefaultDebtParams,
	)
}

// NewCollateralParam returns a new CollateralParam
func NewCollateralParam(
	denom, ctype string, liqRatio sdk.Dec, debtLimit sdk.Coin, stabilityFee sdk.Dec, auctionSize sdk.Int,
	liqPenalty sdk.Dec, prefix byte, spotMarketID, liquidationMarketID string, keeperReward sdk.Dec, checkIndexCount sdk.Int, conversionFactor sdk.Int) CollateralParam {
	return CollateralParam{
		Denom:                            denom,
		Type:                             ctype,
		LiquidationRatio:                 liqRatio,
		DebtLimit:                        debtLimit,
		StabilityFee:                     stabilityFee,
		AuctionSize:                      auctionSize,
		LiquidationPenalty:               liqPenalty,
		Prefix:                           uint32(prefix),
		SpotMarketId:                     spotMarketID,
		LiquidationMarketId:              liquidationMarketID,
		KeeperRewardPercentage:           keeperReward,
		CheckCollateralizationIndexCount: checkIndexCount,
		ConversionFactor:                 conversionFactor,
	}
}

// CollateralParams array of CollateralParam
type CollateralParams []CollateralParam

// String implements fmt.Stringer
func (cps CollateralParams) String() string {
	out := "Collateral Params\n"
	for _, cp := range cps {
		out += fmt.Sprintf("%s\n", cp.String())
	}
	return out
}

// NewDebtParam returns a new DebtParam
func NewDebtParam(denom, refAsset string, conversionFactor, debtFloor sdk.Int, globalDebtLimit sdk.Coin) DebtParam {
	return DebtParam{
		Denom:            denom,
		ReferenceAsset:   refAsset,
		ConversionFactor: conversionFactor,
		DebtFloor:        debtFloor,
		GlobalDebtLimit:  globalDebtLimit,
	}
}

// DebtParams array of DebtParam
type DebtParams []DebtParam

// String implements fmt.Stringer
func (dps DebtParams) String() string {
	out := "Debt Params\n"
	for _, dp := range dps {
		out += fmt.Sprintf("%v\n", dp)
	}
	return out
}

// find the element that returns TRUE in the injected function
func (dps DebtParams) find(f func(DebtParam, string) bool, search_denom string) (DebtParam, bool) {
	for _, debtParam := range dps {
		if (DebtParam{}) != debtParam {
			if f(debtParam, search_denom) {
				return debtParam, true
			}
		}
	}
	return DebtParam{}, false
}

// Returns an element with a denom that matches search_denom
func (dps DebtParams) FindDenom(search_denom string) (DebtParam, bool) {
	findDenom := func(ele DebtParam, target string) bool {
		return ele.Denom == target
	}
	return dps.find(findDenom, search_denom)
}

// Returns an element with a GlobalDebtLimit.Denom that matches search_denom
func (dps DebtParams) FindGlobalDebtLimitDenom(search_denom string) (DebtParam, bool) {
	FindGlobalDebtLimitDenomDenom := func(ele DebtParam, target string) bool {
		return ele.GlobalDebtLimit.Denom == target
	}
	return dps.find(FindGlobalDebtLimitDenomDenom, search_denom)
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{
		paramstype.NewParamSetPair(KeyCollateralParams, &p.CollateralParams, validateCollateralParams),
		paramstype.NewParamSetPair(KeyDebtParam, &p.DebtParams, validateDebtParams),
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {

	if err := validateCollateralParams(p.CollateralParams); err != nil {
		return err
	}

	if err := validateDebtParams(p.DebtParams); err != nil {
		return err
	}

	if len(p.CollateralParams) == 0 { // default value OK
		return nil
	}

	for _, debtParam := range p.DebtParams {
		if (DebtParam{}) != debtParam {
			if debtParam.Denom != debtParam.GlobalDebtLimit.Denom {
				return fmt.Errorf("debt denom %s does not match global debt denom %s",
					debtParam.Denom, debtParam.GlobalDebtLimit.Denom)
			}
		}
	}

	// validate collateral params
	collateralTypeDupMap := make(map[string]bool)
	collateralParamsDebtLimitMap := make(map[string]sdk.Int)

	for _, cp := range p.CollateralParams {
		// Collateral type eg busd-a should be unique, but denom can be same eg busd
		_, typeExists := collateralTypeDupMap[cp.Type]
		if typeExists {
			return fmt.Errorf("duplicate collateral type %s", cp.Denom)
		}
		collateralTypeDupMap[cp.Type] = true

		collateralParamsDebtLimit, denomLimitExists := collateralParamsDebtLimitMap[cp.DebtLimit.Denom]
		if !denomLimitExists {
			collateralParamsDebtLimit = sdk.ZeroInt()
		}

		debtParams := DebtParams(p.DebtParams)
		debtParam, exists := debtParams.FindGlobalDebtLimitDenom(cp.DebtLimit.Denom)

		if !exists {
			return fmt.Errorf("collateral debt limit denom %s does not match global debt limit denoms.", cp.DebtLimit.Denom)
		}

		collateralParamsDebtLimitMap[cp.DebtLimit.Denom] = collateralParamsDebtLimit.Add(cp.DebtLimit.Amount)

		if cp.DebtLimit.Amount.GT(debtParam.GlobalDebtLimit.Amount) {
			return fmt.Errorf("collateral debt limit %s exceeds global debt limit: %s denom: %s",
				cp.DebtLimit, debtParam.GlobalDebtLimit.Amount, debtParam.GlobalDebtLimit.Denom)
		}
	}

	for _, debtParam := range p.DebtParams {
		collateralParamsDebtLimit, exists := collateralParamsDebtLimitMap[debtParam.GlobalDebtLimit.Denom]
		if !exists {
			continue
		}
		if collateralParamsDebtLimit.GT(debtParam.GlobalDebtLimit.Amount) {
			return fmt.Errorf("sum of collateral debt limits %s exceeds global debt limit %s denom: %s",
				collateralParamsDebtLimit, debtParam.GlobalDebtLimit.Amount, debtParam.GlobalDebtLimit.Denom)
		}
	}

	return nil
}

func validateGlobalDebtLimitParam(i interface{}) error {
	globalDebtLimit, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !globalDebtLimit.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "global debt limit %s", globalDebtLimit.String())
	}

	return nil
}

func validateCollateralParams(i interface{}) error {
	collateralParams, ok := i.(CollateralParams)
	if !ok {
		collateralParams, ok = i.([]CollateralParam)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
	}

	prefixDupMap := make(map[int]bool)
	typeDupMap := make(map[string]bool)
	for _, cp := range collateralParams {
		if err := sdk.ValidateDenom(cp.Denom); err != nil {
			return fmt.Errorf("collateral denom invalid %s", cp.Denom)
		}

		if strings.TrimSpace(cp.SpotMarketId) == "" {
			return fmt.Errorf("spot market id cannot be blank %s", cp.String())
		}

		if strings.TrimSpace(cp.Type) == "" {
			return fmt.Errorf("collateral type cannot be blank %s", cp.String())
		}

		if strings.TrimSpace(cp.LiquidationMarketId) == "" {
			return fmt.Errorf("liquidation market id cannot be blank %s", cp.String())
		}

		prefix := int(cp.Prefix)
		if prefix < minCollateralPrefix || prefix > maxCollateralPrefix {
			return fmt.Errorf("invalid prefix for collateral denom %s: %b", cp.Denom, cp.Prefix)
		}

		_, found := prefixDupMap[prefix]
		if found {
			return fmt.Errorf("duplicate prefix for collateral denom %s: %v", cp.Denom, []byte{byte(cp.Prefix)})
		}

		prefixDupMap[prefix] = true

		_, found = typeDupMap[cp.Type]
		if found {
			return fmt.Errorf("duplicate cdp collateral type: %s", cp.Type)
		}
		typeDupMap[cp.Type] = true

		if !cp.DebtLimit.IsValid() {
			return fmt.Errorf("debt limit for all collaterals should be positive, is %s for %s", cp.DebtLimit, cp.Denom)
		}

		if cp.LiquidationPenalty.LT(sdk.ZeroDec()) || cp.LiquidationPenalty.GT(sdk.OneDec()) {
			return fmt.Errorf("liquidation penalty should be between 0 and 1, is %s for %s", cp.LiquidationPenalty, cp.Denom)
		}
		if !cp.AuctionSize.IsPositive() {
			return fmt.Errorf("auction size should be positive, is %s for %s", cp.AuctionSize, cp.Denom)
		}
		if cp.StabilityFee.LT(sdk.OneDec()) || cp.StabilityFee.GT(stabilityFeeMax) {
			return fmt.Errorf("stability fee must be ≥ 1.0, ≤ %s, is %s for %s", stabilityFeeMax, cp.StabilityFee, cp.Denom)
		}
		if cp.KeeperRewardPercentage.IsNegative() || cp.KeeperRewardPercentage.GT(sdk.OneDec()) {
			return fmt.Errorf("keeper reward percentage should be between 0 and 1, is %s for %s", cp.KeeperRewardPercentage, cp.Denom)
		}
		if cp.CheckCollateralizationIndexCount.IsNegative() {
			return fmt.Errorf("keeper reward percentage should be positive, is %s for %s", cp.CheckCollateralizationIndexCount, cp.Denom)
		}
	}

	return nil
}

func validateDebtParams(i interface{}) error {
	debtParams, ok := i.(DebtParams)
	if !ok {
		debtParams, ok = i.([]DebtParam)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
	}
	for _, debtParam := range debtParams {
		if err := sdk.ValidateDenom(debtParam.Denom); err != nil {
			return fmt.Errorf("debtParam Denom invalid %s", debtParam.Denom)
		}

		if err := validateGlobalDebtLimitParam(debtParam.GlobalDebtLimit); err != nil {
			return err
		}

		if err := sdk.ValidateDenom(debtParam.DebtDenom); err != nil {
			return fmt.Errorf("debtParam DebtDenom invalid %s", debtParam.DebtDenom)
		}

		if err := validateCircuitBreakerParam(debtParam.CircuitBreaker); err != nil {
			return err
		}

		if err := validateSurplusAuctionThresholdParam(debtParam.SurplusAuctionThreshold); err != nil {
			return err
		}

		if err := validateSurplusAuctionLotParam(debtParam.SurplusAuctionLot); err != nil {
			return err
		}

		if err := validateDebtAuctionThresholdParam(debtParam.DebtAuctionThreshold); err != nil {
			return err
		}

		if err := validateDebtAuctionLotParam(debtParam.DebtAuctionLot); err != nil {
			return err
		}
	}

	return nil
}

func validateCircuitBreakerParam(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateSurplusAuctionThresholdParam(i interface{}) error {
	sat, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !sat.IsPositive() {
		return fmt.Errorf("surplus auction threshold should be positive: %s", sat)
	}

	return nil
}

func validateSurplusAuctionLotParam(i interface{}) error {
	sal, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !sal.IsPositive() {
		return fmt.Errorf("surplus auction lot should be positive: %s", sal)
	}

	return nil
}

func validateDebtAuctionThresholdParam(i interface{}) error {
	dat, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !dat.IsPositive() {
		return fmt.Errorf("debt auction threshold should be positive: %s", dat)
	}

	return nil
}

func validateDebtAuctionLotParam(i interface{}) error {
	dal, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !dal.IsPositive() {
		return fmt.Errorf("debt auction lot should be positive: %s", dal)
	}

	return nil
}
