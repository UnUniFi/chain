// The implementations especially for the services about the nftbackedloan Frontend reward.
// The reason why it's separated is for achieving the explicity and extensibility of this module.

package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

// RecordIncentiveIdWithNftId is for recording recipientContainerId with nftId
// to know of the receriver of the incentive reward for the frontend creator
// of nftbackedloan in AfterNftPaymentWithCommission method.
func (k Keeper) RecordRecipientContainerIdWithNftId(ctx sdk.Context, nftId nftbackedloantypes.NftIdentifier, recipientContainerId string) {
	// panic if the nftId is already recorded in the store.
	if _, exists := k.GetRecipientContainerIdByNftId(ctx, nftId); exists {
		panic(sdkerrors.Wrap(types.ErrRecordedNftId, nftId.String()))
	}

	// check recipientContainerId is already registered
	if _, exists := k.GetRecipientContainer(ctx, recipientContainerId); !exists {
		k.Logger(ctx).Error(types.ErrNotRegisteredRecipientContainerId.Error())

		// emit event to inform that recording nftid failed because the recipientContainerId is not registered yet.
		_ = ctx.EventManager().EmitTypedEvent(&types.EventNotRegisteredRecipientContainerId{
			RecipientContainerId: recipientContainerId,
			ClassId:              nftId.ClassId,
			NftId:                nftId.NftId,
		})
		return
	}

	if err := k.SetRecipientContainerIdByNftId(ctx, nftId, recipientContainerId); err != nil {
		panic(err)
	}

	// emit event to tell it succeeded.
	_ = ctx.EventManager().EmitTypedEvent(&types.EventRecordedRecipientContainerId{
		RecipientContainerId: recipientContainerId,
		ClassId:              nftId.ClassId,
		NftId:                nftId.NftId,
	})
}

// DeleteFrontendRecord is called in case to clean the record related for frontend incentive
func (k Keeper) DeleteFrontendRecord(ctx sdk.Context, nftId nftbackedloantypes.NftIdentifier) {
	// If the passed NftId doesn't exist in the KVStore, emit the error message
	//  but not panic and just return
	recipientContainerId, exists := k.GetRecipientContainerIdByNftId(ctx, nftId)
	if !exists {
		_ = fmt.Errorf(sdkerrors.Wrap(types.ErrRecipientContainerIdByNftIdDoesntExist, nftId.String()).Error())
		return
	}

	k.DeleteRecipientContainerIdByNftId(ctx, nftId)

	// emit event for telling the nftId is deleted from the KVStore
	_ = ctx.EventManager().EmitTypedEvent(&types.EventDeletedNftIdRecordedForFrontendReward{
		RecipientContainerId: recipientContainerId,
		ClassId:              nftId.ClassId,
		NftId:                nftId.NftId,
	})
}

func (k Keeper) SetRecipientContainerIdByNftId(ctx sdk.Context, nftIdByte nftbackedloantypes.NftIdentifier, recipientContainerId string) error {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainerIdByNftId))
	prefixStore.Set(nftIdByte.IdBytes(), []byte(recipientContainerId))

	return nil
}

func (k Keeper) GetRecipientContainerIdByNftId(ctx sdk.Context, nftId nftbackedloantypes.NftIdentifier) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainerIdByNftId))

	bz := prefixStore.Get(nftId.IdBytes())
	if bz == nil {
		return "", false
	}

	return string(bz), true
}

// DeleteRecipientContainerIdByNftId deletes nftId and recipientContainerId from RecipientContainerIdByNftId KVStore to clean the record.
func (k Keeper) DeleteRecipientContainerIdByNftId(ctx sdk.Context, nftId nftbackedloantypes.NftIdentifier) {
	// Delete incentive unit id by nft id
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainerIdByNftId))

	prefixStore.Delete(nftId.IdBytes())
}

// AccumulateReward is called in AfterNftPaymentWithCommission hook method
// This method updates the reward information for the subject who is associated with the nftId
func (k Keeper) AccumulateRewardForFrontend(ctx sdk.Context, recipientContainerId string, reward sdk.Coin) error {
	recipientContainer, _ := k.GetRecipientContainer(ctx, recipientContainerId)

	// rewardAmountForAll = fee * rewardRate
	rewardsForEach := CalculateRewardsForEachSubject(
		extractWeightsFromSliceOfSubjectInfo(recipientContainer.WeightedAddresses),
		reward,
	)

	for i, subjectInfo := range recipientContainer.WeightedAddresses {
		rewardStore, exists := k.GetRewardStore(ctx, sdk.MustAccAddressFromBech32(subjectInfo.Address))
		if !exists {
			rewardStore = types.NewRewardStore(subjectInfo.Address, nil)
		}

		rewardStore.Rewards = rewardStore.Rewards.Add(sdk.NewCoins(rewardsForEach[i])...)
		if err := k.SetRewardStore(ctx, rewardStore); err != nil {
			panic(err)
		}
	}

	// emit event to inform that the recipientContainer defined by recipientContainerId
	// received new reward
	_ = ctx.EventManager().EmitTypedEvent(&types.EventUpdatedReward{
		RecipientContainerId: recipientContainerId,
		EarnedReward:         reward,
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

func extractWeightsFromSliceOfSubjectInfo(subjectsInfo []types.WeightedAddress) []sdk.Dec {
	var weights []sdk.Dec
	for _, subject := range subjectsInfo {
		weights = append(weights, subject.Weight)
	}
	return weights
}

func (k Keeper) GetnftbackedloanFrontendRewardRate(ctx sdk.Context) sdk.Dec {
	params := k.GetParams(ctx)
	rewardParams := params.RewardParams

	for _, rewardParam := range rewardParams {
		if rewardParam.ModuleName == nftbackedloantypes.ModuleName {
			for _, rewardRate := range rewardParam.RewardRate {
				if rewardRate.RewardType == types.RewardType_FRONTEND_DEVELOPERS {
					return rewardRate.Rate
				}
			}
		}
	}

	// if target param wasn't found somehow, return zero dec
	return sdk.ZeroDec()
}
