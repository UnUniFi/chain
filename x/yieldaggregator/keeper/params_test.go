package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (suite *KeeperTestSuite) TestParamsGetSet() {
	params, err := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)
	suite.Require().NoError(err)
	params.CommissionRate = sdk.NewDecWithPrec(1, 1)
	suite.app.YieldaggregatorKeeper.SetParams(suite.ctx, params)
	newParams, err := suite.app.YieldaggregatorKeeper.GetParams(suite.ctx)
	suite.Require().NoError(err)

	suite.Require().Equal(params, newParams)
}
