package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/tendermint/tendermint/crypto/ed25519"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

func (suite *KeeperTestSuite) TestRecordNftIdWithIncentiveUnitId() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase        string
		classId         string
		nftId           string
		incentiveUnitId string
		subjectAddrs    []ununifitypes.StringAccAddress
		weights         []sdk.Dec
		registerBefore  bool
		expectPass      bool
	}{
		{
			testCase:        "not registered",
			classId:         "class1",
			nftId:           "nft1",
			incentiveUnitId: "id1",
			subjectAddrs:    []ununifitypes.StringAccAddress{sender.Bytes()},
			weights:         []sdk.Dec{sdk.MustNewDecFromStr("1.0")},
			registerBefore:  false,
			expectPass:      false,
		},
		{
			testCase:        "registered",
			classId:         "class2",
			nftId:           "nft2",
			incentiveUnitId: "id2",
			subjectAddrs:    []ununifitypes.StringAccAddress{sender.Bytes()},
			weights:         []sdk.Dec{sdk.MustNewDecFromStr("1.0")},
			registerBefore:  true,
			expectPass:      true,
		},
		{
			testCase:        "already recorded",
			classId:         "class2",
			nftId:           "nft2",
			incentiveUnitId: "id3",
			subjectAddrs:    []ununifitypes.StringAccAddress{sender.Bytes()},
			weights:         []sdk.Dec{sdk.MustNewDecFromStr("1.0")},
			registerBefore:  true,
			expectPass:      false,
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
					Sender:          sender.Bytes(),
					IncentiveUnitId: test.incentiveUnitId,
					SubjectAddrs:    test.subjectAddrs,
					Weights:         test.weights,
				},
			)

			err := suite.RecordNftIdWithIncentiveUnitIdTest(suite.ctx, nftId, test.incentiveUnitId)

			if test.expectPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		} else {
			err := suite.RecordNftIdWithIncentiveUnitIdTest(suite.ctx, nftId, test.incentiveUnitId)
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestAccumulateRewardForFrontend() {
	sender := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	tests := []struct {
		testCase        string
		incentiveUnitId string
		subjectAddrs    []ununifitypes.StringAccAddress
		weights         []sdk.Dec
		fee             sdk.Coin
		nftId           nftmarkettypes.NftIdentifier
		rewardAmount    math.Int
		expect          bool
		record          bool
		multipleSubject bool
		amplify         bool
	}{
		{
			testCase:        "not recorded",
			incentiveUnitId: "failure",
			subjectAddrs:    []ununifitypes.StringAccAddress{sender.Bytes()},
			weights:         []sdk.Dec{sdk.OneDec()},
			fee:             sdk.Coin{},
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
			testCase:        "single success case",
			incentiveUnitId: "id1",
			subjectAddrs:    []ununifitypes.StringAccAddress{sender.Bytes()},
			weights:         []sdk.Dec{sdk.OneDec()},
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
			testCase:        "multiple subject case",
			incentiveUnitId: "id2",
			subjectAddrs: []ununifitypes.StringAccAddress{
				sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()).Bytes(),
				sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()).Bytes(),
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
			testCase:        "amplify case",
			incentiveUnitId: "id1",
			subjectAddrs:    []ununifitypes.StringAccAddress{sender.Bytes()},
			weights:         []sdk.Dec{sdk.OneDec()},
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
				Sender:          sender.Bytes(),
				IncentiveUnitId: test.incentiveUnitId,
				SubjectAddrs:    test.subjectAddrs,
				Weights:         test.weights,
			})

		if test.record {
			suite.app.EcosystemincentiveKeeper.RecordNftIdWithIncentiveUnitId(suite.ctx, test.nftId, test.incentiveUnitId)
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

// RecordNftIdWithIncentiveUnitIdTest is a mehtod to have the exact same logic
// for being used in test cases to return error as return value
// since the normal RecordNftIdWithIncentiveUnitId doesn't return any value by intention
func (suite *KeeperTestSuite) RecordNftIdWithIncentiveUnitIdTest(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, incentiveUnitId string) error {
	// panic if the nftId is already recorded in the store.
	if _, exists := suite.app.EcosystemincentiveKeeper.GetIncentiveUnitIdByNftId(ctx, nftId); exists {
		return types.ErrRecordedNftId
	}

	// check incentiveUnitId is already registered
	if _, exists := suite.app.EcosystemincentiveKeeper.GetIncentiveUnit(ctx, incentiveUnitId); !exists {
		return types.ErrNotRegisteredIncentiveUnitId
	}

	if err := suite.app.EcosystemincentiveKeeper.SetIncentiveUnitIdByNftId(ctx, nftId, incentiveUnitId); err != nil {
		return err
	}

	return nil
}

// Just mock method to use in only test
func (suite *KeeperTestSuite) AccumulateRewardForFrontendTest(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, fee sdk.Coin) error {
	// get incentiveUnitId by nftId from IncentiveUnitIdByNftId KVStore
	incentiveUnitId, exists := suite.app.EcosystemincentiveKeeper.GetIncentiveUnitIdByNftId(ctx, nftId)
	if !exists {
		return types.ErrIncentiveUnitIdByNftIdDoesntExist
	}

	incentiveUnit, exists := suite.app.EcosystemincentiveKeeper.GetIncentiveUnit(ctx, incentiveUnitId)
	if !exists {
		return types.ErrNotRegisteredIncentiveUnitId
	}

	nftmarketFrontendRewardRate := suite.app.EcosystemincentiveKeeper.GetNftmarketFrontendRewardRate(ctx)

	// if the reward rate was not found, emit panic
	if nftmarketFrontendRewardRate == sdk.ZeroDec() {
		return types.ErrRewardRateNotFound
	}

	// rewardAmountForAll = fee * rewardRate
	rewardAmountForAll := nftmarketFrontendRewardRate.MulInt(fee.Amount).RoundInt()

	for _, subjectInfo := range incentiveUnit.SubjectInfoList {
		rewardStore, exists := suite.app.EcosystemincentiveKeeper.GetRewardStore(ctx, subjectInfo.Address.AccAddress())
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
