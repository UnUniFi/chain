package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) SetRemainingMargin(ctx sdk.Context, positionId string, margin sdk.Coin) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshal(&margin)
	store.Set(types.RemainingMarginKeyPrefix(positionId), bz)
}

func (k Keeper) GetRemainingMargin(ctx sdk.Context, positionId string) *sdk.Coin {
	store := ctx.KVStore(k.storeKey)

	margin := sdk.Coin{}
	bz := store.Get(types.RemainingMarginKeyPrefix(positionId))

	if bz == nil {
		return nil
	}

	k.cdc.MustUnmarshal(bz, &margin)

	return &margin
}
