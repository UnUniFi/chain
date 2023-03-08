package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

// GetStrategyCount get the total number of Strategy
func (k Keeper) GetStrategyCount(ctx sdk.Context, vaultDenom string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixStrategyCount(vaultDenom)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetStrategyCount set the total number of Strategy
func (k Keeper) SetStrategyCount(ctx sdk.Context, vaultDenom string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixStrategyCount(vaultDenom)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendStrategy appends a Strategy in the store with a new id and update the count
func (k Keeper) AppendStrategy(
	ctx sdk.Context,
	vaultDenom string,
	Strategy types.Strategy,
) uint64 {
	// Create the Strategy
	count := k.GetStrategyCount(ctx, vaultDenom)

	// Set the ID of the appended value
	Strategy.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	appendedValue := k.cdc.MustMarshal(&Strategy)
	store.Set(GetStrategyIDBytes(Strategy.Id), appendedValue)

	// Update Strategy count
	k.SetStrategyCount(ctx, vaultDenom, count+1)

	return count
}

// SetStrategy set a specific Strategy in the store
func (k Keeper) SetStrategy(ctx sdk.Context, vaultDenom string, Strategy types.Strategy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	b := k.cdc.MustMarshal(&Strategy)
	store.Set(GetStrategyIDBytes(Strategy.Id), b)
}

// GetStrategy returns a Strategy from its id
func (k Keeper) GetStrategy(ctx sdk.Context, vaultDenom string, id uint64) (val types.Strategy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	b := store.Get(GetStrategyIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStrategy removes a Strategy from the store
func (k Keeper) RemoveStrategy(ctx sdk.Context, vaultDenom string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	store.Delete(GetStrategyIDBytes(id))
}

// GetAllStrategy returns all Strategy
func (k Keeper) GetAllStrategy(ctx sdk.Context, vaultDenom string) (list []types.Strategy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Strategy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetStrategyIDBytes returns the byte representation of the ID
func GetStrategyIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetStrategyIDFromBytes returns ID in uint64 format from a byte array
func GetStrategyIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) StakeToStrategy(ctx sdk.Context, vaultDenom string, id uint64, amount sdk.Int) error {
	// call `stake` function of the strategy contract
	panic("not implemented")
}

func (k Keeper) UnstakeFromStrategy(ctx sdk.Context, vaultDenom string, id uint64, amount sdk.Int) error {
	// call `unstake` function of the strategy contract
	panic("not implemented")
}

func (k Keeper) GetAmountFromStrategy(ctx sdk.Context, vaultDenom string, id uint64) sdk.Int {
	// call `amount` function of the strategy contract
	panic("not implemented")
}

func (k Keeper) GetAPRFromStrategy(ctx sdk.Context, vaultDenom string, id uint64) sdk.Dec {
	// call `apr` function of the strategy contract
	panic("not implemented")
}

func (k Keeper) GetInterestFeeRate(vaultDenom string, id uint64) sdk.Dec {
	// call `interest_fee_rate` function of the strategy contract
	panic("not implemented")
}
