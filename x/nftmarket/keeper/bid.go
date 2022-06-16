package keeper

import (
	"time"

	"github.com/UnUniFi/chain/x/nftmarket/types"
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

func (k Keeper) PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid) error {
	// Verify listing is in BIDDING state
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	if listing.State != types.ListingState_BIDDING {
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
	})

	// extend bid if there's bid within gap time
	params := k.GetParamSet(ctx)
	gapTime := ctx.BlockTime().Add(time.Duration(params.NftListingGapTime) * time.Second)
	if listing.EndAt.Before(gapTime) {
		listing.EndAt = gapTime
		k.SetNftListing(ctx, listing)
	}

	// Emit event for placing bid
	ctx.EventManager().EmitTypedEvent(&types.EventPlaceBid{
		Bidder:  msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
		Amount:  msg.Amount.String(),
	})

	return nil
}

func (k Keeper) CancelBid(ctx sdk.Context, msg *types.MsgCancelBid) error {
	// Verify listing is in BIDDING state
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	_ = listing

	// TODO: if you are the only bidder yourself, you cannot cancel
	// TODO: Bidder can cancel bids free of charge if the bidder's bid rank is below the bid_active_rank.
	// TODO: if the bid rank is bid_active_rank or higher, the bid can be cancelled by paying a cancellation fee.
	// TODO: Cancellation Fee Formula: MAX{canceling_bidder's_deposit - (total_deposit - borrowed_lister_amount), 0}
	// TODO: bids can only be cancelled X days after bidding (global_option)
	// TODO: tokens will be reimbursed X days after the bid cancellation is approved (global_option)
	// TODO: Liquidation may occur for sellers whose bids are cancelled.
	// TODO: the bidder and the bid canceller's Sign must match, otherwise the bid will not be accepted and a log will be kept.

	bidder := msg.Sender.AccAddress()

	// check if bid exists by bidder on nft
	bid, err := k.GetBid(ctx, msg.NftId.IdBytes(), bidder)
	if err != nil {
		return types.ErrBidDoesNotExists
	}

	// Delete bid
	k.DeleteBid(ctx, bid)

	// TODO: handle cancel fee
	// params := k.GetParamSet(ctx)
	// params.NftListingCancelFeePercentage

	// Transfer amount of token to bid account
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(listing.BidToken, bid.PaidAmount)})
	if err != nil {
		return err
	}

	// Emit event for cancelling bid
	ctx.EventManager().EmitTypedEvent(&types.EventCancelBid{
		Bidder:  msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) PayFullBid(ctx sdk.Context, msg *types.MsgPayFullBid) error {
	// Verify listing is in SUCCESSFUL_BID state
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	if listing.State != types.ListingState_SUCCESSFUL_BID {
		return types.ErrNftListingNotInSuccessfulBidPhase
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
