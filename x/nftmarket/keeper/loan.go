package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	time "time"

	"github.com/UnUniFi/chain/x/nftmarket/types"
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
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	// check listing token == msg.Amount.Denom
	if listing.BidToken != msg.Amount.Denom {
		return types.ErrInvalidBorrowDenom
	}

	if listing.Owner != msg.Sender.AccAddress().String() {
		return types.ErrNotNftListingOwner
	}

	// calculate maximum borrow amount for the listing
	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	maxDebt := MaxPossibleBorrowAmount(bids)

	currDebt := k.GetDebtByNft(ctx, msg.NftId.IdBytes())
	if !sdk.Coin.IsNil(currDebt.Loan) && msg.Amount.Add(currDebt.Loan).Amount.GT(maxDebt) {
		return types.ErrDebtExceedsMaxDebt
	}
	// todo check is sort lower interest rate
	// todo same deposit re-borrow logic
	requireAmount := sdk.NewCoin(msg.Amount.Denom, msg.Amount.Amount)
	for _, bid := range bids {
		if requireAmount.IsZero() {
			break
		}
		// todo calc borrowed amount on bid
		usableAmount := bid.DepositAmount

		// bigger msg Amount
		if requireAmount.IsGTE(usableAmount) {
			lend := types.Borrowing{
				Amount:             sdk.NewCoin(bid.DepositAmount.Denom, bid.DepositAmount.Amount),
				StartAt:            ctx.BlockTime(),
				PaidInterestAmount: sdk.NewCoin(bid.DepositAmount.Denom, sdk.ZeroInt()),
			}
			bid.Borrowings = append(bid.Borrowings, lend)
			requireAmount.Sub(lend.Amount)
		} else {
			lend := types.Borrowing{
				Amount:             sdk.NewCoin(requireAmount.Denom, requireAmount.Amount),
				StartAt:            ctx.BlockTime(),
				PaidInterestAmount: sdk.NewCoin(requireAmount.Denom, sdk.ZeroInt()),
			}
			bid.Borrowings = append(bid.Borrowings, lend)
			requireAmount.Amount = sdk.ZeroInt()
		}
		k.SetBid(ctx, bid)
	}

	k.IncreaseDebt(ctx, msg.NftId, msg.Amount)

	sender := msg.Sender.AccAddress()
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.Coins{msg.Amount})
	if err != nil {
		return err
	}

	// Emit event for paying full bid
	ctx.EventManager().EmitTypedEvent(&types.EventBorrow{
		Borrower: msg.Sender.AccAddress().String(),
		ClassId:  msg.NftId.ClassId,
		NftId:    msg.NftId.NftId,
		Amount:   msg.Amount.String(),
	})

	return nil
}

func (k Keeper) Repay(ctx sdk.Context, msg *types.MsgRepay) error {
	// todo set interest amount in bid info
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	if listing.Owner != msg.Sender.AccAddress().String() {
		return types.ErrNotNftListingOwner
	}

	// check listing token == msg.Amount.Denom
	if listing.BidToken != msg.Amount.Denom {
		return types.ErrInvalidRepayDenom
	}

	currDebt := k.GetDebtByNft(ctx, msg.NftId.IdBytes())

	// return err if borrowing didn't happen once before
	if currDebt.Loan.IsNil() {
		return types.ErrNotBorrowed
	}

	if msg.Amount.Amount.GT(currDebt.Loan.Amount) {
		return types.ErrRepayAmountExceedsLoanAmount
	}

	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	// todo higher interest rate list
	for i := 0; i < len(bids)/2; i++ {
		bids[i], bids[len(bids)-i-1] = bids[len(bids)-i-1], bids[i]
	}

	for _, bid := range bids {
		if msg.Amount.IsZero() {
			break
		}
		if len(bid.Borrowings) == 0 {
			continue
		}

		len := []types.Borrowing{}
		for _, lend := range bid.Borrowings {
			if msg.Amount.IsZero() {
				break
			}
			principal := lend.Amount
			interest := CalcInterest(principal, bid.DepositLendingRate, lend.StartAt, ctx.BlockTime())
			interest.Sub(lend.PaidInterestAmount)
			total := sdk.NewCoin(principal.Denom, sdk.ZeroInt())
			total.Add(principal)
			total.Add(interest)
			// bigger msg Amount
			if msg.Amount.IsGTE(total) {
				bid.InterestAmount.Add(interest)
				msg.Amount.Sub(total)
				principal.Amount = sdk.ZeroInt()
			} else {
				// bigger total Amount
				if msg.Amount.IsGTE(interest) {
					// can paid interest
					if msg.Amount.Amount.GT(interest.Amount) {
						// all paid interest and part paid principal
						msg.Amount.Sub(interest)
						bid.InterestAmount.Add(interest)
						principal.Sub(msg.Amount)
						lend.PaidInterestAmount.Amount = sdk.ZeroInt()
						lend.StartAt = ctx.BlockTime()
					} else {
						// all paid interest
						bid.InterestAmount.Add(interest)
						lend.PaidInterestAmount.Add(interest)
						msg.Amount.Sub(interest)
					}
				} else {
					// can not paid interest
					bid.InterestAmount.Add(msg.Amount)
					lend.PaidInterestAmount.Add(msg.Amount)
				}
				msg.Amount.Amount = sdk.ZeroInt()
				len = append(len, lend)
			}
		}
		// clean up Borrowings
		bid.Borrowings = len
		k.SetBid(ctx, bid)
	}

	k.DecreaseDebt(ctx, msg.NftId, msg.Amount)

	sender := msg.Sender.AccAddress()
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return err
	}

	// Emit event for paying full bid
	ctx.EventManager().EmitTypedEvent(&types.EventRepay{
		Repayer: msg.Sender.AccAddress().String(),
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

// todo calc
func CalcInterest(lendCoin sdk.Coin, lendingRate string, start, end time.Time) sdk.Coin {
	interest := sdk.ZeroInt()
	return sdk.NewCoin(lendCoin.Denom, interest)
}
