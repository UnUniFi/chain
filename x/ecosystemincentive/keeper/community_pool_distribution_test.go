package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

func (suite *KeeperTestSuite) TestAllocateTokensToCommunityPool() {
	testCases := []struct {
		testCase     string
		rewardAmount sdk.Coin
		expReward    sdk.Coin
		success      bool
	}{
		{
			testCase:     "success case",
			rewardAmount: sdk.NewCoin("uguu", sdk.NewInt(10)),
			expReward:    sdk.NewCoin("uguu", sdk.NewInt(10)),
			success:      true,
		},
	}

	for _, tc := range testCases {
		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.rewardAmount})
		_ = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{tc.rewardAmount})

		if tc.success {
			err := suite.app.EcosystemincentiveKeeper.AllocateTokensToCommunityPool(suite.ctx, tc.rewardAmount)
			suite.Require().NoError(err)

			communityPoolBalance := suite.app.DistrKeeper.GetFeePoolCommunityCoins(suite.ctx).AmountOf(tc.rewardAmount.Denom)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expReward, sdk.NewCoin(tc.rewardAmount.Denom, communityPoolBalance.TruncateInt()))
		}
	}
}
