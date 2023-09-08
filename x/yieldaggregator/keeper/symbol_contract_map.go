package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// SetSymbolSwapContractMap set a specific SymbolSwapContractMap in the store
func (k Keeper) SetSymbolSwapContractMap(ctx sdk.Context, symbol string, contract string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SymbolContractMapKey))
	store.Set([]byte(symbol), []byte(contract))
}

// GetSymbolSwapContractMap returns a SymbolSwapContractMap from its id
func (k Keeper) GetSymbolSwapContractMap(ctx sdk.Context, symbol string) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SymbolContractMapKey))
	bz := store.Get([]byte(symbol))
	if bz == nil {
		return ""
	}
	return string(bz)
}

// GetAllSymbolSwapContractMap returns a SymbolSwapContractMap
func (k Keeper) GetAllSymbolSwapContractMap(ctx sdk.Context) []types.Mapping {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SymbolContractMapKey))

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
