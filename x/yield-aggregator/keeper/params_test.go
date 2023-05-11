package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)
	params.CommissionRate = sdk.NewDecWithPrec(1, 1)
	suite.app.YieldaggregatorKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)
	suite.Require().Equal(params, newParams)
}
