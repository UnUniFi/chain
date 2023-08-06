package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/eventhook/types"
)

// GetHookCount get the total number of Hook
func (k Keeper) GetHookCount(ctx sdk.Context, eventType string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixHookCount(eventType)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHookCount set the total number of Hook
func (k Keeper) SetHookCount(ctx sdk.Context, eventType string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixHookCount(eventType)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendHook appends a Hook in the store with a new id and update the count
func (k Keeper) AppendHook(
	ctx sdk.Context,
	eventType string,
	Hook types.Hook,
) uint64 {
	// Create the Hook
	count := k.GetHookCount(ctx, eventType)

	// Set the ID of the appended value
	Hook.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHook(eventType))
	appendedValue := k.cdc.MustMarshal(&Hook)
	store.Set(GetHookIDBytes(Hook.Id), appendedValue)

	// Update Hook count
	k.SetHookCount(ctx, eventType, count+1)

	return count
}

// SetHook set a specific Hook in the store
func (k Keeper) SetHook(ctx sdk.Context, eventType string, Hook types.Hook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHook(eventType))
	b := k.cdc.MustMarshal(&Hook)
	store.Set(GetHookIDBytes(Hook.Id), b)
}

// GetHook returns a Hook from its id
func (k Keeper) GetHook(ctx sdk.Context, eventType string, id uint64) (val types.Hook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHook(eventType))
	b := store.Get(GetHookIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveHook removes a Hook from the store
func (k Keeper) RemoveHook(ctx sdk.Context, eventType string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHook(eventType))
	store.Delete(GetHookIDBytes(id))
}

// GetAllHook returns all Hook
func (k Keeper) GetAllHook(ctx sdk.Context, eventType string) (list []types.Hook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHook(eventType))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Hook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHookIDBytes returns the byte representation of the ID
func GetHookIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetHookIDFromBytes returns ID in uint64 format from a byte array
func GetHookIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
