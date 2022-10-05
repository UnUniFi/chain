package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

// WithdrawReward is called to execute the actuall operation for MsgWithdrawReward
func (k Keeper) WithdrawReward(ctx sdk.Context, msg *types.MsgWithdrawReward) (sdk.Coin, error) {
	reward, exists := k.GetReward(ctx, msg.Sender.AccAddress())
	if !(exists) {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrRewardNotExists, msg.Sender.AccAddress().String())
	}

	rewardAmount := reward.Rewards.AmountOf(msg.Denom)
	// if reward for specified denom doesn't exist, return err
	if rewardAmount.IsZero() {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrDenomRewardNotExists, msg.Denom)
	}

	// TODO: decide how to define module accout to send token
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

	return rewardCoin, nil
}

// WithdrawAllRewards is called to execute the actuall operation for MsgWithdrawAllRewards
// After sending the all accumulated rewards, delete types.Reward data from KVStore for the subject
func (k Keeper) WithdrawAllRewards(ctx sdk.Context, msg *types.MsgWithdrawAllRewards) (sdk.Coins, error) {
	reward, exists := k.GetReward(ctx, msg.Sender.AccAddress())
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
	k.DeleteReward(ctx, msg.Sender.AccAddress())
	return reward.Rewards, nil
}

// AccumulateReward is called in AfterNftPaymentWithCommission hook method
// This method updates the reward information for the subject who is associated with the nftId
func (k Keeper) AccumulateReward(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, fee sdk.Coin) {

	// Emit Event when to be accumulated reward
}

func (k Keeper) SetReward(ctx sdk.Context, reward types.RewardStore) error {
	bz, err := k.cdc.Marshal(&reward)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixReward))
	prefixStore.Set(reward.SubjectAddr, bz)

	return nil
}

func (k Keeper) GetReward(ctx sdk.Context, subject sdk.AccAddress) (types.RewardStore, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixReward))

	bz := prefixStore.Get(subject)
	if len(bz) == 0 {
		return types.RewardStore{}, false
	}

	var reward types.RewardStore
	k.cdc.MustUnmarshal(bz, &reward)
	return reward, true
}

func (k Keeper) DeleteReward(ctx sdk.Context, subject sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixIncentiveUnit))

	prefixStore.Delete(subject)
}
