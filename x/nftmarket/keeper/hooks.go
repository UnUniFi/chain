package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
)

var _ types.NftmarketHooks = Keeper{}

func (k Keeper) AfterNftListed(ctx sdk.Context, nftIdentifier []byte, txMemo string) {

}

func (k Keeper) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier []byte, fee sdk.Coin) {

}

func (k Keeper) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier []byte) {

}
