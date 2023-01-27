package keeper_test

import (
	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestOpenPerpetualFuturesPosition() {

}

func (suite *KeeperTestSuite) TestClosePerpetualFuturesPosition() {

}

func (suite *KeeperTestSuite) TestReportLiquidationNeededPerpetualFuturesPosition() {

}

func (suite *KeeperTestSuite) TestSetPerpetualFuturesNetPositionOfMarket() {
	// owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	netPosition := sdk.Coin{
		Denom:  "uusd",
		Amount: sdk.NewInt(100),
	}

	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId", netPosition)

	// Check if the netPosition was set
	netPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId")

	suite.Require().Equal(netPositionOfMarket, &netPosition)
}

func (suite *KeeperTestSuite) TestAddPerpetualFuturesNetPositionOfMarket() {
	netPosition := sdk.Coin{
		Denom:  "uusd",
		Amount: sdk.NewInt(100),
	}

	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId", netPosition)

	// Check if the netPosition was set
	netPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId")

	suite.Require().Equal(netPositionOfMarket, &netPosition)

	// Add 50 more
	netPosition.Amount = sdk.NewInt(150)

	suite.keeper.AddPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId", netPosition)

	// Check if the netPosition was set
	netPositionOfMarket = suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId")

	suite.Require().Equal(netPositionOfMarket, &netPosition)
}

func (suite *KeeperTestSuite) TestSubPerpetualFuturesNetPositionOfMarket() {
	netPosition := sdk.Coin{
		Denom:  "uusd",
		Amount: sdk.NewInt(100),
	}

	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId", netPosition)

	// Check if the netPosition was set
	netPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId")

	suite.Require().Equal(netPositionOfMarket, &netPosition)

	// Sub 50 more
	netPosition.Amount = sdk.NewInt(50)

	suite.keeper.SubPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId", netPosition)

	// Check if the netPosition was set
	netPositionOfMarket = suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, "marketId")

	suite.Require().Equal(netPositionOfMarket, &netPosition)
}
