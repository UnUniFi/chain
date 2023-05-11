package keeper_test

import (
	"fmt"

	"github.com/cometbft/cometbft/crypto/ed25519"

	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
)

func (suite *KeeperTestSuite) TestOpenPerpetualFuturesPosition() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	// TODO: add failure case due to the lack of the available asset in the pool
	positions := []struct {
		positionId           string
		margin               sdk.Coin
		instance             types.PerpetualFuturesPositionInstance
		availableAssetInPool sdk.Coin
		expGrossPosition     sdk.Int
	}{
		{
			positionId: "0",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(2000000)),
			expGrossPosition:     sdk.MustNewDecFromStr("2").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "1",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			expGrossPosition:     sdk.MustNewDecFromStr("1").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "2",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     20,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(20000000)),
			expGrossPosition:     sdk.MustNewDecFromStr("4").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "3",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     10,
			},
			availableAssetInPool: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			expGrossPosition:     sdk.MustNewDecFromStr("2").MulInt64(1000000).TruncateInt(),
		},
	}

	for _, testPosition := range positions {
		err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{testPosition.availableAssetInPool})
		suite.Require().NoError(err)

		position, err := suite.keeper.OpenPerpetualFuturesPosition(suite.ctx, testPosition.positionId, owner.Bytes(), testPosition.margin, market, testPosition.instance)
		suite.Require().NoError(err)
		suite.Require().NotNil(position)

		// Check if the position was added
		grossPosition := suite.keeper.GetPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, testPosition.instance.PositionType)

		suite.Require().Equal(testPosition.expGrossPosition, grossPosition.PositionSizeInDenomExponent)
	}
}

func (suite *KeeperTestSuite) TestAddReserveTokensForPosition() {
	testCases := []struct {
		name        string
		reserveCoin sdk.Coin
		expReserve  sdk.Coin
	}{
		{
			name:        "add reserve tokens in uatom",
			reserveCoin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			expReserve:  sdk.NewCoin("uatom", sdk.NewInt(1000000)),
		},
		{
			name:        "add reserve tokens in uatom again",
			reserveCoin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			expReserve:  sdk.NewCoin("uatom", sdk.NewInt(2000000)),
		},
	}

	for _, tc := range testCases {
		err := suite.keeper.AddReserveTokensForPosition(suite.ctx, tc.reserveCoin.Amount, tc.reserveCoin.Denom)
		suite.Require().NoError(err)

		reserve, err := suite.keeper.GetReservedCoin(suite.ctx, tc.reserveCoin.Denom)
		suite.Require().NoError(err)
		suite.Require().Equal(tc.expReserve, reserve)
	}
}

func (suite *KeeperTestSuite) TestSubReserveTokensForPosition() {
	testCases := []struct {
		name        string
		reserveCoin sdk.Coin
		subReserve  sdk.Coin
		expReserve  sdk.Coin
	}{
		{
			name:        "Sub reserve tokens in uatom",
			reserveCoin: sdk.NewCoin("uatom", sdk.NewInt(2000000)),
			subReserve:  sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			expReserve:  sdk.NewCoin("uatom", sdk.NewInt(1000000)),
		},
		{
			name:        "Sub reserve tokens in uatom to zero",
			reserveCoin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			subReserve:  sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			expReserve:  sdk.NewCoin("uatom", sdk.NewInt(0)),
		},
	}

	for _, tc := range testCases {
		err := suite.keeper.SetReservedCoin(suite.ctx, tc.reserveCoin)
		suite.Require().NoError(err)
		err = suite.keeper.SubReserveTokensForPosition(suite.ctx, tc.subReserve.Amount, tc.subReserve.Denom)
		suite.Require().NoError(err)

		reserve, err := suite.keeper.GetReservedCoin(suite.ctx, tc.reserveCoin.Denom)
		suite.Require().NoError(err)
		suite.Require().Equal(tc.expReserve, reserve)
	}
}

func (suite *KeeperTestSuite) TestClosePerpetualFuturesPosition() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	// TODO: Check the returning amount to the owner
	positions := []struct {
		positionId           string
		margin               sdk.Coin
		instance             types.PerpetualFuturesPositionInstance
		availableAssetInPool sdk.Coin
		expGrossPosition     sdk.Int
	}{
		{
			positionId: "0",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(10000000)),
			// 2+2-2 = 2
			expGrossPosition: sdk.MustNewDecFromStr("2").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "1",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			// 2+1-2 = 1
			expGrossPosition: sdk.MustNewDecFromStr("1").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "2",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     20,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(10000000)),
			expGrossPosition:     sdk.MustNewDecFromStr("0").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "3",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     10,
			},
			availableAssetInPool: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			expGrossPosition:     sdk.MustNewDecFromStr("0").MulInt64(1000000).TruncateInt(),
		},
	}

	for _, testPosition := range positions {
		err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{testPosition.availableAssetInPool})
		suite.Require().NoError(err)

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
		grossPosition := suite.keeper.GetPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, testPosition.instance.PositionType)

		suite.Require().Equal(testPosition.expGrossPosition, grossPosition.PositionSizeInDenomExponent)
	}
}

