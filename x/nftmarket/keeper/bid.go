package keeper

import (
	"fmt"
	"time"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
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

	// sort bids by rank
	// todo: sord by lower deposit interest rate
	// sort.SliceStable(bids, func(i, j int) bool {
	// 	if bids[i].Amount.Amount.LT(bids[j].Amount.Amount) {
	// 		return true
	// 	}
	// 	if bids[i].Amount.Amount.GT(bids[j].Amount.Amount) {
	// 		return false
	// 	}
	// 	if bids[i].BidTime.After(bids[j].BidTime) {
	// 		return true
	// 	}
	// 	return false
	// })
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

func (k Keeper) SetBid(ctx sdk.Context, bid types.NftBid) {
	bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
	if err != nil {
		panic(err)
	}
	if bid, err := k.GetBid(ctx, bid.IdBytes(), bidder); err == nil {
		k.DeleteBid(ctx, bid)
	}

	bz := k.cdc.MustMarshal(&bid)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NftBidKey(bid.IdBytes(), bidder), bz)
	store.Set(types.AddressBidKey(bid.IdBytes(), bidder), bz)
	store.Set(append(getTimeKey(types.KeyPrefixEndTimeNftBid, bid.BiddingPeriod), bid.GetIdToByte()...), bid.GetIdToByte())
}

func (k Keeper) DeleteBid(ctx sdk.Context, bid types.NftBid) {
	bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
	if err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NftBidKey(bid.IdBytes(), bidder))
	store.Delete(types.AddressBidKey(bid.IdBytes(), bidder))
	store.Delete(append(getTimeKey(types.KeyPrefixEndTimeNftBid, bid.BiddingPeriod), bid.GetIdToByte()...))
}

func getCancelledBidTimeKey(timestamp time.Time) []byte {
	timeBz := sdk.FormatTimeBytes(timestamp)
	timeBzL := len(timeBz)
	prefixL := len(types.KeyPrefixNftBidCancelled)

	bz := make([]byte, prefixL+8+timeBzL)

	// copy the prefix
	copy(bz[:prefixL], types.KeyPrefixNftBidCancelled)

	// copy the encoded time bytes length
	copy(bz[prefixL:prefixL+8], sdk.Uint64ToBigEndian(uint64(timeBzL)))

	// copy the encoded time bytes
	copy(bz[prefixL+8:prefixL+8+timeBzL], timeBz)
	return bz
}

func (k Keeper) SetCancelledBid(ctx sdk.Context, bid types.NftBid) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&bid)
	store.Set(append(append(getCancelledBidTimeKey(bid.BidTime), bid.IdBytes()...), []byte(bid.Bidder)...), bz)
}

func (k Keeper) GetAllCancelledBids(ctx sdk.Context) []types.NftBid {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixNftBidCancelled))
	defer it.Close()

	bids := []types.NftBid{}
	for ; it.Valid(); it.Next() {
		bid := types.NftBid{}
		k.cdc.MustUnmarshal(it.Value(), &bid)
		bids = append(bids, bid)
	}
	return bids
}

func (k Keeper) GetMaturedCancelledBids(ctx sdk.Context, endTime time.Time) []types.NftBid {
	store := ctx.KVStore(k.storeKey)
	timeKey := getCancelledBidTimeKey(endTime)
	it := store.Iterator([]byte(types.KeyPrefixNftBidCancelled), storetypes.InclusiveEndBytes(timeKey))
	defer it.Close()

	bids := []types.NftBid{}
	for ; it.Valid(); it.Next() {
		bid := types.NftBid{}
		k.cdc.MustUnmarshal(it.Value(), &bid)
		bids = append(bids, bid)
	}
	return bids
}

// todo: implement for all bidder
func (k Keeper) TotalActiveRankDeposit(ctx sdk.Context, nftIdBytes []byte) sdk.Int {
	// listing, err := k.GetNftListingByIdBytes(ctx, nftIdBytes)
	// if err != nil {
	// 	return sdk.ZeroInt()
	// }

	// bids := k.GetBidsByNft(ctx, nftIdBytes)
	// totalActiveRankDeposit := sdk.ZeroInt()

	// winnerCandidateStartIndex := len(bids) - int(listing.BidActiveRank)
	// if winnerCandidateStartIndex < 0 {
	// 	winnerCandidateStartIndex = 0
	// }
	// for _, bid := range bids[winnerCandidateStartIndex:] {
	// 	totalActiveRankDeposit = totalActiveRankDeposit.Add(bid.PaidAmount)
	// }
	// return totalActiveRankDeposit
	return sdk.NewInt(1)
}

