package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

var _ types.NftmarketHooks = Keeper{}

func (k Keeper) AfterNftListed(ctx sdk.Context, nftIdentifier types.NftIdentifier, txMemo string) {
	if k.hooks != nil {
		k.hooks.AfterNftListed(ctx, nftIdentifier, txMemo)
	}
}

func (k Keeper) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier types.NftIdentifier, fee sdk.Coin) {
	if k.hooks != nil {
		k.hooks.AfterNftPaymentWithCommission(ctx, nftIdentifier, fee)
	}
}

func (k Keeper) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier types.NftIdentifier) {
	if k.hooks != nil {
		k.hooks.AfterNftUnlistedWithoutPayment(ctx, nftIdentifier)
	}
}
