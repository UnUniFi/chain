package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// SetDenomInfo set a specific DenomInfo in the store
func (k Keeper) SetDenomInfo(ctx sdk.Context, denomInfo types.DenomInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DenomInfoKey))
	bz := k.cdc.MustMarshal(&denomInfo)
	store.Set([]byte(denomInfo.Denom), bz)
}

// GetDenomInfo returns a DenomInfo from its id
func (k Keeper) GetDenomInfo(ctx sdk.Context, denom string) types.DenomInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DenomInfoKey))
	denomInfo := types.DenomInfo{}
	bz := store.Get([]byte(denom))
	if bz == nil {
		return denomInfo
	}
	k.cdc.MustUnmarshal(bz, &denomInfo)
	return denomInfo
}

// GetAllDenomInfo returns a DenomInfo
func (k Keeper) GetAllDenomInfo(ctx sdk.Context) []types.DenomInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DenomInfoKey))

	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	infos := []types.DenomInfo{}

	for ; iterator.Valid(); iterator.Next() {
		info := types.DenomInfo{}
		k.cdc.MustUnmarshal(iterator.Value(), &info)
		infos = append(infos, info)
	}
	return infos
}
