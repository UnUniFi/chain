package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"

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
		fmt.Println("invalid owner address: %w", err)
		return
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
			fmt.Println("failed to get listing by id bytes: %w", err)
			continue
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
	errorMsg := validateListNftMsg(k, ctx, msg)
	if errorMsg != nil {
		return errorMsg
	}

	// create listing
	// todo: make test
	owner := k.nftKeeper.GetOwner(ctx, msg.NftId.ClassId, msg.NftId.NftId)
	listing := types.NftListing{
		NftId:                msg.NftId,
		Owner:                owner.String(),
		ListingType:          msg.ListingType,
		State:                types.ListingState_LISTING,
		BidToken:             msg.BidToken,
		MinimumDepositRate:   msg.MinimumDepositRate,
		AutomaticRefinancing: msg.AutomaticRefinancing,
		StartedAt:            ctx.BlockTime(),
		CollectedAmount:      sdk.NewCoin(msg.BidToken, sdk.ZeroInt()),
		// todo: add validation.
		// we should to determine maximum bidding period.
		MinimumBiddingPeriod: msg.MinimumBiddingPeriod,
	}
	k.SaveNftListing(ctx, listing)

	// Send ownership to market module
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	err := k.nftKeeper.Transfer(ctx, msg.NftId.ClassId, msg.NftId.NftId, moduleAddr)
	if err != nil {
		k.DeleteNftListing(ctx, listing)
		return err
	}

	// Emit event for nft listing
	ctx.EventManager().EmitTypedEvent(&types.EventListNft{
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
		return types.ErrNotTimeForCancel
	}

	// // check nft is bidding status
	if !listing.IsActive() {
		return types.ErrStatusCannotCancelListing
	}

	currDebt := k.GetDebtByNft(ctx, msg.NftId.IdBytes())

	if !currDebt.Loan.IsNil() {
		return types.ErrCannotCancelListingWithDebt
	}

	// todo implement bid exist cancel listing
	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
	if len(bids) > 0 {
		return types.ErrCannotCancelListingWithBids
	}

	// winnerCandidateStartIndex := len(bids) - int(listing.BidActiveRank)
	// if winnerCandidateStartIndex < 0 {
	// 	winnerCandidateStartIndex = 0
	// }
	// // distribute cancellation fee to winner bidders
	// for _, bid := range bids[winnerCandidateStartIndex:] {
	// 	bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	cancelFee := bid.Amount.Amount.Mul(sdk.NewInt(int64(params.NftListingCancelFeePercentage))).Quo(sdk.NewInt(100))
	// 	if cancelFee.IsPositive() {
	// 		err = k.bankKeeper.SendCoins(ctx, msg.Sender.AccAddress(), bidder, sdk.Coins{sdk.NewCoin(listing.BidToken, cancelFee)})
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// // delete all bids and return funds back
	// for _, bid := range bids {
	// 	bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bidder, sdk.Coins{bid.Amount})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	k.DeleteBid(ctx, bid)
	// }

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	// Send ownership to original owner
	err = k.nftKeeper.Transfer(ctx, msg.NftId.ClassId, msg.NftId.NftId, sender)
	if err != nil {
		return err
	}

	// delete listing
	k.DeleteNftListings(ctx, listing)

	// // Emit event for nft listing cancel
	ctx.EventManager().EmitTypedEvent(&types.EventCancelListNfting{
		Owner:   msg.Sender,
		ClassId: msg.NftId.ClassId,
		NftId:   msg.NftId.NftId,
	})

	return nil
}

// func (k Keeper) ExpandListingPeriod(ctx sdk.Context, msg *types.MsgExpandListingPeriod) error {
// 	// check listing already exists
// 	listing, err := k.GetNftListingByIdBytes(ctx, msg.NftId.IdBytes())
// 	if err != nil {
// 		return types.ErrNftListingDoesNotExist
// 	}

// 	// Check nft exists
// 	_, found := k.nftKeeper.GetNFT(ctx, msg.NftId.ClassId, msg.NftId.NftId)
// 	if !found {
// 		return types.ErrNftDoesNotExists
// 	}

// 	// check ownership of listing
// 	if listing.Owner != msg.Sender.AccAddress().String() {
// 		return types.ErrNotNftListingOwner
// 	}

// 	// check nft is bidding status
// 	if !listing.IsActive() {
// 		return types.ErrListingIsNotInStatusToBid
// 	}

