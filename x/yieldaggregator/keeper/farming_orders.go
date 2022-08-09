package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// AssetManagementKeeper
func (k Keeper) AddFarmingOrder(ctx sdk.Context, obj types.FarmingOrder) error {
	addr, err := sdk.AccAddressFromBech32(obj.FromAddress)
	if err != nil {
		panic(err)
	}

	order := k.GetFarmingOrder(ctx, addr, obj.Id)
	if order.Id != "" {
		return types.ErrFarmingOrderAlreadyExists
	}
	k.SetFarmingOrder(ctx, obj)
	return nil
}

func (k Keeper) GetFarmingOrdersOfAddress(ctx sdk.Context, addr sdk.AccAddress) []types.FarmingOrder {
	store := ctx.KVStore(k.storeKey)

	orders := []types.FarmingOrder{}
	it := sdk.KVStorePrefixIterator(store, append([]byte(types.PrefixKeyFarmingOrder), addr...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		order := types.FarmingOrder{}
		k.cdc.MustUnmarshal(it.Value(), &order)

		orders = append(orders, order)
	}
	return orders
}

func (k Keeper) SetFarmingOrder(ctx sdk.Context, obj types.FarmingOrder) {
	bz := k.cdc.MustMarshal(&obj)
	store := ctx.KVStore(k.storeKey)
	addr, err := sdk.AccAddressFromBech32(obj.FromAddress)
	if err != nil {
		panic(err)
	}
	store.Set(types.FarmingOrderKey(addr, obj.Id), bz)
}

func (k Keeper) DeleteFarmingOrder(ctx sdk.Context, addr sdk.AccAddress, orderId string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.FarmingOrderKey(addr, orderId))
}

func (k Keeper) GetFarmingOrder(ctx sdk.Context, addr sdk.AccAddress, orderId string) types.FarmingOrder {
	order := types.FarmingOrder{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FarmingOrderKey(addr, orderId))
	if bz == nil {
		return order
	}
	k.cdc.MustUnmarshal(bz, &order)
	return order
}

func (k Keeper) GetAllFarmingOrders(ctx sdk.Context) []types.FarmingOrder {
	store := ctx.KVStore(k.storeKey)

	orders := []types.FarmingOrder{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.PrefixKeyFarmingOrder))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		order := types.FarmingOrder{}
		k.cdc.MustUnmarshal(it.Value(), &order)

		orders = append(orders, order)
	}
	return orders
}

func (k Keeper) ActivateFarmingOrder(ctx sdk.Context, addr sdk.AccAddress, farmingOrderId string) error {
	order := k.GetFarmingOrder(ctx, addr, farmingOrderId)
	if order.Id == "" {
		return types.ErrFarmingOrderDoesNotExist
	}

	order.Active = true
	k.SetFarmingOrder(ctx, order)
	return nil
}

func (k Keeper) InactivateFarmingOrder(ctx sdk.Context, addr sdk.AccAddress, farmingOrderId string) error {
	order := k.GetFarmingOrder(ctx, addr, farmingOrderId)
	if order.Id == "" {
		return types.ErrFarmingOrderDoesNotExist
	}

	order.Active = false
	k.SetFarmingOrder(ctx, order)
	return nil
}

func (k Keeper) ExecuteFarmingOrders(ctx sdk.Context, addr sdk.AccAddress) {
	// TODO: create farming units from farming orders execution
	// TODO: reduce users' owned tokens
	// TODO: allocate tokens to farming unit

	orders := k.GetFarmingOrdersOfAddress(ctx, addr)
	for _, order := range orders {
		// order.
	}

	// AssetManagementTarget
	targets := k.GetAssetManagementTargetsOfAccount(ctx, addr)
}
