package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
)

func (k Keeper) GetKeyDataForPoC2(ctx sdk.Context, nftIdBytes []byte, startedAt time.Time) types.KeyDataForPoC2 {
	keyDataForPoC2 := types.KeyDataForPoC2{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyDataForPoC2Key(nftIdBytes, startedAt))
	if bz == nil {
		return keyDataForPoC2
	}

	k.cdc.MustUnmarshal(bz, &keyDataForPoC2)
	return keyDataForPoC2
}

func (k Keeper) GetAllKeyDataForPoC2(ctx sdk.Context) []types.KeyDataForPoC2 {
	store := ctx.KVStore(k.storeKey)

	keyDataForPoC2s := []types.KeyDataForPoC2{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixKeyDataForPoC2))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		keyDataForPoC2 := types.KeyDataForPoC2{}
		k.cdc.MustUnmarshal(it.Value(), &keyDataForPoC2)

		keyDataForPoC2s = append(keyDataForPoC2s, keyDataForPoC2)
	}
	return keyDataForPoC2s
}

func (k Keeper) SetKeyDataForPoC2(ctx sdk.Context, keyData types.KeyDataForPoC2) {
	bz := k.cdc.MustMarshal(&keyData)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyDataForPoC2Key(keyData.NftId.IdBytes(), keyData.StartedAt), bz)
}

func (k Keeper) DeleteKeyDataForPoC2(ctx sdk.Context, nftIdBytes []byte, started_time time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyDataForPoC2Key(nftIdBytes, started_time))
}
