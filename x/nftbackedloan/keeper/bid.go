package keeper

import (
	"fmt"
	"time"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) GetBid(ctx sdk.Context, nftIdBytes []byte, bidder sdk.AccAddress) (types.NftBid, error) {
	bid := types.NftBid{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NftBidKey(nftIdBytes, bidder))
	if bz == nil {
		return bid, types.ErrBidDoesNotExists
	}

	k.cdc.MustUnmarshal(bz, &bid)
	return bid, nil
}

func (k Keeper) GetAllBids(ctx sdk.Context) []types.NftBid {
	store := ctx.KVStore(k.storeKey)

	bids := []types.NftBid{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixNftBid))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		bid := types.NftBid{}
		k.cdc.MustUnmarshal(it.Value(), &bid)

		bids = append(bids, bid)
	}
	return bids
}

func (k Keeper) GetBidsByNft(ctx sdk.Context, nftIdBytes []byte) []types.NftBid {
	store := ctx.KVStore(k.storeKey)

	bids := []types.NftBid{}
	it := sdk.KVStorePrefixIterator(store, append([]byte(types.KeyPrefixNftBid), nftIdBytes...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		bid := types.NftBid{}
		k.cdc.MustUnmarshal(it.Value(), &bid)

		bids = append(bids, bid)
	}
	return bids
}

func (k Keeper) GetBidsByBidder(ctx sdk.Context, bidder sdk.AccAddress) []types.NftBid {
	store := ctx.KVStore(k.storeKey)

	bids := []types.NftBid{}
	it := sdk.KVStorePrefixIterator(store, types.AddressBidKeyPrefix(bidder))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		bid := types.NftBid{}
		k.cdc.MustUnmarshal(it.Value(), &bid)

		bids = append(bids, bid)
	}
	return bids
}

func (k Keeper) SetBid(ctx sdk.Context, bid types.NftBid) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
	if err != nil {
		return err
	}
	if bid, err := k.GetBid(ctx, bid.IdBytes(), bidder); err == nil {
		err = k.DeleteBid(ctx, bid)
		if err != nil {
			return err
		}
	}

	bz := k.cdc.MustMarshal(&bid)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NftBidKey(bid.IdBytes(), bidder), bz)
	store.Set(types.AddressBidKey(bid.IdBytes(), bidder), bz)
	store.Set(append(getTimeKey(types.KeyPrefixEndTimeNftBid, bid.ExpiryAt), bid.GetIdToByte()...), bid.GetIdToByte())
	return nil
}

func (k Keeper) DeleteBid(ctx sdk.Context, bid types.NftBid) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NftBidKey(bid.IdBytes(), bidder))
	store.Delete(types.AddressBidKey(bid.IdBytes(), bidder))
	store.Delete(append(getTimeKey(types.KeyPrefixEndTimeNftBid, bid.ExpiryAt), bid.GetIdToByte()...))
	return nil
}

func (k Keeper) PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid) error {
	// Verify listing is in BIDDING state
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	if !listing.CanBid() {
		return types.ErrNftListingNotInBidState
	}

	if listing.BidDenom != msg.BidAmount.Denom {
		return types.ErrInvalidBidDenom
	}

	// todo add test case
	minimumBiddingPeriodHour := time.Now().Add(listing.MinimumBiddingPeriod)
	if msg.ExpiryAt.Before(minimumBiddingPeriodHour) {
		return types.ErrSmallBiddingPeriod
	}

	bids := types.NftBids(k.GetBidsByNft(ctx, listing.IdBytes()))
	bidder := msg.Sender
	oldBid := bids.GetBidByBidder(bidder)
	newBid := types.NftBid{
		Id: types.BidId{
			NftId:  &msg.NftId,
			Bidder: bidder,
		},
		BidAmount:        msg.BidAmount,
		DepositAmount:    msg.DepositAmount,
		PaidAmount:       sdk.NewCoin(listing.BidDenom, sdk.ZeroInt()),
		ExpiryAt:         msg.ExpiryAt,
		InterestRate:     msg.InterestRate,
		AutomaticPayment: msg.AutomaticPayment,
		CreatedAt:        ctx.BlockTime(),
		Borrow:           types.Borrowing{Amount: sdk.NewCoin(listing.BidDenom, sdk.ZeroInt()), LastRepaidAt: ctx.BlockTime()},
	}

	if !oldBid.IsNil() {
		return k.ReBid(ctx, listing, oldBid, newBid, bids)
	} else {
		return k.FirstBid(ctx, listing, newBid, bids)
	}
}

