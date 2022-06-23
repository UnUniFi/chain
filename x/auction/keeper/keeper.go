package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/UnUniFi/chain/x/auction/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc           codec.Codec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramSpace    paramtypes.Subspace
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(cdc codec.Codec, storeKey, memKey storetypes.StoreKey, paramSpace paramtypes.Subspace, accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramSpace:    paramSpace,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetNextAuctionID stores an ID to be used for the next created auction
func (k Keeper) SetNextAuctionID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.NextAuctionIDKey), types.Uint64ToBytes(id))
}

// GetNextAuctionID reads the next available global ID from store
func (k Keeper) GetNextAuctionID(ctx sdk.Context) (uint64, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefix(types.NextAuctionIDKey))
	if bz == nil {
		return 0, types.ErrInvalidInitialAuctionID
	}
	return types.Uint64FromBytes(bz), nil
}

// IncrementNextAuctionID increments the next auction ID in the store by 1.
func (k Keeper) IncrementNextAuctionID(ctx sdk.Context) error {
	id, err := k.GetNextAuctionID(ctx)
	if err != nil {
		return err
	}
	k.SetNextAuctionID(ctx, id+1)
	return nil
}

// StoreNewAuction stores an auction, adding a new ID
func (k Keeper) StoreNewAuction(ctx sdk.Context, auction types.Auction) (uint64, error) {
	newAuctionID, err := k.GetNextAuctionID(ctx)
	if err != nil {
		return 0, err
	}
	auction = auction.WithID(newAuctionID)

	k.SetAuction(ctx, auction)

	err = k.IncrementNextAuctionID(ctx)
	if err != nil {
		return 0, err
	}
	return newAuctionID, nil
}

// SetAuction puts the auction into the store, and updates any indexes.
func (k Keeper) SetAuction(ctx sdk.Context, auction types.Auction) {
	// remove the auction from the byTime index if it is already in there
	existingAuction, found := k.GetAuction(ctx, auction.GetID())
	if found {
		k.removeFromByTimeIndex(ctx, existingAuction.GetEndTime(), existingAuction.GetID())
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
	auctionAny, _ := codectypes.NewAnyWithValue(auction)
	bz, _ := auctionAny.Marshal()
	store.Set(types.GetAuctionKey(auction.GetID()), bz)

	k.InsertIntoByTimeIndex(ctx, auction.GetEndTime(), auction.GetID())
}

// GetAuction gets an auction from the store.
func (k Keeper) GetAuction(ctx sdk.Context, auctionID uint64) (types.Auction, bool) {
	var auction types.Auction

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
	bz := store.Get(types.GetAuctionKey(auctionID))
	if bz == nil {
		return auction, false
	}

	var auctionAny codectypes.Any
	auctionAny.Unmarshal(bz)
	auction, err := types.UnpackAuction(&auctionAny)
	if err != nil {
		return nil, false
	}
	return auction, true
}

// DeleteAuction removes an auction from the store, and any indexes.
func (k Keeper) DeleteAuction(ctx sdk.Context, auctionID uint64) {
	auction, found := k.GetAuction(ctx, auctionID)
	if found {
		k.removeFromByTimeIndex(ctx, auction.GetEndTime(), auctionID)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
	store.Delete(types.GetAuctionKey(auctionID))
}

// InsertIntoByTimeIndex adds an auction ID and end time into the byTime index.
func (k Keeper) InsertIntoByTimeIndex(ctx sdk.Context, endTime time.Time, auctionID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionByTimeKey))
	store.Set(types.GetAuctionByTimeKey(endTime, auctionID), types.Uint64ToBytes(auctionID))
}

// removeFromByTimeIndex removes an auction ID and end time from the byTime index.
func (k Keeper) removeFromByTimeIndex(ctx sdk.Context, endTime time.Time, auctionID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionByTimeKey))
	store.Delete(types.GetAuctionByTimeKey(endTime, auctionID))
}

// IterateAuctionByTime provides an iterator over auctions ordered by auction.EndTime.
// For each auction cb will be callled. If cb returns true the iterator will close and stop.
func (k Keeper) IterateAuctionsByTime(ctx sdk.Context, inclusiveCutoffTime time.Time, cb func(auctionID uint64) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionByTimeKey))
	iterator := store.Iterator(
		nil, // start at the very start of the prefix store
		sdk.PrefixEndBytes(sdk.FormatTimeBytes(inclusiveCutoffTime)), // include any keys with times equal to inclusiveCutoffTime
	)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		auctionID := types.Uint64FromBytes(iterator.Value())

		if cb(auctionID) {
			break
		}
	}
}

// IterateAuctions provides an iterator over all stored auctions.
// For each auction, cb will be called. If cb returns true, the iterator will close and stop.
func (k Keeper) IterateAuctions(ctx sdk.Context, cb func(auction types.Auction) (stop bool)) {
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var auctionAny codectypes.Any
		auctionAny.Unmarshal(iterator.Value())

		auction, _ := types.UnpackAuction(&auctionAny)

		if cb(auction) {
			break
		}
	}
}

// GetAllAuctions returns all auctions from the store
func (k Keeper) GetAllAuctions(ctx sdk.Context) (auctions types.Auctions) {
	k.IterateAuctions(ctx, func(auction types.Auction) bool {
		auctions = append(auctions, auction)
		return false
	})
	return
}
