package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) TestAfterNftListed() {
	tests := []struct {
		testCase             string
		nftId                nftmarkettypes.NftIdentifier
		recipientContainerId string
		subjectAddrs         []string
		weights              []sdk.Dec
		txMemo               string
		registerBefore       bool
		expectPass           bool
		expectPanic          bool
		memoFormat           bool
	}{
		{
			testCase: "ordinal success case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class1",
				NftId:   "nft1",
			},
			recipientContainerId: "recipientContainerId1",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			txMemo:         `{"version":"v1","incentive_unit_id":"recipientContainerId1"}`,
			registerBefore: true,
			expectPass:     true,
			expectPanic:    false,
			memoFormat:     true,
		},
		{
			testCase: "incentive unit id is not registered",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class2",
				NftId:   "nft2",
			},
			txMemo:         `{"version":"v1","incentive_unit_id":"recipientContainerId2"}`,
			registerBefore: false,
			expectPass:     false,
			expectPanic:    false,
			memoFormat:     true,
		},
		{
			testCase: "panic since incentive unit id is already recorded with nft id",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class1",
				NftId:   "nft1",
			},
			txMemo:         `{"version":"v1","incentive_unit_id":"recipientContainerId2"}`,
			registerBefore: false,
			expectPass:     false,
			expectPanic:    true,
			memoFormat:     true,
		},
		{
			testCase: "invalid memo format",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class3",
				NftId:   "nft3",
			},
			recipientContainerId: "recipientContainerId3",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			txMemo:         `{"error":true,"version":"v1","incentive_unit_id":"recipientContainerId3"}`,
			registerBefore: true,
			expectPass:     false,
			expectPanic:    false,
			memoFormat:     false,
		},
		{
			testCase: "invalid memo format version",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class4",
				NftId:   "nft4",
			},
			recipientContainerId: "recipientContainerId4",
			subjectAddrs: []string{
				suite.addrs[0].String(),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			txMemo:         `{"version":"v0","incentive_unit_id":"recipientContainerId4"}`,
			registerBefore: true,
			expectPass:     false,
			expectPanic:    false,
			memoFormat:     false,
		},
	}

	for _, tc := range tests {
		if tc.registerBefore {
			_, err := suite.app.EcosystemincentiveKeeper.Register(suite.ctx, &types.MsgRegister{
				Sender:               tc.subjectAddrs[0],
				RecipientContainerId: tc.recipientContainerId,
				Addresses:            tc.subjectAddrs,
				Weights:              tc.weights,
			})
			suite.Require().NoError(err)
		}

		if tc.expectPanic {
			suite.Panics(func() {
				suite.app.EcosystemincentiveKeeper.Hooks().AfterNftListed(suite.ctx, tc.nftId, tc.txMemo)
			})
		} else {
			suite.NotPanics(func() {
				suite.app.EcosystemincentiveKeeper.Hooks().AfterNftListed(suite.ctx, tc.nftId, tc.txMemo)
			})
		}

		recipientContainerId, exists := suite.app.EcosystemincentiveKeeper.GetRecipientContainerIdByNftId(suite.ctx, tc.nftId)
		if tc.expectPass {
			suite.Require().True(exists)
			suite.Require().Equal(tc.recipientContainerId, recipientContainerId)
		} else if tc.expectPanic {
			suite.Require().True(exists)
			suite.Require().NotEqual(tc.recipientContainerId, recipientContainerId)
		} else {
			suite.Require().False(exists)
		}
	}
}

