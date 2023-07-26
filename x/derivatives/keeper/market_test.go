package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func (suite *KeeperTestSuite) TestGetAssetPrice() {
	// get uninitialized value
	price, err := suite.keeper.GetAssetPrice(suite.ctx, "uatom")
	suite.Require().NoError(err)
	suite.Require().Equal(price.MarketId, "uatom:usd")
	suite.Require().Equal(price.Price, sdk.MustNewDecFromStr("0.00001"))

	// initialize price keeper
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.MustNewDecFromStr("0.000015"), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:usd", BaseAsset: "uatom", QuoteAsset: "usd", Oracles: []string{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	suite.Require().NoError(err)

	// get initialized value
	price, _ = suite.keeper.GetAssetPrice(suite.ctx, "uatom")
	suite.Require().Equal(price.MarketId, "uatom:usd")
	suite.Require().Equal(price.Price, sdk.MustNewDecFromStr("0.000015"))
}

func (suite *KeeperTestSuite) TestGetPrice() {
	// get uninitialized value
	price, err := suite.keeper.GetPrice(suite.ctx, "uusdc", "usd")
	suite.Require().NoError(err)
	suite.Require().Equal(price.MarketId, "uusdc:usd")
	suite.Require().Equal(price.Price, sdk.MustNewDecFromStr("0.000001"))

	// initialize price keeper
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uusdc:usd", sdk.MustNewDecFromStr("0.0000009"), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uusdc:usd", BaseAsset: "uusdc", QuoteAsset: "usd", Oracles: []string{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uusdc:usd")
	suite.Require().NoError(err)

	// get initialized value
	price, _ = suite.keeper.GetPrice(suite.ctx, "uusdc", "usd")
	suite.Require().Equal(price.MarketId, "uusdc:usd")
	suite.Require().Equal(price.Price, sdk.MustNewDecFromStr("0.0000009"))
}
