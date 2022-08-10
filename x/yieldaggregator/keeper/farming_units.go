package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) AddFarmingUnit(ctx sdk.Context, obj types.FarmingUnit) error {
	addr, err := sdk.AccAddressFromBech32(obj.Owner)
	if err != nil {
		panic(err)
	}

	unit := k.GetFarmingUnit(ctx, addr, obj.Id)
	if unit.Id != "" {
		return types.ErrFarmingUnitAlreadyExists
	}
	k.SetFarmingUnit(ctx, obj)
	return nil
}

func (k Keeper) StopFarmingUnit(ctx sdk.Context, obj types.FarmingUnit) error {
	// TODO: this should perform action to yield farm target
	return nil
}

func (k Keeper) GetFarmingUnitsOfAddress(ctx sdk.Context, addr sdk.AccAddress) []types.FarmingUnit {
	store := ctx.KVStore(k.storeKey)

	units := []types.FarmingUnit{}
	it := sdk.KVStorePrefixIterator(store, append([]byte(types.PrefixKeyFarmingUnit), addr...))
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
	addr, err := sdk.AccAddressFromBech32(obj.Owner)
	if err != nil {
		panic(err)
	}
	store.Set(types.FarmingUnitKey(addr, obj.Id), bz)
}

func (k Keeper) DeleteFarmingUnit(ctx sdk.Context, addr sdk.AccAddress, unitId string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.FarmingUnitKey(addr, unitId))
}

func (k Keeper) GetFarmingUnit(ctx sdk.Context, addr sdk.AccAddress, unitId string) types.FarmingUnit {
	unit := types.FarmingUnit{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FarmingUnitKey(addr, unitId))
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
