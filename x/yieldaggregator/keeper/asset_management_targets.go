package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) GetAssetManagementTargetsOfAccount(ctx sdk.Context, accountId string) []types.AssetManagementTarget {
	store := ctx.KVStore(k.storeKey)

	targets := []types.AssetManagementTarget{}
	it := sdk.KVStorePrefixIterator(store, append([]byte(types.PrefixKeyAssetManagementTarget), accountId...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		target := types.AssetManagementTarget{}
		k.cdc.MustUnmarshal(it.Value(), &target)

		targets = append(targets, target)
	}
	return targets
}

func (k Keeper) GetAssetManagementTargetsOfDenom(ctx sdk.Context, accountId string, denom string) []types.AssetManagementTarget {
	targets := k.GetAssetManagementTargetsOfAccount(ctx, accountId)
	denomTargets := []types.AssetManagementTarget{}
	for _, target := range targets {
		for _, cond := range target.AssetConditions {
			if cond.Denom == denom {
				denomTargets = append(denomTargets, target)
				break
			}
		}
	}
	return denomTargets
}

func (k Keeper) DeleteAssetManagementTargetsOfAccount(ctx sdk.Context, accountId string) {
	targets := k.GetAssetManagementTargetsOfAccount(ctx, accountId)
	for _, target := range targets {
		k.DeleteAssetManagementTarget(ctx, target.AssetManagementAccountId, target.Id)
	}
}

func (k Keeper) UpdateAssetManagementTargetsOfAccount(ctx sdk.Context, accountId string, targets []types.AssetManagementTarget) {
	k.DeleteAssetManagementTargetsOfAccount(ctx, accountId)
	for _, target := range targets {
		k.SetAssetManagementTarget(ctx, target)
	}
}

func (k Keeper) SetAssetManagementTarget(ctx sdk.Context, obj types.AssetManagementTarget) {
	bz := k.cdc.MustMarshal(&obj)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AssetManagementTargetKey(obj.AssetManagementAccountId, obj.Id), bz)
}

func (k Keeper) DeleteAssetManagementTarget(ctx sdk.Context, accountId, targetId string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AssetManagementTargetKey(accountId, targetId))
}

func (k Keeper) GetAssetManagementTarget(ctx sdk.Context, accountId, targetId string) types.AssetManagementTarget {
	acc := types.AssetManagementTarget{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AssetManagementTargetKey(accountId, targetId))
	if bz == nil {
		return acc
	}
	k.cdc.MustUnmarshal(bz, &acc)
	return acc
}

func (k Keeper) GetAllAssetManagementTargets(ctx sdk.Context) []types.AssetManagementTarget {
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
