package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

// Status update to Liquidation
func (k Keeper) SetLiquidation(ctx sdk.Context, msg *types.MsgEndNftListing) error {
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

	// check if listing is already ended
	if listing.IsEnded() {
		return types.ErrStatusEndedListing
	}

	bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
	if len(bids) == 0 {

		// enable NFT transfer
		data, found := k.nftKeeper.GetNftData(ctx, msg.NftId.ClassId, msg.NftId.TokenId)
		if !found {
			return types.ErrNftDoesNotExists
		}
		data.SendDisabled = false
		err := k.nftKeeper.SetNftData(ctx, msg.NftId.ClassId, msg.NftId.TokenId, data)
		if err != nil {
			return err
		}

		k.DeleteNftListings(ctx, listing)

		// Call AfterNftUnlistedWithoutPayment to delete NFT ID from the ecosystem-incentive KVStore
		// since it's unlisted.
		// if _, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes()); err != nil {
		// 	k.AfterNftUnlistedWithoutPayment(ctx, listing.NftId)
		// }

	} else {
		params := k.GetParamSet(ctx)
		listing.State = types.ListingState_LIQUIDATION
		listing.LiquidatedAt = ctx.BlockTime()
		listing.FullPaymentEndAt = ctx.BlockTime().Add(time.Duration(params.NftListingFullPaymentPeriod) * time.Second)
		k.SaveNftListing(ctx, listing)

		// automatic payment after listing ends
		for _, bid := range bids {
			if bid.AutomaticPayment {
				bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
				if err != nil {
					fmt.Println(err)
					continue
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
					continue
				}
			}
		}
	}

	// Emit event for nft listing end
	_ = ctx.EventManager().EmitTypedEvent(&types.EventEndListing{
		Owner:   msg.Sender,
		ClassId: msg.NftId.ClassId,
		TokenId: msg.NftId.TokenId,
	})

	return nil
}

