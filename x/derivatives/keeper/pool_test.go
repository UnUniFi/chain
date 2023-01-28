package keeper_test

import (
	"github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) AddPoolAssets() []types.Pool_Asset {
	assets := []types.Pool_Asset{
		{
			Denom:        "uusdc",
			TargetWeight: sdk.NewDec(1),
		},
		{
			Denom:        "uatom",
			TargetWeight: sdk.NewDec(10),
		},
		{
			Denom:        "uguu",
			TargetWeight: sdk.NewDec(100),
		},
	}

	for _, asset := range assets {
		suite.keeper.AddPoolAsset(suite.ctx, asset)
	}

	return assets
}

func (suite *KeeperTestSuite) TestAddPoolAsset() {
	// owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	assets := suite.AddPoolAssets()

	for _, asset := range assets {
		assetInStore := suite.keeper.GetPoolAssetByDenom(suite.ctx, asset.Denom)
		suite.Require().Equal(asset, assetInStore)
	}

	// Check if the asset was added
	allAssets := suite.keeper.GetPoolAssets(suite.ctx)

	suite.Require().Len(allAssets, len(assets))
}

func (suite *KeeperTestSuite) TestDepositPoolAsset() {
	suite.AddPoolAssets()

	depositors := []sdk.AccAddress{
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
	}

	assets := []sdk.Coin{
		{
			Denom:  "uusdc",
			Amount: sdk.NewInt(100),
		},
		{
			Denom:  "uatom",
			Amount: sdk.NewInt(10),
		},
		{
			Denom:  "uguu",
			Amount: sdk.NewInt(1000),
		},
	}

	for index, asset := range assets {
		suite.keeper.DepositPoolAsset(suite.ctx, depositors[index], asset)
	}

	for _, asset := range assets {
		amount := suite.keeper.GetAssetBalance(suite.ctx, asset.Denom)
		suite.Require().Equal(asset, amount)
	}
}

func (suite *KeeperTestSuite) TestSetPoolMarketCapSnapshot() {
	suite.AddPoolAssets()

	depositors := []sdk.AccAddress{
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
	}

	height := suite.ctx.BlockHeight()

	assets := []sdk.Coin{
		{
			Denom:  "uusdc",
			Amount: sdk.NewInt(100),
		},
		{
			Denom:  "uatom",
			Amount: sdk.NewInt(10),
		},
		{
			Denom:  "uguu",
			Amount: sdk.NewInt(1000),
		},
	}

	for index, asset := range assets {
		suite.keeper.DepositPoolAsset(suite.ctx, depositors[index], asset)
	}

	marketCap := suite.keeper.GetPoolMarketCap(suite.ctx)

	// TODO: it's not working yet as we didn't add the ticker to price feed
	suite.keeper.SetPoolMarketCapSnapshot(suite.ctx, height, marketCap)

	// Check if the market cap was set
	marketCapInStore := suite.keeper.GetPoolMarketCapSnapshot(suite.ctx, height)

	suite.Require().Equal(marketCap, marketCapInStore)
}
