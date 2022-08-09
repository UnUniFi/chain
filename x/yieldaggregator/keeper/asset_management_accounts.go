package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// asset management keeper functions
func (k Keeper) AddAssetManagementAccount(ctx sdk.Context, id string, name string) error {
	acc := k.GetAssetManagementAccount(ctx, id)
	if acc.Id != "" {
		return types.ErrAssetManagementAccountAlreadyExists
	}
	k.SetAssetManagementAccount(ctx, types.AssetManagementAccount{
		Id:   id,
		Name: name,
	})
	return nil
}

func (k Keeper) UpdateAssetManagementAccount(ctx sdk.Context, id string, obj types.AssetManagementAccount) error {
	acc := k.GetAssetManagementAccount(ctx, id)
	if acc.Id == "" {
		return types.ErrAssetManagementAccountDoesNotExists
	}
	k.SetAssetManagementAccount(ctx, obj)
	return nil
}

func (k Keeper) DeleteAssetManagementAccount(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AssetManagementAccountKey(id))
}

func (k Keeper) SetAssetManagementAccount(ctx sdk.Context, obj types.AssetManagementAccount) {
	bz := k.cdc.MustMarshal(&obj)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AssetManagementAccountKey(obj.Id), bz)
}

func (k Keeper) GetAssetManagementAccount(ctx sdk.Context, id string) types.AssetManagementAccount {
	acc := types.AssetManagementAccount{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AssetManagementAccountKey(id))
	if bz == nil {
		return acc
	}
	k.cdc.MustUnmarshal(bz, &acc)
	return acc
}

func (k Keeper) GetAllAssetManagementAccounts(ctx sdk.Context) []types.AssetManagementAccount {
	store := ctx.KVStore(k.storeKey)

	accs := []types.AssetManagementAccount{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.PrefixKeyAssetManagementAccount))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		acc := types.AssetManagementAccount{}
		k.cdc.MustUnmarshal(it.Value(), &acc)

		accs = append(accs, acc)
	}
	return accs
}

// deposit
func (k Keeper) Deposit(ctx sdk.Context, msg *types.MsgDeposit) error {
	acc := k.GetAssetManagementAccount(ctx, msg.AssetManagementAccountId)
	if acc.Id == "" {
		return types.ErrAssetManagementAccountDoesNotExists
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.FromAddress.AccAddress(), types.ModuleName, msg.Amount)
	if err != nil {
		return err
	}

	return nil
}

// withdraw
func (k Keeper) Withdraw(ctx sdk.Context, msg *types.MsgWithdraw) error {
	acc := k.GetAssetManagementAccount(ctx, msg.AssetManagementAccountId)
	if acc.Id == "" {
		return types.ErrAssetManagementAccountDoesNotExists
	}

	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.FromAddress.AccAddress(), msg.Amount)
	if err != nil {
		return err
	}

	return nil
}
