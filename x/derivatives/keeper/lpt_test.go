package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

// TODO: impl more various situations for the test cases
func (suite *KeeperTestSuite) TestInitialLiquidityProviderTokenSupply() {
	mockPrice := sdk.OneDec()
	mockAssetPrice := &pftypes.CurrentPrice{
		MarketId: "uatom:usd",
		Price:    mockPrice,
	}

	mockDepositingTokenAmount := sdk.OneDec()
	mockAssetMarketCap := mockPrice.Mul(mockDepositingTokenAmount)

	initialLPTSupply, fee, err := suite.app.DerivativesKeeper.InitialLiquidityProviderTokenSupply(suite.ctx, mockAssetPrice, mockAssetMarketCap, TestBaseTokenDenom)
	suite.Require().Equal(sdk.NewInt(2), initialLPTSupply.Amount)
	suite.Require().Equal(fee, sdk.NewCoin(types.LiquidityProviderTokenDenom, sdk.ZeroInt()))
	suite.Require().Nil(err)
}

// TODO: impl test
// func (suite *KeeperTestSuite) TestDetermineMintingLPTokenAmount() {}
