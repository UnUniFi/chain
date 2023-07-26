package keeper_test

// import (
// 	"time"

// 	"github.com/cometbft/cometbft/crypto/ed25519"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

// 	"github.com/UnUniFi/chain/deprecated/x/yieldaggregatorv1/types"
// 	"github.com/UnUniFi/chain/deprecated/x/yieldfarm"
// 	yieldfarmtypes "github.com/UnUniFi/chain/deprecated/x/yieldfarm/types"
// )

// func (suite *KeeperTestSuite) TestInvestmentFlow() {
// 	farmer := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

// 	now := time.Now().UTC()
// 	suite.ctx = suite.ctx.WithBlockTime(now)

// 	// create asset management account
// 	amAcc := types.AssetManagementAccount{
// 		Id:   "OsmosisAssetManager",
// 		Name: "Osmosis Asset Management Account",
// 	}
// 	suite.app.YieldaggregatorKeeper.SetAssetManagementAccount(suite.ctx, amAcc)

// 	// create asset management target
// 	amTarget := types.AssetManagementTarget{
// 		Id:                       "GUUStaking",
// 		AssetManagementAccountId: amAcc.Id,
// 		AccountAddress:           "",
// 		AssetConditions:          []types.AssetCondition{},
// 		UnbondingTime:            time.Second,
// 		IntegrateInfo:            types.IntegrateInfo{},
// 	}
// 	suite.app.YieldaggregatorKeeper.SetAssetManagementTarget(suite.ctx, amTarget)

// 	// deposit "uguu" by farmer
// 	coins := sdk.Coins{sdk.NewInt64Coin("uguu", 1000000)}
// 	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
// 	suite.NoError(err)
// 	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, farmer, coins)
// 	suite.NoError(err)
// 	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
// 		FromAddress:   farmer.Bytes(),
// 		Amount:        coins,
// 		ExecuteOrders: false,
// 	})
// 	suite.NoError(err)

// 	// create farming order by farmer
// 	suite.app.YieldaggregatorKeeper.SetFarmingOrder(suite.ctx, types.FarmingOrder{
// 		Id:          "ExecuteGUUStaking",
// 		FromAddress: farmer.String(),
// 		Strategy: types.Strategy{
// 			StrategyType:         "ManualStrategy",
// 			WhitelistedTargetIds: []string{amTarget.Id},
// 		},
// 		MaxUnbondingTime: 0,
// 		OverallRatio:     1,
// 		Min:              "",
// 		Max:              "",
// 		Date:             now,
// 		Active:           true,
// 	})

// 	// execute farming order by farmer
// 	orders := suite.app.YieldaggregatorKeeper.GetFarmingOrdersOfAddress(suite.ctx, farmer)
// 	suite.app.YieldaggregatorKeeper.ExecuteFarmingOrders(suite.ctx, farmer, orders)

// 	// check farm unit created
// 	farmUnits := suite.app.YieldaggregatorKeeper.GetFarmingUnitsOfAddress(suite.ctx, farmer)
// 	suite.Require().GreaterOrEqual(len(farmUnits), 1)

// 	// check deposit result on yieldfarm module
// 	farmUnit := farmUnits[0]
// 	yFarmInfo := suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, farmUnit.GetAddress())
// 	suite.Require().Equal(yFarmInfo, yieldfarmtypes.FarmerInfo{
// 		Account: farmUnit.GetAddress().String(),
// 		Amount:  coins,
// 		Rewards: sdk.Coins(nil),
// 	})

// 	// after a day
// 	future := now.Add(time.Hour * 24)
// 	suite.ctx = suite.ctx.WithBlockTime(future)

// 	// allocate reward
// 	yieldfarm.EndBlocker(suite.ctx, suite.app.YieldfarmKeeper)

// 	// check deposit result on yieldfarm module
// 	yFarmInfo = suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, farmUnit.GetAddress())
// 	suite.Require().Equal(yFarmInfo.Rewards, []sdk.Coin{sdk.NewInt64Coin("uguu", 10000)}) // 1% yield

// 	// claim rewards after a time for all units
// 	suite.app.YieldaggregatorKeeper.ClaimAllFarmUnitRewards(suite.ctx)

// 	// check claim result on yieldfarm module
// 	yFarmInfo = suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, farmUnit.GetAddress())
// 	suite.Require().Equal(yFarmInfo.Rewards, []sdk.Coin(nil)) // amount after claim

// 	// withdraw from farm unit to user deposit
// 	err = suite.app.YieldaggregatorKeeper.ClaimWithdrawFromTarget(suite.ctx, farmer, amTarget)
// 	suite.Require().NoError(err)

// 	// execute farming order once more
// 	suite.app.YieldaggregatorKeeper.ExecuteFarmingOrders(suite.ctx, farmer, orders)

// 	// check update farm unit
// 	farmUnits = suite.app.YieldaggregatorKeeper.GetFarmingUnitsOfAddress(suite.ctx, farmer)
// 	suite.Require().GreaterOrEqual(len(farmUnits), 1)
// 	farmUnit = farmUnits[0]
// 	yFarmInfo = suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, farmUnit.GetAddress())
// 	suite.Require().Equal(yFarmInfo.Amount, []sdk.Coin{sdk.NewInt64Coin("uguu", 1010000)})

// 	// stop farming and close farm unit
// 	err = suite.app.YieldaggregatorKeeper.StopFarmingUnit(suite.ctx, farmUnit)
// 	suite.Require().NoError(err)

// 	// withdraw from farming unit
// 	err = suite.app.YieldaggregatorKeeper.WithdrawFarmingUnit(suite.ctx, farmUnit)
// 	suite.Require().NoError(err)

// 	// check user deposit balance
// 	deposit := suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, farmer)
// 	suite.Require().Equal(deposit, sdk.Coins{sdk.NewInt64Coin("uguu", 1010000)})

// 	// check stop result on yieldfarm module
// 	yFarmInfo = suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, farmUnit.GetAddress())
// 	suite.Require().Equal(yFarmInfo, yieldfarmtypes.FarmerInfo{
// 		Account: farmUnit.GetAddress().String(),
// 		Amount:  sdk.Coins(nil),
// 		Rewards: sdk.Coins(nil),
// 	})

// 	// withdraw whole "uguu" by farmer
// 	err = suite.app.YieldaggregatorKeeper.Withdraw(suite.ctx, &types.MsgWithdraw{
// 		FromAddress: farmer.Bytes(),
// 		Amount:      deposit,
// 	})
// 	suite.Require().NoError(err)

// 	// check user deposit balance
// 	deposit = suite.app.YieldaggregatorKeeper.GetUserDeposit(suite.ctx, farmer)
// 	suite.Require().Equal(deposit, sdk.Coins(nil))

// 	balance := suite.app.BankKeeper.GetBalance(suite.ctx, farmer, "uguu")
// 	suite.Require().Equal(balance, sdk.NewInt64Coin("uguu", 1010000))
// }
