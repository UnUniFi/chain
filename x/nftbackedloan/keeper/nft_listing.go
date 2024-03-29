package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	ecoincentivetypes "github.com/UnUniFi/chain/x/ecosystemincentive/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) GetListedNftByIdBytes(ctx sdk.Context, nftIdBytes []byte) (types.Listing, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NftListingKey(nftIdBytes))
	if bz == nil {
		return types.Listing{}, types.ErrListedNftDoesNotExist
	}
	listing := types.Listing{}
	k.cdc.MustUnmarshal(bz, &listing)
	return listing, nil
}

func (k Keeper) GetListingsByOwner(ctx sdk.Context, owner sdk.AccAddress) []types.Listing {
	store := ctx.KVStore(k.storeKey)

	listings := []types.Listing{}
	it := sdk.KVStorePrefixIterator(store, types.NftAddressNftListingPrefixKey(owner))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		nftIdBytes := it.Value()
		listing, err := k.GetListedNftByIdBytes(ctx, nftIdBytes)
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
func (k Keeper) SaveListedNft(ctx sdk.Context, listing types.Listing) {
	k.SetListedNft(ctx, listing)
	k.UpdateListedClass(ctx, listing)
}

func (k Keeper) SetListedNft(ctx sdk.Context, listing types.Listing) {
	if oldListing, err := k.GetListedNftByIdBytes(ctx, listing.IdBytes()); err == nil {
		k.DeleteListedNfts(ctx, oldListing)
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
func (k Keeper) DeleteListedNfts(ctx sdk.Context, listing types.Listing) {
	k.DeleteListedNft(ctx, listing)
	k.UpdateListedClass(ctx, listing)
}

func (k Keeper) DeleteListedNft(ctx sdk.Context, listing types.Listing) {
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

func (k Keeper) GetFullPaymentListedNftsEndingAt(ctx sdk.Context, endTime time.Time) []types.Listing {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(types.KeyPrefixFullPaymentPeriodListing, endTime)
	it := store.Iterator([]byte(types.KeyPrefixFullPaymentPeriodListing), storetypes.InclusiveEndBytes(timeKey))
	defer it.Close()

	listings := []types.Listing{}
	for ; it.Valid(); it.Next() {
		nftIdBytes := it.Value()
		listing, err := k.GetListedNftByIdBytes(ctx, nftIdBytes)
		if err != nil {
			fmt.Println("failed to get listing by id bytes: %w", err)
			continue
		}

		listings = append(listings, listing)
	}
	return listings
}

func (k Keeper) GetSuccessfulBidListedNftsEndingAt(ctx sdk.Context, endTime time.Time) []types.Listing {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(types.KeyPrefixSuccessfulBidListing, endTime)
	it := store.Iterator([]byte(types.KeyPrefixSuccessfulBidListing), storetypes.InclusiveEndBytes(timeKey))
	defer it.Close()

	listings := []types.Listing{}
	for ; it.Valid(); it.Next() {
		nftIdBytes := it.Value()
		listing, err := k.GetListedNftByIdBytes(ctx, nftIdBytes)
		if err != nil {
			fmt.Println("failed to get listing by id bytes: %w", err)
			continue
		}

		listings = append(listings, listing)
	}
	return listings
}

func (k Keeper) GetAllListedNfts(ctx sdk.Context) []types.Listing {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixNftListing))
	defer it.Close()

	allListings := []types.Listing{}
	for ; it.Valid(); it.Next() {
		var listing types.Listing
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
	owner := k.nftKeeper.GetOwner(ctx, msg.NftId.ClassId, msg.NftId.TokenId)
	listing := types.Listing{
		NftId:                   msg.NftId,
		Owner:                   owner.String(),
		State:                   types.ListingState_LISTING,
		BidDenom:                msg.BidDenom,
		MinDepositRate:          msg.MinDepositRate,
		StartedAt:               ctx.BlockTime(),
		CollectedAmount:         sdk.NewCoin(msg.BidDenom, sdk.ZeroInt()),
		CollectedAmountNegative: false,
		MinBidPeriod:            msg.MinBidPeriod,
	}

	// disable NFT transfer
	data, found := k.nftKeeper.GetNftData(ctx, msg.NftId.ClassId, msg.NftId.TokenId)
	if !found {
		return types.ErrNftDoesNotExists
	}
	data.SendDisabled = true
	err := k.nftKeeper.SetNftData(ctx, msg.NftId.ClassId, msg.NftId.TokenId, data)
	if err != nil {
		return err
	}

	k.SaveListedNft(ctx, listing)

	// get the memo data from Tx contains MsgListNft
	// k.AfterNftListed(ctx, msg.NftId, GetMemo(ctx.TxBytes(), k.txCfg))

	// Emit event for nft listing
	_ = ctx.EventManager().EmitTypedEvent(&types.EventListNft{
		Owner:   msg.Sender,
		ClassId: msg.NftId.ClassId,
		TokenId: msg.NftId.TokenId,
	})

	return nil
}

func (k Keeper) CancelNftListing(ctx sdk.Context, msg *types.MsgCancelListing) error {
	// check listing already exists
	listing, err := k.GetListedNftByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return types.ErrListedNftDoesNotExist
	}

	// // Check nft exists
	_, found := k.nftKeeper.GetNFT(ctx, msg.NftId.ClassId, msg.NftId.TokenId)
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
	data, found := k.nftKeeper.GetNftData(ctx, msg.NftId.ClassId, msg.NftId.TokenId)
	if !found {
		return types.ErrNftDoesNotExists
	}
	data.SendDisabled = false
	err = k.nftKeeper.SetNftData(ctx, msg.NftId.ClassId, msg.NftId.TokenId, data)
	if err != nil {
		return err
	}

	// delete listing
	k.DeleteListedNfts(ctx, listing)

	// Call AfterNftUnlistedWithoutPayment to delete NFT ID from the ecosystem-incentive KVStore
	// since it's unlisted.
	// k.AfterNftUnlistedWithoutPayment(ctx, listing.NftId)

	// Emit event for nft listing cancel
	_ = ctx.EventManager().EmitTypedEvent(&types.EventCancelListing{
		Owner:   msg.Sender,
		ClassId: msg.NftId.ClassId,
		TokenId: msg.NftId.TokenId,
	})

	return nil
}

func (k Keeper) HandleFullPaymentsPeriodEndings(ctx sdk.Context) {
	params := k.GetParamSet(ctx)
	// get listings at the end of the payment period
	listings := k.GetFullPaymentListedNftsEndingAt(ctx, ctx.BlockTime())

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
	listings := k.GetSuccessfulBidListedNftsEndingAt(ctx, ctx.BlockTime())

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
		repayAmount := bid.Loan.Amount.Add(bid.CompoundInterest(listing.LiquidatedAt))
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
		k.DeleteListedNfts(ctx, listing)

		cacheCtx, write := ctx.CacheContext()

		data, found := k.nftKeeper.GetNftData(ctx, listing.NftId.ClassId, listing.NftId.TokenId)
		if !found {
			fmt.Println("nft data not found")
			continue
		}
		data.SendDisabled = false
		err = k.nftKeeper.SetNftData(ctx, listing.NftId.ClassId, listing.NftId.TokenId, data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = k.nftKeeper.Transfer(cacheCtx, listing.NftId.ClassId, listing.NftId.TokenId, bidder)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			write()
		}
	}
}

func (k Keeper) ProcessPaymentWithCommissionFee(ctx sdk.Context, listingOwner sdk.AccAddress, amount sdk.Coin, nftId types.NftId) error {
	params := k.GetParamSet(ctx)
	commissionRate := params.NftListingCommissionRate
	cacheCtx, write := ctx.CacheContext()
	// pay commission fees for nft listing
	fee := sdk.NewDecFromInt(amount.Amount).Mul(commissionRate).RoundInt()
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
	// k.AfterNftPaymentWithCommission(ctx, nftId, sdk.NewCoin(amount.Denom, fee))
	return nil
}
