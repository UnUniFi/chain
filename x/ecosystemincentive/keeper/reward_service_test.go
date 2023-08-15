package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) TestRewardDistributionOfnftbackedloan() {
	testCases := []struct {
		testCase   string
		nftId      nftbackedloantypes.NftId
		reward     sdk.Coin
		validDenom bool
		success    bool
		// use default reward rates for the calculation of each reward
		expRewardFornftbackedloanFrontend sdk.Coin
		expRewardForStakers               sdk.Coin
		expRewardForCommunityPool         sdk.Coin
	}{
		{
			testCase:   "success case",
			nftId:      nftbackedloantypes.NftId{ClassId: "test1", TokenId: "test1"},
			reward:     sdk.NewCoin("uguu", sdk.NewInt(100)),
			validDenom: true,
			success:    true,
		},
		{
			testCase:   "too small amount of reward to not distribute reward",
			nftId:      nftbackedloantypes.NftId{ClassId: "test2", TokenId: "test2"},
			reward:     sdk.NewCoin("uguu", sdk.NewInt(1)),
			validDenom: true,
			success:    true,
		},
	}

	for _, tc := range testCases {
		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.reward})
		_ = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{tc.reward})

		if tc.success {
			err := suite.app.EcosystemincentiveKeeper.RewardDistributionOfNftbackedloan(suite.ctx, tc.nftId, tc.reward)
			suite.Require().NoError(err)

			// TODO: check the reward distribution by seeing the balance of the approriate accounts

			// reward := suite.app.BankKeeper.GetBalance(suite.ctx, suite.app.EcosystemincentiveKeeper.GetnftbackedloanAddress(suite.ctx), tc.reward.Denom)
			// suite.Require().Equal(tc.reward, reward)
		} else {
			err := suite.app.EcosystemincentiveKeeper.RewardDistributionOfNftbackedloan(suite.ctx, tc.nftId, tc.reward)
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestWithdrawReward() {
	testCases := []struct {
		testCase    string
		withdrawer  sdk.AccAddress
		reward      sdk.Coin
		validDenom  bool
		rewardExist bool
		success     bool
	}{
		{
			testCase:    "ordinal success case",
			withdrawer:  suite.addrs[0],
			reward:      sdk.NewCoin("uguu", sdk.NewInt(10)),
			validDenom:  true,
			rewardExist: true,
			success:     true,
		},
		{
			testCase:    "no reward accumulated",
			withdrawer:  suite.addrs[0],
			reward:      sdk.Coin{},
			validDenom:  true,
			rewardExist: false,
			success:     false,
		},
		{
			testCase:    "invalid defined token denom",
			withdrawer:  suite.addrs[0],
			reward:      sdk.NewCoin("uguu", sdk.NewInt(10)),
			validDenom:  false,
			rewardExist: true,
			success:     false,
		},
	}

	for _, tc := range testCases {
		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.reward})
		_ = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{tc.reward})

		if tc.success {
			err := suite.app.EcosystemincentiveKeeper.SetRewardRecord(suite.ctx, types.RewardRecord{
				Address: tc.withdrawer.String(),
				Rewards: sdk.NewCoins(tc.reward),
			})
			suite.Require().NoError(err)

			withdrewReward, err := suite.app.EcosystemincentiveKeeper.WithdrawReward(suite.ctx, &types.MsgWithdrawReward{
				Sender: tc.withdrawer.String(),
				Denom:  tc.reward.Denom,
			})
			suite.Require().NoError(err)
			suite.Require().Equal(withdrewReward, tc.reward)

			_, exists := suite.app.EcosystemincentiveKeeper.GetRewardRecord(suite.ctx, tc.withdrawer)
			suite.Require().False(exists)
		} else if !tc.rewardExist {
			_, err := suite.app.EcosystemincentiveKeeper.WithdrawReward(suite.ctx, &types.MsgWithdrawReward{
				Sender: tc.withdrawer.String(),
				Denom:  tc.reward.Denom,
			})
			suite.Require().Error(err)
			suite.Require().EqualError(err, sdkerrors.Wrap(types.ErrRewardNotExists, tc.withdrawer.String()).Error())

			_, exists := suite.app.EcosystemincentiveKeeper.GetRewardRecord(suite.ctx, tc.withdrawer)
			suite.Require().False(exists)
		} else if !tc.validDenom {
			err := suite.app.EcosystemincentiveKeeper.SetRewardRecord(suite.ctx, types.RewardRecord{
				Address: tc.withdrawer.String(),
				Rewards: sdk.NewCoins(tc.reward),
			})
			suite.Require().NoError(err)

			_, err = suite.app.EcosystemincentiveKeeper.WithdrawReward(suite.ctx, &types.MsgWithdrawReward{
				Sender: tc.withdrawer.String(),
				Denom:  "invalid",
			})
			suite.Require().Error(err)
			suite.Require().EqualError(err, sdkerrors.Wrap(types.ErrDenomRewardNotExists, "invalid").Error())

			RewardRecord, exists := suite.app.EcosystemincentiveKeeper.GetRewardRecord(suite.ctx, tc.withdrawer)
			suite.Require().True(exists)
			rightRewardRecord := types.RewardRecord{
				Address: tc.withdrawer.String(),
				Rewards: sdk.NewCoins(tc.reward),
			}
			suite.Require().Equal(RewardRecord, rightRewardRecord)
		}
	}
}

func (suite *KeeperTestSuite) TestWithdrawAllRewards() {
	testCases := []struct {
		testCase    string
		withdrawer  sdk.AccAddress
		rewards     sdk.Coins
		validDenom  bool
		rewardExist bool
		success     bool
	}{
		{
			testCase:    "ordinal success case",
			withdrawer:  suite.addrs[0],
			rewards:     sdk.NewCoins(sdk.NewCoin("uguu", sdk.NewInt(10))),
			validDenom:  true,
			rewardExist: true,
			success:     true,
		},
		{
			testCase:    "no reward accumulated",
			withdrawer:  suite.addrs[0],
			rewards:     sdk.Coins{},
			validDenom:  true,
			rewardExist: false,
			success:     false,
		},
	}

	for _, tc := range testCases {
		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, tc.rewards)
		_ = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, tc.rewards)

		if tc.success {
			err := suite.app.EcosystemincentiveKeeper.SetRewardRecord(suite.ctx, types.RewardRecord{
				Address: tc.withdrawer.String(),
				Rewards: tc.rewards,
			})
			suite.Require().NoError(err)

			withdrewRewards, err := suite.app.EcosystemincentiveKeeper.WithdrawAllRewards(suite.ctx, &types.MsgWithdrawAllRewards{
				Sender: tc.withdrawer.String(),
			})
			suite.Require().NoError(err)
			suite.Require().Equal(withdrewRewards, tc.rewards)

			_, exists := suite.app.EcosystemincentiveKeeper.GetRewardRecord(suite.ctx, tc.withdrawer)
			suite.Require().False(exists)
		} else {
			_, err := suite.app.EcosystemincentiveKeeper.WithdrawAllRewards(suite.ctx, &types.MsgWithdrawAllRewards{
				Sender: tc.withdrawer.String(),
			})
			suite.Require().Error(err)
			suite.Require().EqualError(err, sdkerrors.Wrap(types.ErrRewardNotExists, tc.withdrawer.String()).Error())

			_, exists := suite.app.EcosystemincentiveKeeper.GetRewardRecord(suite.ctx, tc.withdrawer)
			suite.Require().False(exists)
		}
	}
}
