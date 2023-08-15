package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// MultiNftmarketHooks combine multiple nftbackedloan hooks, all hook functions are run in array sequence
type MultiNftbackedloanHooks []NftbackedloanHooks

// NewMultiNftmarketHooks returns a new MultiNftmarketHooks
func NewMultiNftbackedloanHooks(hooks ...NftbackedloanHooks) MultiNftbackedloanHooks {
	return hooks
}

// AfterNftPaymentWithCommission runs after a nft is sold and paid properly
func (h MultiNftbackedloanHooks) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier NftId, fee sdk.Coin) {
	for i := range h {
		h[i].AfterNftPaymentWithCommission(ctx, nftIdentifier, fee)
	}
}

// AfterNftUnlistedWithoutPayment runs after a nft is unlisted without any payment
func (h MultiNftbackedloanHooks) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier NftId) {
	for i := range h {
		h[i].AfterNftUnlistedWithoutPayment(ctx, nftIdentifier)
	}
}
