package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	ecoincentivetypes "github.com/UnUniFi/chain/x/ecosystemincentive/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) GetNftListingByIdBytes(ctx sdk.Context, nftIdBytes []byte) (types.NftListing, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NftListingKey(nftIdBytes))
	if bz == nil {
		return types.NftListing{}, types.ErrNftListingDoesNotExist
	}
	listing := types.NftListing{}
	k.cdc.MustUnmarshal(bz, &listing)
	return listing, nil
}

func (k Keeper) GetListingsByOwner(ctx sdk.Context, owner sdk.AccAddress) []types.NftListing {
	store := ctx.KVStore(k.storeKey)

	listings := []types.NftListing{}
	it := sdk.KVStorePrefixIterator(store, types.NftAddressNftListingPrefixKey(owner))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		nftIdBytes := it.Value()
		listing, err := k.GetNftListingByIdBytes(ctx, nftIdBytes)
		if err != nil {
			fmt.Println("failed to get listing by id bytes: %w", err)
			continue
		}

		listings = append(listings, listing)
	}
	return listings
}

func getTimeKey(prefix string, timestamp time.Time) []byte {
	timeBz := sdk.FormatTimeBytes(timestamp)
	timeBzL := len(timeBz)
	prefixL := len(prefix)

	bz := make([]byte, prefixL+8+timeBzL)

	// copy the prefix
	copy(bz[:prefixL], prefix)

	// copy the encoded time bytes length
	copy(bz[prefixL:prefixL+8], sdk.Uint64ToBigEndian(uint64(timeBzL)))

	// copy the encoded time bytes
	copy(bz[prefixL+8:prefixL+8+timeBzL], timeBz)
	return bz
}

// call this method when you want to call SetNftListing
func (k Keeper) SaveNftListing(ctx sdk.Context, listing types.NftListing) {
	k.SetNftListing(ctx, listing)
	k.UpdateListedClass(ctx, listing)
}

func (k Keeper) SetNftListing(ctx sdk.Context, listing types.NftListing) {
	if oldListing, err := k.GetNftListingByIdBytes(ctx, listing.IdBytes()); err == nil {
		k.DeleteNftListings(ctx, oldListing)
	}

	nftIdBytes := listing.IdBytes()
	bz := k.cdc.MustMarshal(&listing)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NftListingKey(nftIdBytes), bz)

	owner, err := sdk.AccAddressFromBech32(listing.Owner)
	if err != nil {
		fmt.Println("invalid owner address: %w", err)
		return
	}
	store.Set(types.NftAddressNftListingKey(owner, nftIdBytes), nftIdBytes)

	if listing.IsFullPayment() {
		store.Set(append(getTimeKey(types.KeyPrefixFullPaymentPeriodListing, listing.FullPaymentEndAt), nftIdBytes...), nftIdBytes)
	} else if listing.IsSuccessfulBid() {
		store.Set(append(getTimeKey(types.KeyPrefixSuccessfulBidListing, listing.SuccessfulBidEndAt), nftIdBytes...), nftIdBytes)
	}
}

// call this method when you want to call DeleteNftListing
func (k Keeper) DeleteNftListings(ctx sdk.Context, listing types.NftListing) {
	k.DeleteNftListing(ctx, listing)
	k.UpdateListedClass(ctx, listing)
}

func (k Keeper) DeleteNftListing(ctx sdk.Context, listing types.NftListing) {
	nftIdBytes := listing.IdBytes()
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NftListingKey(nftIdBytes))

	owner, err := sdk.AccAddressFromBech32(listing.Owner)
	if err != nil {
		fmt.Println("invalid owner address: %w", err)
		return
	}
	store.Delete(types.NftAddressNftListingKey(owner, nftIdBytes))

	if listing.IsFullPayment() {
		store.Delete(append(getTimeKey(types.KeyPrefixFullPaymentPeriodListing, listing.FullPaymentEndAt), nftIdBytes...))
	} else if listing.IsSuccessfulBid() {
		store.Delete(append(getTimeKey(types.KeyPrefixSuccessfulBidListing, listing.SuccessfulBidEndAt), nftIdBytes...))
	}
}

func (k Keeper) GetFullPaymentNftListingsEndingAt(ctx sdk.Context, endTime time.Time) []types.NftListing {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(types.KeyPrefixFullPaymentPeriodListing, endTime)
	it := store.Iterator([]byte(types.KeyPrefixFullPaymentPeriodListing), storetypes.InclusiveEndBytes(timeKey))
	defer it.Close()

	listings := []types.NftListing{}
	for ; it.Valid(); it.Next() {
		nftIdBytes := it.Value()
		listing, err := k.GetNftListingByIdBytes(ctx, nftIdBytes)
		if err != nil {
			fmt.Println("failed to get listing by id bytes: %w", err)
			continue
		}

		listings = append(listings, listing)
	}
	return listings
}

