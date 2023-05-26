package keeper_test

// import (
// 	"github.com/cometbft/cometbft/crypto/ed25519"
// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/UnUniFi/chain/deprecated/yieldaggregatorv1/types"
// )

// func (suite *KeeperTestSuite) TestGenesis() {
// 	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
// 	genesisState := types.GenesisState{
// 		Params: types.Params{
// 			RewardRateFeeders: addr1.String(),
// 		},
// 		AssetManagementAccounts: []types.AssetManagementAccount{
// 			{
// 				Id:      "OsmoFarm",
// 				Name:    "Osmosis Farm",
// 				Enabled: true,
// 			},
// 		},
// 		AssetManagementTargets: []types.AssetManagementTarget{
// 			{
// 				Id:                       "OsmoGUUFarm",
// 				AssetManagementAccountId: "OsmoFarm",
// 				Enabled:                  true,
// 				AssetConditions: []types.AssetCondition{
// 					{
// 						Denom: "uguu",
// 						Ratio: 100,
// 						Min:   "",
// 					},
// 				},
// 			},
// 		},
// 		FarmingOrders: []types.FarmingOrder{
// 			{
// 				Id:          "OsmoGUUFarmOrder",
// 				FromAddress: addr1.String(),
// 				Strategy: types.Strategy{
// 					StrategyType:         "Manual",
// 					WhitelistedTargetIds: []string{"OsmoGUUFarm"},
// 				},
// 			},
// 		},
// 		FarmingUnits: []types.FarmingUnit{
// 			{
// 				AccountId: "OsmoFarm",
// 				TargetId:  "OsmoGUUFarm",
// 				Owner:     addr1.String(),
// 			},
// 		},
// 		UserDeposits: []types.UserDeposit{
// 			{
// 				User:   addr1.String(),
// 				Amount: sdk.Coins{sdk.NewInt64Coin("uguu", 1)},
// 			},
// 		},
// 		DailyPercents: []types.DailyPercent{
// 			{
// 				AccountId: "OsmoFarm",
// 				TargetId:  "OsmoGUUFarm",
// 				Rate:      sdk.NewDecWithPrec(1, 1),
// 			},
// 		},
// 	}

// 	suite.app.YieldaggregatorKeeper.InitGenesis(suite.ctx, genesisState)

// 	exportedGenesis := suite.app.YieldaggregatorKeeper.ExportGenesis(suite.ctx)
// 	suite.Require().Equal(genesisState, *exportedGenesis)
// }
