package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/UnUniFi/chain/x/yieldaggregatorv1/types"
	yieldfarmtypes "github.com/UnUniFi/chain/x/yieldfarm/types"
)

// TODO:
// StopFarmingUnit
// WithdrawFarmingUnit

func (suite *KeeperTestSuite) TestFarmingOrderGetSet() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	addr2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// get not available order
	farmingOrder := suite.app.YieldaggregatorKeeper.GetFarmingOrder(suite.ctx, addr1, "OsmoGUUFarm")
	suite.Require().Equal(farmingOrder, types.FarmingOrder{})

	// set farming orders
	order1 := types.FarmingOrder{
		Id:          "OsmoGUUFarmOrder",
		FromAddress: addr1.String(),
		Strategy: types.Strategy{
			StrategyType:         "Manual",
			WhitelistedTargetIds: []string{"OsmoGUUFarm"},
		},
	}
	order2 := types.FarmingOrder{
		Id:          "EvmosGUUFarmOrder",
		FromAddress: addr2.String(),
		Strategy: types.Strategy{
			StrategyType:         "Manual",
			WhitelistedTargetIds: []string{"EvmosGUUFarm"},
		},
	}
	suite.app.YieldaggregatorKeeper.SetFarmingOrder(suite.ctx, order1)
	suite.app.YieldaggregatorKeeper.SetFarmingOrder(suite.ctx, order2)

	// check farming orders
	farmingOrder = suite.app.YieldaggregatorKeeper.GetFarmingOrder(suite.ctx, addr1, "OsmoGUUFarmOrder")
	suite.Require().Equal(farmingOrder, order1)
	farmingOrder = suite.app.YieldaggregatorKeeper.GetFarmingOrder(suite.ctx, addr2, "EvmosGUUFarmOrder")
	suite.Require().Equal(farmingOrder, order2)
	farmingOrders := suite.app.YieldaggregatorKeeper.GetAllFarmingOrders(suite.ctx)
	suite.Require().Len(farmingOrders, 2)
	farmingOrders = suite.app.YieldaggregatorKeeper.GetFarmingOrdersOfAddress(suite.ctx, addr1)
	suite.Require().Len(farmingOrders, 1)

	// delete farming order and check
	suite.app.YieldaggregatorKeeper.DeleteFarmingOrder(suite.ctx, addr2, "EvmosGUUFarmOrder")
	farmingOrder = suite.app.YieldaggregatorKeeper.GetFarmingOrder(suite.ctx, addr2, "EvmosGUUFarmOrder")
	suite.Require().Equal(farmingOrder, types.FarmingOrder{})
	farmingOrders = suite.app.YieldaggregatorKeeper.GetAllFarmingOrders(suite.ctx)
	suite.Require().Len(farmingOrders, 1)

	// add farming order and check
	err := suite.app.YieldaggregatorKeeper.AddFarmingOrder(suite.ctx, order2)
	suite.Require().NoError(err)
	farmingOrders = suite.app.YieldaggregatorKeeper.GetAllFarmingOrders(suite.ctx)
	suite.Require().Len(farmingOrders, 2)

	// try adding once more to check error
	err = suite.app.YieldaggregatorKeeper.AddFarmingOrder(suite.ctx, order2)
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestFarmingOrderActivation() {
	addr1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// activate not available order
	err := suite.app.YieldaggregatorKeeper.ActivateFarmingOrder(suite.ctx, addr1, "OsmoGUUFarmOrder")
	suite.Require().Error(err)

	// inactivate not available order
	err = suite.app.YieldaggregatorKeeper.ActivateFarmingOrder(suite.ctx, addr1, "OsmoGUUFarmOrder")
	suite.Require().Error(err)

	// set farming order
	order1 := types.FarmingOrder{
		Id:          "OsmoGUUFarmOrder",
		FromAddress: addr1.String(),
		Strategy: types.Strategy{
			StrategyType:         "Manual",
			WhitelistedTargetIds: []string{"OsmoGUUFarm"},
		},
	}
	suite.app.YieldaggregatorKeeper.SetFarmingOrder(suite.ctx, order1)

	// inactivate and check
	err = suite.app.YieldaggregatorKeeper.InactivateFarmingOrder(suite.ctx, addr1, order1.Id)
	suite.Require().NoError(err)
	order := suite.app.YieldaggregatorKeeper.GetFarmingOrder(suite.ctx, addr1, order1.Id)
	suite.Require().False(order.Active)

	// activate and check
	err = suite.app.YieldaggregatorKeeper.ActivateFarmingOrder(suite.ctx, addr1, order1.Id)
	suite.Require().NoError(err)
	order = suite.app.YieldaggregatorKeeper.GetFarmingOrder(suite.ctx, addr1, order1.Id)
	suite.Require().True(order.Active)
}

func (suite *KeeperTestSuite) TestExecuteFarmingOrders() {
	farmer := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	now := time.Now().UTC()
	suite.ctx = suite.ctx.WithBlockTime(now)

	// create asset management account
	amAcc := types.AssetManagementAccount{
		Id:   "OsmosisAssetManager",
		Name: "Osmosis Asset Management Account",
	}
	suite.app.YieldaggregatorKeeper.SetAssetManagementAccount(suite.ctx, amAcc)

	// create asset management target
	amTarget := types.AssetManagementTarget{
		Id:                       "GUUStaking",
		AssetManagementAccountId: amAcc.Id,
		AccountAddress:           "",
		AssetConditions:          []types.AssetCondition{},
		UnbondingTime:            time.Second,
		IntegrateInfo:            types.IntegrateInfo{},
	}
	suite.app.YieldaggregatorKeeper.SetAssetManagementTarget(suite.ctx, amTarget)

	// deposit "uguu" by farmer
	coins := sdk.Coins{sdk.NewInt64Coin("uguu", 1000000)}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, farmer, coins)
	suite.NoError(err)
	err = suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
		FromAddress:   farmer.Bytes(),
		Amount:        coins,
		ExecuteOrders: false,
	})
	suite.NoError(err)

	// create farming order by farmer
	suite.app.YieldaggregatorKeeper.SetFarmingOrder(suite.ctx, types.FarmingOrder{
		Id:          "ExecuteGUUStaking",
		FromAddress: farmer.String(),
		Strategy: types.Strategy{
			StrategyType:         "ManualStrategy",
			WhitelistedTargetIds: []string{amTarget.Id},
		},
		MaxUnbondingTime: 0,
		OverallRatio:     1,
		Min:              "",
		Max:              "",
		Date:             now,
		Active:           true,
	})

	// execute farming order by farmer
	orders := suite.app.YieldaggregatorKeeper.GetFarmingOrdersOfAddress(suite.ctx, farmer)
	suite.app.YieldaggregatorKeeper.ExecuteFarmingOrders(suite.ctx, farmer, orders)

	// check farm unit created
	farmUnits := suite.app.YieldaggregatorKeeper.GetFarmingUnitsOfAddress(suite.ctx, farmer)
	suite.Require().GreaterOrEqual(len(farmUnits), 1)

	// check deposit result on yieldfarm module
	farmUnit := farmUnits[0]
	yFarmInfo := suite.app.YieldfarmKeeper.GetFarmerInfo(suite.ctx, farmUnit.GetAddress())
	suite.Require().Equal(yFarmInfo, yieldfarmtypes.FarmerInfo{
		Account: farmUnit.GetAddress().String(),
		Amount:  coins,
		Rewards: sdk.Coins(nil),
	})
}
