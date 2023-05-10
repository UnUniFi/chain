package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/UnUniFi/chain/x/yieldfarm/types"
)

func (k Keeper) SetFarmerInfo(ctx sdk.Context, obj types.FarmerInfo) {
	bz := k.cdc.MustMarshal(&obj)
	store := ctx.KVStore(k.storeKey)
	addr, err := sdk.AccAddressFromBech32(obj.Account)
	if err != nil {
		panic(err)
	}
	store.Set(types.FarmerInfoKey(addr), bz)
}

func (k Keeper) DeleteFarmerInfo(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.FarmerInfoKey(addr))
}

func (k Keeper) GetFarmerInfo(ctx sdk.Context, addr sdk.AccAddress) types.FarmerInfo {
	unit := types.FarmerInfo{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FarmerInfoKey(addr))
	if bz == nil {
		return unit
	}
	k.cdc.MustUnmarshal(bz, &unit)
	return unit
}

func (k Keeper) GetAllFarmerInfos(ctx sdk.Context) []types.FarmerInfo {
	store := ctx.KVStore(k.storeKey)

	units := []types.FarmerInfo{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.PrefixKeyFarmerInfo))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		unit := types.FarmerInfo{}
		k.cdc.MustUnmarshal(it.Value(), &unit)

		units = append(units, unit)
	}
	return units
}

func (k Keeper) Deposit(ctx sdk.Context, user sdk.AccAddress, coins sdk.Coins) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, coins)
	if err != nil {
		return err
	}

	deposit := k.GetFarmerInfo(ctx, user)
	deposit.Account = user.String()
	deposit.Amount = sdk.Coins(deposit.Amount).Add(coins...)
	k.SetFarmerInfo(ctx, deposit)
	return nil
}

func (k Keeper) Withdraw(ctx sdk.Context, user sdk.AccAddress, coins sdk.Coins) error {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, coins)
	if err != nil {
		return err
	}

	deposit := k.GetFarmerInfo(ctx, user)
	deposit.Amount = sdk.Coins(deposit.Amount).Sub(coins...)
	k.SetFarmerInfo(ctx, deposit)
	return nil
}

func (k Keeper) ClaimRewards(ctx sdk.Context, user sdk.AccAddress) sdk.Coins {
	deposit := k.GetFarmerInfo(ctx, user)
	amount := deposit.Rewards
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, amount)
	if err != nil {
		panic(err)
	}

	deposit.Rewards = sdk.Coins{}
	k.SetFarmerInfo(ctx, deposit)

	return amount
}

func (k Keeper) AllocateRewards(ctx sdk.Context, user sdk.AccAddress, amount sdk.Coins) {
	deposit := k.GetFarmerInfo(ctx, user)
	deposit.Rewards = sdk.Coins(deposit.Rewards).Add(amount...)
	k.SetFarmerInfo(ctx, deposit)

	err := k.bankKeeper.MintCoins(ctx, minttypes.ModuleName, amount)
	if err != nil {
		panic(err)
	}
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, types.ModuleName, amount)
	if err != nil {
		panic(err)
	}
}
