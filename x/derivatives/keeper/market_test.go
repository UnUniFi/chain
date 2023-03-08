package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/derivatives/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func (suite *KeeperTestSuite) TestGetPairRate() {
	// get uninitialized value
	rate, err := suite.keeper.GetPairRate(suite.ctx, types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	})
	suite.Require().Error(err)

	// initialize price keeper
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:uusdc", sdk.NewDec(12), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:uusdc")
	suite.Require().NoError(err)

	// get initialized value
	rate, err = suite.keeper.GetPairRate(suite.ctx, types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	})
	suite.Require().NoError(err)
	suite.Require().Equal(rate.String(), "12.000000000000000000")
}

func (suite *KeeperTestSuite) TestGetAssetPrice() {
	// get uninitialized value
	price, err := suite.keeper.GetAssetPrice(suite.ctx, "uatom")
	suite.Require().NoError(err)
	suite.Require().Equal(price.MarketId, "uatom:usd")
	suite.Require().Equal(price.Price.String(), "0.000015280000000000")

	// initialize price keeper
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.NewDec(12), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:usd", BaseAsset: "uatom", QuoteAsset: "usd", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	suite.Require().NoError(err)

	// get initialized value
	price, err = suite.keeper.GetAssetPrice(suite.ctx, "uatom")
	suite.Require().Equal(price.MarketId, "uatom:usd")
	suite.Require().Equal(price.Price.String(), "12.000000000000000000")
}

func (suite *KeeperTestSuite) TestGetPrice() {
	// get uninitialized value
	price, err := suite.keeper.GetPrice(suite.ctx, "uatom", "usd")
	suite.Require().NoError(err)
	suite.Require().Equal(price.MarketId, "uatom:usd")
	suite.Require().Equal(price.Price.String(), "0.000015280000000000")

	// initialize price keeper
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.NewDec(12), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:usd", BaseAsset: "uatom", QuoteAsset: "usd", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	suite.Require().NoError(err)

	// get initialized value
	price, err = suite.keeper.GetPrice(suite.ctx, "uatom", "usd")
	suite.Require().Equal(price.MarketId, "uatom:usd")
	suite.Require().Equal(price.Price.String(), "12.000000000000000000")
}
