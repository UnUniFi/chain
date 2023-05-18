package keeper

import (
	"math"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cdptypes "github.com/UnUniFi/chain/x/deprecated/cdp/types"
	"github.com/UnUniFi/chain/x/deprecated/incentive/types"
)

// AccumulateCdpMintingRewards updates the rewards accumulated for the input reward period
func (k Keeper) AccumulateCdpMintingRewards(ctx sdk.Context, rewardPeriod types.RewardPeriod) error {
	previousAccrualTime, found := k.GetPreviousCdpMintingAccrualTime(ctx, rewardPeriod.CollateralType)
	if !found {
		k.SetPreviousCdpMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
		return nil
	}
	timeElapsed := CalculateTimeElapsed(rewardPeriod.Start, rewardPeriod.End, ctx.BlockTime(), previousAccrualTime)
	if timeElapsed.IsZero() {
		return nil
	}
	if rewardPeriod.RewardsPerSecond.Amount.IsZero() {
		k.SetPreviousCdpMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
		return nil
	}

	denoms, _ := k.GetGenesisDenoms(ctx)

	totalPrincipal := sdk.NewDecFromInt(k.cdpKeeper.GetTotalPrincipal(ctx, rewardPeriod.CollateralType, denoms.PrincipalDenom))
	if totalPrincipal.IsZero() {
		k.SetPreviousCdpMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
		return nil
	}
	newRewards := timeElapsed.Mul(rewardPeriod.RewardsPerSecond.Amount)
	cdpFactor, found := k.cdpKeeper.GetInterestFactor(ctx, rewardPeriod.CollateralType)
	if !found {
		k.SetPreviousCdpMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
		return nil
	}
	rewardFactor := sdk.NewDecFromInt(newRewards).Mul(cdpFactor).Quo(totalPrincipal)

	previousRewardFactor, found := k.GetCdpMintingRewardFactor(ctx, rewardPeriod.CollateralType)
	if !found {
		previousRewardFactor = sdk.ZeroDec()
	}
	newRewardFactor := previousRewardFactor.Add(rewardFactor)
	k.SetCdpMintingRewardFactor(ctx, rewardPeriod.CollateralType, newRewardFactor)
	k.SetPreviousCdpMintingAccrualTime(ctx, rewardPeriod.CollateralType, ctx.BlockTime())
	return nil
}

// InitializeCdpMintingClaim creates or updates a claim such that no new rewards are accrued, but any existing rewards are not lost.
// this function should be called after a cdp is created. If a user previously had a cdp, then closed it, they shouldn't
// accrue rewards during the period the cdp was closed. By setting the reward factor to the current global reward factor,
// any unclaimed rewards are preserved, but no new rewards are added.
func (k Keeper) InitializeCdpMintingClaim(ctx sdk.Context, cdp cdptypes.Cdp) {
	_, found := k.GetCdpMintingRewardPeriod(ctx, cdp.Type)
	if !found {
		// this collateral type is not incentivized, do nothing
		return
	}
	rewardFactor, found := k.GetCdpMintingRewardFactor(ctx, cdp.Type)
	if !found {
		rewardFactor = sdk.ZeroDec()
	}
	claim, found := k.GetCdpMintingClaim(ctx, cdp.Owner.AccAddress())

	denoms, _ := k.GetGenesisDenoms(ctx)

	if !found { // this is the owner's first jpu minting reward claim
		claim = types.NewCdpMintingClaim(cdp.Owner.AccAddress(), sdk.NewCoin(denoms.CdpMintingRewardDenom, sdk.ZeroInt()), types.RewardIndexes{types.NewRewardIndex(cdp.Type, rewardFactor)})
		k.SetCdpMintingClaim(ctx, claim)
		return
	}
	// the owner has an existing jpu minting reward claim
	index, hasRewardIndex := claim.HasRewardIndex(cdp.Type)
	if !hasRewardIndex { // this is the owner's first jpu minting reward for this collateral type
		claim.RewardIndexes = append(claim.RewardIndexes, types.NewRewardIndex(cdp.Type, rewardFactor))
	} else { // the owner has a previous jpu minting reward for this collateral type
		claim.RewardIndexes[index] = types.NewRewardIndex(cdp.Type, rewardFactor)
	}
	k.SetCdpMintingClaim(ctx, claim)
}

// SynchronizeCdpMintingReward updates the claim object by adding any accumulated rewards and updating the reward index value.
// this should be called before a cdp is modified, immediately after the 'SynchronizeInterest' method is called in the cdp module
func (k Keeper) SynchronizeCdpMintingReward(ctx sdk.Context, cdp cdptypes.Cdp) {
	_, found := k.GetCdpMintingRewardPeriod(ctx, cdp.Type)
	if !found {
		// this collateral type is not incentivized, do nothing
		return
	}

	globalRewardFactor, found := k.GetCdpMintingRewardFactor(ctx, cdp.Type)
	if !found {
		globalRewardFactor = sdk.ZeroDec()
	}
	claim, found := k.GetCdpMintingClaim(ctx, cdp.Owner.AccAddress())

	denoms, _ := k.GetGenesisDenoms(ctx)

	if !found {
		claim = types.NewCdpMintingClaim(cdp.Owner.AccAddress(), sdk.NewCoin(denoms.CdpMintingRewardDenom, sdk.ZeroInt()), types.RewardIndexes{types.NewRewardIndex(cdp.Type, globalRewardFactor)})
		k.SetCdpMintingClaim(ctx, claim)
		return
	}

	// the owner has an existing jpu minting reward claim
	index, hasRewardIndex := claim.HasRewardIndex(cdp.Type)
	if !hasRewardIndex { // this is the owner's first jpu minting reward for this collateral type
		claim.RewardIndexes = append(claim.RewardIndexes, types.NewRewardIndex(cdp.Type, globalRewardFactor))
		k.SetCdpMintingClaim(ctx, claim)
		return
	}
	userRewardFactor := claim.RewardIndexes[index].RewardFactor
	rewardsAccumulatedFactor := globalRewardFactor.Sub(userRewardFactor)
	if rewardsAccumulatedFactor.IsZero() {
		return
	}
	claim.RewardIndexes[index].RewardFactor = globalRewardFactor
	newRewardsAmount := rewardsAccumulatedFactor.Mul(sdk.NewDecFromInt(cdp.GetTotalPrincipal().Amount)).RoundInt()
	if newRewardsAmount.IsZero() {
		k.SetCdpMintingClaim(ctx, claim)
		return
	}
	newRewardsCoin := sdk.NewCoin(denoms.CdpMintingRewardDenom, newRewardsAmount)
	claim.Reward = claim.Reward.Add(newRewardsCoin)
	k.SetCdpMintingClaim(ctx, claim)
	return
}

