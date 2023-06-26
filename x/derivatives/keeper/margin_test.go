package keeper_test

import (
	"time"

	"github.com/cometbft/cometbft/crypto/ed25519"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (suite *KeeperTestSuite) TestAddMargin() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testCases := []struct {
		name         string
		positionId   string
		margin       sdk.Coin
		instance     types.PerpetualFuturesPositionInstance
		basedRate    sdk.Dec
		addingMargin sdk.Coin
		expPass      bool
		expMargin    sdk.Coin
	}{
		{
			name:       "success in quote denom",
			positionId: "1",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(100000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.NewDec(1),
				Leverage:     1,
			},
			basedRate:    sdk.MustNewDecFromStr("0.000009"),
			addingMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			expPass:      true,
			expMargin:    sdk.NewCoin("uusdc", sdk.NewInt(100000000).Add(sdk.NewInt(10000000))),
		},
		{
			name:       "success in base denom",
			positionId: "2",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(10000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.NewDec(1),
				Leverage:     1,
			},
			basedRate:    sdk.MustNewDecFromStr("0.000009"),
			addingMargin: sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			expPass:      true,
			expMargin:    sdk.NewCoin("uatom", sdk.NewInt(10000000).Add(sdk.NewInt(1000000))),
		},
		{
			name:       "success in short position",
			positionId: "3",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(100000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.NewDec(1),
				Leverage:     1,
			},
			basedRate:    sdk.MustNewDecFromStr("0.000011"),
			addingMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			expPass:      true,
			expMargin:    sdk.NewCoin("uusdc", sdk.NewInt(100000000).Add(sdk.NewInt(10000000))),
		},
		{
			name:       "fail in liquidation condition",
			positionId: "4",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.NewDec(1),
				Leverage:     10,
			},
			basedRate:    sdk.MustNewDecFromStr("0.000009"),
			addingMargin: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			expPass:      false,
		},
	}

	market := types.Market{BaseDenom: "uatom", QuoteDenom: "uusdc"}
	for _, tc := range testCases {
		openedBaseRate, _ := suite.keeper.GetCurrentPrice(suite.ctx, market.BaseDenom)
		openedQuoteRate, _ := suite.keeper.GetCurrentPrice(suite.ctx, market.QuoteDenom)
		any, _ := codecTypes.NewAnyWithValue(&tc.instance)

		position := types.Position{
			Id:               tc.positionId,
			Market:           market,
			Address:          owner.String(),
			OpenedBaseRate:   openedBaseRate,
			OpenedQuoteRate:  openedQuoteRate,
			RemainingMargin:  tc.margin,
			PositionInstance: *any,
		}
		suite.keeper.SetPosition(suite.ctx, position)

		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.NewCoins(tc.addingMargin))
		_ = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, owner, sdk.NewCoins(tc.addingMargin))

		_, _ = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", tc.basedRate, suite.ctx.BlockTime().Add(time.Hour*3))
		_ = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")

		err := suite.keeper.AddMargin(suite.ctx, owner, tc.positionId, tc.addingMargin)
		if tc.expPass {
			suite.Require().NoError(err)
			position := suite.keeper.GetPositionWithId(suite.ctx, tc.positionId)
			suite.Require().Equal(tc.expMargin, position.RemainingMargin)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestWithdrawMargin() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	testCases := []struct {
		name            string
		positionId      string
		margin          sdk.Coin
		instance        types.PerpetualFuturesPositionInstance
		basedRate       sdk.Dec
		withdrawMargin  sdk.Coin
		expPass         bool
		expMargin       sdk.Coin
		expOwnerBalance sdk.Coin
	}{
		{
			name:       "success in quote denom margin",
			positionId: "1",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.NewDec(1),
				Leverage:     1,
			},
			basedRate:       sdk.MustNewDecFromStr("0.00002"),
			withdrawMargin:  sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			expPass:         true,
			expMargin:       sdk.NewCoin("uusdc", sdk.NewInt(10000000).Sub(sdk.NewInt(1000000))),
			expOwnerBalance: sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
		},
		{
			name:       "success in base denom margin",
			positionId: "2",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.NewDec(1),
				Leverage:     1,
			},
			basedRate:       sdk.MustNewDecFromStr("0.00002"),
			withdrawMargin:  sdk.NewCoin("uatom", sdk.NewInt(100000)),
			expPass:         true,
			expMargin:       sdk.NewCoin("uatom", sdk.NewInt(1000000).Sub(sdk.NewInt(100000))),
			expOwnerBalance: sdk.NewCoin("uatom", sdk.NewInt(100000)),
		},
		{
			name:       "success in base denom margin",
			positionId: "3",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_SHORT,
				Size_:        sdk.NewDec(1),
				Leverage:     1,
			},
			basedRate:       sdk.MustNewDecFromStr("0.000001"),
			withdrawMargin:  sdk.NewCoin("uatom", sdk.NewInt(100000)),
			expPass:         true,
			expMargin:       sdk.NewCoin("uatom", sdk.NewInt(1000000).Sub(sdk.NewInt(100000))),
			expOwnerBalance: sdk.NewCoin("uatom", sdk.NewInt(100000)).AddAmount(sdk.NewInt(100000)),
		},
		{
			name:       "fail in withdrawing margin more than remaining",
			positionId: "4",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.NewDec(1),
				Leverage:     1,
			},
			basedRate:      sdk.MustNewDecFromStr("0.00001"),
			withdrawMargin: sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
			expPass:        false,
		},
		{
			name:       "fail in insufficient margin",
			positionId: "5",
			margin:     sdk.NewCoin("uatom", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.NewDec(1),
				Leverage:     1,
			},
			basedRate:      sdk.MustNewDecFromStr("0.00001"),
			withdrawMargin: sdk.NewCoin("uatom", sdk.NewInt(500000)),
			expPass:        false,
		},
		{
			name:       "fail in withdrawing margin in different denom",
			positionId: "6",
			margin:     sdk.NewCoin("uusdc", sdk.NewInt(1000000)),
			instance: types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.NewDec(1),
				Leverage:     1,
			},
			basedRate:      sdk.MustNewDecFromStr("0.00001"),
			withdrawMargin: sdk.NewCoin("uatom", sdk.NewInt(100000)),
			expPass:        false,
		},
	}

	market := types.Market{BaseDenom: "uatom", QuoteDenom: "uusdc"}
	for _, tc := range testCases {
		openedBaseRate, _ := suite.keeper.GetCurrentPrice(suite.ctx, market.BaseDenom)
		openedQuoteRate, _ := suite.keeper.GetCurrentPrice(suite.ctx, market.QuoteDenom)
		any, _ := codecTypes.NewAnyWithValue(&tc.instance)

		position := types.Position{
			Id:               tc.positionId,
			Market:           market,
			Address:          owner.String(),
			OpenedBaseRate:   openedBaseRate,
			OpenedQuoteRate:  openedQuoteRate,
			RemainingMargin:  tc.margin,
			PositionInstance: *any,
		}
		suite.keeper.SetPosition(suite.ctx, position)

		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.NewCoins(tc.withdrawMargin))
		_ = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.MarginManager, sdk.NewCoins(tc.withdrawMargin))

		_, _ = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:usd", tc.basedRate, suite.ctx.BlockTime().Add(time.Hour*3))
		_ = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:usd")

		err := suite.keeper.WithdrawMargin(suite.ctx, owner, tc.positionId, tc.withdrawMargin)
		if tc.expPass {
			suite.Require().NoError(err)
			position := suite.keeper.GetPositionWithId(suite.ctx, tc.positionId)
			suite.Require().Equal(tc.expMargin, position.RemainingMargin)
			ownerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, owner, tc.withdrawMargin.Denom)
			suite.Require().Equal(tc.expOwnerBalance, ownerBalance)
		} else {
			suite.Require().Error(err)
		}
	}
}
