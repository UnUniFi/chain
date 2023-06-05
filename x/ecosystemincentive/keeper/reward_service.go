package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

func (k Keeper) RewardDistributionOfNftmarket(ctx sdk.Context, nftId nftmarkettypes.NftIdentifier, fee sdk.Coin) error {
	totalReward := sdk.ZeroInt()
	// First, get incentiveUnitId by nftId from IncentiveUnitIdByNftId KVStore
	// If the incentiveUnitId doesn't exist, return nil and distribute the reward for the frontend
	// to the treasury.
	incentiveUnitId, exists := k.GetIncentiveUnitIdByNftId(ctx, nftId)
	if !exists {
		// emit event to inform the nftId is not associated with incentiveUnitId and return
		_ = ctx.EventManager().EmitTypedEvent(&types.EventNotRecordedNftId{
			ClassId: nftId.ClassId,
			NftId:   nftId.NftId,
		})

		// TODO: impl the logic to distribute the reward for the frontend to the treasury
	} else {
		nftmarketFrontendRewardRate := k.GetNftmarketFrontendRewardRate(ctx)
		// if the reward rate was not found or set as zero, just return
		if nftmarketFrontendRewardRate == sdk.ZeroDec() {
			err := fmt.Errorf(sdkerrors.Wrap(types.ErrRewardRateNotFound, "nftmarket frontend").Error())
			return err
		}
		rewardForIncentiveUnit := nftmarketFrontendRewardRate.MulInt(fee.Amount).TruncateInt()
		totalReward = totalReward.Add(rewardForIncentiveUnit)

		// Distribute the reward to the incentive unit
		if err := k.AccumulateRewardForFrontend(ctx, incentiveUnitId, sdk.NewCoin(fee.Denom, rewardForIncentiveUnit)); err != nil {
			return err
		}
	}

	stakersRewardRate := k.GetStakersRewardRate(ctx)
	// if the reward rate was not found or set as zero, just return
	if stakersRewardRate == sdk.ZeroDec() {
		err := fmt.Errorf(sdkerrors.Wrap(types.ErrRewardRateNotFound, "stakers").Error())
		return err
	}

	rewardForStakers := stakersRewardRate.MulInt(fee.Amount).TruncateInt()
	totalReward = totalReward.Add(rewardForStakers)

	// Emit panic if the reward for incentive unit exceeds the fee amount
	if totalReward.GT(fee.Amount) {
		panic(types.ErrRewardExceedsFee)
	}

	// Distribute the reward to the stakers
	if err := k.AllocateTokensToStakers(ctx, sdk.NewCoin(fee.Denom, rewardForStakers)); err != nil {
		return err
	}

	return nil
}

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
	prefixStore.Set(rewardStore.SubjectAddr.AccAddress().Bytes(), bz)

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
