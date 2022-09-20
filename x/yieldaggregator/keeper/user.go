package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// deposit
func (k Keeper) Deposit(ctx sdk.Context, msg *types.MsgDeposit) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.FromAddress.AccAddress(), types.ModuleName, msg.Amount)
	if err != nil {
		return err
	}

	k.IncreaseUserDeposit(ctx, msg.FromAddress.AccAddress(), msg.Amount)

	return nil
}

// withdraw
func (k Keeper) Withdraw(ctx sdk.Context, msg *types.MsgWithdraw) error {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.FromAddress.AccAddress(), msg.Amount)
	if err != nil {
		return err
	}

	k.DecreaseUserDeposit(ctx, msg.FromAddress.AccAddress(), msg.Amount)
	return nil
}

func (k Keeper) GetAllUserDeposits(ctx sdk.Context) []types.UserDeposit {
	store := ctx.KVStore(k.storeKey)

	deposits := []types.UserDeposit{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.PrefixKeyUserDeposit))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		deposit := types.UserDeposit{}
		k.cdc.MustUnmarshal(it.Value(), &deposit)

		deposits = append(deposits, deposit)
	}
	return deposits
}

func (k Keeper) SetUserDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coins) {
	bz := k.cdc.MustMarshal(&types.UserDeposit{
		User:   addr.String(),
		Amount: amount,
	})
	store := ctx.KVStore(k.storeKey)
	store.Set(types.UserDepositKey(addr), bz)
}

func (k Keeper) DeleteUserDeposit(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.UserDepositKey(addr))
}

func (k Keeper) GetUserDeposit(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	deposit := types.UserDeposit{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.UserDepositKey(addr))
	if bz == nil {
		return sdk.Coins{}
	}
	k.cdc.MustUnmarshal(bz, &deposit)
	return deposit.Amount
}

func (k Keeper) IncreaseUserDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coins) {
	deposit := k.GetUserDeposit(ctx, addr)
	deposit = deposit.Add(amount...)
	k.SetUserDeposit(ctx, addr, deposit)
}

func (k Keeper) DecreaseUserDeposit(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coins) {
	deposit := k.GetUserDeposit(ctx, addr)
	deposit = deposit.Sub(amount)
	k.SetUserDeposit(ctx, addr, deposit)
}
