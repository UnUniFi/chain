package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/testutil/sample"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.Params{
					PoolParams: types.PoolParams{
						QuoteTicker:                 "123",
						BaseLptMintFee:              sdk.MustNewDecFromStr("0.001"),
						BaseLptRedeemFee:            sdk.MustNewDecFromStr("0.001"),
						BorrowingFeeRatePerHour:     sdk.MustNewDecFromStr("0.001"),
						ReportLiquidationRewardRate: sdk.MustNewDecFromStr("0.001"),
						ReportLevyPeriodRewardRate:  sdk.MustNewDecFromStr("0.001"),
						AcceptedAssetsConf: []types.PoolAssetConf{
							{
								Denom:        "uatom",
								TargetWeight: sdk.MustNewDecFromStr("0.001"),
							},
						},
						LevyPeriodRequiredSeconds: 3600,
					},
					PerpetualFutures: types.PerpetualFuturesParams{
						CommissionRate:        sdk.MustNewDecFromStr("0.001"),
						MarginMaintenanceRate: sdk.MustNewDecFromStr("0.001"),
						ImaginaryFundingRateProportionalCoefficient: sdk.MustNewDecFromStr("0.001"),
						Markets: []*types.Market{
							{
								BaseDenom:  "uatom",
								QuoteDenom: "uusdc",
							},
						},
						MaxLeverage: 1,
					},
					PerpetualOptions: types.PerpetualOptionsParams{
						PremiumCommissionRate:                       sdk.MustNewDecFromStr("0.001"),
						StrikeCommissionRate:                        sdk.MustNewDecFromStr("0.001"),
						MarginMaintenanceRate:                       sdk.MustNewDecFromStr("0.001"),
						ImaginaryFundingRateProportionalCoefficient: sdk.MustNewDecFromStr("0.001"),
						Markets: []*types.Market{
							{
								BaseDenom:  "uatom",
								QuoteDenom: "uusdc",
							},
						},
					},
				},
				Positions: []types.Position{
					{
						Id: "1",
						Market: types.Market{
							BaseDenom:  "uatom",
							QuoteDenom: "uusdc",
						},
						OpenerAddress:   sample.AccAddress(),
						OpenedAt:        time.Now().UTC(),
						OpenedHeight:    10,
						OpenedBaseRate:  sdk.MustNewDecFromStr("0.001"),
						OpenedQuoteRate: sdk.MustNewDecFromStr("0.001"),
						RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
					},
					{
						Id: "2",
						Market: types.Market{
							BaseDenom:  "uatom",
							QuoteDenom: "uusdc",
						},
						OpenerAddress:   sample.AccAddress(),
						OpenedAt:        time.Now().UTC(),
						OpenedHeight:    10,
						OpenedBaseRate:  sdk.MustNewDecFromStr("0.001"),
						OpenedQuoteRate: sdk.MustNewDecFromStr("0.001"),
						RemainingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000)),
					},
				},
				PoolMarketCap: types.PoolMarketCap{
					QuoteTicker: "ticker",
					Total:       sdk.MustNewDecFromStr("0.111"),
					AssetInfo: []types.PoolMarketCap_AssetInfo{
						{
							Denom:    "denom",
							Amount:   sdk.NewInt(1000),
							Price:    sdk.MustNewDecFromStr("0.001"),
							Reserved: sdk.NewInt(1000),
						},
					},
				},
				PerpetualFuturesGrossPositionOfMarket: []types.PerpetualFuturesGrossPositionOfMarket{
					{
						Market: types.Market{
							BaseDenom:  "uatom",
							QuoteDenom: "uusdc",
						},
						PositionType:                types.PositionType_LONG,
						PositionSizeInDenomExponent: sdk.NewInt(1000),
					},
				},
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
