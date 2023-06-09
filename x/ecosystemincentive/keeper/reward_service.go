package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) RewardDistributionOfNftmarket(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, fee sdk.Coin) error {
	totalReward := sdk.ZeroInt()
	rewardForCommunityPool := sdk.ZeroInt()
	// First, get recipientContainerId by nftId from RecipientContainerIdByNftId KVStore
	// If the recipientContainerId doesn't exist, return nil and distribute the reward for the frontend
	// to the treasury.
	nftmarketFrontendRewardRate := k.GetNftmarketFrontendRewardRate(ctx)
	if nftmarketFrontendRewardRate == sdk.ZeroDec() {
		return sdkerrors.Wrap(types.ErrRewardRateIsZero, "nftmarket frontend")
	}
	rewardForRecipientContainer := nftmarketFrontendRewardRate.MulInt(fee.Amount).TruncateInt()
	totalReward = totalReward.Add(rewardForRecipientContainer)
	recipientContainerId, exists := k.GetRecipientContainerIdByNftId(ctx, nftId)
	if !exists {
		// emit event to inform the nftId is not associated with recipientContainerId and return
		_ = ctx.EventManager().EmitTypedEvent(&types.EventNotRecordedNftId{
			ClassId: nftId.ClassId,
			NftId:   nftId.NftId,
		})

		// Distribute the reward to the community pool if there's no recipientContainerId associated with the nftId
		rewardForCommunityPool = rewardForCommunityPool.Add(rewardForRecipientContainer)
	} else {
		rewardForRecipientContainer = sdk.ZeroInt()
	}

	stakersRewardRate := k.GetStakersRewardRate(ctx)
	// if the reward rate was not found or set as zero, just return
	if stakersRewardRate == sdk.ZeroDec() {
		return sdkerrors.Wrap(types.ErrRewardRateNotFound, "stakers")
	}
	rewardForStakers := stakersRewardRate.MulInt(fee.Amount).TruncateInt()
	totalReward = totalReward.Add(rewardForStakers)

	communityPoolRate := k.GetCommunityPoolRewardRate(ctx)
	if communityPoolRate == sdk.ZeroDec() {
		return sdkerrors.Wrap(types.ErrRewardRateIsZero, communityPoolRate.String())
	}
	rewardForCommunityPool = rewardForCommunityPool.Add(communityPoolRate.MulInt(fee.Amount).TruncateInt())
	totalReward = totalReward.Add(rewardForCommunityPool)

	// TODO: we need better panic handling
	// Emit panic if the reward for incentive unit exceeds the fee amount
	if totalReward.GT(fee.Amount) {
		panic(types.ErrRewardExceedsFee)
	}

	// Distribute the reward to the recipients if the reward exists
	if !rewardForRecipientContainer.IsZero() {
		if err := k.AccumulateRewardForFrontend(ctx, recipientContainerId, sdk.NewCoin(fee.Denom, rewardForRecipientContainer)); err != nil {
			return err
		}
	}

	// Distribute the reward to the stakers if the reward exists
	if !rewardForStakers.IsZero() {
		if err := k.AllocateTokensToStakers(ctx, sdk.NewCoin(fee.Denom, rewardForStakers)); err != nil {
			return err
		}
	}

	// Distribute the reward to the community pool if the reward exists
	if !rewardForCommunityPool.IsZero() {
		if err := k.AllocateTokensToCommunityPool(ctx, sdk.NewCoin(fee.Denom, rewardForCommunityPool)); err != nil {
			return err
		}
	}

	return nil
}

// WithdrawReward is called to execute the actuall operation for MsgWithdrawReward
func (k Keeper) WithdrawReward(ctx sdk.Context, msg *types.MsgWithdrawReward) (sdk.Coin, error) {
	senderAccAddr := sdk.MustAccAddressFromBech32(msg.Sender)
	reward, exists := k.GetRewardStore(ctx, senderAccAddr)
	if !(exists) {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrRewardNotExists, msg.Sender)
	}

	rewardAmount := reward.Rewards.AmountOf(msg.Denom)
	// if reward for specified denom doesn't exist, return err
	if rewardAmount.IsZero() {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrDenomRewardNotExists, msg.Denom)
	}

	// TODO: decide how to define module account to send token
	// send coin from the specific module account which accumulate all fees on UnUniFi(?)
	rewardCoin := sdk.NewCoin(msg.Denom, rewardAmount)
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, senderAccAddr,
		sdk.NewCoins(rewardCoin)); err != nil {
		return sdk.Coin{}, err
	}

	reward.Rewards = reward.Rewards.Sub(rewardCoin)
	if !(reward.Rewards.AmountOf(msg.Denom).IsZero()) {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrCoinAmount, reward.Rewards.AmountOf(msg.Denom).String())
	}

	// If the reward for at least one denom remains, just reset
	// the RewardStore data for the subject.
	// Otherwise, delete the data by key
	if reward.Rewards.Empty() {
		k.DeleteRewardStore(ctx, sdk.MustAccAddressFromBech32(reward.Address))
	} else {
		if err := k.SetRewardStore(ctx, reward); err != nil {
			return sdk.Coin{}, err
		}
	}

	return rewardCoin, nil
}

// WithdrawAllRewards is called to execute the actuall operation for MsgWithdrawAllRewards
// After sending the all accumulated rewards, delete types.Reward data from KVStore for the subject
func (k Keeper) WithdrawAllRewards(ctx sdk.Context, msg *types.MsgWithdrawAllRewards) (sdk.Coins, error) {
	senderAccAddr := sdk.MustAccAddressFromBech32(msg.Sender)
	reward, exists := k.GetRewardStore(ctx, senderAccAddr)
	if !(exists) {
		return sdk.Coins{}, sdkerrors.Wrap(types.ErrRewardNotExists, msg.Sender)
	}

	// TODO: decide how to define module accout to send token
	// send coin from the specific module account which accumulate all fees on UnUniFi(?)
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, senderAccAddr, reward.Rewards); err != nil {
		return sdk.Coins{}, err
	}

	// delete types.Reward data from KVStore since it became none
	k.DeleteRewardStore(ctx, senderAccAddr)
	return reward.Rewards, nil
}

func (k Keeper) SetRewardStore(ctx sdk.Context, rewardStore types.RewardStore) error {
	bz, err := k.cdc.Marshal(&rewardStore)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRewardStore))
	addressKeyBytes := sdk.MustAccAddressFromBech32(rewardStore.Address).Bytes()
	prefixStore.Set(addressKeyBytes, bz)

	return nil
}

func (k Keeper) GetRewardStore(ctx sdk.Context, subject sdk.AccAddress) (types.RewardStore, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRewardStore))

	bz := prefixStore.Get(subject)
	if bz == nil {
		return types.RewardStore{}, false
	}

	var reward types.RewardStore
	k.cdc.MustUnmarshal(bz, &reward)
	return reward, true
}

func (k Keeper) GetAllRewardStores(ctx sdk.Context) []types.RewardStore {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixRewardStore))
	defer it.Close()

	allRewardStores := []types.RewardStore{}
	for ; it.Valid(); it.Next() {
		var rewardStore types.RewardStore
		k.cdc.MustUnmarshal(it.Value(), &rewardStore)

		allRewardStores = append(allRewardStores, rewardStore)
	}

	return allRewardStores
}

func (k Keeper) DeleteRewardStore(ctx sdk.Context, subject sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRewardStore))

	prefixStore.Delete(subject)
}
