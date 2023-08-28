package keeper

import (
	"github.com/UnUniFi/chain/x/kyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetVerification set a specific verification in the store from its index
func (k Keeper) SetVerification(ctx sdk.Context, verification types.Verification) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKeyPrefix))
	b := k.cdc.MustMarshal(&verification)
	store.Set(types.VerificationKey(
		verification.Address,
	), b)
}

// GetVerification returns a verification from its index
func (k Keeper) GetVerification(
	ctx sdk.Context,
	index string,

) (val types.Verification, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKeyPrefix))

	b := store.Get(types.VerificationKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVerification removes a verification from the store
func (k Keeper) RemoveVerification(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKeyPrefix))
	store.Delete(types.VerificationKey(
		index,
	))
}

// GetAllVerification returns all verification
func (k Keeper) GetAllVerification(ctx sdk.Context) (list []types.Verification) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Verification
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
