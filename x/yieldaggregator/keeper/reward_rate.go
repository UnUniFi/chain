package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) SetDailyRewardPercent(ctx sdk.Context, obj types.DailyPercent) {
	bz := k.cdc.MustMarshal(&obj)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.DailyRewardKey(obj.AccountId, obj.TargetId), bz)
}

func (k Keeper) DeleteDailyRewardPercent(ctx sdk.Context, obj types.DailyPercent) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DailyRewardKey(obj.AccountId, obj.TargetId))
}

func (k Keeper) GetDailyRewardPercent(ctx sdk.Context, addr string, accId, targetId string) types.FarmingUnit {
	unit := types.FarmingUnit{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.DailyRewardKey(accId, targetId))
	if bz == nil {
		return unit
	}
	k.cdc.MustUnmarshal(bz, &unit)
	return unit
}

func (k Keeper) GetAllDailyRewardPercents(ctx sdk.Context) []types.DailyPercent {
	store := ctx.KVStore(k.storeKey)

	percents := []types.DailyPercent{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.PrefixKeyDailyPercent))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		percent := types.DailyPercent{}
		k.cdc.MustUnmarshal(it.Value(), &percent)

		percents = append(percents, percent)
	}
	return percents
}
