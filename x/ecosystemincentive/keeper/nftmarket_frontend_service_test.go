package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/cometbft/cometbft/crypto/ed25519"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (suite *KeeperTestSuite) TestRecordIncentiveUnitIdWithNftId() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase             string
		classId              string
		nftId                string
		recipientContainerId string
		subjectAddrs         []string
		weights              []sdk.Dec
		registerBefore       bool
		expectPass           bool
	}{
		{
			testCase:             "not registered",
			classId:              "class1",
			nftId:                "nft1",
			recipientContainerId: "id1",
			subjectAddrs:         []string{sender.String()},
			weights:              []sdk.Dec{sdk.MustNewDecFromStr("1.0")},
			registerBefore:       false,
			expectPass:           false,
		},
		{
			testCase:             "registered",
			classId:              "class2",
			nftId:                "nft2",
			recipientContainerId: "id2",
			subjectAddrs:         []string{sender.String()},
			weights:              []sdk.Dec{sdk.MustNewDecFromStr("1.0")},
			registerBefore:       true,
			expectPass:           true,
		},
		{
			testCase:             "already recorded",
			classId:              "class2",
			nftId:                "nft2",
			recipientContainerId: "id3",
			subjectAddrs:         []string{sender.String()},
			weights:              []sdk.Dec{sdk.MustNewDecFromStr("1.0")},
			registerBefore:       true,
			expectPass:           false,
		},
	}
	for _, test := range tests {
		nftId := nftmarkettypes.NftIdentifier{
			ClassId: test.classId,
			NftId:   test.nftId,
		}

		if test.registerBefore {
			_, _ = suite.app.EcosystemincentiveKeeper.Register(
				suite.ctx,
				&types.MsgRegister{
					Sender:               sender.String(),
					RecipientContainerId: test.recipientContainerId,
					Addresses:            test.subjectAddrs,
					Weights:              test.weights,
				},
			)

			err := suite.RecordIncentiveUnitIdWithNftIdTest(suite.ctx, nftId, test.recipientContainerId)

			if test.expectPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		} else {
			err := suite.RecordIncentiveUnitIdWithNftIdTest(suite.ctx, nftId, test.recipientContainerId)
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestAccumulateRewardForFrontend() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase             string
		recipientContainerId string
		subjectAddrs         []string
		weights              []sdk.Dec
		fee                  sdk.Coin
		nftId                nftmarkettypes.NftIdentifier
		rewardAmount         math.Int
		expect               bool
		record               bool
		multipleSubject      bool
		amplify              bool
	}{
		{
			testCase:             "not recorded",
			recipientContainerId: "failure",
			subjectAddrs:         []string{sender.String()},
			weights:              []sdk.Dec{sdk.OneDec()},
			fee:                  sdk.Coin{},
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class2",
				NftId:   "nft2",
			},
			expect:          false,
			record:          false,
			multipleSubject: false,
			amplify:         false,
		},
		{
			testCase:             "single success case",
			recipientContainerId: "id1",
			subjectAddrs:         []string{sender.String()},
			weights:              []sdk.Dec{sdk.OneDec()},
			fee: sdk.Coin{
				Denom:  "uguu",
				Amount: sdk.NewInt(1000),
			},
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class1",
				NftId:   "nft1",
			},
			expect:          true,
			record:          true,
			multipleSubject: false,
			amplify:         false,
		},
		{
			testCase:             "multiple subject case",
			recipientContainerId: "id2",
			subjectAddrs: []string{
				sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()).String(),
				sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()).String(),
			},
			weights: []sdk.Dec{
				sdk.MustNewDecFromStr("0.5"),
				sdk.MustNewDecFromStr("0.5"),
			},
			fee: sdk.Coin{Denom: "uguu", Amount: math.NewInt(1000)},
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class2",
				NftId:   "nft2",
			},
			expect:          true,
			record:          true,
			multipleSubject: true,
			amplify:         false,
		},
		{
			testCase:             "amplify case",
			recipientContainerId: "id1",
			subjectAddrs:         []string{sender.String()},
			weights:              []sdk.Dec{sdk.OneDec()},
			fee: sdk.Coin{
				Denom:  "uguu",
				Amount: sdk.NewInt(1000),
			},
			nftId: nftmarkettypes.NftIdentifier{
				ClassId: "class3",
				NftId:   "nft3",
			},
			rewardAmount:    math.NewInt(500),
			expect:          true,
			record:          true,
			multipleSubject: false,
			amplify:         true,
		},
	}

	for _, test := range tests {
		_, _ = suite.app.EcosystemincentiveKeeper.Register(
			suite.ctx,
			&types.MsgRegister{
				Sender:               sender.String(),
				RecipientContainerId: test.recipientContainerId,
				Addresses:            test.subjectAddrs,
				Weights:              test.weights,
			})

		if test.record {
			suite.app.EcosystemincentiveKeeper.RecordRecipientContainerIdWithNftId(suite.ctx, test.nftId, test.recipientContainerId)
		}

		if test.expect {
			err := suite.AccumulateRewardForFrontendTest(suite.ctx, test.nftId, test.fee)
			suite.Require().NoError(err)

			// check the actual accumalted reward amount
			feeRate := suite.app.EcosystemincentiveKeeper.GetNftmarketFrontendRewardRate(suite.ctx)
			rewardAmount := feeRate.MulInt(test.fee.Amount).RoundInt()
			if test.multipleSubject {
				for index := range test.weights {
					rewardStore, _ := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, sdk.AccAddress(test.subjectAddrs[index]))
					suite.Require().Equal(test.weights[index].MulInt(rewardAmount).RoundInt(), rewardStore.Rewards.AmountOf(test.fee.Denom))
				}
			} else if test.amplify {
				rewardStore, _ := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, sdk.AccAddress(test.subjectAddrs[0]))
				suite.Require().Equal(test.weights[0].MulInt(rewardAmount).RoundInt().Add(test.rewardAmount), rewardStore.Rewards.AmountOf(test.fee.Denom))
			} else {
				rewardStore, _ := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, sdk.AccAddress(test.subjectAddrs[0]))
				suite.Require().Equal(test.weights[0].MulInt(rewardAmount).RoundInt(), rewardStore.Rewards.AmountOf(test.fee.Denom))
			}
		} else {
			err := suite.AccumulateRewardForFrontendTest(suite.ctx, test.nftId, test.fee)
			suite.Require().Error(err)

			_, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(suite.ctx, sender)
			suite.Require().False(exists)
		}
	}
}

