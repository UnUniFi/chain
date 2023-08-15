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

// RecordRecipientWithNftId is for recording recipient with nftId
// to know of the receriver of the incentive reward for the frontend creator
// of nftbackedloan in AfterNftPaymentWithCommission method.
func (k Keeper) RecordRecipientWithNftId(ctx sdk.Context, nftId nftbackedloantypes.NftId, recipient string) error {
	if _, exists := k.GetRecipientByNftId(ctx, nftId); exists {
		return sdkerrors.Wrap(types.ErrRecordedNftId, nftId.String())
	}

	err := k.SetRecipientByNftId(ctx, nftId, recipient)
	if err != nil {
		return err
	}

	// emit event to tell it succeeded.
	// err = ctx.EventManager().EmitTypedEvent(&types.EventRecordedRecipientContainerId{
	// 	RecipientContainerId: recipientContainerId,
	// 	ClassId:              nftId.ClassId,
	// 	NftId:                nftId.NftId,
	// })

	return err
}

// DeleteFrontendRecord is called in case to clean the record related for frontend incentive
func (k Keeper) DeleteFrontendRecord(ctx sdk.Context, nftId nftbackedloantypes.NftId) error {
	// If the passed NftId doesn't exist in the KVStore, emit the error message
	//  but not panic and just return
	recipient, exists := k.GetRecipientByNftId(ctx, nftId)
	if !exists {
		return fmt.Errorf(sdkerrors.Wrap(types.ErrRecipientContainerIdByNftIdDoesntExist, nftId.String()).Error())

	}

	k.DeleteRecipientByNftId(ctx, nftId)

	// emit event for telling the nftId is deleted from the KVStore
	err := ctx.EventManager().EmitTypedEvent(&types.EventDeletedNftIdRecordedForFrontendReward{
		Recipient: recipient,
		ClassId:   nftId.ClassId,
		TokenId:   nftId.TokenId,
	})

	return err
}

func (k Keeper) SetRecipientByNftId(ctx sdk.Context, nftIdByte nftbackedloantypes.NftId, recipient string) error {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainerIdByNftId))
	prefixStore.Set(nftIdByte.IdBytes(), []byte(recipient))

	return nil
}

func (k Keeper) GetRecipientByNftId(ctx sdk.Context, nftId nftbackedloantypes.NftId) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainerIdByNftId))

	bz := prefixStore.Get(nftId.IdBytes())
	if bz == nil {
		return "", false
	}

	return string(bz), true
}

// DeleteRecipientByNftId deletes nftId and recipient from RecipientByNftId KVStore to clean the record.
func (k Keeper) DeleteRecipientByNftId(ctx sdk.Context, nftId nftbackedloantypes.NftId) {
	// Delete incentive unit id by nft id
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRecipientContainerIdByNftId))

	prefixStore.Delete(nftId.IdBytes())
}

// AccumulateReward is called in AfterNftPaymentWithCommission hook method
// This method updates the reward information for the subject who is associated with the nftId
func (k Keeper) AccumulateRewardForFrontend(ctx sdk.Context, recipient string, reward sdk.Coin) error {
	addr, err := sdk.AccAddressFromBech32(recipient)
	if err != nil {
		return err
	}
	rewardStore, exists := k.GetRewardRecord(ctx, addr)
	if !exists {
		rewardStore = types.NewRewardRecord(recipient, nil)
	}

	rewardStore.Rewards = rewardStore.Rewards.Add(sdk.NewCoins(reward)...)
	if err := k.SetRewardRecord(ctx, rewardStore); err != nil {
		panic(err)
	}

	// emit event to inform the recipient
	// received new reward
	_ = ctx.EventManager().EmitTypedEvent(&types.EventUpdatedReward{
		Recipient:    recipient,
		EarnedReward: reward,
	})
	return nil
}

func (k Keeper) GetNftbackedloanFrontendRewardRate(ctx sdk.Context) sdk.Dec {
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
