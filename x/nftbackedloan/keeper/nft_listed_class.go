package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) UpdateListedClass(ctx sdk.Context, listing types.Listing) {
	// if listing doesn't exist, delete it from listed calss
	if _, err := k.GetListedNftByIdBytes(ctx, listing.IdBytes()); err != nil {
		k.DeleteListingFromListedClass(ctx, listing)
		return
	}

	switch listing.State {
	case types.ListingState_LISTING:
		k.SetListingInListedClass(ctx, listing)
	case types.ListingState_BIDDING:
		k.SetListingInListedClass(ctx, listing)
	case types.ListingState_LIQUIDATION:
		k.DeleteListingFromListedClass(ctx, listing)
	case types.ListingState_SUCCESSFUL_BID:
		k.DeleteListingFromListedClass(ctx, listing)
	case types.ListingState_SELLING_DECISION:
		k.DeleteListingFromListedClass(ctx, listing)
	}
}

func (k Keeper) SetListingInListedClass(ctx sdk.Context, listing types.Listing) {
	store := ctx.KVStore(k.storeKey)
	bzIdlist := store.Get(types.ClassKey(listing.ClassIdBytes()))
	if bzIdlist == nil {
		bz := k.cdc.MustMarshal(
			&types.ListedClass{
				ClassId:  listing.NftId.ClassId,
				TokenIds: []string{listing.NftId.TokenId},
			},
		)
		store.Set(types.ClassKey(listing.ClassIdBytes()), bz)
	} else {
		class := types.ListedClass{}
		k.cdc.MustUnmarshal(bzIdlist, &class)

		// return if the nft_id already exists
		// index := keeper.SliceIndex(class.NftIds, listing.NftId.TokenId)
		// if index != -1 {
		// 	return
		// }
		class.TokenIds = append(class.TokenIds, listing.NftId.TokenId)
		bz := k.cdc.MustMarshal(&class)
		store.Set(types.ClassKey(listing.ClassIdBytes()), bz)
	}
}

func (k Keeper) DeleteListingFromListedClass(ctx sdk.Context, listing types.Listing) {
	store := ctx.KVStore(k.storeKey)
	bzIdlist := store.Get(types.ClassKey(listing.ClassIdBytes()))

	class := types.ListedClass{}
	k.cdc.MustUnmarshal(bzIdlist, &class)

	// removeIndex := keeper.SliceIndex(class.NftIds, listing.NftId.TokenId)
	// if removeIndex == -1 {
	// 	return
	// }
	// class.NftIds = keeper.RemoveIndex(class.NftIds, removeIndex)
	// if class doesn't have any listed nft, just delete class id key from KVStore
	if len(class.TokenIds) == 1 && class.TokenIds[0] == listing.NftId.TokenId {
		store.Delete(types.ClassKey(listing.ClassIdBytes()))
		return
	}

	bz := k.cdc.MustMarshal(&class)
	store.Set(types.ClassKey(listing.ClassIdBytes()), bz)
}

func (k Keeper) GetListedClassByClassIdBytes(ctx sdk.Context, classIdByte []byte) (types.ListedClass, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(classIdByte)
	if bz == nil {
		return types.ListedClass{}, types.ErrListedNftDoesNotExist
	}
	class := types.ListedClass{}
	k.cdc.MustUnmarshal(bz, &class)
	return class, nil
}

func (k Keeper) GetListedClasses(ctx sdk.Context) ([]types.ListedClass, error) {
	store := ctx.KVStore(k.storeKey)
	classes := []types.ListedClass{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixClass))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var class types.ListedClass
		k.cdc.MustUnmarshal(it.Value(), &class)

		classes = append(classes, class)
	}

	return classes, nil
}
