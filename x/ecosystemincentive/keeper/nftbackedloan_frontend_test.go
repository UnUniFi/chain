package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) RecordRecipientWithNftIdTest(ctx sdk.Context, nftId nftbackedloantypes.NftId, recipientContainerId string) error {
	// panic if the nftId is already recorded in the store.
	if _, exists := suite.app.EcosystemincentiveKeeper.GetRecipientByNftId(ctx, nftId); exists {
		return types.ErrRecordedNftId
	}

	if err := suite.app.EcosystemincentiveKeeper.SetRecipientByNftId(ctx, nftId, recipientContainerId); err != nil {
		return err
	}

	return nil
}

// Just mock method to use in only test
func (suite *KeeperTestSuite) AccumulateRewardForFrontendTest(ctx sdk.Context, nftId nftbackedloantypes.NftId, fee sdk.Coin) error {
	// get recipient by nftId
	recipient, exists := suite.app.EcosystemincentiveKeeper.GetRecipientByNftId(ctx, nftId)
	if !exists {
		return types.ErrRegisteredIncentiveId
	}

	nftbackedloanFrontendRewardRate := suite.app.EcosystemincentiveKeeper.GetNftbackedloanFrontendRewardRate(ctx)

	// if the reward rate was not found, emit panic
	if nftbackedloanFrontendRewardRate == sdk.ZeroDec() {
		return types.ErrRewardRateNotFound
	}

	reward, exists := suite.app.EcosystemincentiveKeeper.GetRewardRecord(ctx, sdk.AccAddress(recipient))
	if !exists {
		reward = types.NewRewardRecord(recipient, nil)
	}
	if err := suite.app.EcosystemincentiveKeeper.SetRewardRecord(ctx, reward); err != nil {
		return err
	}

	return nil
}
