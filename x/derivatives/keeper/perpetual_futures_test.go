package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"

	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
			positionId: "-1",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("1"),
				Leverage:     1,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(1)),
			expGrossPosition:     sdk.MustNewDecFromStr("0").MulInt64(1000000).TruncateInt(),
		},
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
		if testPosition.positionId == "-1" {
			suite.Require().Error(err)
			suite.Require().Nil(position)
		} else {
			suite.Require().NoError(err)
			suite.Require().NotNil(position)
		}

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

func (suite *KeeperTestSuite) TestReportLiquidationNeededPerpetualFuturesPosition() {
	suite.SetParams()
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	market := types.Market{
		BaseDenom:  "uatom",
		QuoteDenom: "uusdc",
	}
	_, err := suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.MustNewDecFromStr("0.00002"), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uusdc:usd", sdk.MustNewDecFromStr("0.000001"), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	suite.Require().NoError(err)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uusdc:usd")
	suite.Require().NoError(err)

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
			// margin rate 125% = margin 10usd / require 8usd
			// => 69% = (margin 9usd + loss 4usd = 5usd) / require 7.2 usd
			expGrossPosition: sdk.MustNewDecFromStr("4").MulInt64(1000000).TruncateInt(),
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
			// margin rate 125%
			// => 180% = (margin 9usd + profit 4usd = 13usd) / require 7.2 usd
			expGrossPosition: sdk.MustNewDecFromStr("2").MulInt64(1000000).TruncateInt(),
		},
		{
			positionId: "2",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(5000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.MustNewDecFromStr("2"),
				Leverage:     10,
			},
			availableAssetInPool: sdk.NewCoin("uatom", sdk.NewInt(20000000)),
			// margin rate 125% = margin 5usd / require 4usd
			// => 27% = (margin 5usd - loss 4usd = 1usd) / require 3.6 usd
			// Close position#2
			expGrossPosition: sdk.MustNewDecFromStr("2").MulInt64(1000000).TruncateInt(),
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

	// 10% price down
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", sdk.MustNewDecFromStr("0.000018"), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")
	suite.Require().NoError(err)

	for _, testPosition := range positions {
		position := suite.keeper.GetPositionWithId(suite.ctx, testPosition.positionId)
		positionInstance, err := types.UnpackPositionInstance(position.PositionInstance)
		suite.Require().NoError(err)
		switch positionInstance := positionInstance.(type) {
		case *types.PerpetualFuturesPositionInstance:
			perpetualFuturesPosition := types.NewPerpetualFuturesPosition(*position, *positionInstance)
			err = suite.keeper.ReportLiquidationNeededPerpetualFuturesPosition(suite.ctx, owner.Bytes(), perpetualFuturesPosition)
		}
		suite.Require().NoError(err)

		// Check if the position was closed
		grossPosition := suite.keeper.GetPerpetualFuturesGrossPositionOfMarket(suite.ctx, market, testPosition.instance.PositionType)
		suite.Require().Equal(testPosition.expGrossPosition, grossPosition.PositionSizeInDenomExponent)
	}
}

func (suite *KeeperTestSuite) TestReportLevyPeriodPerpetualFuturesPosition() {
	suite.SetParams()
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
			// Levy 500000 * 0.05 * 2 / 6 = 8333
			// 500000 - 8333 (funding) - 500(commission) = 491167
			expMargin: sdk.MustNewDecFromStr("491167").TruncateInt(),
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
			// 500000 + 8333(funding) - 500(commission) = 507833
			expMargin: sdk.MustNewDecFromStr("507833").TruncateInt(),
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
			// 1000000 - 16667(funding) - 1000(commission) = 982334
			expMargin: sdk.MustNewDecFromStr("982333").TruncateInt(),
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
			// 1000000 + 16667(funding) - 1000(commission) = 1015666
			expMargin: sdk.MustNewDecFromStr("1015667").TruncateInt(),
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
		positionInstance, err := types.UnpackPositionInstance(position.PositionInstance)
		suite.Require().NoError(err)
		switch positionInstance := positionInstance.(type) {
		case *types.PerpetualFuturesPositionInstance:
			err = suite.keeper.ReportLevyPeriodPerpetualFuturesPosition(suite.ctx, owner.Bytes(), *position, *positionInstance)
		}
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

func (suite *KeeperTestSuite) SetParams() {
	params := suite.app.DerivativesKeeper.GetParams(suite.ctx)
	params.PoolParams = types.PoolParams{
		QuoteTicker:                 "usd",
		BaseLptMintFee:              sdk.MustNewDecFromStr("0.001"),
		BaseLptRedeemFee:            sdk.MustNewDecFromStr("0.001"),
		BorrowingFeeRatePerHour:     sdk.MustNewDecFromStr("0.000001"),
		ReportLiquidationRewardRate: sdk.MustNewDecFromStr("0.3"),
		ReportLevyPeriodRewardRate:  sdk.MustNewDecFromStr("0.3"),
		AcceptedAssetsConf: []types.PoolAssetConf{
			{
				Denom:        "uatom",
				TargetWeight: sdk.OneDec(),
			},
		},
	}
	params.PerpetualFutures = types.PerpetualFuturesParams{
		CommissionRate:        sdk.MustNewDecFromStr("0.001"),
		MarginMaintenanceRate: sdk.MustNewDecFromStr("0.5"),
		ImaginaryFundingRateProportionalCoefficient: sdk.MustNewDecFromStr("0.05"),
		Markets: []*types.Market{
			{
				BaseDenom:  "uatom",
				QuoteDenom: "uusdc",
			},
		},
		MaxLeverage: 30,
	}
	suite.app.DerivativesKeeper.SetParams(suite.ctx, params)
}