func (k Keeper) DeleteCancelledBid(ctx sdk.Context, bid types.NftBid) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(append(getCancelledBidTimeKey(bid.BidTime), bid.IdBytes()...), []byte(bid.Bidder)...))
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

	if listing.BidToken != msg.BidAmount.Denom {
		return types.ErrInvalidBidDenom
	}
	depositLendingRate, err := sdk.NewDecFromStr(msg.DepositLendingRate)
	if err != nil {
		return types.ErrCannotParseDec
	}
	// todo we not decided rebid spec
	// re-bidのとき
	// 自動借り換えを行う
	// 自動借り換えは
	// 以前の利息を計算して引き継ぐ

	bids := types.NftBids(k.GetBidsByNft(ctx, listing.IdBytes()))
	err = CheckBidParams(listing, msg.BidAmount, msg.DepositAmount, bids)
	if err != nil {
		return err
	}

	bidder := msg.Sender.AccAddress()
	bid := bids.GetBidByBidder(bidder.String())

	if !bid.IsNil() && !bid.CanReBid() {
		return types.ErrBorrowedDeposit
	}

	// if !bid.IsNil() {
	// 	// k.FirstBid(ctx, msg, listing, )
	// } else {
	// 	// k.ReBid(ctx, msg, listing, bid)
	// }

	// todo we not decided rebid spec
	// todo update for v2
	// increaseAmount := msg.BidAmount.Amount
	// paidAmount := sdk.ZeroInt()
	// if previous bid exists add more on top of existings
	// bid, err := k.GetBid(ctx, msg.NftId.IdBytes(), bidder)
	// if err == nil {
	// todo change check logic
	// if bid.Amount.Amount.GTE(msg.Amount.Amount) {
	// 	return types.ErrInvalidBidAmount
	// }

	// todo: update new logic
	// paidAmount = bid.DepositAmount
	// increaseAmount = msg.BidAmount.Amount.Sub(bid.BidAmount.Amount)
	// }

	// todo update for v2
	// Transfer required amount of token from bid account to module
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidder, types.ModuleName, sdk.Coins{msg.DepositAmount})
	if err != nil {
		return err
	}

	// todo update for v2
	// Add new bid on the listing
	k.SetBid(ctx, types.NftBid{
		NftId:            msg.NftId,
		Bidder:           msg.Sender.AccAddress().String(),
		BidAmount:        msg.BidAmount,
		AutomaticPayment: msg.AutomaticPayment,
		DepositAmount:    msg.DepositAmount,
		BidTime:          ctx.BlockTime(),
		BiddingPeriod:    msg.BiddingPeriod,
		Id: types.BidId{
			NftId:  &msg.NftId,
			Bidder: msg.Sender.AccAddress().String(),
		},
		PaidAmount:         sdk.NewCoin(listing.BidToken, sdk.ZeroInt()),
		DepositLendingRate: depositLendingRate,
		InterestAmount:     sdk.NewCoin(listing.BidToken, sdk.ZeroInt()),
	})

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

	// Emit event for placing bid
	ctx.EventManager().EmitTypedEvent(&types.EventPlaceBid{
		Bidder:  msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
		Amount:  msg.BidAmount.String(),
	})

	return nil
}

func (k Keeper) FirstBid(ctx sdk.Context, listing types.NftListing, bid types.NftBid, bidder sdk.AccAddress, msg *types.MsgPlaceBid) error {
	depositLendingRate, err := sdk.NewDecFromStr(msg.DepositLendingRate)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidder, types.ModuleName, sdk.Coins{msg.DepositAmount})
	if err != nil {
		return err
	}

	// todo update for v2
	// Add new bid on the listing
	k.SetBid(ctx, types.NftBid{
		NftId:            msg.NftId,
		Bidder:           msg.Sender.AccAddress().String(),
		BidAmount:        msg.BidAmount,
		AutomaticPayment: msg.AutomaticPayment,
		DepositAmount:    msg.DepositAmount,
		BidTime:          ctx.BlockTime(),
		BiddingPeriod:    msg.BiddingPeriod,
		Id: types.BidId{
			NftId:  &msg.NftId,
			Bidder: msg.Sender.AccAddress().String(),
		},
		PaidAmount:         sdk.NewCoin(listing.BidToken, sdk.ZeroInt()),
		DepositLendingRate: depositLendingRate,
		InterestAmount:     sdk.NewCoin(listing.BidToken, sdk.ZeroInt()),
	})

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

	// Emit event for placing bid
	ctx.EventManager().EmitTypedEvent(&types.EventPlaceBid{
		Bidder:  msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
		Amount:  msg.BidAmount.String(),
	})
	return nil
}

func higherBids(bids []types.NftBid, amount sdk.Int) uint64 {
	higherBids := uint64(0)
	// todo: update for v2
	// for _, bid := range bids {
	// 	if bid.Amount.Amount.GTE(amount) {
	// 		higherBids++
	// 	}
	// }
	return higherBids
}

func (k Keeper) SafeCloseBid(ctx sdk.Context, bid types.NftBid) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
	if err != nil {
		return err
	}

	// Delete bid
	k.DeleteBid(ctx, bid)

	// todo send paid amount to bidder
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(bid.DepositAmount.Denom, bid.DepositAmount.Amount)})
}

