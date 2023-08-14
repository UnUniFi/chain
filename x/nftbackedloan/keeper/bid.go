package keeper

import (
	"fmt"
	"time"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) GetBid(ctx sdk.Context, nftIdBytes []byte, bidder sdk.AccAddress) (types.Bid, error) {
	bid := types.Bid{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NftBidKey(nftIdBytes, bidder))
	if bz == nil {
		return bid, types.ErrBidDoesNotExists
	}

	k.cdc.MustUnmarshal(bz, &bid)
	return bid, nil
}

func (k Keeper) GetAllBids(ctx sdk.Context) []types.Bid {
	store := ctx.KVStore(k.storeKey)

	bids := []types.Bid{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixNftBid))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		bid := types.Bid{}
		k.cdc.MustUnmarshal(it.Value(), &bid)

		bids = append(bids, bid)
	}
	return bids
}

func (k Keeper) GetBidsByNft(ctx sdk.Context, nftIdBytes []byte) []types.Bid {
	store := ctx.KVStore(k.storeKey)

	bids := []types.Bid{}
	it := sdk.KVStorePrefixIterator(store, append([]byte(types.KeyPrefixNftBid), nftIdBytes...))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		bid := types.Bid{}
		k.cdc.MustUnmarshal(it.Value(), &bid)

		bids = append(bids, bid)
	}
	return bids
}

func (k Keeper) GetBidsByBidder(ctx sdk.Context, bidder sdk.AccAddress) []types.Bid {
	store := ctx.KVStore(k.storeKey)

	bids := []types.Bid{}
	it := sdk.KVStorePrefixIterator(store, types.AddressBidKeyPrefix(bidder))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		bid := types.Bid{}
		k.cdc.MustUnmarshal(it.Value(), &bid)

		bids = append(bids, bid)
	}
	return bids
}

func (k Keeper) SetBid(ctx sdk.Context, bid types.Bid) error {
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
	store.Set(append(getTimeKey(types.KeyPrefixEndTimeNftBid, bid.Expiry), bid.GetIdToByte()...), bid.GetIdToByte())
	return nil
}

func (k Keeper) DeleteBid(ctx sdk.Context, bid types.Bid) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NftBidKey(bid.IdBytes(), bidder))
	store.Delete(types.AddressBidKey(bid.IdBytes(), bidder))
	store.Delete(append(getTimeKey(types.KeyPrefixEndTimeNftBid, bid.Expiry), bid.GetIdToByte()...))
	return nil
}

func (k Keeper) PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid) error {
	// Verify listing is in BIDDING state
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	if !listing.CanBid() {
		return types.ErrStatusCannotPlaceBid
	}

	if listing.Owner == msg.Sender {
		return types.ErrOwnerCannotPlaceBid
	}

	if listing.BidDenom != msg.Price.Denom {
		return types.ErrInvalidPriceDenom
	}
	if !msg.Price.IsPositive() {
		return types.ErrInvalidPriceAmount
	}
	if listing.BidDenom != msg.Deposit.Denom {
		return types.ErrInvalidDepositDenom
	}

	// check over min_deposit_rate
	minDeposit := listing.MinDepositRate.Mul(sdk.NewDecFromInt(msg.Price.Amount)).TruncateInt()
	if msg.Deposit.Amount.LT(minDeposit) {
		return types.ErrInvalidDepositAmount
	}

	// todo add test case
	minimumBiddingPeriodHour := time.Now().Add(listing.MinBidPeriod)
	if msg.Expiry.Before(minimumBiddingPeriodHour) {
		return types.ErrSmallExpiryPeriod
	}

	bids := types.NftBids(k.GetBidsByNft(ctx, listing.IdBytes()))
	bidder := msg.Sender
	oldBid := bids.GetBidByBidder(bidder)
	newBid := types.Bid{
		Id: types.BidId{
			NftId:  &msg.NftId,
			Bidder: bidder,
		},
		Price:            msg.Price,
		Deposit:          msg.Deposit,
		PaidAmount:       sdk.NewCoin(listing.BidDenom, sdk.ZeroInt()),
		Expiry:           msg.Expiry,
		InterestRate:     msg.InterestRate,
		AutomaticPayment: msg.AutomaticPayment,
		CreatedAt:        ctx.BlockTime(),
		Loan:             types.Loan{Amount: sdk.NewCoin(listing.BidDenom, sdk.ZeroInt()), LastRepaidAt: ctx.BlockTime()},
	}

	if !oldBid.IsNil() {
		return k.ReBid(ctx, listing, oldBid, newBid, bids)
	} else {
		return k.FirstBid(ctx, listing, newBid)
	}
}

func (k Keeper) ManualBid(ctx sdk.Context, listing types.Listing, newBid types.Bid) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.MustAccAddressFromBech32(newBid.Id.Bidder), types.ModuleName, sdk.Coins{newBid.Deposit})
	if err != nil {
		return err
	}

	// Add new bid on the listing
	err = k.SetBid(ctx, newBid)
	if err != nil {
		return err
	}

	if listing.State == types.ListingState_LISTING {
		listing.State = types.ListingState_BIDDING
	}
	k.SaveNftListing(ctx, listing)

	return nil
}

