package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) Borrow(ctx sdk.Context, msg *types.MsgBorrow) error {
	bids := types.NftBids(k.GetBidsByNft(ctx, msg.NftId.IdBytes()))
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
	if !types.IsAbleToBorrow(bids, msg.BorrowBids, ctx.BlockTime()) {
		return types.ErrCannotBorrow
	}
	err := k.ManualBorrow(ctx, msg.NftId, msg.BorrowBids, msg.Sender)
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
		depositAmount := bid.DepositAmount
		if borrow.Amount.IsGTE(depositAmount) {
			bid.Borrow.Amount = depositAmount
			bid.Borrow.LastRepaidAt = ctx.BlockTime()
			borrowedAmount = borrowedAmount.Add(depositAmount)
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
	return k.ManualRepay(ctx, msg.NftId, msg.RepayBids, msg.Sender, msg.Sender)
}

func (k Keeper) ManualRepay(ctx sdk.Context, nft types.NftIdentifier, repays []types.BorrowBid, borrower, receiver string) error {
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

		repaidResult := bid.RepaidResult(repay.Amount, ctx.BlockTime())
		bid.Borrow.Amount = repaidResult.RemainingBorrowAmount
		bid.Borrow.LastRepaidAt = repaidResult.LastRepaidAt
		repaidAmount = repaidAmount.Add(repaidResult.RepaidAmount)

		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
	}

	// todo: pay interest to bidder
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{repaidAmount})
	if err != nil {
		return err
	}

	// Emit event for repay
	_ = ctx.EventManager().EmitTypedEvent(&types.EventRepay{
		Repayer: borrower,
		ClassId: nft.ClassId,
		NftId:   nft.NftId,
		Amount:  repaidAmount.String(),
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
	repayAmount, err := types.ExistRepayAmountAtTime(bids, ctx.BlockTime())
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

		repaidResult := bid.FullRepaidResult(ctx.BlockTime())
		bid.Borrow.Amount = repaidResult.RemainingBorrowAmount
		bid.Borrow.LastRepaidAt = repaidResult.LastRepaidAt
		repaidAmount = repaidAmount.Add(repaidResult.RepaidAmount)

		err = k.SetBid(ctx, bid)
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
		Repayer: borrower,
		ClassId: nft.ClassId,
		NftId:   nft.NftId,
		Amount:  repaidAmount.String(),
	})
	return nil
}