// RecordIncentiveUnitIdWithNftIdTest is a mehtod to have the exact same logic
// for being used in test cases to return error as return value
// since the normal RecordIncentiveUnitIdWithNftId doesn't return any value by intention
func (suite *KeeperTestSuite) RecordIncentiveUnitIdWithNftIdTest(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, recipientContainerId string) error {
	// panic if the nftId is already recorded in the store.
	if _, exists := suite.app.EcosystemincentiveKeeper.GetRecipientContainerIdByNftId(ctx, nftId); exists {
		return types.ErrRecordedNftId
	}

	// check recipientContainerId is already registered
	if _, exists := suite.app.EcosystemincentiveKeeper.GetRecipientContainer(ctx, recipientContainerId); !exists {
		return types.ErrNotRegisteredRecipientContainerId
	}

	if err := suite.app.EcosystemincentiveKeeper.SetRecipientContainerIdByNftId(ctx, nftId, recipientContainerId); err != nil {
		return err
	}

	return nil
}

// Just mock method to use in only test
func (suite *KeeperTestSuite) AccumulateRewardForFrontendTest(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, fee sdk.Coin) error {
	// get recipientContainerId by nftId from IncentiveUnitIdByNftId KVStore
	recipientContainerId, exists := suite.app.EcosystemincentiveKeeper.GetRecipientContainerIdByNftId(ctx, nftId)
	if !exists {
		return types.ErrRegisteredIncentiveId
	}

	incentiveUnit, exists := suite.app.EcosystemincentiveKeeper.GetRecipientContainer(ctx, recipientContainerId)
	if !exists {
		return types.ErrNotRegisteredRecipientContainerId
	}

	nftmarketFrontendRewardRate := suite.app.EcosystemincentiveKeeper.GetNftmarketFrontendRewardRate(ctx)

	// if the reward rate was not found, emit panic
	if nftmarketFrontendRewardRate == sdk.ZeroDec() {
		return types.ErrRewardRateNotFound
	}

	// rewardAmountForAll = fee * rewardRate
	rewardAmountForAll := nftmarketFrontendRewardRate.MulInt(fee.Amount).RoundInt()

	for _, subjectInfo := range incentiveUnit.WeightedAddresses {
		rewardStore, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(ctx, sdk.AccAddress(subjectInfo.Address))
		if !exists {
			rewardStore = types.NewRewardStore(subjectInfo.Address, nil)
		}

		weight := subjectInfo.Weight

		// calculate actual reward to distribute for the subject addr by considering
		// its weight defined in IncentivenUnit
		// newRewardAmount = weight * rewardAmountForAll
		newRewardAmount := weight.MulInt(rewardAmountForAll).RoundInt()
		rewardCoin := sdk.NewCoin(fee.Denom, newRewardAmount)
		rewardStore.Rewards = rewardStore.Rewards.Add(sdk.NewCoins(rewardCoin)...)
		if err := suite.app.EcosystemincentiveKeeper.SetRewardStore(ctx, rewardStore); err != nil {
			return err
		}
	}

	return nil
}
