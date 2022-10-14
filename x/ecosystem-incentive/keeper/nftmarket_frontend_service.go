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
func (k Keeper) RecordNftIdWithIncentiveUnitId(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, incentiveUnitId string) {
	// panic if the nftId is already recorded in the store.
	if _, exists := k.GetNftIdForFrontend(ctx, nftId); exists {
		panic(sdkerrors.Wrap(types.ErrRecordedNftId, nftId.String()))
	}

	// check incentiveUnitId is already registered
	if _, exists := k.GetIncentiveUnit(ctx, incentiveUnitId); !exists {
		_ = fmt.Errorf(sdkerrors.Wrap(types.ErrNotRegisteredIncentiveUnitId, incentiveUnitId).Error())

		// emit event to inform that recording nftid failed because the incentiveUnitId is not registered yet.
		_ = ctx.EventManager().EmitTypedEvent(&types.EventNotRegisteredIncentiveUnitId{
			IncentiveUnitId: incentiveUnitId,
			ClassId:         nftId.ClassId,
			NftId:           nftId.NftId,
		})
		return
	}

	if err := k.SetNftIdForFrontend(ctx, nftId, incentiveUnitId); err != nil {
		panic(err)
	}

	// emit event to tell it succeeded.
	_ = ctx.EventManager().EmitTypedEvent(&types.EventRecordedIncentiveUnitId{
		IncentiveUnitId: incentiveUnitId,
		ClassId:         nftId.ClassId,
		NftId:           nftId.NftId,
	})
}

// DeleteNftId deletes nftId and incentiveUnitId from NftIdForFrontend KVStore to clean the record.
func (k Keeper) DeleteNftId(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier) {
	// If the passed NftId doesn't exist in the KVStore, emit the error message
	//  but not panic and just return
	incentiveUnitId, exists := k.GetNftIdForFrontend(ctx, nftId)
	if !exists {
		_ = fmt.Errorf(sdkerrors.Wrap(types.ErrNftIdForFrontendDoesntExist, nftId.String()).Error())
		return
	}

	k.DeleteNftIdForFrontend(ctx, nftId)

	// emit event for telling the nftId is deleted from the KVStore
	_ = ctx.EventManager().EmitTypedEvent(&types.EventDeletedNftIdRecordedForFrontendReward{
		IncentiveUnitId: incentiveUnitId,
		ClassId:         nftId.ClassId,
		NftId:           nftId.NftId,
	})
}

func (k Keeper) SetNftIdForFrontend(ctx sdk.Context, nftIdByte nftmarkettypes.NftIdentifier, incentiveUnitId string) error {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixNftIdForFrontend))
	prefixStore.Set(nftIdByte.IdBytes(), []byte(incentiveUnitId))

	return nil
}

func (k Keeper) GetNftIdForFrontend(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixNftIdForFrontend))

	bz := prefixStore.Get(nftId.IdBytes())
	if len(bz) == 0 {
		return "", false
	}

	return string(bz), true
}

func (k Keeper) DeleteNftIdForFrontend(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixNftIdForFrontend))

	prefixStore.Delete(nftId.IdBytes())
}

// AccumulateReward is called in AfterNftPaymentWithCommission hook method
// This method updates the reward information for the subject who is associated with the nftId
func (k Keeper) AccumulateRewardForFrontend(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, fee sdk.Coin) {
	// get incentiveUnitId by nftId from NftIdForFrontend KVStore
	incentiveUnitId, exists := k.GetNftIdForFrontend(ctx, nftId)
	if !exists {
		// emit event to inform the nftId is not associated with incentiveUnitId and return
		_ = ctx.EventManager().EmitTypedEvent(&types.EventNotRecordedNftId{
			ClassId: nftId.ClassId,
			NftId:   nftId.NftId,
		})
		return
	}

	incentiveUnit, exists := k.GetIncentiveUnit(ctx, incentiveUnitId)
	if !exists {
		// emit event to inform the incentiveUnit is not registered with incentiveUnitId and return
		_ = ctx.EventManager().EmitTypedEvent(&types.EventNotRegisteredIncentiveUnitId{
			IncentiveUnitId: incentiveUnitId,
		})
		return
	}

	nftmarketFrontendRewardRate := k.GetNftmarketFrontendRewardRate(ctx)

	// if the reward rate was not found, emit panic
	if nftmarketFrontendRewardRate == sdk.ZeroDec() {
		panic(sdkerrors.Wrap(types.ErrRewardRateNotFound, "nftmarket frontend"))
	}

	// rewardAmountForAll = fee * rewardRate
	rewardAmountForAll := nftmarketFrontendRewardRate.MulInt(fee.Amount).RoundInt()

	for _, subjectInfo := range incentiveUnit.SubjectInfoList {
		rewardStore, exists := k.GetRewardStore(ctx, subjectInfo.Address.AccAddress())
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
		if err := k.SetRewardStore(ctx, rewardStore); err != nil {
			panic(err)
		}
	}

	// emit event to inform that the incentiveUnit defined by incentiveUnitId
	// received new reward
	_ = ctx.EventManager().EmitTypedEvent(&types.EventUpdatedReward{
		IncentiveUnitId: incentiveUnitId,
		Reward:          sdk.NewCoin(fee.Denom, rewardAmountForAll),
	})
}

func (k Keeper) GetNftmarketFrontendRewardRate(ctx sdk.Context) sdk.Dec {
	params := k.GetParams(ctx)
	rewardParams := params.RewardParams

	for _, rewardParam := range rewardParams {
		if rewardParam.ModuleName == "nftmarket" {
			for _, rewardRate := range rewardParam.RewardRate {
				if rewardRate.RewardType == types.RewardType_NFTMARKET_FRONTEND {
					return rewardRate.Rate
				}
			}
		}
	}

	return sdk.ZeroDec()
}
