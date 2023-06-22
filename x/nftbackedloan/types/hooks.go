package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// MultiNftmarketHooks combine multiple nftmarket hooks, all hook functions are run in array sequence
type MultiNftbackedloanHooks []NftbackedloanHooks

// NewMultiNftmarketHooks returns a new MultiNftmarketHooks
func NewMultiNftbackedloanHooks(hooks ...NftbackedloanHooks) MultiNftbackedloanHooks {
	return hooks
}

// AfterNftListed runs after a nft is listed
func (h MultiNftbackedloanHooks) AfterNftListed(ctx sdk.Context, nftIdentifier NftIdentifier, txMemo string) {
	for i := range h {
		h[i].AfterNftListed(ctx, nftIdentifier, txMemo)
	}
}

// AfterNftPaymentWithCommission runs after a nft is sold and paid properly
func (h MultiNftbackedloanHooks) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier NftIdentifier, fee sdk.Coin) {
	for i := range h {
		h[i].AfterNftPaymentWithCommission(ctx, nftIdentifier, fee)
	}
}

// AfterNftUnlistedWithoutPayment runs after a nft is unlisted without any payment
func (h MultiNftbackedloanHooks) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier NftIdentifier) {
	for i := range h {
		h[i].AfterNftUnlistedWithoutPayment(ctx, nftIdentifier)
	}
}
