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
	return k.ManualBorrow(ctx, msg.NftId, msg.Amount, msg.Sender, msg.Sender)
}

func (k Keeper) ManualBorrow(ctx sdk.Context, nft types.NftIdentifier, require sdk.Coin, borrower, receiver string) error {
	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
	if err != nil {
		return err
	}

	// check listing token == msg.Amount.Denom
	if listing.BidToken != require.Denom {
		return types.ErrInvalidBorrowDenom
	}

	if listing.Owner != borrower {
		return types.ErrNotNftListingOwner
	}

	// calculate maximum borrow amount for the listing
	// bids := k.GetBidsByNft(ctx, nft.IdBytes())

	bids := types.NftBids(k.GetBidsByNft(ctx, nft.IdBytes()))

	maxDebt := listing.MaxPossibleBorrowAmount(bids, []types.NftBid{})

	if require.Amount.GT(maxDebt) {
		return types.ErrDebtExceedsMaxDebt
	}
	// todo same deposit re-borrow logic
	requireAmount := sdk.NewCoin(require.Denom, require.Amount)
	borrowingOrderBids := bids.SortBorrowing()
	// todo use BorrowFromBids
	for _, bid := range borrowingOrderBids {
		if requireAmount.IsZero() {
			break
		}

		usableAmount := bid.BorrowableAmount()
		// bigger msg Amount
		if requireAmount.IsGTE(usableAmount) {
			borrow := types.Borrowing{
				Amount:             sdk.NewCoin(usableAmount.Denom, usableAmount.Amount),
				StartAt:            ctx.BlockTime(),
				PaidInterestAmount: sdk.NewCoin(usableAmount.Denom, sdk.ZeroInt()),
			}
			bid.Borrowings = append(bid.Borrowings, borrow)
			requireAmount = requireAmount.Sub(borrow.Amount)
		} else {
			borrow := types.Borrowing{
				Amount:             sdk.NewCoin(requireAmount.Denom, requireAmount.Amount),
				StartAt:            ctx.BlockTime(),
				PaidInterestAmount: sdk.NewCoin(requireAmount.Denom, sdk.ZeroInt()),
			}
			bid.Borrowings = append(bid.Borrowings, borrow)
			requireAmount.Amount = sdk.ZeroInt()
		}
		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
	}

	k.IncreaseDebt(ctx, nft, require)

	receiverAddress, err := sdk.AccAddressFromBech32(receiver)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddress, sdk.Coins{require})
	if err != nil {
		return err
	}

	// Emit event for paying full bid
	ctx.EventManager().EmitTypedEvent(&types.EventBorrow{
		Borrower: borrower,
		ClassId:  nft.ClassId,
		NftId:    nft.NftId,
		Amount:   require.String(),
	})

	return nil
}

func (k Keeper) Refinancings(ctx sdk.Context, listing types.NftListing, liquidationBids []types.NftBid) {
	for _, v := range liquidationBids {
		k.Refinancing(ctx, listing, v)
	}
}

func (k Keeper) Refinancing(ctx sdk.Context, listing types.NftListing, bid types.NftBid) {
	err := k.DeleteBid(ctx, bid)
	if err != nil {
		panic(err)
	}
	// todo delete not depend on Debt
	k.DecreaseDebt(ctx, listing.NftId, bid.BorrowingAmount())
	liquidationAmount := bid.LiquidationAmount(ctx.BlockTime())
	err = k.ManualBorrow(ctx, listing.NftId, liquidationAmount, listing.Owner, bid.Bidder)
	if err != nil {
		panic(err)
	}
}

func (k Keeper) Repay(ctx sdk.Context, msg *types.MsgRepay) error {
	// todo set interest amount in bid info
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	if listing.Owner != msg.Sender {
		return types.ErrNotNftListingOwner
	}

	// check listing token == msg.Amount.Denom
	if listing.BidToken != msg.Amount.Denom {
		return types.ErrInvalidRepayDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	listerAmount := k.bankKeeper.GetBalance(ctx, sender, msg.Amount.Denom)
	if listerAmount.Amount.LT(msg.Amount.Amount) {
		return types.ErrInsufficientBalance
	}

	currDebt := k.GetDebtByNft(ctx, msg.NftId.IdBytes())

	// return err if borrowing didn't happen once before
	if currDebt.Loan.IsNil() {
		return types.ErrNotBorrowed
	}

	bids := types.NftBids(k.GetBidsByNft(ctx, msg.NftId.IdBytes())).SortRepay()
	repaidAmount := sdk.NewCoin(msg.Amount.Denom, msg.Amount.Amount)
	for _, bid := range bids {
		if repaidAmount.IsZero() {
			break
		}
		if len(bid.Borrowings) == 0 {
			continue
		}

		borrowings := []types.Borrowing{}
		for _, borrow := range bid.Borrowings {
			if repaidAmount.IsZero() {
				break
			}

			receipt := borrow.RepayThenGetReceipt(repaidAmount, ctx.BlockTime(), bid.CalcInterestF())
			repaidAmount.Amount = receipt.Charge.Amount
			bid.InterestAmount = bid.InterestAmount.Add(receipt.PaidInterestAmount)

			if !borrow.IsAllRepaid() {
				borrowings = append(borrowings, borrow)
			}
		}
		// clean up Borrowings
		bid.Borrowings = borrowings
		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
	}

	debitAmount := msg.Amount.Sub(repaidAmount)
	k.DecreaseDebt(ctx, msg.NftId, debitAmount)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{debitAmount})
	if err != nil {
		return err
	}

	// Emit event for paying full bid
	ctx.EventManager().EmitTypedEvent(&types.EventRepay{
		Repayer: msg.Sender,
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
		Amount:  msg.Amount.String(),
	})

	return nil
}

func MaxPossibleBorrowAmount(bids []types.NftBid) sdk.Int {
	maxPossibleBorrowAmount := sdk.ZeroInt()
	for _, bid := range bids {
		maxPossibleBorrowAmount = maxPossibleBorrowAmount.Add(bid.DepositAmount.Amount)
	}
	return maxPossibleBorrowAmount
}
