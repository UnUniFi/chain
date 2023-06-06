package keeper_test

import (
	"time"

	ununifitypes "github.com/UnUniFi/chain/deprecated/types"
	"github.com/UnUniFi/chain/x/derivatives/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestGetAssetBalanceInPoolByDenom() {
	assets := []sdk.Coin{
		{
			Denom:  "uusdc",
			Amount: sdk.NewInt(100),
		},
		{
			Denom:  "uatom",
			Amount: sdk.NewInt(10),
		},
	}

	for _, asset := range assets {
		err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(asset))
		suite.NoError(err)

		amount := suite.keeper.GetAssetBalanceInPoolByDenom(suite.ctx, asset.Denom)

		suite.Require().Equal(asset, amount)
	}
}

func (suite *KeeperTestSuite) TestSetPoolMarketCapSnapshot() {
	height := suite.ctx.BlockHeight()

	assets := []sdk.Coin{
		{
			Denom:  "uusdc",
			Amount: sdk.NewInt(10000000),
		},
		{
			Denom:  "uatom",
			Amount: sdk.NewInt(1000000),
		},
	}

	for _, asset := range assets {
		err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(asset))
		suite.NoError(err)
	}

	marketCap, err := suite.keeper.GetPoolMarketCap(suite.ctx)
	suite.Require().NoError(err)

	err = suite.keeper.SetPoolMarketCapSnapshot(suite.ctx, height, marketCap)
	suite.Require().NoError(err)

	// Check if the market cap was set
	marketCapInStore := suite.keeper.GetPoolMarketCapSnapshot(suite.ctx, height)

	suite.Require().Equal(marketCap, marketCapInStore)

	suite.Require().Equal(sdk.MustNewDecFromStr("20"), marketCapInStore.Total)
}

func (suite *KeeperTestSuite) TestIsAssetValid() {
	poolAssets := suite.keeper.GetPoolAcceptedAssetsConf(suite.ctx)
	suite.Require().Len(poolAssets, 2)

	isValid := suite.keeper.IsAssetAcceptable(suite.ctx, "uatom")
	suite.Require().True(isValid)

	isValid = suite.keeper.IsAssetAcceptable(suite.ctx, "xxxx")
	suite.Require().False(isValid)
}

// TODO: implement additional test cases
func (suite *KeeperTestSuite) TestGetAssetTargetAmount() {
	// get target amount at initial
	targetAmount, err := suite.keeper.GetAssetTargetAmount(suite.ctx, "uatom")
	suite.Require().NoError(err)
	suite.Require().Equal(targetAmount.String(), "0uatom")

	// set lp token supply
	err = suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.Coins{sdk.NewInt64Coin("uusdc", 1000000)})
	suite.Require().NoError(err)

	// get target amount after data set
	targetAmount, err = suite.keeper.GetAssetTargetAmount(suite.ctx, "uusdc")
	suite.Require().NoError(err)
	suite.Require().Equal("500000uusdc", targetAmount.String())
}

func (suite *KeeperTestSuite) TestIsPriceReady() {
	suite.SetupTest()
	// get the value when nothing is set
	isReady := suite.keeper.IsPriceReady(suite.ctx)
	suite.Require().True(isReady)

	// get value after adding one pool asset
	isReady = suite.keeper.IsPriceReady(suite.ctx)
	suite.Require().True(isReady)

	// get value after configuring price
	_, err := suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:uusdc", sdk.NewDec(13), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:uusdc")
	suite.Require().NoError(err)

	isReady = suite.keeper.IsPriceReady(suite.ctx)
	suite.Require().True(isReady)
}