// 	// pay nft listing extend fee
// 	params := k.GetParamSet(ctx)
// 	feeAmount := params.NftListingPeriodExtendFeePerHour.Amount.Mul(sdk.NewInt(int64(params.NftListingExtendSeconds))).Quo(sdk.NewInt(3600))

// 	// distribute nft listing extend fee to winner bidders
// 	bids := k.GetBidsByNft(ctx, msg.NftId.IdBytes())
// 	totalBidAmount := sdk.ZeroInt()

// 	winnerCandidateStartIndex := len(bids) - int(listing.BidActiveRank)
// 	if winnerCandidateStartIndex < 0 {
// 		winnerCandidateStartIndex = 0
// 	}

// 	for _, bid := range bids[winnerCandidateStartIndex:] {
// 		totalBidAmount = totalBidAmount.Add(bid.Amount.Amount)
// 	}

// 	if totalBidAmount.IsPositive() {
// 		for _, bid := range bids[winnerCandidateStartIndex:] {
// 			bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
// 			if err != nil {
// 				return err
// 			}
// 			bidderCommission := bid.Amount.Amount.Mul(feeAmount).Quo(totalBidAmount)
// 			if bidderCommission.IsPositive() {
// 				commmission := sdk.NewCoin(params.NftListingPeriodExtendFeePerHour.Denom, bidderCommission)
// 				err = k.bankKeeper.SendCoins(ctx, msg.Sender.AccAddress(), bidder, sdk.Coins{commmission})
// 				if err != nil {
// 					return err
// 				}
// 			}
// 		}
// 	}

// 	// update listing end time
// 	listing.EndAt = listing.EndAt.Add(time.Second * time.Duration(params.NftListingExtendSeconds))
// 	k.SaveNftListing(ctx, listing)

// 	// Emit event for nft listing cancel
// 	ctx.EventManager().EmitTypedEvent(&types.EventExpandListingPeriod{
// 		Owner:   msg.Sender.AccAddress().String(),
// 		ClassId: msg.NftId.ClassId,
// 		NftId:   msg.NftId.NftId,
// 	})

