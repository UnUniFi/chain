package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (p PoolParams) Validate() error {
	if p.QuoteTicker != "USD" {
		return fmt.Errorf("USD is only allowed for quote ticker")
	}

	if p.BaseLptMintFee.IsNil() || p.BaseLptMintFee.IsNegative() || p.BaseLptMintFee.GT(sdk.OneDec()) {
		return fmt.Errorf("BaseLptMintFee should be between 0-1")
	}

	if p.BaseLptRedeemFee.IsNil() || p.BaseLptRedeemFee.IsNegative() || p.BaseLptRedeemFee.GT(sdk.OneDec()) {
		return fmt.Errorf("BaseLptRedeemFee should be between 0-1")
	}

	if p.BorrowingFeeRatePerHour.IsNil() || p.BorrowingFeeRatePerHour.IsNegative() || p.BorrowingFeeRatePerHour.GT(sdk.OneDec()) {
		return fmt.Errorf("BorrowingFeeRatePerHour should be between 0-1")
	}

	if p.ReportLiquidationRewardRate.IsNil() || p.ReportLiquidationRewardRate.IsNegative() || p.ReportLiquidationRewardRate.GT(sdk.OneDec()) {
		return fmt.Errorf("ReportLiquidationRewardRate should be between 0-1")
	}

	if p.ReportLevyPeriodRewardRate.IsNil() || p.ReportLevyPeriodRewardRate.IsNegative() || p.ReportLevyPeriodRewardRate.GT(sdk.OneDec()) {
		return fmt.Errorf("ReportLevyPeriodRewardRate should be between 0-1")
	}

	if len(p.AcceptedAssetsConf) == 0 {
		return fmt.Errorf("Empty AcceptedAssets")
	}

	usedDenom := make(map[string]bool)
	sumWeight := sdk.ZeroDec()
	for _, asset := range p.AcceptedAssetsConf {
		if usedDenom[asset.Denom] {
			return fmt.Errorf("Duplication in accepted denom: %s", asset.Denom)
		}
		if asset.TargetWeight.IsNil() || asset.TargetWeight.IsNegative() || asset.TargetWeight.GT(sdk.OneDec()) {
			return fmt.Errorf("Asset weight should be between 0-1")
		}
		usedDenom[asset.Denom] = true
		sumWeight = sumWeight.Add(asset.TargetWeight)
	}

	if !sumWeight.Equal(sdk.OneDec()) {
		return fmt.Errorf("Sum of accepted assets weight should be 1")
	}
	return nil
}

func IsValidDepositForPool(deposit sdk.Coin, acceptableAssets []PoolAssetConf) bool {
	for _, asset := range acceptableAssets {
		if deposit.Denom == asset.Denom {
			return deposit.Amount.IsPositive()
		}
	}
	return false
}
