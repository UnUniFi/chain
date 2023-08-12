package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

var _ types.NftbackedloanHooks = Keeper{}

func (k Keeper) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier types.NftId, fee sdk.Coin) {
	if k.hooks != nil {
		k.hooks.AfterNftPaymentWithCommission(ctx, nftIdentifier, fee)
	}
}

func (k Keeper) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier types.NftId) {
	if k.hooks != nil {
		k.hooks.AfterNftUnlistedWithoutPayment(ctx, nftIdentifier)
	}
}
