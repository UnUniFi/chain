package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

// SetVault set a specific vault in the store
func (k Keeper) SetVault(ctx sdk.Context, vault types.InterestRateSwapVault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	b := k.cdc.MustMarshal(&vault)
	store.Set([]byte(vault.StrategyContract), b)
}

// GetVault returns a vault from its identifier
func (k Keeper) GetVault(ctx sdk.Context, strategyContract string) (val types.InterestRateSwapVault, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	b := store.Get([]byte(strategyContract))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVault removes a vault from the store
func (k Keeper) RemoveVault(ctx sdk.Context, strategyContract string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	store.Delete([]byte(strategyContract))
}

// GetAllVault returns all vault
func (k Keeper) GetAllVault(ctx sdk.Context) (list []types.InterestRateSwapVault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VaultKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InterestRateSwapVault
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
