package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) GetDebtByNft(ctx sdk.Context, nftIdBytes []byte) types.Loan {
	loan := types.Loan{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NftLoanKey(nftIdBytes))
	if bz == nil {
		return loan
	}

	k.cdc.MustUnmarshal(bz, &loan)
	return loan
}

func (k Keeper) GetAllDebts(ctx sdk.Context) []types.Loan {
	store := ctx.KVStore(k.storeKey)

	loans := []types.Loan{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixNftLoan))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		loan := types.Loan{}
		k.cdc.MustUnmarshal(it.Value(), &loan)

		loans = append(loans, loan)
	}
	return loans
}

func (k Keeper) SetDebt(ctx sdk.Context, loan types.Loan) {
	bz := k.cdc.MustMarshal(&loan)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NftLoanKey(loan.NftId.IdBytes()), bz)
}

func (k Keeper) DeleteDebt(ctx sdk.Context, nftBytes []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NftLoanKey(nftBytes))
}

// remove debt (loan) from KVStore by using DeleteDebt method with the feature
// to judge if it exists before calling it
func (k Keeper) RemoveDebt(ctx sdk.Context, nftBytes []byte) {
	loan := k.GetDebtByNft(ctx, nftBytes)
	if !loan.Loan.Amount.IsNil() {
		k.DeleteDebt(ctx, nftBytes)
	}
}

func (k Keeper) IncreaseDebt(ctx sdk.Context, nftId types.NftIdentifier, amount sdk.Coin) {
	currDebt := k.GetDebtByNft(ctx, nftId.IdBytes())
	if sdk.Coin.IsNil(currDebt.Loan) {
		currDebt.NftId = nftId
		currDebt.Loan = amount
	} else {
		currDebt.Loan = currDebt.Loan.Add(amount)
	}
	k.SetDebt(ctx, currDebt)
}

func (k Keeper) DecreaseDebt(ctx sdk.Context, nftId types.NftIdentifier, amount sdk.Coin) {
	currDebt := k.GetDebtByNft(ctx, nftId.IdBytes())
	currDebt.Loan = currDebt.Loan.Sub(amount)
	k.SetDebt(ctx, currDebt)
}

func (k Keeper) Borrow(ctx sdk.Context, msg *types.MsgBorrow) error {
	bids := types.NftBids(k.GetBidsByNft(ctx, msg.NftId.IdBytes()))
	// todo impl re-borrow
	for _, bid := range bids {
		if len(bid.Borrowings) != 0 {
			return types.ErrAlreadyBorrowed
		}
	}
	if types.IsAbleToBorrow(bids, msg.BorrowBids, ctx.BlockTime()) {
		return k.ManualBorrow(ctx, msg.NftId, msg.BorrowBids, msg.Sender, msg.Sender)
	} else {
		return types.ErrCannotBorrow
	}
}

func (k Keeper) ManualBorrow(ctx sdk.Context, nft types.NftIdentifier, borrows []types.BorrowBid, borrower, receiver string) error {
	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
	if err != nil {
		return err
	}
	if listing.Owner != borrower {
		return types.ErrNotNftListingOwner
	}

	borrowedAmount := sdk.NewCoin(listing.BidToken, sdk.ZeroInt())
	for _, borrow := range borrows {
		bidderAddress, err := sdk.AccAddressFromBech32(borrow.Bidder)
		if err != nil {
			return err
		}
		bid, err := k.GetBid(ctx, nft.IdBytes(), bidderAddress)
		if err != nil {
			return err
		}
		usableAmount := bid.BorrowableAmount()
		if borrow.Amount.IsGTE(usableAmount) {
			borrowing := types.Borrowing{
				Amount:             usableAmount,
				StartAt:            ctx.BlockTime(),
				PaidInterestAmount: sdk.NewCoin(borrow.Amount.Denom, sdk.ZeroInt()),
			}
			borrowedAmount = borrowedAmount.Add(usableAmount)
			bid.Borrowings = append(bid.Borrowings, borrowing)
		} else {
			borrowing := types.Borrowing{
				Amount:             borrow.Amount,
				StartAt:            ctx.BlockTime(),
				PaidInterestAmount: sdk.NewCoin(borrow.Amount.Denom, sdk.ZeroInt()),
			}
			borrowedAmount = borrowedAmount.Add(borrow.Amount)
			bid.Borrowings = append(bid.Borrowings, borrowing)
		}
		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
	}

	k.IncreaseDebt(ctx, nft, borrowedAmount)

	receiverAddress, err := sdk.AccAddressFromBech32(receiver)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddress, sdk.Coins{borrowedAmount})
	if err != nil {
		return err
	}

	// Emit event for borrow from bids
	ctx.EventManager().EmitTypedEvent(&types.EventBorrow{
		Borrower: borrower,
		ClassId:  nft.ClassId,
		NftId:    nft.NftId,
		Amount:   borrowedAmount.String(),
	})

	return nil
}

