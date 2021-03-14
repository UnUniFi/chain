package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetAuctionCount get the total number of auction
func (k Keeper) GetAuctionCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionCountKey))
	byteKey := types.KeyPrefix(types.AuctionCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetAuctionCount set the total number of auction
func (k Keeper) SetAuctionCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionCountKey))
	byteKey := types.KeyPrefix(types.AuctionCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateAuction creates a auction with a new id and update the count
func (k Keeper) CreateAuction(ctx sdk.Context, msg types.MsgCreateAuction) {
	// Create the auction
    count := k.GetAuctionCount(ctx)
    var auction = types.Auction{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
    key := types.KeyPrefix(types.AuctionKey + auction.Id)
    value := k.cdc.MustMarshalBinaryBare(&auction)
    store.Set(key, value)

    // Update auction count
    k.SetAuctionCount(ctx, count+1)
}

// SetAuction set a specific auction in the store
func (k Keeper) SetAuction(ctx sdk.Context, auction types.Auction) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
	b := k.cdc.MustMarshalBinaryBare(&auction)
	store.Set(types.KeyPrefix(types.AuctionKey + auction.Id), b)
}

// GetAuction returns a auction from its id
func (k Keeper) GetAuction(ctx sdk.Context, key string) types.Auction {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
	var auction types.Auction
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.AuctionKey + key)), &auction)
	return auction
}

// HasAuction checks if the auction exists
func (k Keeper) HasAuction(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
	return store.Has(types.KeyPrefix(types.AuctionKey + id))
}

// GetAuctionOwner returns the creator of the auction
func (k Keeper) GetAuctionOwner(ctx sdk.Context, key string) string {
    return k.GetAuction(ctx, key).Creator
}

// DeleteAuction deletes a auction
func (k Keeper) DeleteAuction(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
	store.Delete(types.KeyPrefix(types.AuctionKey + key))
}

// GetAllAuction returns all auction
func (k Keeper) GetAllAuction(ctx sdk.Context) (msgs []types.Auction) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.AuctionKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Auction
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
