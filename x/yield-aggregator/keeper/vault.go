package keeper

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
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
	// vault.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	appendedValue := k.cdc.MustMarshal(&vault)
	store.Set(GetVaultDenomBytes(vault.Denom), appendedValue)

	// Update vault count
	k.SetVaultCount(ctx, count+1)

	return count
}

// SetVault set a specific vault in the store
func (k Keeper) SetVault(ctx sdk.Context, vault types.Vault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	b := k.cdc.MustMarshal(&vault)
	store.Set(GetVaultDenomBytes(vault.Denom), b)
}

// GetVault returns a vault from its id
func (k Keeper) GetVault(ctx sdk.Context, denom string) (val types.Vault, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	b := store.Get(GetVaultDenomBytes(denom))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVault removes a vault from the store
func (k Keeper) RemoveVault(ctx sdk.Context, denom string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	store.Delete(GetVaultDenomBytes(denom))
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

// GetVaultDenomBytes returns the byte representation of the Denom
func GetVaultDenomBytes(denom string) []byte {
	return []byte(denom)
}

func (k Keeper) GetAPY(ctx sdk.Context, denom string) sdk.Dec {
	strategies := k.GetAllStrategy(ctx, denom)
	sum := sdk.ZeroDec()

	for _, strategy := range strategies {
		sum = sum.Add(strategy.Weight.Mul(strategy.Metrics.Apr))
	}

	return sum
}

func (k Keeper) DepositToVault(ctx sdk.Context, sender string, amount sdk.Coin) {
	strategies := k.GetAllStrategy(ctx, amount.Denom)

	for _, strategy := range strategies {
		allocation := strategy.Weight.Mul(sdk.NewDecFromInt(amount.Amount)).TruncateInt()
		k.StakeToStrategy(amount.Denom, strategy.Id, allocation)
	}
}

func (k Keeper) WithdrawFromVault(ctx sdk.Context, sender string, vaultDenom string, LpTokenAmount sdk.Int) {

}