func (k Keeper) GetSuccessfulBidNftListingsEndingAt(ctx sdk.Context, endTime time.Time) []types.NftListing {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(types.KeyPrefixSuccessfulBidListing, endTime)
	it := store.Iterator([]byte(types.KeyPrefixSuccessfulBidListing), storetypes.InclusiveEndBytes(timeKey))
	defer it.Close()

	listings := []types.NftListing{}
	for ; it.Valid(); it.Next() {
		nftIdBytes := it.Value()
		listing, err := k.GetNftListingByIdBytes(ctx, nftIdBytes)
		if err != nil {
			fmt.Println("failed to get listing by id bytes: %w", err)
			continue
		}

		listings = append(listings, listing)
	}
	return listings
}

func (k Keeper) GetAllNftListings(ctx sdk.Context) []types.NftListing {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixNftListing))
	defer it.Close()

	allListings := []types.NftListing{}
	for ; it.Valid(); it.Next() {
		var listing types.NftListing
		k.cdc.MustUnmarshal(it.Value(), &listing)

		allListings = append(allListings, listing)
	}

	return allListings
}

func (k Keeper) ListNft(ctx sdk.Context, msg *types.MsgListNft) error {
	errorMsg := ValidateListNftMsg(k, ctx, msg)
	if errorMsg != nil {
		return errorMsg
	}

	// create listing
	// todo: make test
	owner := k.nftKeeper.GetOwner(ctx, msg.NftId.ClassId, msg.NftId.NftId)
	listing := types.NftListing{
		NftId:                   msg.NftId,
		Owner:                   owner.String(),
		State:                   types.ListingState_LISTING,
		BidDenom:                msg.BidDenom,
		MinimumDepositRate:      msg.MinimumDepositRate,
		StartedAt:               ctx.BlockTime(),
		CollectedAmount:         sdk.NewCoin(msg.BidDenom, sdk.ZeroInt()),
		CollectedAmountNegative: false,
		MinimumBiddingPeriod:    msg.MinimumBiddingPeriod,
	}

	// disable NFT transfer
	data, found := k.nftKeeper.GetNftData(ctx, msg.NftId.ClassId, msg.NftId.NftId)
	if !found {
		return types.ErrNftDoesNotExists
	}
	data.SendDisabled = true
	err := k.nftKeeper.SetNftData(ctx, msg.NftId.ClassId, msg.NftId.NftId, data)
	if err != nil {
		return err
	}

	k.SaveNftListing(ctx, listing)

	// get the memo data from Tx contains MsgListNft
	k.AfterNftListed(ctx, msg.NftId, GetMemo(ctx.TxBytes(), k.txCfg))

	// Emit event for nft listing
	_ = ctx.EventManager().EmitTypedEvent(&types.EventListingNft{
		Owner:   msg.Sender,
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) CancelNftListing(ctx sdk.Context, msg *types.MsgCancelNftListing) error {
	// check listing already exists
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return types.ErrNftListingDoesNotExist
	}

	// // Check nft exists
	_, found := k.nftKeeper.GetNFT(ctx, msg.NftId.ClassId, msg.NftId.NftId)
	if !found {
		return types.ErrNftDoesNotExists
	}

	// check ownership of listing
	if listing.Owner != msg.Sender {
		return types.ErrNotNftListingOwner
	}

	// The listing of items can only be cancelled after N seconds have elapsed from the time it was placed on the marketplace
	params := k.GetParamSet(ctx)
	if listing.StartedAt.Add(time.Duration(params.NftListingCancelRequiredSeconds) * time.Second).After(ctx.BlockTime()) {
		return types.ErrCancelAfterSomeTime
	}

	// check bidding status
	if !listing.IsActive() {
		return types.ErrStatusCannotCancelListing
	}

	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	for _, bid := range bids {
		if bid.IsBorrowed() {
			return types.ErrCannotCancelBorrowedListing
		}
	}

	// enable NFT transfer
	data, found := k.nftKeeper.GetNftData(ctx, msg.NftId.ClassId, msg.NftId.NftId)
	if !found {
		return types.ErrNftDoesNotExists
	}
	data.SendDisabled = false
	err = k.nftKeeper.SetNftData(ctx, msg.NftId.ClassId, msg.NftId.NftId, data)
	if err != nil {
		return err
	}

	// delete listing
	k.DeleteNftListings(ctx, listing)

	// Call AfterNftUnlistedWithoutPayment to delete NFT ID from the ecosystem-incentive KVStore
	// since it's unlisted.
	k.AfterNftUnlistedWithoutPayment(ctx, listing.NftId)

	// Emit event for nft listing cancel
	_ = ctx.EventManager().EmitTypedEvent(&types.EventCancelListingNft{
		Owner:   msg.Sender,
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) HandleFullPaymentsPeriodEndings(ctx sdk.Context) {
	params := k.GetParamSet(ctx)
	// get listings at the end of the payment period
	listings := k.GetFullPaymentNftListingsEndingAt(ctx, ctx.BlockTime())

	// handle not fully paid bids
	for _, listing := range listings {
		bids := types.NftBids(k.GetBidsByNft(ctx, listing.NftId.IdBytes()))
		if listing.State == types.ListingState_SELLING_DECISION {
			err := k.RunSellingDecisionProcess(ctx, bids, listing, params)
			if err != nil {
				fmt.Println("failed to selling decision process: %w", err)
				continue
			}
		} else if listing.State == types.ListingState_LIQUIDATION {
			err := k.RunLiquidationProcess(ctx, bids, listing, params)
			if err != nil {
				fmt.Println("failed to liquidation process: %w", err)
				continue
			}
		}
	}
}

func (k Keeper) DeliverSuccessfulBids(ctx sdk.Context) {
	params := k.GetParamSet(ctx)
	// get listings ended earlier
	listings := k.GetSuccessfulBidNftListingsEndingAt(ctx, ctx.BlockTime())

	_, _ = params, listings
	for _, listing := range listings {
		bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
		if len(bids) != 1 {
			continue
		}
		bid := bids[0]
		bidder, err := sdk.AccAddressFromBech32(bid.Id.Bidder)
		if err != nil {
			continue
		}

		listingOwner, err := sdk.AccAddressFromBech32(listing.Owner)
		if err != nil {
			continue
		}

		listerProfit := sdk.ZeroInt()
		repayAmount := bid.Borrow.Amount.Add(bid.CompoundInterest(listing.LiquidatedAt))
		bidderPaidAmount := bid.Price
		listerProfit = listerProfit.Add(bidderPaidAmount.Amount).Sub(repayAmount.Amount)
		if listing.CollectedAmountNegative {
			listerProfit = listerProfit.Sub(listing.CollectedAmount.Amount)
		} else {
			listerProfit = listerProfit.Add(listing.CollectedAmount.Amount)
		}

		if listerProfit.IsNegative() {
			fmt.Println("lister profit is negative")
			continue
		}
		err = k.ProcessPaymentWithCommissionFee(ctx, listingOwner, sdk.NewCoin(listing.BidDenom, listerProfit), listing.NftId)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = k.DeleteBid(ctx, bid)
		if err != nil {
			fmt.Println(err)
			continue
		}
		k.DeleteNftListings(ctx, listing)

		cacheCtx, write := ctx.CacheContext()

		data, found := k.nftKeeper.GetNftData(ctx, listing.NftId.ClassId, listing.NftId.NftId)
		if !found {
			fmt.Println("nft data not found")
			continue
		}
		data.SendDisabled = false
		err = k.nftKeeper.SetNftData(ctx, listing.NftId.ClassId, listing.NftId.NftId, data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = k.nftKeeper.Transfer(cacheCtx, listing.NftId.ClassId, listing.NftId.NftId, bidder)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			write()
		}
	}
}

func (k Keeper) ProcessPaymentWithCommissionFee(ctx sdk.Context, listingOwner sdk.AccAddress, amount sdk.Coin, nftId types.NftIdentifier) error {
	params := k.GetParamSet(ctx)
	commissionFee := params.NftListingCommissionFee
	cacheCtx, write := ctx.CacheContext()
	// pay commission fees for nft listing
	fee := amount.Amount.Mul(sdk.NewInt(int64(commissionFee))).Quo(sdk.NewInt(100))
	if fee.IsPositive() {
		feeCoins := sdk.Coins{sdk.NewCoin(amount.Denom, fee)}
		err := k.bankKeeper.SendCoinsFromModuleToModule(cacheCtx, types.ModuleName, ecoincentivetypes.ModuleName, feeCoins)
		if err != nil {
			return err
		} else {
			write()
		}
	}

	listerProfit := amount.Amount.Sub(fee)
	if listerProfit.IsPositive() {
		err := k.bankKeeper.SendCoinsFromModuleToAccount(cacheCtx, types.ModuleName, listingOwner, sdk.Coins{sdk.NewCoin(amount.Denom, listerProfit)})
		if err != nil {
			return err
		} else {
			write()
		}
	}

	// Call AfterNftPaymentWithCommission hook method to inform the payment is successfully
	// executed.
	k.AfterNftPaymentWithCommission(ctx, nftId, sdk.NewCoin(amount.Denom, fee))
	return nil
}
