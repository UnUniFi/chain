package keeper_test

import (
	// "github.com/UnUniFi/chain/x/derivatives/types"

	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/derivatives/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func (suite *KeeperTestSuite) TestAnnualizeYieldRate() {
	// calculation of APR without timestamp set
	annualizedYieldRate := suite.keeper.AnnualizeYieldRate(suite.ctx, sdk.NewDec(4), 22, 44)
	suite.Require().Equal(annualizedYieldRate, sdk.ZeroDec())

	// calculation of APR with timestamp set
	now := time.Now()
	future := time.Now().Add(time.Second * 43200)
	suite.keeper.SaveBlockTimestamp(suite.ctx, 22, now)
	suite.keeper.SaveBlockTimestamp(suite.ctx, 42, future)
	annualizedYieldRate = suite.keeper.AnnualizeYieldRate(suite.ctx, sdk.NewDec(1), 22, 42) // 1% per half day

	// Check if the annualizedYieldRate was calculated
	suite.Require().Equal(annualizedYieldRate, sdk.NewDec(730))
}

func (suite *KeeperTestSuite) TestBlockTimestampGetSet() {
	unsavedTime := suite.keeper.GetBlockTimestamp(suite.ctx, 1)
	suite.Require().Equal(unsavedTime, time.Time{})

	now := time.Now()
	future := time.Now().Add(time.Second * 43200)
	suite.keeper.SaveBlockTimestamp(suite.ctx, 1, now)
	savedTime := suite.keeper.GetBlockTimestamp(suite.ctx, 1)
	suite.Require().Equal(savedTime.Unix(), now.Unix())

	suite.keeper.SaveBlockTimestamp(suite.ctx, 1, future)
	savedTime = suite.keeper.GetBlockTimestamp(suite.ctx, 1)
	suite.Require().Equal(savedTime.Unix(), future.Unix())
}

func (suite *KeeperTestSuite) TestGetLPNominalYieldRate() {
	// check value when no value is initialized
	uninitializedRate := suite.keeper.GetLPNominalYieldRate(suite.ctx, 1, 11)
	suite.Require().Equal(uninitializedRate, sdk.ZeroDec())

	// setup snapshot for block height 1
	suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 1, sdk.NewDec(1000000))
	suite.keeper.SetPoolMarketCapSnapshot(suite.ctx, 1, types.PoolMarketCap{
		QuoteTicker: "uatom",
		Total:       sdk.NewDec(100000000),
		Breakdown:   []types.PoolMarketCap_Breakdown{},
	})

	// setup snapshot for block height 11
	suite.keeper.SetLPTokenSupplySnapshot(suite.ctx, 11, sdk.NewDec(2000000))
	suite.keeper.SetPoolMarketCapSnapshot(suite.ctx, 11, types.PoolMarketCap{
		QuoteTicker: "uatom",
		Total:       sdk.NewDec(300000000),
		Breakdown:   []types.PoolMarketCap_Breakdown{},
	})

	// setup data for current block height
	suite.ctx = suite.ctx.WithBlockHeight(20)
	initializedRate := suite.keeper.GetLPNominalYieldRate(suite.ctx, 1, 11)
	suite.Require().Equal(initializedRate, sdk.NewDecWithPrec(5, 1))

	// check value when height is for future
	uninitializedRate = suite.keeper.GetLPNominalYieldRate(suite.ctx, 1, 21)
	suite.Require().Equal(uninitializedRate, sdk.ZeroDec())

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

	// add pool asset and balance
	suite.keeper.AddPoolAsset(suite.ctx, types.PoolParams_Asset{
		Denom:        "uatom",
		TargetWeight: sdk.OneDec(),
	})
	suite.keeper.SetAssetBalance(suite.ctx, sdk.NewInt64Coin("uatom", 1000000))

	// set lp token supply
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(types.LiquidityProviderTokenDenom, 1000000)})
	suite.Require().NoError(err)

	// check current height rate
	pastLptPrice := suite.keeper.GetLPTPriceFromSnapshot(suite.ctx, 1)
	suite.Require().Equal(pastLptPrice.String(), "0.000000000000000100")
	currLptPrice := suite.keeper.GetLPTokenPrice(suite.ctx)
	suite.Require().Equal(currLptPrice.String(), "0.000015280000000000")
	currentRate := suite.keeper.GetLPNominalYieldRate(suite.ctx, 1, suite.ctx.BlockHeight())
	suite.Require().Equal(currentRate, sdk.NewDec(152799999999))
}

// TODO: add test for
// func (k Keeper) GetInflationRateOfAssetsInPool(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
// func (k Keeper) GetLPRealYieldRate(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
