package keeper_test

import (
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestOpenPerpetualFuturesPosition() {
	positionId := "0"
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	margin := sdk.NewCoin("uatom", sdk.NewInt(50))

	market := types.Market{
		Denom:      "uatom",
		QuoteDenom: "uusdc",
	}

	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, sdk.NewDec(0))

	positionInst := types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     sdk.NewInt(5),
	}

	suite.keeper.OpenPerpetualFuturesPosition(suite.ctx, positionId, owner.Bytes(), margin, market, positionInst)

	// Check if the position was added
	netPosition := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(netPosition, sdk.NewDecWithPrec(100, 0))
}

func (suite *KeeperTestSuite) TestClosePerpetualFuturesPosition() {

}

func (suite *KeeperTestSuite) TestReportLiquidationNeededPerpetualFuturesPosition() {

}

func (suite *KeeperTestSuite) TestSetPerpetualFuturesNetPositionOfMarket() {
	market := types.Market{
		Denom:      "uatom",
		QuoteDenom: "uusdc",
	}

	netPosition := sdk.NewDec(100)

	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, netPosition)

	// Check if the netPosition was set
	netPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(netPositionOfMarket, netPosition)
}

func (suite *KeeperTestSuite) TestAddPerpetualFuturesNetPositionOfMarket() {
	market := types.Market{
		Denom:      "uatom",
		QuoteDenom: "uusdc",
	}

	netPosition := sdk.NewDec(100)

	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, netPosition)

	// Check if the netPosition was set
	netPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(netPositionOfMarket, netPosition)

	// Add 50 more
	netAddPosition := sdk.NewDec(50)

	suite.keeper.AddPerpetualFuturesNetPositionOfMarket(suite.ctx, market, netAddPosition)

	// Check if the netPosition was set
	netPositionOfMarket = suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(netPositionOfMarket, netPosition.Add(netAddPosition))
}

func (suite *KeeperTestSuite) TestSubPerpetualFuturesNetPositionOfMarket() {
	market := types.Market{
		Denom:      "uatom",
		QuoteDenom: "uusdc",
	}

	netPosition := sdk.NewDec(100)

	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, netPosition)

	// Check if the netPosition was set
	netPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(netPositionOfMarket, netPosition)

	// Sub 50 more
	netSubPosition := sdk.NewDec(50)

	suite.keeper.SubPerpetualFuturesNetPositionOfMarket(suite.ctx, market, netSubPosition)

	// Check if the netPosition was set
	netPositionOfMarket = suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market)

	suite.Require().Equal(netPositionOfMarket, netPosition.Sub(netSubPosition))
}
