package keeper_test

import (
	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestGetAllPositions() {
	// owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	positions := []types.Position{
		{
			PositionId: "positionId",
			Owner:      "owner",
			MarketId:   "marketId",
			Amount:     sdk.NewInt(100),
			EntryPrice: sdk.NewDec(100),
			EntryTime:  100,
			IsLong:     true,
		},
	}

	for _, position := range positions {
		suite.keeper.SetPosition(suite.ctx, position)
	}

	// Check if the position was added
	allPositions := suite.keeper.GetAllPositions(suite.ctx)

	suite.Require().Len(allPositions, len(positions))
}

func (suite *KeeperTestSuite) TestDeletePosition() {
	// owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	positions := []types.Position{
		{
			PositionId: "positionId",
			Owner:      "owner",
			MarketId:   "marketId",
			Amount:     sdk.NewInt(100),
			EntryPrice: sdk.NewDec(100),
			EntryTime:  100,
			IsLong:     true,
		},
	}

	for _, position := range positions {
		suite.keeper.CreatePosition(suite.ctx, position)
	}

	// Check if the position was added
	allPositions := suite.keeper.GetAllPositions(suite.ctx)

	suite.Require().Len(allPositions, len(positions))

	// Delete the position
	suite.keeper.DeletePosition(suite.ctx, "positionId")

	// Check if the position was deleted
	allPositions = suite.keeper.GetAllPositions(suite.ctx)

	suite.Require().Len(allPositions, 0)
}
