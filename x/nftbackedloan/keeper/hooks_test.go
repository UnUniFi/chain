package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

var statusAfterNftListed bool
var statusAfterNftPaymentWithCommission bool
var statusAfterNftUnlistedWithoutPayment bool

type dummyNftmarketHook struct{}

func (hook *dummyNftmarketHook) AfterNftListed(ctx sdk.Context, nftId types.NftIdentifier, txMemo string) {
	statusAfterNftListed = true
}

func (hook *dummyNftmarketHook) AfterNftPaymentWithCommission(ctx sdk.Context, nftId types.NftIdentifier, fee sdk.Coin) {
	statusAfterNftPaymentWithCommission = true
}

func (hook *dummyNftmarketHook) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftId types.NftIdentifier) {
	statusAfterNftUnlistedWithoutPayment = true
}
