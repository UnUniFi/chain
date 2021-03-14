package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/x/hard/types"
)

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetMinimumBorrowJPYValue returns the minimum borrow JPY value
func (k Keeper) GetMinimumBorrowJPYValue(ctx sdk.Context) sdk.Dec {
	params := k.GetParams(ctx)
	return params.MinimumBorrowJPYValue
}