func (k Keeper) ManualBid(ctx sdk.Context, listing types.NftListing, newBid types.NftBid, bids types.NftBids) error {
	// delete kick out bid
	// err := CheckBidParams(listing, newBid.BidAmount, newBid.DepositAmount, bids)
	// if err != nil {
	// 	kickOutBid := bids.FindKickOutBid(newBid, ctx.BlockTime())
	// 	if kickOutBid.IsNil() {
	// 		// cannot kick out bid
	// 		return err
	// 	} else {
	// 		bids = bids.RemoveBids(types.NftBids{kickOutBid})
	// 		err = CheckBidParams(listing, newBid.BidAmount, newBid.DepositAmount, bids)
	// 		if err != nil {
	// 			return err
	// 		} else {
	// 			err = k.SafeCloseBidWithAllInterest(ctx, kickOutBid)
	// 			if err != nil {
	// 				return err
	// 			}
	// 		}
	// 	}
	// }

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.MustAccAddressFromBech32(newBid.Id.Bidder), types.ModuleName, sdk.Coins{newBid.DepositAmount})
	if err != nil {
		return err
	}

	// Add new bid on the listing
	err = k.SetBid(ctx, newBid)
	if err != nil {
		return err
	}

	// extend bid if there's bid within gap time
	params := k.GetParamSet(ctx)
	if listing.State == types.ListingState_LISTING {
		listing.State = types.ListingState_BIDDING
	}
	// todo implement listing end
	gapTime := ctx.BlockTime().Add(time.Duration(params.NftListingGapTime) * time.Second)
	if listing.EndAt.Before(gapTime) {
		listing.EndAt = gapTime
	}
	k.SaveNftListing(ctx, listing)

	return nil
}

func (k Keeper) FirstBid(ctx sdk.Context, listing types.NftListing, newBid types.NftBid, bids types.NftBids) error {
	err := k.ManualBid(ctx, listing, newBid, bids)
	if err != nil {
		return err
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPlaceBid{
		Bidder:  newBid.Id.Bidder,
		ClassId: newBid.Id.NftId.ClassId,
		NftId:   newBid.Id.NftId.NftId,
		Amount:  newBid.BidAmount.String(),
	})
	return nil
}

func (k Keeper) ReBid(ctx sdk.Context, listing types.NftListing, oldBid, newBid types.NftBid, bids types.NftBids) error {
	// check no borrow
	if !oldBid.CanReBid() {
		return types.ErrBorrowedDeposit
	}
	// check for liquidation
	if !types.IsAbleToReBid(bids, oldBid.Id, newBid) {
		return types.ErrCannotRebid
	}
	bids = bids.RemoveBid(oldBid)
	err := k.SafeCloseBid(ctx, oldBid)
	if err != nil {
		return err
	}
	err = k.ManualBid(ctx, listing, newBid, bids)
	if err != nil {
		return err
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPlaceBid{
		Bidder:  newBid.Id.Bidder,
		ClassId: newBid.Id.NftId.ClassId,
		NftId:   newBid.Id.NftId.NftId,
		Amount:  newBid.BidAmount.String(),
	})
	return nil
}

func (k Keeper) ManualSafeCloseBid(ctx sdk.Context, bid types.NftBid, bidder sdk.AccAddress) error {
	err := k.DeleteBid(ctx, bid)
	if err != nil {
		return err
	}
	if bid.PaidAmount.Amount.GT(sdk.ZeroInt()) {
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(bid.PaidAmount.Denom, bid.PaidAmount.Amount)})
		if err != nil {
			return err
		}
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(bid.DepositAmount.Denom, bid.DepositAmount.Amount)})
}

func (k Keeper) SafeCloseBid(ctx sdk.Context, bid types.NftBid) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
	if err != nil {
		return err
	}
	return k.ManualSafeCloseBid(ctx, bid, bidder)
}

func (k Keeper) SafeCloseBidCollectDeposit(ctx sdk.Context, bid types.NftBid) (sdk.Coin, error) {
	CollectedAmount := bid.DepositAmount
	err := k.DeleteBid(ctx, bid)
	if err != nil {
		return sdk.Coin{}, err
	}
	return CollectedAmount, nil
}

// todo make unit test
func (k Keeper) SafeCloseBidWithAllInterest(ctx sdk.Context, bid types.NftBid) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
	if err != nil {
		return err
	}
	interestAmount := bid.CompoundInterest(ctx.BlockTime())
	if interestAmount.Amount.IsPositive() {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(interestAmount.Denom, interestAmount.Amount)})
		if err != nil {
			return err
		}
	}
	return k.ManualSafeCloseBid(ctx, bid, bidder)
}

// implement SafeCloseBidWithPartInterest
// func (k Keeper) SafeCloseBidWithPartInterest(ctx sdk.Context, bid types.NftBid, interest sdk.Coin) error {
// 	bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
// 	if err != nil {
// 		return err
// 	}
// 	// check total interest amount is greater than interest
// 	if bid.TotalInterestAmount(ctx.BlockTime()).Amount.LT(interest.Amount) {
// 		return types.ErrInterestAmountTooLarge
// 	}
// 	if interest.Amount.GT(sdk.ZeroInt()) {
// 		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(interest.Denom, interest.Amount)})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return k.ManualSafeCloseBid(ctx, bid, bidder)
// }

