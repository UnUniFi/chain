package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// SetSymbolInfo set a specific SymbolInfo in the store
func (k Keeper) SetSymbolInfo(ctx sdk.Context, symbolInfo types.SymbolInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SymbolInfoKey))
	bz := k.cdc.MustMarshal(&symbolInfo)
	store.Set([]byte(symbolInfo.Symbol), bz)
}

// GetSymbolInfo returns a SymbolInfo from its id
func (k Keeper) GetSymbolInfo(ctx sdk.Context, symbol string) types.SymbolInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SymbolInfoKey))
	bz := store.Get([]byte(symbol))
	symbolInfo := types.SymbolInfo{}
	if bz == nil {
		return symbolInfo
	}

	k.cdc.MustUnmarshal(bz, &symbolInfo)
	return symbolInfo
}

// GetAllSymbolInfo returns a SymbolInfo
func (k Keeper) GetAllSymbolInfo(ctx sdk.Context) []types.SymbolInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SymbolInfoKey))

	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	infos := []types.SymbolInfo{}

	for ; iterator.Valid(); iterator.Next() {
		info := types.SymbolInfo{}
		k.cdc.MustUnmarshal(iterator.Value(), &info)
		infos = append(infos, info)
	}
	return infos
}
