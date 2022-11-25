package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
)

var successAfterNftListedCounter uint8
var successAfterNftPaymentWithCommissionCounter uint8
var successAfterNftUnlistedWithoutPaymentCounter uint8

type dummyNftmarketHook struct {
}

func (hook *dummyNftmarketHook) AfterNftListed(ctx sdk.Context, nftId types.NftIdentifier, txMemo string) {
	successAfterNftListedCounter += 1
}

func (hook *dummyNftmarketHook) AfterNftPaymentWithCommission(ctx sdk.Context, nftId types.NftIdentifier, fee sdk.Coin) {
	successAfterNftPaymentWithCommissionCounter += 1
}

func (hook *dummyNftmarketHook) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftId types.NftIdentifier) {
	successAfterNftUnlistedWithoutPaymentCounter += 1
}