func (k Keeper) LiquidateExpiredBids(ctx sdk.Context) {
	fmt.Println("---Block time---")
	fmt.Println(ctx.BlockTime())
	bids := k.GetExpiredBids(ctx, ctx.BlockTime())
	fmt.Println("---expired bids---")
	fmt.Println(bids)
	k.DeleteBidsWithoutBorrowing(ctx, bids)
	checkListingsWithBorrowedBids := map[types.Listing][]types.Bid{}
	for _, bid := range bids {
		if !bid.IsBorrowed() {
			continue
		}

		listing, err := k.GetNftListingByIdBytes(ctx, bid.Id.NftId.IdBytes())
		if err != nil {
			fmt.Println("failed to get listing by id bytes: %w", err)
			continue
		}
		if listing.IsEnded() {
			continue
		}
		checkListingsWithBorrowedBids[listing] = append(checkListingsWithBorrowedBids[listing], bid)
	}

	for listing := range checkListingsWithBorrowedBids {
		// check if listing is already ended
		fmt.Println("---occur liquidation---")
		listingOwner, err := sdk.AccAddressFromBech32(listing.Owner)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = k.SetLiquidation(ctx, &types.MsgEndNftListing{
			Sender: listingOwner.String(),
			NftId:  listing.NftId,
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func (k Keeper) RunLiquidationProcess(ctx sdk.Context, bids types.NftBids, listing types.Listing, params types.Params) error {
	// loop to find winner bid (collect deposits + bid amount > repay amount)
	bidsSortedByDeposit := bids.SortHigherDeposit()
	winnerBid, forfeitedBids, refundBids, err := types.LiquidationBid(bidsSortedByDeposit, listing, listing.LiquidatedAt)
	if err != nil {
		return err
	}

	cacheCtx, write := ctx.CacheContext()
	if winnerBid.IsNil() {
		// No one has PayRemainder.
		err := k.LiquidateWithoutWinner(cacheCtx, bidsSortedByDeposit, listing)
		if err != nil {
			fmt.Println("failed to liquidation process with no winner: %w", err)
			return err
		}
		k.DeleteNftListings(ctx, listing)
	} else {
		// forfeitedBids, refundBids := types.ForfeitedBidsAndRefundBids(bidsSortedByDeposit, winnerBid)
		err := k.LiquidateWithWinner(cacheCtx, forfeitedBids, refundBids, listing)
		if err != nil {
			fmt.Println("failed to liquidation process with winner: %w", err)
			return err
		}
		// schedule NFT & token send after X days
		listing.SuccessfulBidEndAt = ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingNftDeliveryPeriod))
		listing.State = types.ListingState_SUCCESSFUL_BID
		k.SaveNftListing(ctx, listing)
	}
	write()
	return nil
}

// todo add test
func (k Keeper) LiquidateWithoutWinner(ctx sdk.Context, bids types.NftBids, listing types.Listing) error {
	listingOwner, err := sdk.AccAddressFromBech32(listing.Owner)
	if err != nil {
		return err
	}

	// collect deposit from all bids
	forfeitedDeposit, err := k.ForfeitDepositsFromBids(ctx, bids, listing)
	if err != nil {
		return err
	}
	listing = listing.AddCollectedAmount(forfeitedDeposit)

	borrowAmount := bids.TotalBorrowedAmount()
	// pay fee
	if listing.IsNegativeCollectedAmount() {
		return types.ErrNegativeCollectedAmount
	}
	listerProfit := listing.CollectedAmount.Amount.Sub(borrowAmount.Amount)
	if listerProfit.IsNegative() {
		return types.ErrNegativeProfit
	}
	err = k.ProcessPaymentWithCommissionFee(ctx, listingOwner, sdk.NewCoin(listing.BidDenom, listerProfit), listing.NftId)
	if err != nil {
		return err
	}

	// enable NFT transfer
	data, found := k.nftKeeper.GetNftData(ctx, listing.NftId.ClassId, listing.NftId.TokenId)
	if !found {
		return types.ErrNftDoesNotExists
	}
	data.SendDisabled = false
	err = k.nftKeeper.SetNftData(ctx, listing.NftId.ClassId, listing.NftId.TokenId, data)
	if err != nil {
		return err
	}
	return nil
}

// todo add test
func (k Keeper) LiquidateWithWinner(ctx sdk.Context, forfeitedBids, refundBids types.NftBids, listing types.Listing) error {
	forfeitedDeposit, err := k.ForfeitDepositsFromBids(ctx, forfeitedBids, listing)
	if err != nil {
		fmt.Println("failed to collect deposit from bids: %w", err)
		return err
	}
	listing = listing.AddCollectedAmount(forfeitedDeposit)
	if listing.IsNegativeCollectedAmount() {
		return types.ErrNegativeCollectedAmount
	}

	totalSubAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())

	// refund bids
	if len(refundBids) > 0 {
		refundInterestAmount := refundBids.TotalCompoundInterest(listing.LiquidatedAt)
		refundBorrowedAmount := refundBids.TotalBorrowedAmount()
		totalSubAmount = totalSubAmount.Add(refundInterestAmount).Add(refundBorrowedAmount)
	}

	// lister's profit (without winner)
	// = collected amount - (refund's interest + refund's borrowed amount)
	listing = listing.SubCollectedAmount(totalSubAmount)

	// pay interest to winner & refund to other bids
	err = k.RefundBids(ctx, refundBids, listing.LiquidatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) RefundBids(ctx sdk.Context, refundBids types.NftBids, time time.Time) error {
	for _, bid := range refundBids {
		err := k.SafeCloseBidWithAllInterest(ctx, bid, time)
		if err != nil {
			return err
		}
	}
	return nil
}

// todo add test
func (k Keeper) ForfeitDepositsFromBids(ctx sdk.Context, bids types.NftBids, listing types.Listing) (sdk.Coin, error) {
	result := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	for _, bid := range bids {
		// not pay bidder amount, collected deposit
		collectedAmount, err := k.SafeCloseBidCollectDeposit(ctx, bid)
		if err != nil {
			return result, err
		}
		if collectedAmount.IsPositive() {
			result = result.Add(collectedAmount)
		}
	}
	return result, nil
}
