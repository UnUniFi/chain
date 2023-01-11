package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

// WithdrawReward is called to execute the actuall operation for MsgWithdrawReward
func (k Keeper) WithdrawReward(ctx sdk.Context, msg *types.MsgWithdrawReward) (sdk.Coin, error) {
	reward, exists := k.GetRewardStore(ctx, msg.Sender.AccAddress())
	if !(exists) {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrRewardNotExists, msg.Sender.AccAddress().String())
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
		ctx, types.ModuleName, msg.Sender.AccAddress(),
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
		k.DeleteRewardStore(ctx, reward.SubjectAddr.AccAddress())
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
	reward, exists := k.GetRewardStore(ctx, msg.Sender.AccAddress())
	if !(exists) {
		return sdk.Coins{}, sdkerrors.Wrap(types.ErrRewardNotExists, msg.Sender.AccAddress().String())
	}

	// TODO: decide how to define module accout to send token
	// send coin from the specific module account which accumulate all fees on UnUniFi(?)
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, msg.Sender.AccAddress(), reward.Rewards); err != nil {
		return sdk.Coins{}, err
	}

	// delete types.Reward data from KVStore since it became none
	k.DeleteRewardStore(ctx, msg.Sender.AccAddress())
	return reward.Rewards, nil
}

func (k Keeper) SetRewardStore(ctx sdk.Context, rewardStore types.RewardStore) error {
	bz, err := k.cdc.Marshal(&rewardStore)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRewardStore))
	prefixStore.Set(rewardStore.SubjectAddr.AccAddress(), bz)

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

func (k Keeper) DeleteRewardStore(ctx sdk.Context, subject sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRewardStore))

	prefixStore.Delete(subject)
}
