package keeper_test

import (
	// "github.com/UnUniFi/chain/x/derivatives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestAnnualizeYieldRate() {
	annualizedYieldRate := suite.keeper.AnnualizeYieldRate(suite.ctx, sdk.NewDec(4), 22, 44)

	// Check if the annualizedYieldRate was calculated
	suite.Require().Equal(annualizedYieldRate, sdk.NewDec(100))
}

// TODO: add test for
// func (k Keeper) SaveBlockTimestamp(ctx sdk.Context, height int64, blockTime time.Time) {
// func (k Keeper) GetBlockTimestamp(ctx sdk.Context, height int64) time.Time {
// func (k Keeper) GetLPNominalYieldRate(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
// func (k Keeper) GetInflationRateOfAssetsInPool(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
// func (k Keeper) GetLPRealYieldRate(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec {
// func (k Keeper) AnnualizeYieldRate(ctx sdk.Context, yieldRate sdk.Dec, beforeHeight int64, afterHeight int64) sdk.Dec {
