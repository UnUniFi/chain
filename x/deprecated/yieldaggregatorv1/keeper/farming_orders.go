package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregatorv1/types"
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

func (k Keeper) ExecuteFarmingOrders(ctx sdk.Context, addr sdk.AccAddress, orders []types.FarmingOrder) error {
	overallRatio := uint32(0)
	for _, order := range orders {
		overallRatio = order.OverallRatio
	}

	deposit := k.GetUserDeposit(ctx, addr)
	for _, order := range orders {
		orderAlloc := sdk.Coins{}
		for _, coin := range deposit {
			orderAlloc = orderAlloc.Add(sdk.NewCoin(coin.Denom, coin.Amount.Mul(sdk.NewInt(int64(order.OverallRatio))).Quo(sdk.NewInt(int64(overallRatio)))))
		}

		// move tokens to asset management targets based on strategy
		strategy := order.Strategy
		switch strategy.StrategyType {
		case "recent30DaysHighDPRStrategy": // Invest in the best DPR destination in the last 30 days on average
			// TODO: implement individual strategy once historical info calcuator is ready
			fallthrough
		case "recent1DayHighDPRStrategy": // Invest in the best DPR destination in the last average day
			// TODO: implement individual strategy once historical info calcuator is ready
			fallthrough
		case "notHaveDPRStrategy": // Invest in something that does not have a DPR
			// TODO: implement individual strategy once historical info calcuator is ready
			fallthrough
		case "ManualStrategy": // Manual investment, whiteTargetIdlist required
			targets := k.GetAllAssetManagementTargets(ctx)
			if len(targets) == 0 {
				return types.ErrNoAssetManagementTargetExists
			}
			target := targets[0]

			cacheCtx, write := ctx.CacheContext()
			err := k.InvestOnTarget(cacheCtx, addr, target, orderAlloc)
			if err != nil {
				fmt.Println("ERROR on InvestOnTarget", err)
				return err
			} else {
				write()
			}
		}
	}

	// reduce user owned tokens since its allocated to farming units
	k.DeleteUserDeposit(ctx, addr)
	return nil
}

func (k Keeper) StopFarmingUnit(ctx sdk.Context, obj types.FarmingUnit) error {
	fmt.Println("DEBUG StopFarmingUnit", obj)

	target := k.GetAssetManagementTarget(ctx, obj.AccountId, obj.TargetId)
	addr, err := sdk.AccAddressFromBech32(obj.Owner)
	if err != nil {
		return err
	}

	cacheCtx, write := ctx.CacheContext()
	err = k.BeginWithdrawFromTarget(cacheCtx, addr, target, obj.Amount)
	if err != nil {
		fmt.Println("ERROR stopping farming unit", err)
		return err
	} else {
		fmt.Println("SUCCESS BeginWithdrawFromTarget", obj)
		write()
	}

	return nil
}

func (k Keeper) WithdrawFarmingUnit(ctx sdk.Context, obj types.FarmingUnit) error {
	addr := obj.GetAddress()
	balances := k.bankKeeper.GetAllBalances(ctx, addr)
	if balances.IsAllPositive() {
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, balances)
		if err != nil {
			return err
		}

		unitOwner, err := sdk.AccAddressFromBech32(obj.Owner)
		if err != nil {
			return err
		}
		k.IncreaseUserDeposit(ctx, unitOwner, balances)
	}

	k.DeleteFarmingUnit(ctx, obj)
	return nil
}
