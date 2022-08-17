package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (suite *KeeperTestSuite) TestInvestmentFlow() {
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
	err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, coins)
	suite.NoError(err)
	suite.app.YieldaggregatorKeeper.Deposit(suite.ctx, &types.MsgDeposit{
		FromAddress:   farmer.Bytes(),
		Amount:        coins,
		ExecuteOrders: false,
	})

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
	suite.app.YieldaggregatorKeeper.ExecuteFarmingOrders(suite.ctx, farmer)

	// check farm unit created
	farmUnits := suite.app.YieldaggregatorKeeper.GetFarmingUnitsOfAddress(suite.ctx, farmer)
	suite.Require().GreaterOrEqual(len(farmUnits), 1)

	// after a month
	future := now.Add(time.Hour * 24 * 30)
	suite.ctx = suite.ctx.WithBlockTime(future)

	// claim rewards after a time for all units
	suite.app.YieldaggregatorKeeper.ClaimAllFarmUnitRewards(suite.ctx)

	// stop farming and close farm unit
	err = suite.app.YieldaggregatorKeeper.StopFarmingUnit(suite.ctx, farmUnits[0])
	suite.Require().NoError(err)

	// withdraw whole "uguu" by farmer
	suite.app.YieldaggregatorKeeper.Withdraw(suite.ctx, &types.MsgWithdraw{
		FromAddress: farmer.Bytes(),
		Amount:      coins,
	})
}
