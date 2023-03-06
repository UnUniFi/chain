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
			OpenedRate:       sdk.NewDec(10),
			PositionInstance: *position0Inst,
			RemainingMargin:  sdk.NewCoin("uusdc", sdk.NewInt(1000)),
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
			OpenedRate:       sdk.NewDec(10),
			PositionInstance: *position1Inst,
			RemainingMargin:  sdk.NewCoin("uatom", sdk.NewInt(1000)),
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
			OpenedRate:       sdk.NewDec(10),
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
			OpenedRate:       sdk.NewDec(10),
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
		suite.Require().Equal(p.OpenedRate, position.OpenedRate)
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

	suite.Require().Equal(suite.keeper.GetLastPositionId(suite.ctx), string(types.GetPositionIdBytes(1)))

	suite.keeper.IncreaseLastPositionId(suite.ctx)

	suite.Require().Equal(suite.keeper.GetLastPositionId(suite.ctx), string(types.GetPositionIdBytes(2)))
}

// TODO: add test for
// func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpenPosition) error {
// func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClosePosition) error {
// func (k Keeper) ReportLiquidation(ctx sdk.Context, msg *types.MsgReportLiquidation) error {
// func (k Keeper) ReportLevyPeriod(ctx sdk.Context, msg *types.MsgReportLevyPeriod) error {
