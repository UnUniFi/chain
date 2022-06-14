package nftmarket

import (
	"fmt"
	"time"

	"github.com/UnUniFi/chain/x/nftmarket/keeper"
	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker updates the current pricefeed
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParamSet(ctx)
	listings := k.GetActiveNftListingsEndingAt(ctx, ctx.BlockTime())
	for _, listing := range listings {
		if listing.AutoRelistedCount < params.AutoRelistingCountIfNoBid {
			bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
			if len(bids) == 0 {
				listing.EndAt = listing.EndAt.Add(time.Duration(params.NftListingExtendSeconds) * time.Second)
				listing.AutoRelistedCount++
				k.SetNftListing(ctx, listing)
			}
		} else {
			listingOwner, err := sdk.AccAddressFromBech32(listing.Owner)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = k.EndNftListing(ctx, &types.MsgEndNftListing{
				Sender: listingOwner.Bytes(),
				NftId:  listing.NftId,
			})
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