// TODO: Implement this test
func (suite *KeeperTestSuite) TestReportLiquidationNeededPerpetualFuturesPosition() {}

// TODO: Fix param & work on this test
func (suite *KeeperTestSuite) TestReportLevyPeriodPerpetualFuturesPosition() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}
	positions := []struct {
		positionId           string
		margin               sdk.Coin
		instance             types.PerpetualFuturesPositionInstance
		availableAssetInPool sdk.Coin
		expMargin            sdk.Int
	}{
		{
			positionId: "0",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(2000000)),
			expMargin:            sdk.MustNewDecFromStr("500000").TruncateInt(),
		},
		{
			positionId: "1",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(500000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     5,
			},
			availableAssetInPool: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			expMargin:            sdk.MustNewDecFromStr("500000").TruncateInt(),
		},
		{
			positionId: "2",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     20,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(20000000)),
			expMargin:            sdk.MustNewDecFromStr("500000").TruncateInt(),
		},
		{
			positionId: "3",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     10,
			},
			availableAssetInPool: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			expMargin:            sdk.MustNewDecFromStr("500000").TruncateInt(),
		},
	}

	for _, testPosition := range positions {
		err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{testPosition.availableAssetInPool})
		suite.Require().NoError(err)

		position, err := suite.keeper.OpenPerpetualFuturesPosition(suite.ctx, testPosition.positionId, owner.Bytes(), testPosition.margin, market, testPosition.instance)
		suite.Require().NoError(err)
		suite.Require().NotNil(position)
	}

	for _, testPosition := range positions {
		position := suite.keeper.GetPositionWithId(suite.ctx, testPosition.positionId)
		err := suite.keeper.ReportLevyPeriodPerpetualFuturesPosition(suite.ctx, ununifitypes.StringAccAddress(owner), *position, testPosition.instance)
		suite.Require().NoError(err)

		// Check if the position was changed
		updatedPosition := suite.keeper.GetPositionWithId(suite.ctx, testPosition.positionId)

		suite.Require().Equal(testPosition.expMargin, updatedPosition.RemainingMargin.Amount)
	}
}

func (suite *KeeperTestSuite) TestSetPerpetualFuturesGrossPositionOfMarket() {
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	grossPosition := sdk.NewInt(100)
	grossPositionOfMarket := types.NewPerpetualFuturesGrossPositionOfMarket(market, types.PositionType_LONG, grossPosition)
	suite.keeper.SetPerpetualFuturesGrossPositionOfMarket(suite.ctx, grossPositionOfMarket)

	// Check if the grossPosition was set
	gotGrossPositionOfMarket := suite.keeper.GetPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(grossPosition, gotGrossPositionOfMarket.PositionSizeInDenomExponent)
}

func (suite *KeeperTestSuite) TestAddPerpetualFuturesGrossPositionOfMarket() {
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	grossPosition := sdk.NewInt(100)

	grossPositionOfMarket := types.NewPerpetualFuturesGrossPositionOfMarket(market, types.PositionType_LONG, grossPosition)
	suite.keeper.SetPerpetualFuturesGrossPositionOfMarket(suite.ctx, grossPositionOfMarket)

	// Check if the grossPosition was set
	gotGrossPositionOfMarket := suite.keeper.GetPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(grossPosition, gotGrossPositionOfMarket.PositionSizeInDenomExponent)

	// Add 50 more
	netAddPosition := sdk.NewInt(50)

	suite.keeper.AddPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, types.PositionType_LONG, netAddPosition)

	// Check if the grossPosition was set
	positionSizeGrossPositionOfMarket := suite.keeper.GetPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(positionSizeGrossPositionOfMarket.PositionSizeInDenomExponent, grossPosition.Add(netAddPosition))
}

func (suite *KeeperTestSuite) TestSubPerpetualFuturesGrossPositionOfMarket() {
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}

	grossPosition := sdk.NewInt(100)
	grossPositionOfMarket := types.NewPerpetualFuturesGrossPositionOfMarket(market, types.PositionType_LONG, grossPosition)
	suite.keeper.SetPerpetualFuturesGrossPositionOfMarket(suite.ctx, grossPositionOfMarket)

	// Check if the grossPosition was set
	positionSizeGrossPositionOfMarket := suite.keeper.GetPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(positionSizeGrossPositionOfMarket.PositionSizeInDenomExponent, grossPosition)

	// Sub 50 more
	netSubPosition := sdk.NewInt(50)

	suite.keeper.SubPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, types.PositionType_LONG, netSubPosition)

	// Check if the grossPosition was set
	positionSizeGrossPositionOfMarket = suite.keeper.GetPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, types.PositionType_LONG)

	suite.Require().Equal(positionSizeGrossPositionOfMarket.PositionSizeInDenomExponent, grossPosition.Sub(netSubPosition))
}
