package nftmarket

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/keeper"
)

// EndBlocker updates the current pricefeed
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	// process ending nft listings
	k.ProcessEndingNftListings(ctx)

	// handle full payment period endings
	k.HandleFullPaymentsPeriodEndings(ctx)

	// deliver successful bids
	k.DeliverSuccessfulBids(ctx)

	// process matured nft bids cancel
	err := k.HandleMaturedCancelledBids(ctx)
	if err != nil {
		panic(err)
	}
}