// func (k Keeper) AutoBorrow(ctx sdk.Context, nft types.NftIdentifier, require sdk.Coin, borrower, receiver string) error {
// 	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
// 	if err != nil {
// 		return err
// 	}

// 	// check listing token == msg.Amount.Denom
// 	if listing.BidToken != require.Denom {
// 		return types.ErrInvalidBorrowDenom
// 	}

// 	if listing.Owner != borrower {
// 		return types.ErrNotNftListingOwner
// 	}

// 	// calculate maximum borrow amount for the listing
// 	// bids := k.GetBidsByNft(ctx, nft.IdBytes())

// 	bids := types.NftBids(k.GetBidsByNft(ctx, nft.IdBytes()))

// 	maxDebt := listing.MaxPossibleBorrowAmount(bids, []types.NftBid{})

// 	if require.Amount.GT(maxDebt) {
// 		return types.ErrDebtExceedsMaxDebt
// 	}
// 	// todo same deposit re-borrow logic
// 	requireAmount := sdk.NewCoin(require.Denom, require.Amount)
// 	borrowingOrderBids := bids.SortBorrowing()
// 	// todo use BorrowFromBids
// 	for _, bid := range borrowingOrderBids {
// 		if requireAmount.IsZero() {
// 			break
// 		}

// 		usableAmount := bid.BorrowableAmount()
// 		// bigger msg Amount
// 		if requireAmount.IsGTE(usableAmount) {
// 			borrow := types.Borrowing{
// 				Amount:             sdk.NewCoin(usableAmount.Denom, usableAmount.Amount),
// 				StartAt:            ctx.BlockTime(),
// 				PaidInterestAmount: sdk.NewCoin(usableAmount.Denom, sdk.ZeroInt()),
// 			}
// 			bid.Borrowings = append(bid.Borrowings, borrow)
// 			requireAmount = requireAmount.Sub(borrow.Amount)
// 		} else {
// 			borrow := types.Borrowing{
// 				Amount:             sdk.NewCoin(requireAmount.Denom, requireAmount.Amount),
// 				StartAt:            ctx.BlockTime(),
// 				PaidInterestAmount: sdk.NewCoin(requireAmount.Denom, sdk.ZeroInt()),
// 			}
// 			bid.Borrowings = append(bid.Borrowings, borrow)
// 			requireAmount.Amount = sdk.ZeroInt()
// 		}
// 		err = k.SetBid(ctx, bid)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	k.IncreaseDebt(ctx, nft, require)

// 	receiverAddress, err := sdk.AccAddressFromBech32(receiver)
// 	if err != nil {
// 		return err
// 	}
// 	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddress, sdk.Coins{require})
// 	if err != nil {
// 		return err
// 	}

// 	// Emit event for paying full bid
// 	ctx.EventManager().EmitTypedEvent(&types.EventBorrow{
// 		Borrower: borrower,
// 		ClassId:  nft.ClassId,
// 		NftId:    nft.NftId,
// 		Amount:   require.String(),
// 	})

// 	return nil
// }

// func (k Keeper) Refinancings(ctx sdk.Context, listing types.NftListing, liquidationBids []types.NftBid) {
// 	for _, v := range liquidationBids {
// 		err := k.Refinancing(ctx, listing, v)
// 		if err != nil {
// 			fmt.Println("Refinancing error: %w", err)
// 			continue
// 		}
// 	}
// }

// func (k Keeper) Refinancing(ctx sdk.Context, listing types.NftListing, bid types.NftBid) error {
// 	err := k.DeleteBid(ctx, bid)
// 	if err != nil {
// 		return err
// 	}
// 	// todo delete not depend on Debt
// 	k.DecreaseDebt(ctx, listing.NftId, bid.BorrowingAmount())
// 	liquidationAmount := bid.LiquidationAmount(ctx.BlockTime())
// 	err = k.AutoBorrow(ctx, listing.NftId, liquidationAmount, listing.Owner, bid.Id.Bidder)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (k Keeper) Repay(ctx sdk.Context, msg *types.MsgRepay) error {
	return k.ManualRepay(ctx, msg.NftId, msg.RepayBids, msg.Sender, msg.Sender)
}

