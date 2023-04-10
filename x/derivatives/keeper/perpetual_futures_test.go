package keeper_test

import (
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestOpenPerpetualFuturesPosition() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	positions := []struct {
		positionId string
		margin     sdk.Coin
		instance   types.PerpetualFuturesPositionInstance
	}{
		{
			positionId: "0",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     5,
			},
		},
		{
			positionId: "1",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     5,
			},
		},
		{
			positionId: "2",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     20,
			},
		},
		{
			positionId: "3",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     10,
			},
		},
	}

	expectNetPosition := sdk.ZeroInt()
	for _, testPosition := range positions {
		position, err := suite.keeper.OpenPerpetualFuturesPosition(suite.ctx, testPosition.positionId, owner.Bytes(), testPosition.margin, market, testPosition.instance)
		suite.Require().NoError(err)
		suite.Require().NotNil(position)

		// Check if the position was added
		netPosition := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, testPosition.instance.PositionType)

		// any, _ := codecTypes.NewAnyWithValue(&testPosition.instance)
		pfPosition, _ := types.NewPerpetualFuturesPositionFromPosition(*position)
		if testPosition.instance.PositionType == types.PositionType_LONG {
			expectNetPosition = expectNetPosition.Add(pfPosition.PositionInstance.SizeInDenomUnit(types.OneMillionInt))
		} else {
			expectNetPosition = expectNetPosition.Sub(pfPosition.PositionInstance.SizeInDenomUnit(types.OneMillionInt))
		}
		suite.Require().Equal(expectNetPosition, netPosition.PositionSizeInDenomUnit)
	}
}

// TODO: Implement this test
func (suite *KeeperTestSuite) TestClosePerpetualFuturesPosition() {}

// TODO: Implement this test
func (suite *KeeperTestSuite) TestReportLiquidationNeededPerpetualFuturesPosition() {}

func (suite *KeeperTestSuite) TestSetPerpetualFuturesNetPositionOfMarket() {
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	netPosition := sdk.NewInt(100)
	netPositionOfMarket := types.NewPerpetualFuturesNetPositionOfMarket(market, types.PositionType_LONG, netPosition)
	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, netPositionOfMarket)

	// Check if the netPosition was set
	gotNetPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(netPosition, gotNetPositionOfMarket.PositionSizeInDenomUnit)
}

func (suite *KeeperTestSuite) TestAddPerpetualFuturesNetPositionOfMarket() {
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	netPosition := sdk.NewInt(100)

	netPositionOfMarket := types.NewPerpetualFuturesNetPositionOfMarket(market, types.PositionType_LONG, netPosition)
	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, netPositionOfMarket)

	// Check if the netPosition was set
	gotNetPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(netPosition, gotNetPositionOfMarket.PositionSizeInDenomUnit)

	// Add 50 more
	netAddPosition := sdk.NewInt(50)

	suite.keeper.AddPerpetualFuturesNetPositionOfMarket(suite.ctx, market, types.PositionType_LONG, netAddPosition)

	// Check if the netPosition was set
	positionSizeNetPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(positionSizeNetPositionOfMarket.PositionSizeInDenomUnit, netPosition.Add(netAddPosition))
}

func (suite *KeeperTestSuite) TestSubPerpetualFuturesNetPositionOfMarket() {
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	netPosition := sdk.NewInt(100)
	netPositionOfMarket := types.NewPerpetualFuturesNetPositionOfMarket(market, types.PositionType_LONG, netPosition)
	suite.keeper.SetPerpetualFuturesNetPositionOfMarket(suite.ctx, netPositionOfMarket)

	// Check if the netPosition was set
	positionSizeNetPositionOfMarket := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(positionSizeNetPositionOfMarket.PositionSizeInDenomUnit, netPosition)

	// Sub 50 more
	netSubPosition := sdk.NewInt(50)

	suite.keeper.SubPerpetualFuturesNetPositionOfMarket(suite.ctx, market, types.PositionType_LONG, netSubPosition)

	// Check if the netPosition was set
	positionSizeNetPositionOfMarket = suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(positionSizeNetPositionOfMarket.PositionSizeInDenomUnit, netPosition.Sub(netSubPosition))
}
