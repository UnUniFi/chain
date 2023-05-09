package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

// FIXME: This test is not working
func (suite *KeeperTestSuite) TestParamsGetSet() {
	params := suite.app.DerivativesKeeper.GetParams(suite.ctx)
	params.PoolParams = types.PoolParams{
		QuoteTicker:                 "usd",
		BaseLptMintFee:              sdk.MustNewDecFromStr("0.001"),
		BaseLptRedeemFee:            sdk.MustNewDecFromStr("0.001"),
		BorrowingFeeRatePerHour:     sdk.MustNewDecFromStr("0.001"),
		ReportLiquidationRewardRate: sdk.MustNewDecFromStr("0.3"),
		ReportLevyPeriodRewardRate:  sdk.MustNewDecFromStr("0.3"),
		AcceptedAssetsConf: []types.PoolAssetConf{
			{
				Denom:        "uatom",
				TargetWeight: sdk.OneDec(),
			},
		},
	}

	params.PerpetualFutures = types.PerpetualFuturesParams{
		CommissionRate:        sdk.MustNewDecFromStr("0.001"),
		MarginMaintenanceRate: sdk.MustNewDecFromStr("0.5"),
		ImaginaryFundingRateProportionalCoefficient: sdk.MustNewDecFromStr("0.0005"),
		Markets: []*types.Market{
			{
				BaseDenom:  "uatom",
				QuoteDenom: "uusdc",
			},
		},
	}

	params.PerpetualOptions = types.PerpetualOptionsParams{
		PremiumCommissionRate:                       sdk.ZeroDec(),
		StrikeCommissionRate:                        sdk.MustNewDecFromStr("0.001"),
		MarginMaintenanceRate:                       sdk.ZeroDec(),
		ImaginaryFundingRateProportionalCoefficient: sdk.ZeroDec(),
		Markets: []*types.Market{
			{
				BaseDenom:  "uatom",
				QuoteDenom: "uusdc",
			},
		},
	}

	suite.app.DerivativesKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.DerivativesKeeper.GetParams(suite.ctx)
	suite.Require().Equal(newParams, params)
}
