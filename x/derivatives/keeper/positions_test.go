package keeper_test

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cometbft/cometbft/crypto/ed25519"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (suite *KeeperTestSuite) TestIncreaseLastPositionId() {
	suite.keeper.IncreaseLastPositionId(suite.ctx)

	suite.Require().Equal(suite.keeper.GetLastPositionId(suite.ctx), uint64(1))

	suite.keeper.IncreaseLastPositionId(suite.ctx)

	suite.Require().Equal(suite.keeper.GetLastPositionId(suite.ctx), uint64(2))
}

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
			Id:            "0",
			OpenerAddress: owner.String(),
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
			LeviedAmount:     sdk.NewCoin("uusdc", sdk.NewInt(100)),
		},
		{
			Id:            "1",
			OpenerAddress: owner.String(),
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
			LeviedAmount:     sdk.NewCoin("uatom", sdk.NewInt(100)),
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
			Id:            "0",
			OpenerAddress: owner.String(),
			Market: types.Market{
				BaseDenom:  "uatom",
				QuoteDenom: "uusdc",
			},
			OpenedAt:         time.Now().UTC(),
			OpenedHeight:     1,
			PositionInstance: *position0Inst,
		},
		{
			Id:            "1",
			OpenerAddress: owner2.String(),
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
		address, err := sdk.AccAddressFromBech32(position.OpenerAddress)
		suite.Require().NoError(err)
		p := suite.keeper.GetAddressPositionWithId(suite.ctx, address, position.Id)
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

// close, liquidate, and levy test implemented in perpetual_futures_test.go
