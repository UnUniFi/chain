package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

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
			err := suite.app.EcosystemincentiveKeeper.SetRewardStore(suite.ctx, types.RewardStore{
				SubjectAddr: tc.withdrawer.Bytes(),
				Rewards:     sdk.NewCoins(tc.reward),
			})
			suite.Require().NoError(err)

			withdrewReward, err := suite.app.EcosystemincentiveKeeper.WithdrawReward(suite.ctx, &types.MsgWithdrawReward{
				Sender: tc.withdrawer.Bytes(),
				Denom:  tc.reward.Denom,
			})
			suite.Require().NoError(err)
			suite.Require().Equal(withdrewReward, tc.reward)

			_, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, tc.withdrawer.Bytes())
			suite.Require().False(exists)
		} else if !tc.rewardExist {
			_, err := suite.app.EcosystemincentiveKeeper.WithdrawReward(suite.ctx, &types.MsgWithdrawReward{
				Sender: tc.withdrawer.Bytes(),
				Denom:  tc.reward.Denom,
			})
			suite.Require().Error(err)
			suite.Require().EqualError(err, sdkerrors.Wrap(types.ErrRewardNotExists, tc.withdrawer.String()).Error())

			_, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, tc.withdrawer.Bytes())
			suite.Require().False(exists)
		} else if !tc.validDenom {
			err := suite.app.EcosystemincentiveKeeper.SetRewardStore(suite.ctx, types.RewardStore{
				SubjectAddr: tc.withdrawer.Bytes(),
				Rewards:     sdk.NewCoins(tc.reward),
			})
			suite.Require().NoError(err)

			_, err = suite.app.EcosystemincentiveKeeper.WithdrawReward(suite.ctx, &types.MsgWithdrawReward{
				Sender: tc.withdrawer.Bytes(),
				Denom:  "invalid",
			})
			suite.Require().Error(err)
			suite.Require().EqualError(err, sdkerrors.Wrap(types.ErrDenomRewardNotExists, "invalid").Error())

			rewardStore, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, tc.withdrawer.Bytes())
			suite.Require().True(exists)
			rightRewardStore := types.RewardStore{
				SubjectAddr: tc.withdrawer.Bytes(),
				Rewards:     sdk.NewCoins(tc.reward),
			}
			suite.Require().Equal(rewardStore, rightRewardStore)
		}
	}
}

// TODO
func (suite *KeeperTestSuite) TestWithdrawAllReward() {

}
