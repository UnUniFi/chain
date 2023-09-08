package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// SetDenomSymbolMap set a specific DenomSymbolMap in the store
func (k Keeper) SetDenomSymbolMap(ctx sdk.Context, denom string, symbol string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DenomSymbolMapKey))
	store.Set([]byte(denom), []byte(symbol))
}

// GetDenomSymbolMap returns a DenomSymbolMap from its id
func (k Keeper) GetDenomSymbolMap(ctx sdk.Context, denom string) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DenomSymbolMapKey))
	bz := store.Get([]byte(denom))
	if bz == nil {
		return ""
	}
	return string(bz)
}

// GetAllDenomSymbolMap returns a DenomSymbolMap
func (k Keeper) GetAllDenomSymbolMap(ctx sdk.Context) []types.Mapping {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DenomSymbolMapKey))

	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	mappings := []types.Mapping{}

	for ; iterator.Valid(); iterator.Next() {
		mappings = append(mappings, types.Mapping{
			Key:   string(iterator.Key()),
			Value: string(iterator.Value()),
		})
	}
	return mappings
}