func (k Keeper) ManualRepay(ctx sdk.Context, nft types.NftIdentifier, repays []types.BorrowBid, borrower, receiver string) error {
	// todo set interest amount in bid info
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

	listerAmount := k.bankKeeper.GetBalance(ctx, sender, listing.BidToken)
	repayAmount := sdk.NewCoin(listing.BidToken, sdk.ZeroInt())
	for _, repay := range repays {
		repayAmount = repayAmount.Add(repay.Amount)
	}
	if listerAmount.Amount.LT(repayAmount.Amount) {
		return types.ErrInsufficientBalance
	}

	currDebt := k.GetDebtByNft(ctx, nft.IdBytes())

	// return err if borrowing didn't happen once before
	if currDebt.Loan.IsNil() {
		return types.ErrNotBorrowed
	}

	repaidAmount := sdk.NewCoin(listing.BidToken, sdk.ZeroInt())
	for _, repay := range repays {
		bidderAddress, err := sdk.AccAddressFromBech32(repay.Bidder)
		if err != nil {
			return err
		}
		bid, err := k.GetBid(ctx, nft.IdBytes(), bidderAddress)
		if err != nil {
			return err
		}

		if len(bid.Borrowings) == 0 {
			continue
		}
		repaidAmount = repaidAmount.Add(repay.Amount)
		borrowings := []types.Borrowing{}
		for _, borrowing := range bid.Borrowings {
			if repay.Amount.IsZero() {
				break
			}
			receipt := borrowing.RepayThenGetReceipt(repay.Amount, ctx.BlockTime(), bid.CalcInterestF())
			repay.Amount = receipt.Charge
			bid.InterestAmount = bid.InterestAmount.Add(receipt.PaidInterestAmount)

			if !borrowing.IsAllRepaid() {
				borrowings = append(borrowings, borrowing)
			}
		}
		// clean up Borrowings
		bid.Borrowings = borrowings
		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
		// Subtract when surplus repay amount exists
		repaidAmount = repaidAmount.Sub(repay.Amount)
	}

	k.DecreaseDebt(ctx, nft, repaidAmount)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{repaidAmount})
	if err != nil {
		return err
	}

	// Emit event for paying full bid
	ctx.EventManager().EmitTypedEvent(&types.EventRepay{
		Repayer: borrower,
		ClassId: nft.ClassId,
		NftId:   nft.NftId,
		Amount:  repaidAmount.String(),
	})
	return nil
}

// func (k Keeper) AutoRepay(ctx sdk.Context, nft types.NftIdentifier, require sdk.Coin, borrower, receiver string) error {
// 	// todo set interest amount in bid info
// 	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
// 	if err != nil {
// 		return err
// 	}

// 	if listing.Owner != borrower {
// 		return types.ErrNotNftListingOwner
// 	}

// 	// check listing token == msg.Amount.Denom
// 	if listing.BidToken != require.Denom {
// 		return types.ErrInvalidBorrowDenom
// 	}

// 	sender, err := sdk.AccAddressFromBech32(borrower)
// 	if err != nil {
// 		return err
// 	}

// 	listerAmount := k.bankKeeper.GetBalance(ctx, sender, require.Denom)
// 	if listerAmount.Amount.LT(require.Amount) {
// 		return types.ErrInsufficientBalance
// 	}

// 	currDebt := k.GetDebtByNft(ctx, nft.IdBytes())

// 	// return err if borrowing didn't happen once before
// 	if currDebt.Loan.IsNil() {
// 		return types.ErrNotBorrowed
// 	}

// 	bids := types.NftBids(k.GetBidsByNft(ctx, nft.IdBytes())).SortRepay()
// 	repaidAmount := require
// 	for _, bid := range bids {
// 		if repaidAmount.IsZero() {
// 			break
// 		}
// 		if len(bid.Borrowings) == 0 {
// 			continue
// 		}

// 		borrowings := []types.Borrowing{}
// 		for _, borrow := range bid.Borrowings {
// 			if repaidAmount.IsZero() {
// 				break
// 			}

// 			receipt := borrow.RepayThenGetReceipt(repaidAmount, ctx.BlockTime(), bid.CalcInterestF())
// 			repaidAmount.Amount = receipt.Charge.Amount
// 			bid.InterestAmount = bid.InterestAmount.Add(receipt.PaidInterestAmount)

// 			if !borrow.IsAllRepaid() {
// 				borrowings = append(borrowings, borrow)
// 			}
// 		}
// 		// clean up Borrowings
// 		bid.Borrowings = borrowings
// 		err = k.SetBid(ctx, bid)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	debitAmount := require.Sub(repaidAmount)
// 	k.DecreaseDebt(ctx, nft, debitAmount)
// 	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{debitAmount})
// 	if err != nil {
// 		return err
// 	}

// 	// Emit event for paying full bid
// 	ctx.EventManager().EmitTypedEvent(&types.EventRepay{
// 		Repayer: borrower,
// 		ClassId: nft.ClassId,
// 		NftId:   nft.NftId,
// 		Amount:  require.String(),
// 	})

// 	return nil
// }

func MaxPossibleBorrowAmount(bids []types.NftBid) sdk.Int {
	maxPossibleBorrowAmount := sdk.ZeroInt()
	for _, bid := range bids {
		maxPossibleBorrowAmount = maxPossibleBorrowAmount.Add(bid.DepositAmount.Amount)
	}
	return maxPossibleBorrowAmount
}
