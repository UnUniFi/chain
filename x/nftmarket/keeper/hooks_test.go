package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
)

type dummyNftmarketHook struct {
	successCounter int
}

func (hook *dummyNftmarketHook) AfterNftListed(ctx sdk.Context, nftId types.NftIdentifier, txMemo string) {
	hook.successCounter += 1
}

func (hook *dummyNftmarketHook) AfterNftPaymentWithCommission(ctx sdk.Context, nftId types.NftIdentifier, fee sdk.Coin) {
	hook.successCounter += 1
}

func (hook *dummyNftmarketHook) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftId types.NftIdentifier) {
	hook.successCounter += 1
}
