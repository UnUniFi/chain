package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// MultiNftmarketHooks combine multiple nftmarket hooks, all hook functions are run in array sequence
type MultiNftmarketHooks []NftmarketHooks

// NewMultiNftmarketHooks returns a new MultiNftmarketHooks
func NewMultiNftmarketHooks(hooks ...NftmarketHooks) MultiNftmarketHooks {
	return hooks
}

// AfterNftListed runs after a nft is listed
func (h MultiNftmarketHooks) AfterNftListed(ctx sdk.Context, nftIdentifier []byte, txMemo []byte) {
	for i := range h {
		h[i].AfterNftListed(ctx, nftIdentifier, txMemo)
	}
}

// AfterNftPaymentWithCommission runs after a nft is sold and paid properly
func (h MultiNftmarketHooks) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier []byte, fee sdk.Coin) {
	for i := range h {
		h[i].AfterNftPaymentWithCommission(ctx, nftIdentifier, fee)
	}
}

// AfterNftUnlistedWithoutPayment runs after a nft is unlisted without any payment
func (h MultiNftmarketHooks) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier []byte) {
	for i := range h {
		h[i].AfterNftUnlistedWithoutPayment(ctx, nftIdentifier)
	}
}
