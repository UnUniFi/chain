package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) Borrow(ctx sdk.Context, msg *types.MsgBorrow) error {
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}
	bids := types.NftBids(k.GetBidsByNft(ctx, msg.NftId.IdBytes()))
	if len(bids) == 0 {
		return types.ErrBidDoesNotExists
	}
	// if re-borrow, repay all borrowed bids
	borrowedBid := types.NftBids{}
	for _, bid := range bids {
		if bid.Borrow.Amount.IsPositive() {
			borrowedBid = append(borrowedBid, bid)
		}
	}
	if len(borrowedBid) != 0 {
		err := k.AutoRepay(ctx, msg.NftId, borrowedBid, msg.Sender, msg.Sender)
		if err != nil {
			return err
		}
	}
	if !types.IsAbleToBorrow(bids, msg.BorrowBids, listing, ctx.BlockTime()) {
		return types.ErrCannotBorrowForLiquidation
	}
	err = k.ManualBorrow(ctx, msg.NftId, msg.BorrowBids, msg.Sender)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) ManualBorrow(ctx sdk.Context, nft types.NftIdentifier, borrows []types.BorrowBid, borrower string) error {
	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
	if err != nil {
		return err
	}
	if listing.Owner != borrower {
		return types.ErrNotNftListingOwner
	}

	borrowedAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	for _, borrow := range borrows {
		if borrow.Amount.Denom != listing.BidDenom {
			return types.ErrInvalidBorrowDenom
		}
		bidderAddress, err := sdk.AccAddressFromBech32(borrow.Bidder)
		if err != nil {
			return err
		}
		bid, err := k.GetBid(ctx, nft.IdBytes(), bidderAddress)
		if err != nil {
			return err
		}
		deposit := bid.Deposit
		if borrow.Amount.IsGTE(deposit) {
			bid.Borrow.Amount = deposit
			bid.Borrow.LastRepaidAt = ctx.BlockTime()
			borrowedAmount = borrowedAmount.Add(deposit)
		} else {
			bid.Borrow.Amount = borrow.Amount
			bid.Borrow.LastRepaidAt = ctx.BlockTime()
			borrowedAmount = borrowedAmount.Add(borrow.Amount)
		}
		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
	}

	if !borrowedAmount.IsPositive() {
		return types.ErrInvalidBorrowAmount
	}

	borrowerAddress, err := sdk.AccAddressFromBech32(borrower)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrowerAddress, sdk.Coins{borrowedAmount})
	if err != nil {
		return err
	}

	// Emit event for borrow from bids
	_ = ctx.EventManager().EmitTypedEvent(&types.EventBorrow{
		Borrower: borrower,
		ClassId:  nft.ClassId,
		NftId:    nft.NftId,
		Amount:   borrowedAmount.String(),
	})

	return nil
}

func (k Keeper) Repay(ctx sdk.Context, msg *types.MsgRepay) error {
	return k.ManualRepay(ctx, msg.NftId, msg.RepayBids, msg.Sender)
}

func (k Keeper) ManualRepay(ctx sdk.Context, nft types.NftIdentifier, repays []types.BorrowBid, borrower string) error {
	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
	if err != nil {
		return err
	}
	if listing.Owner != borrower {
		return types.ErrNotNftListingOwner
	}

	sender, err := sdk.AccAddressFromBech32(borrower)
	if err != nil {
		return err
	}

	listerAmount := k.bankKeeper.GetBalance(ctx, sender, listing.BidDenom)
	repayAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	for _, repay := range repays {
		if repay.Amount.Denom != listing.BidDenom {
			return types.ErrInvalidRepayDenom
		}
		repayAmount = repayAmount.Add(repay.Amount)
	}
	if listerAmount.Amount.LT(repayAmount.Amount) {
		return types.ErrInsufficientBalance
	}

	repaidAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())

	for _, repay := range repays {
		bidderAddress, err := sdk.AccAddressFromBech32(repay.Bidder)
		if err != nil {
			return err
		}
		bid, err := k.GetBid(ctx, nft.IdBytes(), bidderAddress)
		if err != nil {
			return err
		}

		if bid.Borrow.Amount.IsZero() {
			continue
		}

		repaidResult := bid.RepayInfo(repay.Amount, ctx.BlockTime())
		bid.Borrow.Amount = repaidResult.RemainingAmount
		bid.Borrow.LastRepaidAt = repaidResult.LastRepaidAt
		repaidAmount = repaidAmount.Add(repaidResult.RepaidAmount)

		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}

		// send interest to bidder
		err = k.SendInterestToBidder(ctx, bid, repaidResult.RepaidInterestAmount)
		if err != nil {
			return err
		}
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{repaidAmount})
	if err != nil {
		return err
	}

	// Emit event for repay
	_ = ctx.EventManager().EmitTypedEvent(&types.EventRepay{
		Borrower: borrower,
		ClassId:  nft.ClassId,
		NftId:    nft.NftId,
		Amount:   repaidAmount.String(),
	})
	return nil
}

func (k Keeper) AutoRepay(ctx sdk.Context, nft types.NftIdentifier, bids types.NftBids, borrower, receiver string) error {
	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
	if err != nil {
		return err
	}
	if listing.Owner != borrower {
		return types.ErrNotNftListingOwner
	}

	sender, err := sdk.AccAddressFromBech32(borrower)
	if err != nil {
		return err
	}

	listerAmount := k.bankKeeper.GetBalance(ctx, sender, listing.BidDenom)
	repayAmount, err := types.ExistRepayAmountAtTime(bids, listing, ctx.BlockTime())
	if err != nil {
		return err
	}
	if listerAmount.Amount.LT(repayAmount.Amount) {
		return types.ErrInsufficientBalance
	}

	repaidAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	for _, bid := range bids {
		if bid.Borrow.Amount.IsZero() {
			continue
		}

		repaidInfo := bid.RepayInfoInFull(ctx.BlockTime())
		bid.Borrow.Amount = repaidInfo.RemainingAmount
		bid.Borrow.LastRepaidAt = repaidInfo.LastRepaidAt
		repaidAmount = repaidAmount.Add(repaidInfo.RepaidAmount)

		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}

		// send interest to bidder
		err = k.SendInterestToBidder(ctx, bid, repaidInfo.RepaidInterestAmount)
		if err != nil {
			return err
		}
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{repaidAmount})
	if err != nil {
		return err
	}

	// Emit event for repay
	_ = ctx.EventManager().EmitTypedEvent(&types.EventRepay{
		Borrower: borrower,
		ClassId:  nft.ClassId,
		NftId:    nft.NftId,
		Amount:   repaidAmount.String(),
	})
	return nil
}

func (k Keeper) SendInterestToBidder(ctx sdk.Context, bid types.NftBid, interestAmount sdk.Coin) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
	if err != nil {
		return err
	}
	if interestAmount.IsNil() {
		return types.ErrInvalidInterestAmount
	}
	if interestAmount.IsPositive() {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(interestAmount.Denom, interestAmount.Amount)})
		if err != nil {
			return err
		}
	}
	return nil
}
