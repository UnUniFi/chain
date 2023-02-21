package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

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
	suite.Require().Condition(initialLPTSupply.Amount.IsPositive)
	suite.Require().Equal(fee, sdk.NewCoin(types.LiquidityProviderTokenDenom, sdk.ZeroInt()))
	suite.Require().Nil(err)
}

// TODO: impl test
// func (suite *KeeperTestSuite) TestGetLPTokenAmount() {}
