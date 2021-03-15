package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/cdp/types"
	"strconv"
)

// GetCdpCount get the total number of cdp
func (k Keeper) GetCdpCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpCountKey))
	byteKey := types.KeyPrefix(types.CdpCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetCdpCount set the total number of cdp
func (k Keeper) SetCdpCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpCountKey))
	byteKey := types.KeyPrefix(types.CdpCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateCdp creates a cdp with a new id and update the count
func (k Keeper) CreateCdp(ctx sdk.Context, msg types.MsgCreateCdp) {
	// Create the cdp
	count := k.GetCdpCount(ctx)
	var cdp = types.Cdp{
		Creator: msg.Creator,
		Id:      strconv.FormatInt(count, 10),
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpKey))
	key := types.KeyPrefix(types.CdpKey + cdp.Id)
	value := k.cdc.MustMarshalBinaryBare(&cdp)
	store.Set(key, value)

	// Update cdp count
	k.SetCdpCount(ctx, count+1)
}

// SetCdp set a specific cdp in the store
func (k Keeper) SetCdp(ctx sdk.Context, cdp types.Cdp) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpKey))
	b := k.cdc.MustMarshalBinaryBare(&cdp)
	store.Set(types.KeyPrefix(types.CdpKey+cdp.Id), b)
}

// GetCdp returns a cdp from its id
func (k Keeper) GetCdp(ctx sdk.Context, key string) types.Cdp {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpKey))
	var cdp types.Cdp
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.CdpKey+key)), &cdp)
	return cdp
}

// HasCdp checks if the cdp exists
func (k Keeper) HasCdp(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpKey))
	return store.Has(types.KeyPrefix(types.CdpKey + id))
}

// GetCdpOwner returns the creator of the cdp
func (k Keeper) GetCdpOwner(ctx sdk.Context, key string) string {
	return k.GetCdp(ctx, key).Creator
}

// DeleteCdp deletes a cdp
func (k Keeper) DeleteCdp(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpKey))
	store.Delete(types.KeyPrefix(types.CdpKey + key))
}

// GetAllCdp returns all cdp
func (k Keeper) GetAllCdp(ctx sdk.Context) (msgs []types.Cdp) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.CdpKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Cdp
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