// ZeroCdpMintingClaim zeroes out the claim object's rewards and returns the updated claim object
func (k Keeper) ZeroCdpMintingClaim(ctx sdk.Context, claim types.CdpMintingClaim) types.CdpMintingClaim {
	claim.Reward = sdk.NewCoin(claim.Reward.Denom, sdk.ZeroInt())
	k.SetCdpMintingClaim(ctx, claim)
	return claim
}

// SynchronizeCdpMintingClaim updates the claim object by adding any rewards that have accumulated.
// Returns the updated claim object
func (k Keeper) SynchronizeCdpMintingClaim(ctx sdk.Context, claim types.CdpMintingClaim) (types.CdpMintingClaim, error) {
	for _, ri := range claim.RewardIndexes {
		cdp, found := k.cdpKeeper.GetCdpByOwnerAndCollateralType(ctx, claim.Owner.AccAddress(), ri.CollateralType)
		if !found {
			// if the cdp for this collateral type has been closed, no updates are needed
			continue
		}
		claim = k.synchronizeRewardAndReturnClaim(ctx, cdp)
	}
	return claim, nil
}

// this function assumes a claim already exists, so don't call it if that's not the case
func (k Keeper) synchronizeRewardAndReturnClaim(ctx sdk.Context, cdp cdptypes.Cdp) types.CdpMintingClaim {
	k.SynchronizeCdpMintingReward(ctx, cdp)
	claim, _ := k.GetCdpMintingClaim(ctx, cdp.Owner.AccAddress())
	return claim
}

// CalculateTimeElapsed calculates the number of reward-eligible seconds that have passed since the previous
// time rewards were accrued, taking into account the end time of the reward period
func CalculateTimeElapsed(start, end, blockTime time.Time, previousAccrualTime time.Time) sdk.Int {
	if (end.Before(blockTime) &&
		(end.Before(previousAccrualTime) || end.Equal(previousAccrualTime))) ||
		(start.After(blockTime)) ||
		(start.Equal(blockTime)) {
		return sdk.ZeroInt()
	}
	if start.After(previousAccrualTime) && start.Before(blockTime) {
		previousAccrualTime = start
	}

	if end.Before(blockTime) {
		return sdk.MaxInt(sdk.ZeroInt(), sdk.NewInt(int64(math.RoundToEven(
			end.Sub(previousAccrualTime).Seconds(),
		))))
	}
	return sdk.MaxInt(sdk.ZeroInt(), sdk.NewInt(int64(math.RoundToEven(
		blockTime.Sub(previousAccrualTime).Seconds(),
	))))
}

// SimulateCdpMintingSynchronization calculates a user's outstanding Cdp minting rewards by simulating reward synchronization
func (k Keeper) SimulateCdpMintingSynchronization(ctx sdk.Context, claim types.CdpMintingClaim) types.CdpMintingClaim {
	for _, ri := range claim.RewardIndexes {
		_, found := k.GetCdpMintingRewardPeriod(ctx, ri.CollateralType)
		if !found {
			continue
		}

		globalRewardFactor, found := k.GetCdpMintingRewardFactor(ctx, ri.CollateralType)
		if !found {
			globalRewardFactor = sdk.ZeroDec()
		}

		// the owner has an existing jpu minting reward claim
		index, hasRewardIndex := claim.HasRewardIndex(ri.CollateralType)
		if !hasRewardIndex { // this is the owner's first jpu minting reward for this collateral type
			claim.RewardIndexes = append(claim.RewardIndexes, types.NewRewardIndex(ri.CollateralType, globalRewardFactor))
		}
		userRewardFactor := claim.RewardIndexes[index].RewardFactor
		rewardsAccumulatedFactor := globalRewardFactor.Sub(userRewardFactor)
		if rewardsAccumulatedFactor.IsZero() {
			continue
		}

		claim.RewardIndexes[index].RewardFactor = globalRewardFactor

		cdp, found := k.cdpKeeper.GetCdpByOwnerAndCollateralType(ctx, claim.GetOwner(), ri.CollateralType)
		if !found {
			continue
		}
		newRewardsAmount := rewardsAccumulatedFactor.Mul(sdk.NewDecFromInt(cdp.GetTotalPrincipal().Amount)).RoundInt()
		if newRewardsAmount.IsZero() {
			continue
		}

		denoms, _ := k.GetGenesisDenoms(ctx)

		newRewardsCoin := sdk.NewCoin(denoms.CdpMintingRewardDenom, newRewardsAmount)
		claim.Reward = claim.Reward.Add(newRewardsCoin)
	}

	return claim
}
