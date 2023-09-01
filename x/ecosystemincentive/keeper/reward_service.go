package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) RewardDistributionOfNftbackedloan(ctx sdk.Context, nftId nftbackedloantypes.NftId, fee sdk.Coin) error {
	totalReward := sdk.ZeroInt()
	rewardForCommunityPool := sdk.ZeroInt()
	// First, get recipient by nftId from RecipientByNftId KVStore
	// If the recipient doesn't exist, return nil and distribute the reward for the frontend
	// to the treasury.
	nftbackedloanFrontendRewardRate := k.GetNftbackedloanFrontendRewardRate(ctx)
	if nftbackedloanFrontendRewardRate == sdk.ZeroDec() {
		return sdkerrors.Wrap(types.ErrRewardRateIsZero, "nftbackedloan frontend")
	}
	rewardForRecipient := nftbackedloanFrontendRewardRate.MulInt(fee.Amount).TruncateInt()
	totalReward = totalReward.Add(rewardForRecipient)
	recipient, exist := k.GetRecipientByNftId(ctx, nftId)
	if !exist {
		// emit event to inform the nftId is not associated with recipient and return
		_ = ctx.EventManager().EmitTypedEvent(&types.EventNotRecordedNftId{
			ClassId: nftId.ClassId,
			TokenId: nftId.TokenId,
		})

		// Distribute the reward to the community pool if there's no recipient associated with the nftId
		rewardForCommunityPool = rewardForCommunityPool.Add(rewardForRecipient)
		rewardForRecipient = sdk.ZeroInt()
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
	if !rewardForRecipient.IsZero() {
		if err := k.AccumulateRewardForFrontend(ctx, recipient, sdk.NewCoin(fee.Denom, rewardForRecipient)); err != nil {
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
	reward, exists := k.GetRewardRecord(ctx, senderAccAddr)
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
	// the RewardRecord data for the subject.
	// Otherwise, delete the data by key
	if reward.Rewards.Empty() {
		k.DeleteRewardRecord(ctx, sdk.MustAccAddressFromBech32(reward.Address))
	} else {
		if err := k.SetRewardRecord(ctx, reward); err != nil {
			return sdk.Coin{}, err
		}
	}

	return rewardCoin, nil
}

// WithdrawAllRewards is called to execute the actuall operation for MsgWithdrawAllRewards
// After sending the all accumulated rewards, delete types.Reward data from KVStore for the subject
func (k Keeper) WithdrawAllRewards(ctx sdk.Context, msg *types.MsgWithdrawAllRewards) (sdk.Coins, error) {
	senderAccAddr := sdk.MustAccAddressFromBech32(msg.Sender)
	reward, exists := k.GetRewardRecord(ctx, senderAccAddr)
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
	k.DeleteRewardRecord(ctx, senderAccAddr)
	return reward.Rewards, nil
}

func (k Keeper) SetRewardRecord(ctx sdk.Context, RewardRecord types.RewardRecord) error {
	bz, err := k.cdc.Marshal(&RewardRecord)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRewardStore))
	addressKeyBytes := sdk.MustAccAddressFromBech32(RewardRecord.Address).Bytes()
	prefixStore.Set(addressKeyBytes, bz)

	return nil
}

func (k Keeper) GetRewardRecord(ctx sdk.Context, subject sdk.AccAddress) (types.RewardRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRewardStore))

	bz := prefixStore.Get(subject)
	if bz == nil {
		return types.RewardRecord{}, false
	}

	var reward types.RewardRecord
	k.cdc.MustUnmarshal(bz, &reward)
	return reward, true
}

func (k Keeper) GetAllRewardRecords(ctx sdk.Context) []types.RewardRecord {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixRewardStore))
	defer it.Close()

	allRewardRecords := []types.RewardRecord{}
	for ; it.Valid(); it.Next() {
		var RewardRecord types.RewardRecord
		k.cdc.MustUnmarshal(it.Value(), &RewardRecord)

		allRewardRecords = append(allRewardRecords, RewardRecord)
	}

	return allRewardRecords
}

func (k Keeper) DeleteRewardRecord(ctx sdk.Context, subject sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixRewardStore))

	prefixStore.Delete(subject)
}
