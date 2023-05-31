// The implementations especially for the services about the Nftmarket Frontend reward.
// The reason why it's separated is for achieving the explicity and extensibility of this module.

package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

// RecordIncentiveIdWithNftId is for recording incentiveUnitId with nftId
// to know of the receriver of the incentive reward for the frontend creator
// of Nftmarket in AfterNftPaymentWithCommission method.
func (k Keeper) RecordIncentiveUnitIdWithNftId(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, incentiveUnitId string) {
	// panic if the nftId is already recorded in the store.
	if _, exists := k.GetIncentiveUnitIdByNftId(ctx, nftId); exists {
		panic(sdkerrors.Wrap(types.ErrRecordedNftId, nftId.String()))
	}

	// check incentiveUnitId is already registered
	if _, exists := k.GetIncentiveUnit(ctx, incentiveUnitId); !exists {
		k.Logger(ctx).Error(types.ErrNotRegisteredIncentiveUnitId.Error())

		// emit event to inform that recording nftid failed because the incentiveUnitId is not registered yet.
		_ = ctx.EventManager().EmitTypedEvent(&types.EventNotRegisteredIncentiveUnitId{
			IncentiveUnitId: incentiveUnitId,
			ClassId:         nftId.ClassId,
			NftId:           nftId.NftId,
		})
		return
	}

	if err := k.SetIncentiveUnitIdByNftId(ctx, nftId, incentiveUnitId); err != nil {
		panic(err)
	}

	// emit event to tell it succeeded.
	_ = ctx.EventManager().EmitTypedEvent(&types.EventRecordedIncentiveUnitId{
		IncentiveUnitId: incentiveUnitId,
		ClassId:         nftId.ClassId,
		NftId:           nftId.NftId,
	})
}

// DeleteFrontendRecord is called in case to clean the record related for frontend incentive
func (k Keeper) DeleteFrontendRecord(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier) {
	// If the passed NftId doesn't exist in the KVStore, emit the error message
	//  but not panic and just return
	incentiveUnitId, exists := k.GetIncentiveUnitIdByNftId(ctx, nftId)
	if !exists {
		_ = fmt.Errorf(sdkerrors.Wrap(types.ErrIncentiveUnitIdByNftIdDoesntExist, nftId.String()).Error())
		return
	}

	k.DeleteIncentiveUnitIdByNftId(ctx, nftId)

	// emit event for telling the nftId is deleted from the KVStore
	_ = ctx.EventManager().EmitTypedEvent(&types.EventDeletedNftIdRecordedForFrontendReward{
		IncentiveUnitId: incentiveUnitId,
		ClassId:         nftId.ClassId,
		NftId:           nftId.NftId,
	})
}

func (k Keeper) SetIncentiveUnitIdByNftId(ctx sdk.Context, nftIdByte nftmarkettypes.NftIdentifier, incentiveUnitId string) error {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnitIdByNftId))
	prefixStore.Set(nftIdByte.IdBytes(), []byte(incentiveUnitId))

	return nil
}

func (k Keeper) GetIncentiveUnitIdByNftId(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnitIdByNftId))

	bz := prefixStore.Get(nftId.IdBytes())
	if bz == nil {
		return "", false
	}

	return string(bz), true
}

// DeleteIncentiveUnitIdByNftId deletes nftId and incentiveUnitId from IncentiveUnitIdByNftId KVStore to clean the record.
func (k Keeper) DeleteIncentiveUnitIdByNftId(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier) {
	// Delete incentive unit id by nft id
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnitIdByNftId))

	prefixStore.Delete(nftId.IdBytes())
}

// AccumulateReward is called in AfterNftPaymentWithCommission hook method
// This method updates the reward information for the subject who is associated with the nftId
func (k Keeper) AccumulateRewardForFrontend(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, reward sdk.Coin) error {
	// get incentiveUnitId by nftId from IncentiveUnitIdByNftId KVStore
	incentiveUnitId, exists := k.GetIncentiveUnitIdByNftId(ctx, nftId)
	if !exists {
		// emit event to inform the nftId is not associated with incentiveUnitId and return
		_ = ctx.EventManager().EmitTypedEvent(&types.EventNotRecordedNftId{
			ClassId: nftId.ClassId,
			NftId:   nftId.NftId,
		})
		return types.ErrIncentiveUnitIdByNftIdDoesntExist
	}

	incentiveUnit, _ := k.GetIncentiveUnit(ctx, incentiveUnitId)

	// rewardAmountForAll = fee * rewardRate
	rewardsForEach := CalculateRewardsForEachSubject(
		extractWeightsFromSliceOfSubjectInfo(incentiveUnit.SubjectInfoLists),
		reward,
	)

	for i, subjectInfo := range incentiveUnit.SubjectInfoLists {
		rewardStore, exists := k.GetRewardStore(ctx, subjectInfo.SubjectAddr.AccAddress())
		if !exists {
			rewardStore = types.NewRewardStore(subjectInfo.SubjectAddr, nil)
		}

		rewardStore.Rewards = rewardStore.Rewards.Add(sdk.NewCoins(rewardsForEach[i])...)
		if err := k.SetRewardStore(ctx, rewardStore); err != nil {
			panic(err)
		}
	}

	// emit event to inform that the incentiveUnit defined by incentiveUnitId
	// received new reward
	_ = ctx.EventManager().EmitTypedEvent(&types.EventUpdatedReward{
		IncentiveUnitId: incentiveUnitId,
		EarnedReward:    reward,
	})
	return nil
}

// calculate actual reward to distribute for the subject addr by considering
// its weight defined in IncentivenUnit
// newRewardAmount = weight * rewardAmountForAll
func CalculateRewardsForEachSubject(weights []sdk.Dec, reward sdk.Coin) []sdk.Coin {
	var rewardsForEach []sdk.Coin

	for _, weight := range weights {
		newRewardAmount := weight.MulInt(reward.Amount).TruncateInt()
		rewardCoin := sdk.NewCoin(reward.Denom, newRewardAmount)
		rewardsForEach = append(rewardsForEach, rewardCoin)
	}

	return rewardsForEach
}

func extractWeightsFromSliceOfSubjectInfo(subjectsInfo []types.SubjectInfo) []sdk.Dec {
	var weights []sdk.Dec
	for _, subject := range subjectsInfo {
		weights = append(weights, subject.Weight)
	}
	return weights
}

func (k Keeper) GetNftmarketFrontendRewardRate(ctx sdk.Context) sdk.Dec {
	params := k.GetParams(ctx)
	rewardParams := params.RewardParams

	for _, rewardParam := range rewardParams {
		if rewardParam.ModuleName == nftmarkettypes.ModuleName {
			for _, rewardRate := range rewardParam.RewardRate {
				if rewardRate.RewardType == types.RewardType_NFTMARKET_FRONTEND {
					return rewardRate.Rate
				}
			}
		}
	}

	// if target param wasn't found somehow, return zero dec
	return sdk.ZeroDec()
}
