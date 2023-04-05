package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/incentive/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetCdpMintingRewardPeriod returns the reward period with the specified collateral type if it's found in the params
func (k Keeper) GetCdpMintingRewardPeriod(ctx sdk.Context, collateralType string) (types.RewardPeriod, bool) {
	params := k.GetParams(ctx)
	for _, rp := range params.CdpMintingRewardPeriods {
		if rp.CollateralType == collateralType {
			return rp, true
		}
	}
	return types.RewardPeriod{}, false
}

// GetMultiplier returns the multiplier with the specified name if it's found in the params
func (k Keeper) GetMultiplier(ctx sdk.Context, name string) (types.Multiplier, bool) {
	params := k.GetParams(ctx)
	for _, m := range params.ClaimMultipliers {
		if m.Name == name {
			return m, true
		}
	}
	return types.Multiplier{}, false
}

// GetClaimEnd returns the claim end time for the params
func (k Keeper) GetClaimEnd(ctx sdk.Context) time.Time {
	params := k.GetParams(ctx)
	return params.ClaimEnd
}
