package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) AddFarmingUnit(ctx sdk.Context, obj types.FarmingUnit) error {
	unit := k.GetFarmingUnit(ctx, obj.Owner, obj.AccountId, obj.TargetId)
	if unit.AccountId != "" {
		return types.ErrFarmingUnitAlreadyExists
	}
	k.SetFarmingUnit(ctx, obj)
	return nil
}

func (k Keeper) GetFarmingUnitsOfAddress(ctx sdk.Context, addr sdk.AccAddress) []types.FarmingUnit {
	store := ctx.KVStore(k.storeKey)

	units := []types.FarmingUnit{}
	it := sdk.KVStorePrefixIterator(store, append([]byte(types.PrefixKeyFarmingUnit), addr.String()...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		unit := types.FarmingUnit{}
		k.cdc.MustUnmarshal(it.Value(), &unit)

		units = append(units, unit)
	}
	return units
}

func (k Keeper) SetFarmingUnit(ctx sdk.Context, obj types.FarmingUnit) {
	bz := k.cdc.MustMarshal(&obj)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.FarmingUnitKey(obj.Owner, obj.AccountId, obj.TargetId), bz)
}

func (k Keeper) DeleteFarmingUnit(ctx sdk.Context, obj types.FarmingUnit) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.FarmingUnitKey(obj.Owner, obj.AccountId, obj.TargetId))
}

func (k Keeper) GetFarmingUnit(ctx sdk.Context, addr string, accId, targetId string) types.FarmingUnit {
	unit := types.FarmingUnit{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FarmingUnitKey(addr, accId, targetId))
	if bz == nil {
		return unit
	}
	k.cdc.MustUnmarshal(bz, &unit)
	return unit
}

func (k Keeper) GetAllFarmingUnits(ctx sdk.Context) []types.FarmingUnit {
	store := ctx.KVStore(k.storeKey)

	units := []types.FarmingUnit{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.PrefixKeyFarmingUnit))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		unit := types.FarmingUnit{}
		k.cdc.MustUnmarshal(it.Value(), &unit)

		units = append(units, unit)
	}
	return units
}
