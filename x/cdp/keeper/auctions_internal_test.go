package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) TestGetTotalDenom(ctx sdk.Context, accountName string, denom string) sdk.Int {
	return k.getTotalDenom(ctx, accountName, denom)
}
