package keeper_test

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (suite *KeeperTestSuite) TestGetAllPositions() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	position0Inst, err := codecTypes.NewAnyWithValue(&types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     sdk.NewInt(5),
	})

	if err != nil {
		panic("failed to create any value")
	}

	position1Inst, err := codecTypes.NewAnyWithValue(&types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     sdk.NewInt(5),
	})

	if err != nil {
		panic("failed to create any value")
	}

	positions := []types.Position{
		{
			Id:      "0",
			Address: owner.Bytes(),
			Market: types.Market{
				Denom:      "uatom",
				QuoteDenom: "uusdc",
			},
			OpenedAt:         time.Now().UTC(),
			OpenedHeight:     1,
			OpenedRate:       sdk.NewDec(10),
			PositionInstance: *position0Inst,
		},
		{
			Id:      "1",
			Address: owner.Bytes(),
			Market: types.Market{
				Denom:      "uatom",
				QuoteDenom: "uusdc",
			},
			OpenedAt:         time.Now().UTC(),
			OpenedHeight:     2,
			OpenedRate:       sdk.NewDec(10),
			PositionInstance: *position1Inst,
		},
	}

	for _, position := range positions {
		suite.keeper.CreatePosition(suite.ctx, position)
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
}

func (suite *KeeperTestSuite) TestDeletePosition() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	position0Inst, err := codecTypes.NewAnyWithValue(&types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     sdk.NewInt(5),
	})

	if err != nil {
		panic("failed to create any value")
	}

	position1Inst, err := codecTypes.NewAnyWithValue(&types.PerpetualFuturesPositionInstance{
		PositionType: types.PositionType_LONG,
		Size_:        sdk.NewDecWithPrec(100, 0),
		Leverage:     sdk.NewInt(5),
	})

	if err != nil {
		panic("failed to create any value")
	}

	positions := []types.Position{
		{
			Id:      "0",
			Address: owner.Bytes(),
			Market: types.Market{
				Denom:      "uatom",
				QuoteDenom: "uusdc",
			},
			OpenedAt:         time.Now(),
			OpenedHeight:     1,
			OpenedRate:       sdk.NewDec(10),
			PositionInstance: *position0Inst,
		},
		{
			Id:      "1",
			Address: owner2.Bytes(),
			Market: types.Market{
				Denom:      "uatom",
				QuoteDenom: "uusdc",
			},
			OpenedAt:         time.Now(),
			OpenedHeight:     2,
			OpenedRate:       sdk.NewDec(10),
			PositionInstance: *position1Inst,
		},
	}

	for _, position := range positions {
		suite.keeper.CreatePosition(suite.ctx, position)
	}

	// Check if the position was added
	allPositions := suite.keeper.GetAllPositions(suite.ctx)

	suite.Require().Len(allPositions, len(positions))

	// Delete the position
	suite.keeper.DeletePosition(suite.ctx, owner, "0")

	// Check if the position was deleted
	allPositions = suite.keeper.GetAllPositions(suite.ctx)

	suite.Require().Len(allPositions, 1)

	// Delete the position
	suite.keeper.DeletePosition(suite.ctx, owner2, "1")

	// Check if the position was deleted
	allPositions = suite.keeper.GetAllPositions(suite.ctx)

	suite.Require().Len(allPositions, 0)
}

func (suite *KeeperTestSuite) TestIncreaseLastPositionId() {
	suite.keeper.IncreaseLastPositionId(suite.ctx)

	suite.Require().Equal(suite.keeper.GetLastPositionId(suite.ctx), string(types.GetPositionIdBytes(1)))

	suite.keeper.IncreaseLastPositionId(suite.ctx)

	suite.Require().Equal(suite.keeper.GetLastPositionId(suite.ctx), string(types.GetPositionIdBytes(2)))
}
