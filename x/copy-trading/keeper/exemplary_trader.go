package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/copy-trading/types"
)

// SetExemplaryTrader set a specific exemplaryTrader in the store from its index
func (k Keeper) SetExemplaryTrader(ctx sdk.Context, exemplaryTrader types.ExemplaryTrader) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExemplaryTraderKeyPrefix))
	b := k.cdc.MustMarshal(&exemplaryTrader)
	store.Set(types.ExemplaryTraderKey(
		exemplaryTrader.Address,
	), b)
}

// GetExemplaryTrader returns a exemplaryTrader from its index
func (k Keeper) GetExemplaryTrader(
	ctx sdk.Context,
	index string,

) (val types.ExemplaryTrader, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExemplaryTraderKeyPrefix))

	b := store.Get(types.ExemplaryTraderKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExemplaryTrader removes a exemplaryTrader from the store
func (k Keeper) RemoveExemplaryTrader(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExemplaryTraderKeyPrefix))
	store.Delete(types.ExemplaryTraderKey(
		index,
	))
}

// GetAllExemplaryTrader returns all exemplaryTrader
func (k Keeper) GetAllExemplaryTrader(ctx sdk.Context) (list []types.ExemplaryTrader) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExemplaryTraderKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExemplaryTrader
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
