package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

// SetTranchePool set a specific TranchePool in the store
func (k Keeper) SetTranchePool(ctx sdk.Context, tranchePool types.TranchePool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	b := k.cdc.MustMarshal(&tranchePool)
	store.Set(sdk.Uint64ToBigEndian(tranchePool.Id), b)

	store = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrancheByStrategyKey))
	store.Set(types.KeyTrancheByStrategy(tranchePool), sdk.Uint64ToBigEndian(tranchePool.Id))
}

// GetTranchePool returns a TranchePool from its identifier
func (k Keeper) GetTranchePool(ctx sdk.Context, id uint64) (val types.TranchePool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	b := store.Get(sdk.Uint64ToBigEndian(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTranchePool removes a TranchePool from the store
func (k Keeper) RemoveTranchePool(ctx sdk.Context, tranchePool types.TranchePool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	store.Delete(sdk.Uint64ToBigEndian(tranchePool.Id))

	store = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrancheByStrategyKey))
	store.Delete(types.KeyTrancheByStrategy(tranchePool))
}

func (k Keeper) GetTranchesByStrategy(ctx sdk.Context, strategyContract string) (list []types.TranchePool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrancheByStrategyKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte(strategyContract))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		id := sdk.BigEndianToUint64(iterator.Value())
		pool, found := k.GetTranchePool(ctx, id)
		if found {
			list = append(list, pool)
		}
	}

	return
}

// GetAllTranchePool returns all TranchePool
func (k Keeper) GetAllTranchePool(ctx sdk.Context) (list []types.TranchePool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TranchePool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetLastTrancheId(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	iterator := sdk.KVStoreReversePrefixIterator(store, []byte{})

	defer iterator.Close()
	if iterator.Valid() {
		var val types.TranchePool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		return val.Id
	}

	return 0
}
