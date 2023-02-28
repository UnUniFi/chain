package keeper_test

import (
	// "github.com/UnUniFi/chain/x/derivatives/types"

	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

// TODO: add test for
// func (k Keeper) GetLPNominalYieldRate(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
// func (k Keeper) GetInflationRateOfAssetsInPool(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
// func (k Keeper) GetLPRealYieldRate(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
