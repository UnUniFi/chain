package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) TestAfterNftPaymentWithCommission() {
	tests := []struct {
		testCase             string
		nftId                nftbackedloantypes.NftId
		reward               sdk.Coin
		recipientContainerId string
		subjectAddrs         []string
		weights              []sdk.Dec
		recordedBefore       bool
		expectPass           bool
		// calculate the reward for recipientContainer using the default rate
		expRewardForRecipientContainer sdk.Int
		// calculate the reward for stakers using the default rate
		expRewardForStakers sdk.Coin
	}{
		{
			testCase: "failure case since incentive unit id was not recorded with nft id",
			nftId: nftbackedloantypes.NftId{
				ClassId: "class1",
				TokenId: "nft1",
			},
			reward:               sdk.NewCoin("uguu", sdk.NewInt(100)),
			recipientContainerId: "recipientContainerId1",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore: false,
			expectPass:     false,
		},
		{
			testCase: "failure case since there's no fee to distribute",
			nftId: nftbackedloantypes.NftId{
				ClassId: "class2",
				TokenId: "nft2",
			},
			reward:               sdk.NewCoin("uguu", sdk.NewInt(0)),
			recipientContainerId: "recipientContainerId1",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore: true,
			expectPass:     false,
		},
		{
			testCase: "ordinal case",
			nftId: nftbackedloantypes.NftId{
				ClassId: "class3",
				TokenId: "nft3",
			},
			reward:               sdk.NewCoin("uguu", sdk.NewInt(100)),
			recipientContainerId: "recipientContainerId3",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights:                        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore:                 true,
			expectPass:                     true,
			expRewardForRecipientContainer: sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt(),
			expRewardForStakers:            sdk.NewCoin("uguu", sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt()),
		},
		{
			testCase: "ordinal case",
			nftId: nftbackedloantypes.NftId{
				ClassId: "class4",
				TokenId: "nft4",
			},
			reward:               sdk.NewCoin("uguu", sdk.NewInt(100)),
			recipientContainerId: "recipientContainerId4",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights:                        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore:                 true,
			expectPass:                     true,
			expRewardForRecipientContainer: sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt(),
			expRewardForStakers:            sdk.NewCoin("uguu", sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt()),
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		if tc.recordedBefore {
			suite.app.EcosystemincentiveKeeper.RecordRecipientWithNftId(suite.ctx, tc.nftId, tc.recipientContainerId)
		}

		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.reward})
		_ = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{tc.reward})

		suite.app.EcosystemincentiveKeeper.Hooks().AfterNftPaymentWithCommission(suite.ctx, tc.nftId, tc.reward)

		if tc.expectPass {
			totalRewardForRecipientContainer := sdk.ZeroInt()
			for _, subject := range tc.subjectAddrs {
				rewardStore, exists := suite.app.EcosystemincentiveKeeper.GetRewardRecord(suite.ctx, sdk.AccAddress(subject))
				totalRewardForRecipientContainer = totalRewardForRecipientContainer.Add(rewardStore.Rewards.AmountOf(tc.reward.Denom))
				suite.Require().True(exists)
			}
			suite.Require().Equal(tc.expRewardForRecipientContainer, totalRewardForRecipientContainer)

			// check the reward distribution for stakers by checking the balance of the fee_collector module account
			feeCollectorAcc := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
			feeCollectorAccBalance := suite.app.BankKeeper.GetBalance(suite.ctx, feeCollectorAcc, tc.reward.Denom)
			suite.Require().Equal(tc.expRewardForStakers, feeCollectorAccBalance)
		} else {
			for _, subject := range tc.subjectAddrs {
				_, exists := suite.app.EcosystemincentiveKeeper.GetRewardRecord(suite.ctx, sdk.AccAddress(subject))
				suite.Require().False(exists)
			}
		}
	}
}
