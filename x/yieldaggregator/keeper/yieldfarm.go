package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// TODO: implementation should be following the type of asset management target
func (k Keeper) InvestOnTarget(ctx sdk.Context, target types.AssetManagementTarget, unit types.FarmingUnit) {
	// TODO: set farming unit
	// TODO: move tokens to farm target
}

func (k Keeper) BeginWithdrawFromTarget(ctx sdk.Context, target types.AssetManagementTarget, unit types.FarmingUnit) {
	// TODO: request withdrawal from target by unit amount
}

func (k Keeper) ClaimWithdrawFromTarget(ctx sdk.Context, target types.AssetManagementTarget, unit types.FarmingUnit) {
	// TODO: check unbonding time passed
	// TODO: destroy farming unit and increase users' deposit balance
}

func (k Keeper) ClaimRewardsFromTarget(ctx sdk.Context, target types.AssetManagementTarget) {
	// TODO: claim and assign rewards to farm units
}
