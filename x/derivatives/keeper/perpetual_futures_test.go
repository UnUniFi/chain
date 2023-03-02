package keeper_test

import (
	"time"

	ununifitypes "github.com/UnUniFi/chain/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"

	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestOpenPerpetualFuturesPosition() {
	positionId := "0"
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	margin := sdk.NewCoin("uatom", sdk.NewInt(50))

	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	netPositionOfMarket := types.NewPerpetualFuturesNetPositionOfMarket(market, sdk.ZeroDec())
	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, netPositionOfMarket)
	_, err := suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:uusdc", sdk.NewDec(13), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:uusdc")
	suite.Require().NoError(err)

	positionInst := types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     5,
	}

	position, err := suite.keeper.OpenPerpetualFuturesPosition(suite.ctx, positionId, owner.Bytes(), margin, market, positionInst)
	suite.Require().NoError(err)
	suite.Require().NotNil(position)

	// Check if the position was added
	netPosition := suite.keeper.GetPositionSizeOfNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(netPosition, sdk.NewDecWithPrec(100, 0))
}

func (suite *KeeperTestSuite) TestClosePerpetualFuturesPosition() {

}

func (suite *KeeperTestSuite) TestReportLiquidationNeededPerpetualFuturesPosition() {

}

func (suite *KeeperTestSuite) TestSetPerpetualFuturesNetPositionOfMarket() {
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	netPosition := sdk.NewDec(100)
	netPositionOfMarket := types.NewPerpetualFuturesNetPositionOfMarket(market, netPosition)
	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, netPositionOfMarket)

	// Check if the netPosition was set
	gotNetPositionOfMarket := suite.keeper.GetPositionSizeOfNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(netPosition, gotNetPositionOfMarket)
}

func (suite *KeeperTestSuite) TestAddPerpetualFuturesNetPositionOfMarket() {
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	netPosition := sdk.NewDec(100)

	netPositionOfMarket := types.NewPerpetualFuturesNetPositionOfMarket(market, netPosition)
	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, netPositionOfMarket)

	// Check if the netPosition was set
	gotNetPositionOfMarket := suite.keeper.GetPositionSizeOfNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(netPosition, gotNetPositionOfMarket)

	// Add 50 more
	netAddPosition := sdk.NewDec(50)

	suite.keeper.AddPerpetualFuturesNetPositionOfMarket(suite.ctx, market, netAddPosition)

	// Check if the netPosition was set
	positionSizeNetPositionOfMarket := suite.keeper.GetPositionSizeOfNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(positionSizeNetPositionOfMarket, netPosition.Add(netAddPosition))
}

func (suite *KeeperTestSuite) TestSubPerpetualFuturesNetPositionOfMarket() {
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	netPosition := sdk.NewDec(100)
	netPositionOfMarket := types.NewPerpetualFuturesNetPositionOfMarket(market, netPosition)
	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, netPositionOfMarket)

	// Check if the netPosition was set
	positionSizeNetPositionOfMarket := suite.keeper.GetPositionSizeOfNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(positionSizeNetPositionOfMarket, netPosition)

	// Sub 50 more
	netSubPosition := sdk.NewDec(50)

	suite.keeper.SubPerpetualFuturesNetPositionOfMarket(suite.ctx, market, netSubPosition)

	// Check if the netPosition was set
	positionSizeNetPositionOfMarket = suite.keeper.GetPositionSizeOfNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(positionSizeNetPositionOfMarket, netPosition.Sub(netSubPosition))
}
