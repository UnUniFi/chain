package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
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
			panic(err)
		}

		listings = append(listings, listing)
	}
	return listings
}

func getTimeKey(timestamp time.Time) []byte {
	timeBz := sdk.FormatTimeBytes(timestamp)
	timeBzL := len(timeBz)
	prefixL := len(types.KeyPrefixEndTimeNftListing)

	bz := make([]byte, prefixL+8+timeBzL)

	// copy the prefix
	copy(bz[:prefixL], types.KeyPrefixEndTimeNftListing)

	// copy the encoded time bytes length
	copy(bz[prefixL:prefixL+8], sdk.Uint64ToBigEndian(uint64(timeBzL)))

	// copy the encoded time bytes
	copy(bz[prefixL+8:prefixL+8+timeBzL], timeBz)
	return bz
}

func (k Keeper) SetNftListing(ctx sdk.Context, listing types.NftListing) {
	if oldListing, err := k.GetNftListingByIdBytes(ctx, listing.IdBytes()); err == nil {
		k.DeleteNftListing(ctx, oldListing)
	}

	nftIdBytes := listing.IdBytes()
	bz := k.cdc.MustMarshal(&listing)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NftListingKey(nftIdBytes), bz)

	owner, err := sdk.AccAddressFromBech32(listing.Owner)
	if err != nil {
		panic(err)
	}
	store.Set(types.NftAddressNftListingKey(owner, nftIdBytes), nftIdBytes)

	if listing.IsActive() {
		store.Set(append(getTimeKey(listing.EndAt), nftIdBytes...), nftIdBytes)
	}
}

func (k Keeper) DeleteNftListing(ctx sdk.Context, listing types.NftListing) {
	nftIdBytes := listing.IdBytes()
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NftListingKey(nftIdBytes))

	owner, err := sdk.AccAddressFromBech32(listing.Owner)
	if err != nil {
		panic(err)
	}
	store.Delete(types.NftAddressNftListingKey(owner, nftIdBytes))

	if listing.IsActive() {
		store.Delete(getTimeKey(listing.EndAt))
	}
}

func (k Keeper) GetActiveNftListingsEndingAt(ctx sdk.Context, endTime time.Time) []types.NftListing {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(endTime)
	it := store.Iterator([]byte(types.KeyPrefixEndTimeNftListing), storetypes.InclusiveEndBytes(timeKey))
	defer it.Close()

	listings := []types.NftListing{}
	for ; it.Valid(); it.Next() {
		nftIdBytes := it.Value()
		listing, err := k.GetNftListingByIdBytes(ctx, nftIdBytes)
		if err != nil {
			panic(err)
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
	// check listing already exists
	_, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err == nil {
		return types.ErrNftListingAlreadyExists
	}

	// Check nft exists
	_, found := k.nftKeeper.GetNFT(ctx, msg.NftId.ClassId, msg.NftId.NftId)
	if !found {
		return types.ErrNftDoesNotExists
	}

	// check ownership of nft
	owner := k.nftKeeper.GetOwner(ctx, msg.NftId.ClassId, msg.NftId.NftId)
	if owner.String() != msg.Sender.AccAddress().String() {
		return types.ErrNotNftOwner
	}

	params := k.GetParamSet(ctx)
	for !Contains(params.BidTokens, msg.BidToken) {
		return types.ErrNotSupportedBidToken
	}

	// pay fees for nft listing
	listingFee := params.NftListingCommissionFee
	if listingFee.IsPositive() {
		feeCoins := sdk.Coins{listingFee}
		sender := sdk.AccAddress(msg.Sender)
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, feeCoins)
		if err != nil {
			return err
		}
	}

	// Send ownership to market module
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	err = k.nftKeeper.Transfer(ctx, msg.NftId.ClassId, msg.NftId.NftId, moduleAddr)
	if err != nil {
		return err
	}

	// create listing
	listing := types.NftListing{
		NftId:       msg.NftId,
		Owner:       owner.String(),
		ListingType: msg.ListingType,
		State:       types.ListingState_SELLING,
		BidToken:    msg.BidToken,
		MinBid:      msg.MinBid,
		BidHook:     msg.BidHook,
		EndAt:       ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingPeriodInitial)),
	}
	k.SetNftListing(ctx, listing)

	// Emit event for nft listing
	ctx.EventManager().EmitTypedEvent(&types.EventListNft{
		Owner:   msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) CancelNftListing(ctx sdk.Context, msg *types.MsgCancelNftListing) error {
	// check listing already exists
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err == nil {
		return types.ErrNftListingAlreadyExists
	}

	// Check nft exists
	_, found := k.nftKeeper.GetNFT(ctx, msg.NftId.ClassId, msg.NftId.NftId)
	if !found {
		return types.ErrNftDoesNotExists
	}

	// check ownership of listing
	if listing.Owner != msg.Sender.AccAddress().String() {
		return types.ErrNotNftListingOwner
	}

	// check bids exists
	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	if len(bids) > 0 {
		return types.ErrNftBidAlreadyExists
	}

	// Send ownership to original owner
	err = k.nftKeeper.Transfer(ctx, msg.NftId.ClassId, msg.NftId.NftId, msg.Sender.AccAddress())
	if err != nil {
		return err
	}

	// delete listing
	k.DeleteNftListing(ctx, listing)

	// Emit event for nft listing cancel
	ctx.EventManager().EmitTypedEvent(&types.EventCancelListNfting{
		Owner:   msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) NftBuyBack(ctx sdk.Context, msg *types.MsgNftBuyBack) error {
	// check listing already exists
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err == nil {
		return types.ErrNftListingAlreadyExists
	}

	// Check nft exists
	_, found := k.nftKeeper.GetNFT(ctx, msg.NftId.ClassId, msg.NftId.NftId)
	if !found {
		return types.ErrNftDoesNotExists
	}

	// check ownership of listing
	if listing.Owner != msg.Sender.AccAddress().String() {
		return types.ErrNotNftListingOwner
	}

	// check bids exists
	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	if len(bids) > 0 {
		return types.ErrNftBidDoesNotExists
	}

	// send extra buy back funds to winner bidders
	params := k.GetParamSet(ctx)
	for _, bid := range bids[len(bids)-int(listing.BidHook):] {
		bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
		if err != nil {
			return err
		}
		bidderCancelFee := bid.Amount.Amount.Int64() * int64(params.NftListingBuyBackExtraPercentage) / 100
		err = k.bankKeeper.SendCoins(ctx, msg.Sender.AccAddress(), bidder, sdk.Coins{sdk.NewInt64Coin(listing.BidToken, bidderCancelFee)})
		if err != nil {
			return err
		}
	}

	// delete all bids and return funds back
	for _, bid := range bids {
		bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
		if err != nil {
			return err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{bid.Amount})
		if err != nil {
			return err
		}
		k.DeleteBid(ctx, bid)
	}

	// Send nft ownership to original owner
	err = k.nftKeeper.Transfer(ctx, msg.NftId.ClassId, msg.NftId.NftId, msg.Sender.AccAddress())
	if err != nil {
		return err
	}

	// delete listing
	k.DeleteNftListing(ctx, listing)

	// Emit event for nft listing cancel
	ctx.EventManager().EmitTypedEvent(&types.EventNftBuyBack{
		Owner:   msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}
