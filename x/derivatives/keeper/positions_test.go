package keeper_test

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/derivatives/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func (suite *KeeperTestSuite) TestGetAllPositions() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	position0Inst, err := codecTypes.NewAnyWithValue(&types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     5,
	})

	if err != nil {
		panic("failed to create any value")
	}

	position1Inst, err := codecTypes.NewAnyWithValue(&types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     5,
	})

	if err != nil {
		panic("failed to create any value")
	}

	positions := []types.Position{
		{
			Id:      "0",
			Address: owner.Bytes(),
			Market: types.Market{
				BaseDenom:  "uatom",
				QuoteDenom: "uusdc",
			},
			OpenedAt:         time.Now().UTC(),
			OpenedHeight:     1,
			OpenedBaseRate:   sdk.NewDec(10),
			OpenedQuoteRate:  sdk.NewDec(10),
			PositionInstance: *position0Inst,
			RemainingMargin:  sdk.NewCoin("uusdc", sdk.NewInt(1000)),
			LastLeviedAt:     time.Now().UTC(),
		},
		{
			Id:      "1",
			Address: owner.Bytes(),
			Market: types.Market{
				BaseDenom:  "uatom",
				QuoteDenom: "uusdc",
			},
			OpenedAt:         time.Now().UTC(),
			OpenedHeight:     2,
			OpenedBaseRate:   sdk.NewDec(10),
			OpenedQuoteRate:  sdk.NewDec(10),
			PositionInstance: *position1Inst,
			RemainingMargin:  sdk.NewCoin("uatom", sdk.NewInt(1000)),
			LastLeviedAt:     time.Now().UTC(),
		},
	}

	for _, position := range positions {
		suite.keeper.SetPosition(suite.ctx, position)
	}

	for index, position := range positions {
		positionId := fmt.Sprintf("%d", index)
		positionInStore := suite.keeper.GetPositionWithId(suite.ctx, positionId)

		positionInstance, _ := types.UnpackPositionInstance(position.PositionInstance)
		positionInstanceInstore, _ := types.UnpackPositionInstance(positionInStore.PositionInstance)

		position.PositionInstance.Reset()
		positionInStore.PositionInstance.Reset()

		suite.Require().Equal(position, *positionInStore)
		suite.Require().Equal(positionInstance, positionInstanceInstore)
	}

	// Check if the position was added
	allPositions := suite.keeper.GetAllPositions(suite.ctx)
	suite.Require().Len(allPositions, len(positions))

	// Check address positions
	addrPositions := suite.keeper.GetAddressPositions(suite.ctx, owner)
	suite.Require().Len(addrPositions, len(positions))
}

func (suite *KeeperTestSuite) TestDeletePosition() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	position0Inst, err := codecTypes.NewAnyWithValue(&types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     5,
	})

	if err != nil {
		panic("failed to create any value")
	}

	position1Inst, err := codecTypes.NewAnyWithValue(&types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     5,
	})

	if err != nil {
		panic("failed to create any value")
	}

	positions := []types.Position{
		{
			Id:      "0",
			Address: owner.Bytes(),
			Market: types.Market{
				BaseDenom:  "uatom",
				QuoteDenom: "uusdc",
			},
			OpenedAt:         time.Now().UTC(),
			OpenedHeight:     1,
			PositionInstance: *position0Inst,
		},
		{
			Id:      "1",
			Address: owner2.Bytes(),
			Market: types.Market{
				BaseDenom:  "uatom",
				QuoteDenom: "uusdc",
			},
			OpenedAt:         time.Now().UTC(),
			OpenedHeight:     2,
			PositionInstance: *position1Inst,
		},
	}

	for _, position := range positions {
		suite.keeper.SetPosition(suite.ctx, position)
	}

	// Check if the position was added
	allPositions := suite.keeper.GetAllPositions(suite.ctx)
	suite.Require().Len(allPositions, len(positions))

	// check per id
	for _, position := range positions {
		p := suite.keeper.GetAddressPositionWithId(suite.ctx, position.Address.AccAddress(), position.Id)
		suite.Require().NotNil(p)
		suite.Require().Equal(p.Id, position.Id)
		suite.Require().Equal(p.Market, position.Market)
	}

	// Delete the position
	suite.keeper.DeletePosition(suite.ctx, owner, "0")

	// Check if the position was deleted
	allPositions = suite.keeper.GetAllPositions(suite.ctx)
	suite.Require().Len(allPositions, 1)

	// Check last position
	lastPosition := suite.keeper.GetLastPosition(suite.ctx)
	suite.Require().Equal(lastPosition, allPositions[0])

	// Delete the position
	suite.keeper.DeletePosition(suite.ctx, owner2, "1")

	// Check if the position was deleted
	allPositions = suite.keeper.GetAllPositions(suite.ctx)
	suite.Require().Len(allPositions, 0)

	// Check last position
	lastPosition = suite.keeper.GetLastPosition(suite.ctx)
	suite.Require().Equal(lastPosition, types.Position{})
}

