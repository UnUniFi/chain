package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	ecoincentivetypes "github.com/UnUniFi/chain/x/ecosystem-incentive/types"

	"github.com/UnUniFi/chain/x/nftmarketv1/types"
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
		panic(err)
	}
	store.Set(types.NftAddressNftListingKey(owner, nftIdBytes), nftIdBytes)

	if listing.IsActive() {
		store.Set(append(getTimeKey(types.KeyPrefixEndTimeNftListing, listing.EndAt), nftIdBytes...), nftIdBytes)
	} else if listing.IsFullPayment() {
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
		panic(err)
	}
	store.Delete(types.NftAddressNftListingKey(owner, nftIdBytes))

	if listing.IsActive() {
		store.Delete(append(getTimeKey(types.KeyPrefixEndTimeNftListing, listing.EndAt), nftIdBytes...))
	} else if listing.IsFullPayment() {
		store.Delete(append(getTimeKey(types.KeyPrefixFullPaymentPeriodListing, listing.FullPaymentEndAt), nftIdBytes...))
	} else if listing.IsSuccessfulBid() {
		store.Delete(append(getTimeKey(types.KeyPrefixSuccessfulBidListing, listing.SuccessfulBidEndAt), nftIdBytes...))
	}
}

func (k Keeper) GetActiveNftListingsEndingAt(ctx sdk.Context, endTime time.Time) []types.NftListing {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(types.KeyPrefixEndTimeNftListing, endTime)
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
			panic(err)
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

	// todo:delete
	k.TestMint(ctx, msg.Sender.AccAddress(), msg.NftId.ClassId, msg.NftId.NftId)

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
	// todo: delete
	params.BidTokens = append(params.BidTokens, "uguu")
	for !Contains(params.BidTokens, msg.BidToken) {
		return types.ErrNotSupportedBidToken
	}

	// Send ownership to market module
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	err = k.nftKeeper.Transfer(ctx, msg.NftId.ClassId, msg.NftId.NftId, moduleAddr)
	if err != nil {
		return err
	}

	// create listing
	bidActiveRank := msg.BidActiveRank
	if bidActiveRank == 0 {
		bidActiveRank = params.DefaultBidActiveRank
	}
	listing := types.NftListing{
		NftId:         msg.NftId,
		Owner:         owner.String(),
		ListingType:   msg.ListingType,
		State:         types.ListingState_LISTING,
		BidToken:      msg.BidToken,
		MinBid:        msg.MinBid,
		BidActiveRank: bidActiveRank,
		StartedAt:     ctx.BlockTime(),
		EndAt:         ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingPeriodInitial)),
	}
	k.SaveNftListing(ctx, listing)

	// get the memo data from Tx contains MsgListNft
	k.AfterNftListed(ctx, msg.NftId, GetMemo(ctx.TxBytes(), k.txCfg))

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
	if err != nil {
		return types.ErrNftListingDoesNotExist
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

	// The listing of items can only be cancelled after N seconds have elapsed from the time it was placed on the marketplace
	params := k.GetParamSet(ctx)
	if listing.StartedAt.Add(time.Duration(params.NftListingCancelRequiredSeconds) * time.Second).After(ctx.BlockTime()) {
		return types.ErrNotTimeForCancel
	}

	// check nft is bidding status
	if !listing.IsActive() {
		return types.ErrStatusCannotCancelListing
	}

	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())

	winnerCandidateStartIndex := len(bids) - int(listing.BidActiveRank)
	if winnerCandidateStartIndex < 0 {
		winnerCandidateStartIndex = 0
	}
	// distribute cancellation fee to winner bidders
	for _, bid := range bids[winnerCandidateStartIndex:] {
		bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
		if err != nil {
			return err
		}
		cancelFee := bid.Amount.Amount.Mul(sdk.NewInt(int64(params.NftListingCancelFeePercentage))).Quo(sdk.NewInt(100))
		if cancelFee.IsPositive() {
			err = k.bankKeeper.SendCoins(ctx, msg.Sender.AccAddress(), bidder, sdk.Coins{sdk.NewCoin(listing.BidToken, cancelFee)})
			if err != nil {
				return err
			}
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

	// Send ownership to original owner
	err = k.nftKeeper.Transfer(ctx, msg.NftId.ClassId, msg.NftId.NftId, msg.Sender.AccAddress())
	if err != nil {
		return err
	}

	// delete listing
	k.DeleteNftListings(ctx, listing)

	// Call AfterNftUnlistedWithoutPayment to delete NFT ID from the ecosystem-incentive KVStore
	// since it's unlisted.
	k.AfterNftUnlistedWithoutPayment(ctx, listing.NftId)

	// Emit event for nft listing cancel
	ctx.EventManager().EmitTypedEvent(&types.EventCancelListNfting{
		Owner:   msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) ExpandListingPeriod(ctx sdk.Context, msg *types.MsgExpandListingPeriod) error {
	// check listing already exists
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return types.ErrNftListingDoesNotExist
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

	// check nft is bidding status
	if !listing.IsActive() {
		return types.ErrListingIsNotInStatusToBid
	}

	// pay nft listing extend fee
	params := k.GetParamSet(ctx)
	feeAmount := params.NftListingPeriodExtendFeePerHour.Amount.Mul(sdk.NewInt(int64(params.NftListingExtendSeconds))).Quo(sdk.NewInt(3600))

	// distribute nft listing extend fee to winner bidders
	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	totalBidAmount := sdk.ZeroInt()

	winnerCandidateStartIndex := len(bids) - int(listing.BidActiveRank)
	if winnerCandidateStartIndex < 0 {
		winnerCandidateStartIndex = 0
	}

	for _, bid := range bids[winnerCandidateStartIndex:] {
		totalBidAmount = totalBidAmount.Add(bid.Amount.Amount)
	}

	if totalBidAmount.IsPositive() {
		for _, bid := range bids[winnerCandidateStartIndex:] {
			bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
			if err != nil {
				return err
			}
			bidderCommission := bid.Amount.Amount.Mul(feeAmount).Quo(totalBidAmount)
			if bidderCommission.IsPositive() {
				commmission := sdk.NewCoin(params.NftListingPeriodExtendFeePerHour.Denom, bidderCommission)
				err = k.bankKeeper.SendCoins(ctx, msg.Sender.AccAddress(), bidder, sdk.Coins{commmission})
				if err != nil {
					return err
				}
			}
		}
	}

	// update listing end time
	listing.EndAt = listing.EndAt.Add(time.Second * time.Duration(params.NftListingExtendSeconds))
	k.SaveNftListing(ctx, listing)

	// Emit event for nft listing cancel
	ctx.EventManager().EmitTypedEvent(&types.EventExpandListingPeriod{
		Owner:   msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) SellingDecision(ctx sdk.Context, msg *types.MsgSellingDecision) error {
	// check listing already exists
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return types.ErrNftListingDoesNotExist
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

	// check if listing is already ended or on selling decision status
	if listing.State != types.ListingState_BIDDING {
		return types.ErrListingNeedsToBeBiddingStatus
	}

	params := k.GetParamSet(ctx)
	listing.FullPaymentEndAt = ctx.BlockTime().Add(time.Duration(params.NftListingFullPaymentPeriod) * time.Second)
	listing.State = types.ListingState_SELLING_DECISION
	k.SaveNftListing(ctx, listing)

	// automatic payment if enabled
	bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
	if len(bids) > 0 {
		winnerIndex := len(bids) - 1
		bid := bids[winnerIndex]
		if bid.AutomaticPayment {
			bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
			if err != nil {
				fmt.Println(err)
				return err
			}

			cacheCtx, write := ctx.CacheContext()
			err = k.PayFullBid(cacheCtx, &types.MsgPayFullBid{
				Sender: bidder.Bytes(),
				NftId:  listing.NftId,
			})
			if err == nil {
				write()
			} else {
				fmt.Println(err)
			}
		}
	}

	// Emit event for nft listing end
	ctx.EventManager().EmitTypedEvent(&types.EventSellingDecision{
		Owner:   msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

func (k Keeper) EndNftListing(ctx sdk.Context, msg *types.MsgEndNftListing) error {
	// check listing already exists
	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
	if err != nil {
		return types.ErrNftListingDoesNotExist
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

	// check if listing is already ended
	if listing.State == types.ListingState_END_LISTING || listing.State == types.ListingState_SELLING_DECISION {
		return types.ErrListingAlreadyEnded
	}

	bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
	if len(bids) == 0 {
		err = k.nftKeeper.Transfer(ctx, listing.NftId.ClassId, listing.NftId.NftId, msg.Sender.AccAddress())
		if err != nil {
			panic(err)
		}
		k.DeleteNftListings(ctx, listing)
	} else {
		params := k.GetParamSet(ctx)
		listing.FullPaymentEndAt = ctx.BlockTime().Add(time.Duration(params.NftListingFullPaymentPeriod) * time.Second)
		listing.State = types.ListingState_END_LISTING
		k.SaveNftListing(ctx, listing)

		// automatic payment after listing ends
		winnerCandidateStartIndex := len(bids) - int(listing.BidActiveRank)
		if winnerCandidateStartIndex < 0 {
			winnerCandidateStartIndex = 0
		}
		for _, bid := range bids[winnerCandidateStartIndex:] {
			if bid.AutomaticPayment {
				bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
				if err != nil {
					fmt.Println(err)
					continue
				}

				cacheCtx, write := ctx.CacheContext()
				err = k.PayFullBid(cacheCtx, &types.MsgPayFullBid{
					Sender: bidder.Bytes(),
					NftId:  listing.NftId,
				})
				if err == nil {
					write()
				} else {
					fmt.Println(err)
					continue
				}
			}
		}

		// automatically cancel bids for not active rank
		for _, bid := range bids[:winnerCandidateStartIndex] {
			bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
			if err != nil {
				fmt.Println(err)
				continue
			}
			// Delete bid
			k.DeleteBid(ctx, bid)
			cacheCtx, write := ctx.CacheContext()
			err = k.bankKeeper.SendCoinsFromModuleToAccount(cacheCtx, types.ModuleName, bidder, sdk.Coins{sdk.NewCoin(bid.Amount.Denom, bid.PaidAmount)})
			if err != nil {
				return err
			}
			if err == nil {
				write()
			} else {
				fmt.Println(err)
				continue
			}
		}
	}

	// Emit event for nft listing end
	ctx.EventManager().EmitTypedEvent(&types.EventEndListNfting{
		Owner:   msg.Sender.AccAddress().String(),
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	// Call AfterNftUnlistedWithoutPayment to delete NFT ID from the ecosystem-incentive KVStore
	// since it's unlisted.
	if _, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes()); err != nil {
		k.AfterNftUnlistedWithoutPayment(ctx, listing.NftId)
	}

	return nil
}

func (k Keeper) ProcessEndingNftListings(ctx sdk.Context) {
	params := k.GetParamSet(ctx)
	listings := k.GetActiveNftListingsEndingAt(ctx, ctx.BlockTime())
	for _, listing := range listings {
		bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
		if listing.AutoRelistedCount < params.AutoRelistingCountIfNoBid && len(bids) == 0 {
			listing.EndAt = listing.EndAt.Add(time.Duration(params.NftListingExtendSeconds) * time.Second)
			listing.AutoRelistedCount++
			k.SaveNftListing(ctx, listing)
		} else {
			listingOwner, err := sdk.AccAddressFromBech32(listing.Owner)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = k.EndNftListing(ctx, &types.MsgEndNftListing{
				Sender: listingOwner.Bytes(),
				NftId:  listing.NftId,
			})
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func (k Keeper) HandleFullPaymentsPeriodEndings(ctx sdk.Context) {
	params := k.GetParamSet(ctx)
	// get listings ended earlier
	listings := k.GetFullPaymentNftListingsEndingAt(ctx, ctx.BlockTime())

	// handle not fully paid bids
	for _, listing := range listings {
		bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
		if listing.State == types.ListingState_SELLING_DECISION {
			i := len(bids) - 1
			bid := bids[i]

			// if winner bidder did not pay full bid, nft is listed again after deleting winner bidder
			if bid.PaidAmount.LT(bid.Amount.Amount) {
				k.DeleteBid(ctx, bid)
				if len(bids) == 1 {
					listing.State = types.ListingState_LISTING
				} else {
					listing.State = types.ListingState_BIDDING
				}
				listing.EndAt = ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingExtendSeconds))

				// Reset the loan data for a lister
				// If the bid.PaidAmount is more than loan.Coin.Amount, then just delete the loan data for lister.
				// Otherwise, subtract bid.PaidAmount from loaning amount
				loan := k.GetDebtByNft(ctx, listing.IdBytes())
				if !loan.Loan.Amount.IsNil() {
					if loan.Loan.Amount.LTE(bid.PaidAmount) {
						k.DeleteDebt(ctx, listing.IdBytes())
					} else {
						renewedLoanAmount := loan.Loan.Amount.Sub(bid.PaidAmount)
						loan.Loan.Amount = renewedLoanAmount
						k.SetDebt(ctx, loan)
					}
				}
			} else {
				// schedule NFT / token send after X days
				listing.SuccessfulBidEndAt = ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingNftDeliveryPeriod))
				listing.State = types.ListingState_SUCCESSFUL_BID
				// delete the loan data for the nftId which is deleted from the market
				k.RemoveDebt(ctx, listing.IdBytes())
			}
			k.SaveNftListing(ctx, listing)
		} else if listing.State == types.ListingState_END_LISTING {
			index := len(bids) - 1
			for ; index >= 0; index-- {
				bid := bids[index]
				if bid.PaidAmount.Equal(bid.Amount.Amount) {
					break
				}
			}

			if index >= 0 { // if winner bidder exists who paid full amount
				// schedule NFT / token send after X days
				listing.SuccessfulBidEndAt = ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingNftDeliveryPeriod))
				listing.State = types.ListingState_SUCCESSFUL_BID
				k.SaveNftListing(ctx, listing)

				for i, bid := range bids {
					if index != i {
						cacheCtx, write := ctx.CacheContext()
						err := k.SafeCloseBid(cacheCtx, bid)
						if err == nil {
							write()
						} else {
							fmt.Println(err)
						}
					}
				}
				// TODO: shouldn't we handle winning bidder candidates above successful bidder that didn't pay full amount?
			} else { // if all winning bidder candidates do not pay
				// the amount of the collected deposit plus NFT to be listed will be given to the lister
				listingOwner, err := sdk.AccAddressFromBech32(listing.Owner)
				if err != nil {
					continue
				}

				depositCollected := sdk.ZeroInt()
				for _, bid := range bids {
					depositCollected = depositCollected.Add(bid.PaidAmount)
					k.DeleteBid(ctx, bid)
				}

				// pay fee
				loan := k.GetDebtByNft(ctx, listing.IdBytes())
				k.ProcessPaymentWithCommissionFee(ctx, listingOwner, listing.BidToken, depositCollected, loan.Loan.Amount, listing.NftId)

				// transfer nft to listing owner
				cacheCtx, write := ctx.CacheContext()
				err = k.nftKeeper.Transfer(cacheCtx, listing.NftId.ClassId, listing.NftId.NftId, listingOwner)
				if err != nil {
					fmt.Println(err)
				} else {
					write()
				}

				// remove listing
				k.DeleteNftListings(ctx, listing)
			}
			// delete the loan data for the nftId which is deleted from the market anyway
			k.RemoveDebt(ctx, listing.IdBytes())
		}
	}
}

func (k Keeper) DelieverSuccessfulBids(ctx sdk.Context) {
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
		bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
		if err != nil {
			continue
		}

		listingOwner, err := sdk.AccAddressFromBech32(listing.Owner)
		if err != nil {
			continue
		}

		cacheCtx, write := ctx.CacheContext()
		err = k.nftKeeper.Transfer(cacheCtx, listing.NftId.ClassId, listing.NftId.NftId, bidder)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			write()
		}

		loan := k.GetDebtByNft(ctx, listing.IdBytes())
		k.ProcessPaymentWithCommissionFee(ctx, listingOwner, bid.Amount.Denom, bid.PaidAmount, loan.Loan.Amount, listing.NftId)

		k.DeleteBid(ctx, bid)
		k.DeleteNftListings(ctx, listing)
	}
}

func (k Keeper) ProcessPaymentWithCommissionFee(ctx sdk.Context, listingOwner sdk.AccAddress, denom string, amount sdk.Int, loanAmount sdk.Int, nftId types.NftIdentifier) {
	params := k.GetParamSet(ctx)
	commissionFee := params.NftListingCommissionFee
	cacheCtx, write := ctx.CacheContext()
	// pay commission fees for nft listing
	fee := amount.Mul(sdk.NewInt(int64(commissionFee))).Quo(sdk.NewInt(100))
	if fee.IsPositive() {
		feeCoins := sdk.Coins{sdk.NewCoin(denom, fee)}
		err := k.bankKeeper.SendCoinsFromModuleToModule(cacheCtx, types.ModuleName, ecoincentivetypes.ModuleName, feeCoins)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			write()
		}
	}

	if loanAmount.IsNil() {
		loanAmount = sdk.ZeroInt()
	}
	listerPayment := amount.Sub(fee)
	listerPayment = listerPayment.Sub(loanAmount)
	if !listerPayment.IsZero() {
		err := k.bankKeeper.SendCoinsFromModuleToAccount(cacheCtx, types.ModuleName, listingOwner, sdk.Coins{sdk.NewCoin(denom, listerPayment)})
		if err != nil {
			fmt.Println(err)
			return
		} else {
			write()
		}
	}

	// Call AfterNftPaymentWithCommission hook method to inform the payment is successfuly
	// executed.
	k.AfterNftPaymentWithCommission(ctx, nftId, sdk.NewCoin(denom, fee))
}

// todo: delete
func (k Keeper) TestMint(ctx sdk.Context, addr sdk.AccAddress, classId, nftId string) {
	_, exists := k.nftKeeper.GetNFT(ctx, classId, nftId)
	if exists {
		return
	}
	const (
		testClassName        = "Crypto Kitty"
		testClassSymbol      = "kitty"
		testClassDescription = "Crypto Kitty"
		testClassURI         = "class uri"
		testClassURIHash     = "ae702cefd6b6a65fe2f991ad6d9969ed"
		testURI              = "kitty uri"
		testURIHash          = "229bfd3c1b431c14a526497873897108"
	)

	_, hasId := k.nftKeeper.GetClass(ctx, classId)
	if !hasId {
		class := nft.Class{
			Id:          classId,
			Name:        testClassName,
			Symbol:      testClassSymbol,
			Description: testClassDescription,
			Uri:         testClassURI,
			UriHash:     testClassURIHash,
		}
		k.nftKeeper.SaveClass(ctx, class)
		fmt.Println("save class")
	}

	expNFT := nft.NFT{
		ClassId: classId,
		Id:      nftId,
		Uri:     testURI,
	}
	err := k.nftKeeper.Mint(ctx, expNFT, addr)
	if err != nil {
		fmt.Println("err occur")
	}
}
