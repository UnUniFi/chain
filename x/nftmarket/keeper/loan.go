package keeper

import (
	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	// calculate maximum borrow amount for the listing
	maxDebt := k.TotalActiveRankDeposit(ctx, msg.NftId.IdBytes())

	currDebt := k.GetDebtByNft(ctx, msg.NftId.IdBytes())
	if !sdk.Coin.IsNil(currDebt.Loan) && msg.Amount.Add(currDebt.Loan).Amount.GT(maxDebt) {
		return types.ErrDebtExceedsMaxDebt
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
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	// check listing token == msg.Amount.Denom
	if listing.BidToken != msg.Amount.Denom {
		return types.ErrInvalidRepayDenom
	}

	currDebt := k.GetDebtByNft(ctx, msg.NftId.IdBytes())
	if msg.Amount.Amount.GT(currDebt.Loan.Amount) {
		return types.ErrRepayAmountExceedsLoanAmount
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

func (k Keeper) Liquidate(ctx sdk.Context, msg *types.MsgLiquidate) error {
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	if listing.State != types.ListingState_LIQUIDATION {
		return types.ErrNftListingNotInLiquidation
	}

	// TODO: handle nft sending
	// TODO: handle token flow

	// Emit event for liquidation
	ctx.EventManager().EmitTypedEvent(&types.EventLiquidate{
		Liquidator: msg.Sender.AccAddress().String(),
		ClassId:    msg.NftId.ClassId,
		NftId:      msg.NftId.NftId,
	})

	return nil
}
