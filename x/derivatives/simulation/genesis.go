package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func RandomGenesisBool(r *rand.Rand) bool {
	// 90% chance
	return r.Int63n(100) < 90
}

func RandomizedGenState(simState *module.SimulationState) {
	sdk.NewCoins()
	// numAccs := int64(len(simState.Accounts))

	derivativesGenesis := types.GenesisState{
		Params: types.Params{
			Pool: types.Pool{
				QuoteTicker:                       "USD",
				BaseLptMintFee:                    sdk.NewDecWithPrec(1, 2),
				BaseLptRedeemFee:                  sdk.NewDecWithPrec(1, 2),
				BorrowingFeeRatePerHour:           sdk.NewDecWithPrec(1, 6),
				LiquidationNeededReportRewardRate: sdk.NewDecWithPrec(1, 6),
				AcceptedAssets: []*types.Pool_Asset{
					{
						Denom:        "btc",
						TargetWeight: sdk.NewDecWithPrec(1, 2),
					},
					{
						Denom:        "eth",
						TargetWeight: sdk.NewDecWithPrec(1, 3),
					},
				},
			},
			PerpetualFutures: types.PerpetualFuturesParams{
				CommissionRate:        sdk.NewDecWithPrec(1, 6),
				MarginMaintenanceRate: sdk.NewDecWithPrec(5, 1),
				ImaginaryFundingRateProportionalCoefficient: sdk.NewDecWithPrec(1, 4),
				Markets: []*types.Market{
					{
						BaseDenom:  "btc",
						QuoteDenom: "usd",
					},
					{
						BaseDenom:  "eth",
						QuoteDenom: "usd",
					},
				},
			},
		},
	}

	paramsBytes, err := json.MarshalIndent(&derivativesGenesis.Params, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated derivatives parameters:\n%s\n", paramsBytes)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&derivativesGenesis)
}
