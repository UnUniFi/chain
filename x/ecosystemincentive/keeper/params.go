package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetMaxIncentiveUnitIdLen(ctx sdk.Context) uint64 {
	return k.GetParams(ctx).MaxIncentiveUnitIdLen
}

func (k Keeper) GetMaxSubjectInfoNumInUnitParam(ctx sdk.Context) uint64 {
	return k.GetParams(ctx).MaxSubjectInfoNumInUnit
}
