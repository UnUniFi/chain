package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) GetIntermediaryAccountInfo(ctx sdk.Context) types.IntermediaryAccountInfo {
	store := ctx.KVStore(k.storeKey)
	byteKey := types.KeyPrefix(types.ChainReceiverKey)
	bz := store.Get(byteKey)

	info := types.IntermediaryAccountInfo{}
	if bz == nil {
		return info
	}

	k.cdc.MustUnmarshal(bz, &info)
	return info
}

func (k Keeper) SetIntermediaryAccountInfo(ctx sdk.Context, addrs []types.ChainAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.IntermediaryAccountInfo{
		Addrs: addrs,
	})
	store.Set(types.KeyPrefix(types.ChainReceiverKey), bz)
}

func (k Keeper) GetIntermediaryReceiver(ctx sdk.Context, chainId string) string {
	info := k.GetIntermediaryAccountInfo(ctx)
	for _, ca := range info.Addrs {
		if ca.ChainId == chainId {
			return ca.Address
		}
	}
	return ""
}