func (k Keeper) CancelBid(ctx sdk.Context, msg *types.MsgCancelBid) error {
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	if !listing.CanCancelBid() {
		return types.ErrCannotCancelBid
	}

	bidder := msg.Sender.AccAddress()

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
	if bid.BidTime.Add(time.Duration(params.BidCancelRequiredSeconds) * time.Second).After(ctx.BlockTime()) {
		return types.ErrBidCancelIsAllowedAfterSomeTime
	}

	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	if len(bids) == 1 {
		return types.ErrCannotCancelListingSingleBid
	}

	k.DeleteBid(ctx, bid)

	// tokens will be reimbursed X days after the bid cancellation is approved
	bid.BidTime = ctx.BlockTime().Add(time.Duration(params.BidTokenDisburseSecondsAfterCancel) * time.Second)
	k.SetCancelledBid(ctx, bid)

	// TODO: Liquidation may occur for sellers whose bids are cancelled.

	// Emit event for cancelling bid
	ctx.EventManager().EmitTypedEvent(&types.EventCancelBid{
		Bidder:  msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

// todo test for pay part amount and full amount
func (k Keeper) PayFullBid(ctx sdk.Context, msg *types.MsgPayFullBid) error {
	// todo update for v2
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	bidder := msg.Sender.AccAddress()

	// check if bid exists by bidder on nft
	bid, err := k.GetBid(ctx, msg.NftId.IdBytes(), bidder)
	if err != nil {
		return types.ErrBidDoesNotExists
	}

	// Transfer unpaid amount of token from bid account
	unpaidAmount := bid.BidAmount.Sub(bid.DepositAmount).Sub(bid.PaidAmount)
	if unpaidAmount.IsPositive() {
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidder, types.ModuleName, sdk.Coins{sdk.NewCoin(listing.BidToken, unpaidAmount.Amount)})
		if err != nil {
			return err
		}

		bid.PaidAmount = bid.PaidAmount.Add(unpaidAmount)
		k.SetBid(ctx, bid)
	}
	// Emit event for paying full bid
	ctx.EventManager().EmitTypedEvent(&types.EventPayFullBid{
		Bidder:  msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) HandleMaturedCancelledBids(ctx sdk.Context) error {
	bids := k.GetMaturedCancelledBids(ctx, ctx.BlockTime())
	for _, bid := range bids {
		// transfer amount of token except fee to bid account
		bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(bid.BidAmount.Denom, bid.DepositAmount.Amount)})
		if err != nil {
			return err
		}
		k.DeleteCancelledBid(ctx, bid)
	}

	return nil
}

// todo add unit test
func CheckBidParams(listing types.NftListing, bid, deposit sdk.Coin, bids []types.NftBid) error {
	c, err := sdk.NewDecFromStr(listing.MinimumDepositRate)
	if err != nil {
		return err
	}
	p := sdk.NewDecFromInt(bid.Amount)
	cp := c.Mul(p)
	depositDec := sdk.NewDecFromInt(deposit.Amount)
	if cp.GT(depositDec) {
		return types.ErrNotEnoughDeposit
	}
	q := bid
	s := deposit
	for _, bid := range bids {
		q.Add(bid.BidAmount)
		s.Add(bid.DepositAmount)
	}
	if q.IsLTE(s) {
		return types.ErrBidParamInvalid
	}
	q_s := q.Sub(s)
	q_s_dec := sdk.NewDecFromInt(q_s.Amount)
	if depositDec.GT(q_s_dec) {
		// deposit amount bigger
		return types.ErrBidParamInvalid
	}
	return nil
}

func (k Keeper) GetActiveNftBiddingsEndingAt(ctx sdk.Context, endTime time.Time) []types.NftBid {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(types.KeyPrefixEndTimeNftBid, endTime)
	it := store.Iterator([]byte(types.KeyPrefixEndTimeNftBid), storetypes.InclusiveEndBytes(timeKey))
	defer it.Close()

	bids := []types.NftBid{}
	for ; it.Valid(); it.Next() {
		bidId := types.NftBidBytesToBidId(it.Value())
		fmt.Println("GetActiveNftBiddingsEndingAt")
		fmt.Println(bidId)
		bidder, _ := sdk.AccAddressFromBech32(bidId.Bidder)
		bid, err := k.GetBid(ctx, bidId.NftId.IdBytes(), bidder)
		if err != nil {
			panic(err)
		}
		bids = append(bids, bid)
	}
	return bids
}

func (k Keeper) DeleteBidsWithoutBorrowing(ctx sdk.Context, bids []types.NftBid) {
	for _, bid := range bids {
		if !bid.IsBorrowing() {
			k.SafeCloseBid(ctx, bid)
		}
	}
}
