package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

//   GetAssetManagementTargetsOfAccount(ctx sdk.Context, accountId string)
//   AddAssetManagementTargetsOfAccount(ctx sdk.Context, account_id string, obj types.AssetManagementTarget)
//   UpdateAssetManagementTargetsOfAccount(ctx sdk.Context, targetId string, obj types.AssetManagementTarget)
//   DeleteAssetManagementTargetsOfAccount(ctx sdk.Context, targetId string)
//   GetAssetManagementTargetsOfDenom(ctx sdk.Context, accountId string, denom string)

func (k Keeper) SetAssetManagementTarget(ctx sdk.Context, obj types.AssetManagementTarget) {
	bz := k.cdc.MustMarshal(&obj)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AssetManagementTargetKey(obj.Id), bz)
}

func (k Keeper) DeleteAssetManagementTarget(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AssetManagementTargetKey(id))
}

func (k Keeper) GetAssetManagementTarget(ctx sdk.Context, id string) types.AssetManagementTarget {
	acc := types.AssetManagementTarget{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AssetManagementTargetKey(id))
	if bz == nil {
		return acc
	}
	k.cdc.MustUnmarshal(bz, &acc)
	return acc
}

func (k Keeper) GetAllAssetManagementTarget(ctx sdk.Context) []types.AssetManagementTarget {
	store := ctx.KVStore(k.storeKey)

	targets := []types.AssetManagementTarget{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.PrefixKeyAssetManagementTarget))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		target := types.AssetManagementTarget{}
		k.cdc.MustUnmarshal(it.Value(), &target)

		targets = append(targets, target)
	}
	return targets
}

// // AssetManagementAccountBankKeeper
//   PayBack(ctx sdk.Context, targetId string, farmingUnit FarmingUnit)
