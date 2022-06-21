package nftmarket

import (
	"github.com/UnUniFi/chain/x/nftmarket/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker updates the current pricefeed
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	// process ending nft listings
	k.ProcessEndingNftListings(ctx)

	// process matured nft bids cancel
	err := k.HandleMaturedCancelledBids(ctx)
	if err != nil {
		panic(err)
	}
}
