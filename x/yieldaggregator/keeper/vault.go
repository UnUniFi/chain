package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// GetVaultCount get the total number of vault
func (k Keeper) GetVaultCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VaultCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVaultCount set the total number of vault
func (k Keeper) SetVaultCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VaultCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVault appends a vault in the store with a new id and update the count
func (k Keeper) AppendVault(
	ctx sdk.Context,
	vault types.Vault,
) uint64 {
	// Create the vault
	count := k.GetVaultCount(ctx)

	// Set the ID of the appended value
	vault.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	appendedValue := k.cdc.MustMarshal(&vault)
	store.Set(GetVaultIDBytes(vault.Id), appendedValue)

	// Update vault count
	k.SetVaultCount(ctx, count+1)

	return count
}

// SetVault set a specific vault in the store
func (k Keeper) SetVault(ctx sdk.Context, vault types.Vault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	b := k.cdc.MustMarshal(&vault)
	store.Set(GetVaultIDBytes(vault.Id), b)
}

// GetVault returns a vault from its id
func (k Keeper) GetVault(ctx sdk.Context, id uint64) (val types.Vault, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	b := store.Get(GetVaultIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVault removes a vault from the store
func (k Keeper) RemoveVault(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	store.Delete(GetVaultIDBytes(id))
}

// GetAllVault returns all vault
func (k Keeper) GetAllVault(ctx sdk.Context) (list []types.Vault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Vault
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetStrategyIDBytes returns the byte representation of the ID
func GetVaultIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetStrategyIDFromBytes returns ID in uint64 format from a byte array
func GetVaultIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) GetAPY(ctx sdk.Context, vaultId uint64) (*sdk.Dec, error) {
	vault, found := k.GetVault(ctx, vaultId)
	if !found {
		return nil, types.ErrInvalidVaultId
	}

	sum := sdk.ZeroDec()
	for _, weight := range vault.StrategyWeights {
		strategy, _ := k.GetStrategy(ctx, vault.Denom, weight.StrategyId)
		apr, err := k.GetAPRFromStrategy(ctx, strategy)
		if err != nil {
			return nil, err
		}
		sum = sum.Add(apr.Mul(weight.Weight))
	}

	return &sum, nil
}
