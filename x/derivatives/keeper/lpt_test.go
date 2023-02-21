package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

// TODO: impl more various situations for the test cases
func (suite *KeeperTestSuite) TestInitialLiquidityProviderTokenSupply() {
	assetPrice, err := suite.app.DerivativesKeeper.GetAssetPrice(suite.ctx, TestBaseTokenDenom)
	if err != nil {
		suite.Fail("failed to get asset price")
	}
	assetMarketCap := assetPrice.Price.Mul(sdk.NewDecFromInt(sdk.OneInt()))

	initialLPTSupply, fee, err := suite.app.DerivativesKeeper.InitialLiquidityProviderTokenSupply(suite.ctx, assetPrice, assetMarketCap, TestBaseTokenDenom)
	if err != nil {
		suite.Fail("failed to get initial LPT supply")
	}
	suite.Condition(initialLPTSupply.Amount.IsPositive)
	suite.Condition(fee.Amount.IsZero)
}

// TODO: impl test
// func (suite *KeeperTestSuite) TestGetLPTokenAmount() {}
