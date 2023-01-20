package keeper

import (
	"fmt"
	"math/big"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

func (k Keeper) GetLPTokenMarketCapBreakdownAtLastRedemption(ctx sdk.Context, provider sdk.AccAddress) types.PoolMarketCap {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.AddressLPTokenMarketCapBreakdownAtTimeOfLastRedemptionKeyPrefix(provider))
	marketCap := types.PoolMarketCap{}
	k.cdc.Unmarshal(bz, &marketCap)

	return marketCap
}

func (k Keeper) SetLPTokenMarketCapBreakdownAtLastRedemption(ctx sdk.Context, provider sdk.AccAddress, marketCap types.PoolMarketCap) error {

	return nil
}

func (k Keeper) GetLiquidityProviderRewardsSinceLastRedemption(ctx sdk.Context, provider sdk.AccAddress) error {

	return nil
}