func (k Keeper) CancelBid(ctx sdk.Context, msg *types.MsgCancelBid) error {
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	if !listing.CanCancelBid() {
		return types.ErrCannotCancelBid
	}

	bidder, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	// check if bid exists by bidder on nft
	bid, err := k.GetBid(ctx, msg.NftId.IdBytes(), bidder)
	if err != nil {
		return types.ErrBidDoesNotExists
	}

	if !bid.CanCancel() {
		return types.ErrBorrowedDeposit
	}

	// bids can only be cancelled X days after bidding
	params := k.GetParamSet(ctx)
	if bid.CreatedAt.Add(time.Duration(params.BidCancelRequiredSeconds) * time.Second).After(ctx.BlockTime()) {
		return types.ErrBidCancelIsAllowedAfterSomeTime
	}

	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	if len(bids) == 1 {
		return types.ErrCannotCancelListingSingleBid
	}
	// for liquidation validation
	if !types.IsAbleToCancelBid(types.BidId{Bidder: msg.Sender, NftId: bid.Id.NftId}, bids) {
		return types.ErrCannotCancelBid
	}

	err = k.DeleteBid(ctx, bid)
	if err != nil {
		return err
	}

	// tokens will be reimbursed X days after the bid cancellation is approved
	// bid.CreateAt = ctx.BlockTime().Add(time.Duration(params.BidDenomDisburseSecondsAfterCancel) * time.Second)
	// k.SetCancelledBid(ctx, bid)

	// return deposit amount to bidder
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{bid.DepositAmount})
	if err != nil {
		return err
	}

	// Emit event for cancelling bid
	_ = ctx.EventManager().EmitTypedEvent(&types.EventCancelBid{
		Bidder:  msg.Sender,
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) PayFullBid(ctx sdk.Context, msg *types.MsgPayFullBid) error {
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	bidder, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	// check if bid exists by bidder on nft
	bid, err := k.GetBid(ctx, msg.NftId.IdBytes(), bidder)
	if err != nil {
		return types.ErrBidDoesNotExists
	}

	// Transfer unpaid amount of token from bid account
	unpaidAmount := bid.BidAmount.Sub(bid.DepositAmount).Sub(bid.PaidAmount)
	if unpaidAmount.IsPositive() {
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidder, types.ModuleName, sdk.Coins{sdk.NewCoin(listing.BidDenom, unpaidAmount.Amount)})
		if err != nil {
			return err
		}

		bid.PaidAmount = bid.PaidAmount.Add(unpaidAmount)
		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
	}
	// Emit event for paying full bid
	_ = ctx.EventManager().EmitTypedEvent(&types.EventPayFullBid{
		Bidder:  msg.Sender,
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

// todo add unit test
func CheckBidParams(listing types.NftListing, bid, deposit sdk.Coin, bids []types.NftBid) error {
	c := listing.MinimumDepositRate
	p := sdk.NewDecFromInt(bid.Amount)
	cp := c.Mul(p)
	depositDec := sdk.NewDecFromInt(deposit.Amount)
	if cp.GT(depositDec) {
		return types.ErrNotEnoughDeposit
	}
	q := bid
	s := deposit
	for _, bid := range bids {
		q = q.Add(bid.BidAmount)
		s = s.Add(bid.DepositAmount)
	}

	bidLen := 1
	if len(bids) > 0 {
		// sum new bid and old bids
		bidLen = len(bids) + 1
	}
	q.Amount = q.Amount.Quo(sdk.NewInt(int64(bidLen)))
	if q.IsLTE(s) {
		fmt.Println("q", q.String())
		fmt.Println("s", s.String())
		return types.ErrBidParamInvalid
	}
	q_s := q.Sub(s)
	q_s_dec := sdk.NewDecFromInt(q_s.Amount)
	// todo implement min{bid, q-s}

	if depositDec.GT(q_s_dec) {
		fmt.Println("depositDec", depositDec.String())
		fmt.Println("q_s_dec", q_s_dec.String())
		// deposit amount bigger
		return types.ErrBidParamInvalid
	}
	return nil
}

func (k Keeper) GetActiveNftBiddingsExpired(ctx sdk.Context, endTime time.Time) []types.NftBid {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(types.KeyPrefixEndTimeNftBid, endTime)
	it := store.Iterator([]byte(types.KeyPrefixEndTimeNftBid), storetypes.InclusiveEndBytes(timeKey))
	defer it.Close()

	bids := []types.NftBid{}
	for ; it.Valid(); it.Next() {
		bidId := types.NftBidBytesToBidId(it.Value())
		fmt.Println("GetActiveNftBiddingsExpired")
		fmt.Println(bidId)
		bidder, _ := sdk.AccAddressFromBech32(bidId.Bidder)
		bid, err := k.GetBid(ctx, bidId.NftId.IdBytes(), bidder)
		if err != nil {
			fmt.Println(err)
			continue
		}
		bids = append(bids, bid)
	}
	return bids
}

func (k Keeper) DeleteBidsWithoutBorrowing(ctx sdk.Context, bids []types.NftBid) {
	for _, bid := range bids {
		if !bid.IsBorrowing() {
			err := k.SafeCloseBid(ctx, bid)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