func (suite *KeeperTestSuite) TestIncreaseLastPositionId() {
	suite.keeper.IncreaseLastPositionId(suite.ctx)

	suite.Require().Equal(suite.keeper.GetLastPositionId(suite.ctx), uint64(1))

	suite.keeper.IncreaseLastPositionId(suite.ctx)

	suite.Require().Equal(suite.keeper.GetLastPositionId(suite.ctx), uint64(2))
}

// FIXME: fix this test
func (suite *KeeperTestSuite) TestOpenCloseLiquidatePosition() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// set price for asset
	_, err := suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:uusdc", sdk.NewDec(13), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:uusdc")
	suite.Require().NoError(err)

	// initial atom balance
	coins := sdk.Coins{sdk.NewInt64Coin("uatom", 1000000), sdk.NewInt64Coin("uusdc", 1000000)}
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, owner, coins)
	suite.Require().NoError(err)

	// open a new long future position when no previous position exists
	positionInstVal := types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     5,
	}

	positionAny, err := codectypes.NewAnyWithValue(&positionInstVal)
	suite.Require().Nil(err)
	err = suite.keeper.OpenPosition(suite.ctx, &types.MsgOpenPosition{
		Sender: owner.Bytes(),
		Margin: sdk.NewInt64Coin("uatom", 1000), // long -> uatom, short -> uusdc
		Market: types.Market{
			BaseDenom:  "uatom",
			QuoteDenom: "uusdc",
		},
		PositionInstance: *positionAny,
	})
	suite.Require().NoError(err)
	allPositions := suite.keeper.GetAllPositions(suite.ctx)
	suite.Require().Len(allPositions, 1)

	// open another position with same size
	err = suite.keeper.OpenPosition(suite.ctx, &types.MsgOpenPosition{
		Sender: owner.Bytes(),
		Margin: sdk.NewInt64Coin("uatom", 1000), // long -> uatom, short -> uusdc
		Market: types.Market{
			BaseDenom:  "uatom",
			QuoteDenom: "uusdc",
		},
		PositionInstance: *positionAny,
	})
	suite.Require().NoError(err)
	allPositions = suite.keeper.GetAllPositions(suite.ctx)
	suite.Require().Len(allPositions, 2)

	// open short future position
	positionInstVal = types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_SHORT,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     5,
	}

	positionAny, err = codectypes.NewAnyWithValue(&positionInstVal)
	suite.Require().Nil(err)
	err = suite.keeper.OpenPosition(suite.ctx, &types.MsgOpenPosition{
		Sender: owner.Bytes(),
		Margin: sdk.NewInt64Coin("uusdc", 100), // long -> uatom, short -> uusdc
		Market: types.Market{
			BaseDenom:  "uatom",
			QuoteDenom: "uusdc",
		},
		PositionInstance: *positionAny,
	})
	suite.Require().NoError(err)
	allPositions = suite.keeper.GetAllPositions(suite.ctx)
	suite.Require().Len(allPositions, 3)

	// TODO: open long options position after implementation
	// positionIns := types.PerpetualOptionsPositionInstance{
	// 		OptionType   OptionType
	// 		PositionType PositionType
	// 		StrikePrice  sdk.Dec
	// 		Premium      sdk.Dec
	// }

	balances := suite.app.BankKeeper.GetAllBalances(suite.ctx, owner)
	suite.Require().Equal(balances.String(), "998000uatom,999900uusdc")

	// ReportLiquidation
	cacheCtx, _ := suite.ctx.CacheContext()
	err = suite.keeper.ReportLiquidation(cacheCtx, &types.MsgReportLiquidation{
		Sender:          owner.Bytes(),
		PositionId:      allPositions[0].Id,
		RewardRecipient: owner.Bytes(),
	})
	suite.Require().Error(err)

	// ReportLevyPeriod
	cacheCtx, _ = suite.ctx.CacheContext()
	err = suite.keeper.ReportLevyPeriod(cacheCtx, &types.MsgReportLevyPeriod{
		Sender:          owner.Bytes(),
		PositionId:      allPositions[0].Id,
		RewardRecipient: owner.Bytes(),
	})
	suite.Require().Error(err)

	err = suite.keeper.ClosePosition(suite.ctx, &types.MsgClosePosition{
		Sender:     owner.Bytes(),
		PositionId: allPositions[0].Id,
	})
	suite.Require().NoError(err)
	balances = suite.app.BankKeeper.GetAllBalances(suite.ctx, owner)
	suite.Require().Equal(balances.String(), "999000uatom,999900uusdc")

	err = suite.keeper.ClosePosition(suite.ctx, &types.MsgClosePosition{
		Sender:     owner.Bytes(),
		PositionId: allPositions[1].Id,
	})
	suite.Require().NoError(err)
	balances = suite.app.BankKeeper.GetAllBalances(suite.ctx, owner)
	suite.Require().Equal(balances.String(), "1000000uatom,999900uusdc")

	err = suite.keeper.ClosePosition(suite.ctx, &types.MsgClosePosition{
		Sender:     owner.Bytes(),
		PositionId: allPositions[2].Id,
	})
	suite.Require().NoError(err)
	balances = suite.app.BankKeeper.GetAllBalances(suite.ctx, owner)
	suite.Require().Equal(balances.String(), "1000000uatom,1000000uusdc")
}