func (suite *KeeperTestSuite) TestAfterNftPaymentWithCommission() {
	tests := []struct {
		testCase             string
		nftId                nftmarkettypes.NftIdentifier
		reward               sdk.Coin
		recipientContainerId string
		subjectAddrs         []ununifitypes.StringAccAddress
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
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class1",
				NftId:   "nft1",
			},
			reward:               sdk.Coin{"uguu", sdk.NewInt(100)},
			recipientContainerId: "recipientContainerId1",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore: false,
			expectPass:     false,
		},
		{
			testCase: "failure case since there's no fee to distribute",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class2",
				NftId:   "nft2",
			},
			reward:               sdk.Coin{"uguu", sdk.NewInt(0)},
			recipientContainerId: "recipientContainerId1",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore: true,
			expectPass:     false,
		},
		{
			testCase: "ordinal case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class3",
				NftId:   "nft3",
			},
			reward:               sdk.Coin{"uguu", sdk.NewInt(100)},
			recipientContainerId: "recipientContainerId3",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:                        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			recordedBefore:                 true,
			expectPass:                     true,
			expRewardForRecipientContainer: sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt(),
			expRewardForStakers:            sdk.NewCoin("uguu", sdk.MustNewDecFromStr("0.5").MulInt(sdk.NewInt(100)).TruncateInt()),
		},
		{
			testCase: "ordinal case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class4",
				NftId:   "nft4",
			},
			reward:               sdk.Coin{"uguu", sdk.NewInt(100)},
			recipientContainerId: "recipientContainerId4",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
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
		_, err := suite.app.EcosystemincentiveKeeper.Register(suite.ctx, &types.MsgRegister{
			Sender:               tc.subjectAddrs[0],
			RecipientContainerId: tc.recipientContainerId,
			Addresses:            tc.subjectAddrs,
			Weights:              tc.weights,
		})
		suite.Require().NoError(err)
		if tc.recordedBefore {
			suite.app.EcosystemincentiveKeeper.RecordRecipientContainerIdWithNftId(suite.ctx, tc.nftId, tc.recipientContainerId)
		}

		_ = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{tc.reward})
		_ = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{tc.reward})

		suite.app.EcosystemincentiveKeeper.Hooks().AfterNftPaymentWithCommission(suite.ctx, tc.nftId, tc.reward)

		if tc.expectPass {
			totalRewardForRecipientContainer := sdk.ZeroInt()
			for _, subject := range tc.subjectAddrs {
				rewardStore, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, sdk.AccAddress(subject))
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
				_, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, sdk.AccAddress(subject))
				suite.Require().False(exists)
			}
		}

		_, exists := suite.app.EcosystemincentiveKeeper.GetRecipientContainerIdByNftId(suite.ctx, tc.nftId)
		suite.Require().False(exists)
	}
}

func (suite *KeeperTestSuite) TestAfterNftUnlistedWithoutPayment() {
	tests := []struct {
		testCase             string
		nftId                nftmarkettypes.NftIdentifier
		recipientContainerId string
		subjectAddrs         []ununifitypes.StringAccAddress
		weights              []sdk.Dec
		registerBefore       bool
		expectPass           bool
	}{
		{
			testCase: "ordinal case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class1",
				NftId:   "nft1",
			},
			recipientContainerId: "recipientContainerId1",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			registerBefore: true,
		},
		{
			testCase: "not recorded case",
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class2",
				NftId:   "nft2",
			},
			recipientContainerId: "recipientContainerId2",
			subjectAddrs: []ununifitypes.StringAccAddress{
				ununifitypes.StringAccAddress(suite.addrs[0]),
			},
			weights:        []sdk.Dec{sdk.MustNewDecFromStr("1")},
			registerBefore: true,
		},
	}

	for _, tc := range tests {
		_, err := suite.app.EcosystemincentiveKeeper.Register(suite.ctx, &types.MsgRegister{
			Sender:               tc.subjectAddrs[0],
			RecipientContainerId: tc.recipientContainerId,
			Addresses:            tc.subjectAddrs,
			Weights:              tc.weights,
		})
		suite.Require().NoError(err)
		if tc.registerBefore {
			suite.app.EcosystemincentiveKeeper.RecordRecipientContainerIdWithNftId(suite.ctx, tc.nftId, tc.recipientContainerId)
		}

		suite.NotPanics(func() {
			suite.app.EcosystemincentiveKeeper.Hooks().AfterNftUnlistedWithoutPayment(suite.ctx, tc.nftId)
		})

		_, exists := suite.app.EcosystemincentiveKeeper.GetRecipientContainerIdByNftId(suite.ctx, tc.nftId)
		suite.Require().False(exists)
	}
}
