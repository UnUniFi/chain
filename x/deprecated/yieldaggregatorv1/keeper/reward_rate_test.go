package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregatorv1/types"
)

func (suite *KeeperTestSuite) TestDailyRewardPercentGetSet() {
	// get not available rate
	percent := suite.app.YieldaggregatorKeeper.GetDailyRewardPercent(suite.ctx, "OsmoFarm", "OsmoGUUFarm")
	suite.Require().Equal(percent, types.DailyPercent{})

	// set rates
	percent1 := types.DailyPercent{
		AccountId: "OsmoFarm",
		TargetId:  "OsmoGUUFarm",
		Rate:      sdk.NewDecWithPrec(1, 1),
	}
	percent2 := types.DailyPercent{
		AccountId: "EvmosFarm",
		TargetId:  "EvmosGUUFarm",
		Rate:      sdk.NewDecWithPrec(1, 1),
	}
	suite.app.YieldaggregatorKeeper.SetDailyRewardPercent(suite.ctx, percent1)
	suite.app.YieldaggregatorKeeper.SetDailyRewardPercent(suite.ctx, percent2)

	// check rates
	percent = suite.app.YieldaggregatorKeeper.GetDailyRewardPercent(suite.ctx, "OsmoFarm", "OsmoGUUFarm")
	suite.Require().Equal(percent, percent1)
	percent = suite.app.YieldaggregatorKeeper.GetDailyRewardPercent(suite.ctx, "EvmosFarm", "EvmosGUUFarm")
	suite.Require().Equal(percent, percent2)
	percents := suite.app.YieldaggregatorKeeper.GetAllDailyRewardPercents(suite.ctx)
	suite.Require().Len(percents, 2)

	// delete rate and check
	suite.app.YieldaggregatorKeeper.DeleteDailyRewardPercent(suite.ctx, percent2)
	percent = suite.app.YieldaggregatorKeeper.GetDailyRewardPercent(suite.ctx, "EvmosFarm", "EvmosGUUFarm")
	suite.Require().Equal(percent, types.DailyPercent{})
	percents = suite.app.YieldaggregatorKeeper.GetAllDailyRewardPercents(suite.ctx)
	suite.Require().Len(percents, 1)
}
