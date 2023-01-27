package keeper_test

import (
	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestAddPoolAsset() {
	// owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	assets := []types.Pool_Asset{
		{
			Denom:        "uusd",
			TargetWeight: sdk.NewDec(1),
		},
	}

	for _, asset := range assets {
		suite.keeper.AddPoolAsset(suite.ctx, asset)
	}

	// Check if the asset was added
	allAssets := suite.keeper.GetPoolAssets(suite.ctx)

	suite.Require().Len(allAssets, len(assets))
}
