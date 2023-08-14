package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) SetSellingDecision(ctx sdk.Context, msg *types.MsgSellingDecision) error {
	// check listing already exists
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return types.ErrNftListingDoesNotExist
	}

	// Check nft exists
	_, found := k.nftKeeper.GetNFT(ctx, msg.NftId.ClassId, msg.NftId.TokenId)
	if !found {
		return types.ErrNftDoesNotExists
	}

	// check ownership of listing
	if listing.Owner != msg.Sender {
		return types.ErrNotNftListingOwner
	}

	// check if listing is already ended or on selling decision status
	if !listing.IsBidding() {
		return types.ErrStatusCannotSelling
	}

	// check bid exists
	bids := types.NftBids(k.GetBidsByNft(ctx, listing.NftId.IdBytes()))
	if len(bids) == 0 {
		return types.ErrBidDoesNotExists
	}

	// check no borrowing bid
	for _, bid := range bids {
		if bid.IsBorrowed() {
			return types.ErrCannotSellingBorrowedListing
		}
	}

	params, err := k.GetParams(ctx)
	if err != nil {
		return err
	}
	listing.State = types.ListingState_SELLING_DECISION
	listing.LiquidatedAt = ctx.BlockTime()
	listing.FullPaymentEndAt = ctx.BlockTime().Add(time.Duration(params.NftListingFullPaymentPeriod) * time.Second)
	k.SaveNftListing(ctx, listing)

	// automatic payment if enabled
	if len(bids) > 0 {
		highestBid, err := bids.GetHighestBid()
		if err != nil {
			return err
		}
		if highestBid.AutomaticPayment {
			bidder, err := sdk.AccAddressFromBech32(highestBid.Id.Bidder)
			if err != nil {
				return err
			}

			cacheCtx, write := ctx.CacheContext()
			err = k.PayRemainder(cacheCtx, &types.MsgPayRemainder{
				Sender: bidder.String(),
				NftId:  listing.NftId,
			})
			if err == nil {
				write()
			} else {
				fmt.Println(err)
			}
		}
	}

	// Emit event for nft listing end
	_ = ctx.EventManager().EmitTypedEvent(&types.EventSellingDecision{
		Owner:   msg.Sender,
		ClassId: msg.NftId.ClassId,
		TokenId: msg.NftId.TokenId,
	})

	return nil
}

func (k Keeper) RunSellingDecisionProcess(ctx sdk.Context, bids types.NftBids, listing types.Listing, params types.Params) error {
	highestBid, err := bids.GetHighestBid()
	if err != nil {
		return err
	}
	// if winner bidder did not pay remainder, nft is listed again after deleting winner bidder
	if !highestBid.IsPaidSalePrice() {
		borrowedAmount := highestBid.Loan.Amount
		forfeitedDeposit, err := k.SafeCloseBidCollectDeposit(ctx, highestBid)
		if err != nil {
			return err
		}
		collectedAmount := forfeitedDeposit.Sub(borrowedAmount)
		listing = listing.AddCollectedAmount(collectedAmount)

		if len(bids) == 1 {
			listing.State = types.ListingState_LISTING
		} else {
			listing.State = types.ListingState_BIDDING
		}
	} else {
		// close other bids
		otherBids := bids.RemoveBid(highestBid)
		for _, bid := range otherBids {
			err := k.SafeCloseBid(ctx, bid)
			if err != nil {
				return err
			}
		}
		// schedule NFT & token send after X days
		listing.SuccessfulBidEndAt = ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingNftDeliveryPeriod))
		listing.State = types.ListingState_SUCCESSFUL_BID
	}
	k.SaveNftListing(ctx, listing)
	return nil
}
