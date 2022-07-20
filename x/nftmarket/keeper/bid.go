package keeper

import (
	"sort"
	"time"

	"github.com/UnUniFi/chain/x/nftmarket/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	sort.SliceStable(bids, func(i, j int) bool {
		if bids[i].Amount.Amount.LT(bids[j].Amount.Amount) {
			return true
		}
		if bids[i].Amount.Amount.GT(bids[j].Amount.Amount) {
			return false
		}
		if bids[i].BidTime.After(bids[j].BidTime) {
			return true
		}
		return false
	})
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
}

func (k Keeper) DeleteBid(ctx sdk.Context, bid types.NftBid) {
	bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
	if err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NftBidKey(bid.IdBytes(), bidder))
	store.Delete(types.AddressBidKey(bid.IdBytes(), bidder))
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

func (k Keeper) TotalActiveRankDeposit(ctx sdk.Context, nftIdBytes []byte) sdk.Int {
	listing, err := k.GetNftListingByIdBytes(ctx, nftIdBytes)
	if err != nil {
		return sdk.ZeroInt()
	}

	bids := k.GetBidsByNft(ctx, nftIdBytes)
	totalActiveRankDeposit := sdk.ZeroInt()

	winnerCandidateStartIndex := len(bids) - int(listing.BidActiveRank)
	if winnerCandidateStartIndex < 0 {
		winnerCandidateStartIndex = 0
	}
	for _, bid := range bids[winnerCandidateStartIndex:] {
		totalActiveRankDeposit = totalActiveRankDeposit.Add(bid.PaidAmount)
	}
	return totalActiveRankDeposit
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

	if listing.State != types.ListingState_LISTING && listing.State != types.ListingState_BIDDING {
		return types.ErrNftListingNotInBidState
	}

	if listing.BidToken != msg.Amount.Denom {
		return types.ErrInvalidBidDenom
	}

	if listing.MinBid.GT(msg.Amount.Amount) {
		return types.ErrInvalidBidAmount
	}

	bidder := msg.Sender.AccAddress()
	increaseAmount := msg.Amount.Amount
	paidAmount := sdk.ZeroInt()

	// if previous bid exists add more on top of existings
	bid, err := k.GetBid(ctx, msg.NftId.IdBytes(), bidder)
	if err == nil {
		if bid.Amount.Amount.GTE(msg.Amount.Amount) {
			return types.ErrInvalidBidAmount
		}
		paidAmount = bid.PaidAmount
		increaseAmount = msg.Amount.Amount.Sub(bid.Amount.Amount)
	}

	// Transfer required amount of token from bid account to module
	initialDeposit := increaseAmount.Quo(sdk.NewInt(int64(listing.BidActiveRank)))
	if initialDeposit.IsPositive() {
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidder, types.ModuleName, sdk.Coins{sdk.NewCoin(listing.BidToken, initialDeposit)})
		if err != nil {
			return err
		}
	}

	// Add new bid on the listing
	k.SetBid(ctx, types.NftBid{
		NftId:            msg.NftId,
		Bidder:           msg.Sender.AccAddress().String(),
		Amount:           msg.Amount,
		AutomaticPayment: msg.AutomaticPayment,
		PaidAmount:       paidAmount.Add(initialDeposit),
		BidTime:          ctx.BlockTime(),
	})

	// extend bid if there's bid within gap time
	params := k.GetParamSet(ctx)
	if listing.State == types.ListingState_LISTING {
		listing.State = types.ListingState_BIDDING
	}
	gapTime := ctx.BlockTime().Add(time.Duration(params.NftListingGapTime) * time.Second)
	if listing.EndAt.Before(gapTime) {
		listing.EndAt = gapTime
	}
	k.SetNftListing(ctx, listing)

	// Emit event for placing bid
	ctx.EventManager().EmitTypedEvent(&types.EventPlaceBid{
		Bidder:  msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
		Amount:  msg.Amount.String(),
	})

	return nil
}

func higherBids(bids []types.NftBid, amount sdk.Int) uint64 {
	higherBids := uint64(0)
	for _, bid := range bids {
		if bid.Amount.Amount.GTE(amount) {
			higherBids++
		}
	}
	return higherBids
}

func (k Keeper) SafeCloseBid(ctx sdk.Context, bid types.NftBid) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
	if err != nil {
		return err
	}

	// Delete bid
	k.DeleteBid(ctx, bid)

	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(bid.Amount.Denom, bid.PaidAmount)})
}

func (k Keeper) CancelBid(ctx sdk.Context, msg *types.MsgCancelBid) error {
	// Verify listing is in BIDDING state
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

	// bids can only be cancelled X days after bidding
	params := k.GetParamSet(ctx)
	if bid.BidTime.Add(time.Duration(params.BidCancelRequiredSeconds) * time.Second).After(ctx.BlockTime()) {
		return types.ErrBidCancelIsAllowedAfterSomeTime
	}

	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	if len(bids) == 1 {
		return types.ErrCannotCancelListingSingleBid
	}

	cancelFee := sdk.ZeroInt()
	if higherBids(bids, bid.Amount.Amount) <= listing.BidActiveRank {
		// Cancellation Fee Formula: MAX{canceling_bidder's_deposit - (total_deposit - borrowed_lister_amount), 0}
		listingDebt := k.GetDebtByNft(ctx, msg.NftId.IdBytes())
		totalDeposit := sdk.ZeroInt()
		for _, b := range bids {
			totalDeposit = totalDeposit.Add(b.PaidAmount)
		}

		// PANIC HERE XXXXlistingDebt
		loanAmount := sdk.ZeroInt()
		if listingDebt.NftId.NftId != "" {
			loanAmount = listingDebt.Loan.Amount
		}
		if bid.PaidAmount.Add(loanAmount).GT(totalDeposit) {
			cancelFee = bid.PaidAmount.Add(loanAmount).Sub(totalDeposit)
		}
	}

	// Delete bid
	k.DeleteBid(ctx, bid)

	// tokens will be reimbursed X days after the bid cancellation is approved
	bid.BidTime = ctx.BlockTime().Add(time.Duration(params.BidTokenDisburseSecondsAfterCancel) * time.Second)
	bid.PaidAmount = bid.PaidAmount.Sub(cancelFee)
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

func (k Keeper) PayFullBid(ctx sdk.Context, msg *types.MsgPayFullBid) error {
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
	unpaidAmount := bid.Amount.Amount.Sub(bid.PaidAmount)
	if unpaidAmount.IsPositive() {
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidder, types.ModuleName, sdk.Coins{sdk.NewCoin(listing.BidToken, unpaidAmount)})
		if err != nil {
			return err
		}

		bid.PaidAmount = bid.Amount.Amount
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
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(bid.Amount.Denom, bid.PaidAmount)})
		if err != nil {
			return err
		}
		k.DeleteCancelledBid(ctx, bid)
	}

	return nil
}
