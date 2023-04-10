package keeper_test

import (
	"fmt"

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
		positionId     string
		margin         sdk.Coin
		instance       types.PerpetualFuturesPositionInstance
		expNetPosition sdk.Int
	}{
		{
			positionId: "0",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     5,
			},
			expNetPosition: sdk.MustNewDecFromStr("2").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "1",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     5,
			},
			expNetPosition: sdk.MustNewDecFromStr("1").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "2",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     20,
			},
			expNetPosition: sdk.MustNewDecFromStr("4").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "3",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     10,
			},
			expNetPosition: sdk.MustNewDecFromStr("2").MulInt64(1000000).TruncateInt(),
		},
	}

	for _, testPosition := range positions {
		position, err := suite.keeper.OpenPerpetualFuturesPosition(suite.ctx, testPosition.positionId, owner.Bytes(), testPosition.margin, market, testPosition.instance)
		suite.Require().NoError(err)
		suite.Require().NotNil(position)

		// Check if the position was added
		netPosition := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, testPosition.instance.PositionType)
		fmt.Println(netPosition)

		suite.Require().Equal(testPosition.expNetPosition, netPosition.PositionSizeInDenomUnit)
	}
}

// TODO: Implement this test
func (suite *KeeperTestSuite) TestClosePerpetualFuturesPosition() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	positions := []struct {
		positionId     string
		margin         sdk.Coin
		instance       types.PerpetualFuturesPositionInstance
		expNetPosition sdk.Int
	}{
		{
			positionId: "0",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     5,
			},
			expNetPosition: sdk.MustNewDecFromStr("2").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "1",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     5,
			},
			expNetPosition: sdk.MustNewDecFromStr("1").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "2",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     20,
			},
			expNetPosition: sdk.MustNewDecFromStr("0").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "3",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     10,
			},
			expNetPosition: sdk.MustNewDecFromStr("0").MulInt64(1000000).TruncateInt(),
		},
	}

	for _, testPosition := range positions {
		position, err := suite.keeper.OpenPerpetualFuturesPosition(suite.ctx, testPosition.positionId, owner.Bytes(), testPosition.margin, market, testPosition.instance)
		suite.Require().NoError(err)
		suite.Require().NotNil(position)

		suite.keeper.SetPosition(suite.ctx, *position)

		_ = suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{testPosition.margin})
	}

	for _, testPosition := range positions {
		position := suite.keeper.GetPositionWithId(suite.ctx, testPosition.positionId)
		err := suite.keeper.ClosePerpetualFuturesPosition(suite.ctx, types.NewPerpetualFuturesPosition(*position, testPosition.instance))
		suite.Require().NoError(err)

		// Check if the position was added
		netPosition := suite.keeper.GetPerpetualFuturesNetPositionOfMarket(suite.ctx, market, testPosition.instance.PositionType)

		suite.Require().Equal(testPosition.expNetPosition, netPosition.PositionSizeInDenomUnit)
	}
}

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