// 	return nil
// }

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
	if listing.Owner != msg.Sender {
		return types.ErrNotNftListingOwner
	}

	// check if listing is already ended or on selling decision status
	if listing.State != types.ListingState_BIDDING {
		return types.ErrListingNeedsToBeBiddingStatus
	}

	// check bid exists
	bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
	if len(bids) == 0 {
		return types.ErrNotExistsBid
	}

	params := k.GetParamSet(ctx)
	listing.FullPaymentEndAt = ctx.BlockTime().Add(time.Duration(params.NftListingFullPaymentPeriod) * time.Second)
	listing.State = types.ListingState_SELLING_DECISION
	k.SaveNftListing(ctx, listing)

	// automatic payment if enabled
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
				Sender: bidder.String(),
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
		Owner:   msg.Sender,
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
	if listing.Owner != msg.Sender {
		return types.ErrNotNftListingOwner
	}

	// check if listing is already ended
	if listing.State == types.ListingState_END_LISTING || listing.State == types.ListingState_SELLING_DECISION {
		return types.ErrListingAlreadyEnded
	}

	bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
	if len(bids) == 0 {
		sender, err := sdk.AccAddressFromBech32(msg.Sender)
		if err != nil {
			return err
		}
		err = k.nftKeeper.Transfer(ctx, listing.NftId.ClassId, listing.NftId.NftId, sender)
		if err != nil {
			return err
		}
		k.DeleteNftListings(ctx, listing)
	} else {
		params := k.GetParamSet(ctx)
		listing.FullPaymentEndAt = ctx.BlockTime().Add(time.Duration(params.NftListingFullPaymentPeriod) * time.Second)
		listing.State = types.ListingState_END_LISTING
		k.SaveNftListing(ctx, listing)

		// automatic payment after listing ends
		for _, bid := range bids {
			if bid.AutomaticPayment {
				bidder, err := sdk.AccAddressFromBech32(bid.Bidder)
				if err != nil {
					fmt.Println(err)
					continue
				}

				cacheCtx, write := ctx.CacheContext()
				err = k.PayFullBid(cacheCtx, &types.MsgPayFullBid{
					Sender: bidder.String(),
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
	}

	// Emit event for nft listing end
	ctx.EventManager().EmitTypedEvent(&types.EventEndListNfting{
		Owner:   msg.Sender,
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
	// params := k.GetParamSet(ctx)
	// listings := k.GetActiveNftListingsEndingAt(ctx, ctx.BlockTime())
	fmt.Println("---Block time---")
	fmt.Println(ctx.BlockTime())
	bids := k.GetActiveNftBiddingsEndingAt(ctx, ctx.BlockTime())
	fmt.Println("---bids---")
	fmt.Println(bids)
	k.DeleteBidsWithoutBorrowing(ctx, bids)
	checkListingsWithBorrowedBids := map[types.NftListing][]types.NftBid{}
	for _, bid := range bids {
		if !bid.IsBorrowing() {
			continue
		}

		listing, err := k.GetNftListingByIdBytes(ctx, bid.NftId.IdBytes())
		if err != nil {
			fmt.Println("failed to get listing by id bytes: %w", err)
			continue
		}
		if !listing.IsSelling() {
			continue
		}
		checkListingsWithBorrowedBids[listing] = append(checkListingsWithBorrowedBids[listing], bid)
	}

	for listing, expiredBorrowedBids := range checkListingsWithBorrowedBids {
		bids := k.GetBidsByNft(ctx, listing.NftId.IdBytes())
		if listing.CanRefinancing(bids, expiredBorrowedBids, ctx.BlockTime()) {
			fmt.Println("---occur refinaing---")
			k.Refinancings(ctx, listing, expiredBorrowedBids)
		} else if listing.State == types.ListingState_END_LISTING {
			continue
		} else {
			fmt.Println("---occur endlisting---")
			listingOwner, err := sdk.AccAddressFromBech32(listing.Owner)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = k.EndNftListing(ctx, &types.MsgEndNftListing{
				Sender: listingOwner.String(),
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
	// todo update for v2
	params := k.GetParamSet(ctx)
	// get listings ended earlier
	listings := k.GetFullPaymentNftListingsEndingAt(ctx, ctx.BlockTime())

	// handle not fully paid bids
	for _, listing := range listings {
		// todo change bids order which bids order should be ordered first
		bids := types.NftBids(k.GetBidsByNft(ctx, listing.NftId.IdBytes()))
		if listing.State == types.ListingState_SELLING_DECISION {
			// todo get higher bidding price Bid
			HighestBid := bids.GetHighestBid()
			// if winner bidder did not pay full bid, nft is listed again after deleting winner bidder
			if !HighestBid.IsPaidBidAmount() {
				err := k.DeleteBid(ctx, HighestBid)
				if err != nil {
					fmt.Println(err)
					continue
				}
				if len(bids) == 1 {
					listing.State = types.ListingState_LISTING
				} else {
					listing.State = types.ListingState_BIDDING
				}
				listing.EndAt = ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingExtendSeconds))

				// Reset the loan data for a lister
				// If the bid.PaidAmount is more than loan.Coin.Amount, then just delete the loan data for lister.
				// Otherwise, subtract bid.PaidAmount from loaning amount

				// todo change logic
				// deposit - borrwoings amount to lister and interst aount to lister
				if HighestBid.IsBorrowing() {
					loan := k.GetDebtByNft(ctx, listing.IdBytes())
					if loan.Loan.Equal(HighestBid.BorrowingAmount()) {
						k.DeleteDebt(ctx, listing.IdBytes())
					} else {
						renewedLoanAmount := loan.Loan.Sub(HighestBid.BorrowingAmount())
						loan.Loan.Amount = renewedLoanAmount.Amount
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
			err := k.LiquidationProcess(ctx, bids, listing, params)
			if err != nil {
				fmt.Println("failed to liquidation process: %w", err)
				continue
			}
		}
	}
}

func (k Keeper) LiquidationProcess(ctx sdk.Context, bids types.NftBids, listing types.NftListing, params types.Params) error {
	bids = bids.SortLiquidation()
	winnerBid := bids.GetWinnerBid()

	cacheCtx, write := ctx.CacheContext()
	if winnerBid.IsNil() {
		err := k.LiquidationProcessNotExitsWinner(cacheCtx, bids, listing)
		if err != nil {
			return err
		}
		k.DeleteNftListings(ctx, listing)
	} else {
		collectBids, refundBids := bids.MakeCollectBidsAndRefundBids()
		err := k.LiquidationProcessExitsWinner(cacheCtx, collectBids, refundBids, listing, winnerBid, ctx.BlockTime(), k.RefundBids)
		if err != nil {
			return err
		}
		listing.SuccessfulBidEndAt = ctx.BlockTime().Add(time.Second * time.Duration(params.NftListingNftDeliveryPeriod))
		listing.State = types.ListingState_SUCCESSFUL_BID
		k.SaveNftListing(ctx, listing)
	}
	write()

	// delete the loan data for the nftId which is deleted from the market anyway
	k.RemoveDebt(ctx, listing.IdBytes())
	return nil
}

// todo add test
func (k Keeper) LiquidationProcessNotExitsWinner(ctx sdk.Context, bids types.NftBids, listing types.NftListing) error {
	listingOwner, err := sdk.AccAddressFromBech32(listing.Owner)
	if err != nil {
		return err
	}

	collectedDeposit, err := k.CollectedDepositFromBids(ctx, bids)
	if err != nil {
		return nil
	}
	listing.CollectedAmount = listing.CollectedAmount.Add(collectedDeposit)

	depositCollected := listing.CollectedAmount
	// pay fee
	loan := k.GetDebtByNft(ctx, listing.IdBytes())
	k.ProcessPaymentWithCommissionFee(ctx, listingOwner, listing.BidToken, depositCollected.Amount, loan.Loan.Amount, listing.NftId)
	// transfer nft to listing owner
	err = k.nftKeeper.Transfer(ctx, listing.NftId.ClassId, listing.NftId.NftId, listingOwner)
	if err != nil {
		return err
	}
	return nil
}

// todo add test
func (k Keeper) LiquidationProcessExitsWinner(
	ctx sdk.Context, collectBids, refundBids types.NftBids,
	listing types.NftListing, winnerBid types.NftBid,
	now time.Time,
	refundF func(ctx sdk.Context, refundBids types.NftBids, totalInterest, surplusAmount sdk.Coin, listing types.NftListing) error) error {

	collectedDeposit, err := k.CollectedDepositFromBids(ctx, collectBids)
	if err != nil {
		return err
	}
	if collectedDeposit.IsPositive() {
		listing.CollectedAmount = listing.CollectedAmount.Add(collectedDeposit)
	}

	surplusAmount := k.GetSurplusAmount(refundBids, winnerBid).Add(listing.CollectedAmount)
	totalInterest := refundBids.TotalInterestAmount(now)
	if totalInterest.IsNil() {
		totalInterest = sdk.NewCoin(listing.BidToken, sdk.ZeroInt())
	}
	err = refundF(ctx, refundBids, totalInterest, surplusAmount, listing)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) RefundBids(ctx sdk.Context, refundBids types.NftBids, totalInterest, surplusAmount sdk.Coin, listing types.NftListing) error {
	if totalInterest.IsLTE(surplusAmount) {
		listing.CollectedAmount = listing.CollectedAmount.Sub(totalInterest)
		for _, bid := range refundBids {
			err := k.SafeCloseBidWithAllInterest(ctx, bid)
			if err != nil {
				return err
			}
		}
	} else {
		for _, bid := range refundBids {
			bidderGetInterest := types.CalcPartInterest(totalInterest.Amount, surplusAmount.Amount, bid.TotalInterestAmountDec(ctx.BlockTime()))
			err := k.SafeCloseBidWithPartInterest(ctx, bid, sdk.NewCoin(bid.DepositAmount.Denom, bidderGetInterest))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// todo add test
func (k Keeper) CollectedDepositFromBids(ctx sdk.Context, bids types.NftBids) (sdk.Coin, error) {
	result := sdk.Coin{
		Denom:  "",
		Amount: sdk.ZeroInt(),
	}
	for _, bid := range bids {
		// not pay bidder amount, collected deposit
		CollectedAmount, err := k.SafeCloseBidCollectDeposit(ctx, bid)
		if err != nil {
			return result, err
		}
		if CollectedAmount.IsPositive() {
			if !result.IsPositive() {
				result.Amount = CollectedAmount.Amount
				result.Denom = CollectedAmount.Denom
				continue
			}
			result = result.Add(CollectedAmount)
		}
	}
	return result, nil
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
		totalPayAmount := listing.CollectedAmount.Add(bid.BidAmount)
		k.ProcessPaymentWithCommissionFee(ctx, listingOwner, totalPayAmount.Denom, totalPayAmount.Amount, loan.Loan.Amount, listing.NftId)

		err = k.DeleteBid(ctx, bid)
		if err != nil {
			fmt.Println(err)
			continue
		}
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

// get surplus amount
func (k Keeper) GetSurplusAmount(bids types.NftBids, winBid types.NftBid) sdk.Coin {
	if len(bids) == 0 {
		return sdk.NewCoin(winBid.BidAmount.Denom, sdk.ZeroInt())
	}
	return winBid.BidAmount.Sub(bids.TotalDeposit())
}