func (k Keeper) FirstBid(ctx sdk.Context, listing types.Listing, newBid types.Bid) error {
	err := k.ManualBid(ctx, listing, newBid)
	if err != nil {
		return err
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPlaceBid{
		Bidder:  newBid.Id.Bidder,
		ClassId: newBid.Id.NftId.ClassId,
		TokenId: newBid.Id.NftId.TokenId,
		Amount:  newBid.Price.String(),
	})
	return nil
}

func (k Keeper) ReBid(ctx sdk.Context, listing types.Listing, oldBid, newBid types.Bid, bids types.NftBids) error {
	// check no borrow
	if !oldBid.CanReBid() {
		return types.ErrCannotChangeBidBorrowed
	}
	// check for liquidation
	if !types.IsAbleToReBid(bids, oldBid.Id, newBid, listing) {
		return types.ErrCannotReBidForLiquidation
	}

	err := k.SafeCloseBid(ctx, oldBid)
	if err != nil {
		return err
	}
	err = k.ManualBid(ctx, listing, newBid)
	if err != nil {
		return err
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPlaceBid{
		Bidder:  newBid.Id.Bidder,
		ClassId: newBid.Id.NftId.ClassId,
		TokenId: newBid.Id.NftId.TokenId,
		Amount:  newBid.Price.String(),
	})
	return nil
}

func (k Keeper) ManualSafeCloseBid(ctx sdk.Context, bid types.Bid, bidder sdk.AccAddress) error {
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
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(bid.Deposit.Denom, bid.Deposit.Amount)})
}

func (k Keeper) SafeCloseBid(ctx sdk.Context, bid types.Bid) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
	if err != nil {
		return err
	}
	return k.ManualSafeCloseBid(ctx, bid, bidder)
}

func (k Keeper) SafeCloseBidCollectDeposit(ctx sdk.Context, bid types.Bid) (sdk.Coin, error) {
	CollectedAmount := bid.Deposit
	err := k.DeleteBid(ctx, bid)
	if err != nil {
		return sdk.Coin{}, err
	}
	return CollectedAmount, nil
}

// todo make unit test
func (k Keeper) SafeCloseBidWithAllInterest(ctx sdk.Context, bid types.Bid, time time.Time) error {
	bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
	if err != nil {
		return err
	}
	interestAmount := bid.CompoundInterest(time)
	if interestAmount.Amount.IsPositive() {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(interestAmount.Denom, interestAmount.Amount)})
		if err != nil {
			return err
		}
	}
	return k.ManualSafeCloseBid(ctx, bid, bidder)
}

func (k Keeper) CancelBid(ctx sdk.Context, msg *types.MsgCancelBid) error {
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return err
	}

	// check status
	if !listing.CanCancelBid() {
		return types.ErrStatusCannotCancelBid
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

	// check borrow
	if !bid.CanCancel() {
		return types.ErrCannotChangeBidBorrowed
	}

	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	if len(bids) == 1 {
		return types.ErrCannotCancelListingSingleBid
	}
	// for liquidation validation
	if !types.IsAbleToCancelBid(types.BidId{Bidder: msg.Sender, NftId: bid.Id.NftId}, bids, listing) {
		return types.ErrCannotCancelBidForLiquidation
	}

	err = k.SafeCloseBid(ctx, bid)
	if err != nil {
		return err
	}

	// Emit event for cancelling bid
	_ = ctx.EventManager().EmitTypedEvent(&types.EventCancelBid{
		Bidder:  msg.Sender,
		ClassId: msg.NftId.ClassId,
		TokenId: msg.NftId.TokenId,
	})

	return nil
}

func (k Keeper) PayRemainder(ctx sdk.Context, msg *types.MsgPayRemainder) error {
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
	unpaidAmount := bid.Price.Sub(bid.Deposit).Sub(bid.PaidAmount)
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
	_ = ctx.EventManager().EmitTypedEvent(&types.EventPayRemainder{
		Bidder:  msg.Sender,
		ClassId: msg.NftId.ClassId,
		TokenId: msg.NftId.TokenId,
	})

	return nil
}

func (k Keeper) GetExpiredBids(ctx sdk.Context, endTime time.Time) []types.Bid {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(types.KeyPrefixEndTimeNftBid, endTime)
	it := store.Iterator([]byte(types.KeyPrefixEndTimeNftBid), storetypes.InclusiveEndBytes(timeKey))
	defer it.Close()

	bids := []types.Bid{}
	for ; it.Valid(); it.Next() {
		bidId := types.NftBidBytesToBidId(it.Value())
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

func (k Keeper) DeleteBidsWithoutBorrowing(ctx sdk.Context, bids []types.Bid) {
	for _, bid := range bids {
		listing, err := k.GetNftListingByIdBytes(ctx, bid.Id.NftId.IdBytes())
		if err != nil {
			fmt.Println(err)
			continue
		}
		if listing.IsBidding() {
			if !bid.IsBorrowed() {
				err := k.SafeCloseBid(ctx, bid)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}
	}
}
