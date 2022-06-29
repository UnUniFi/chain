package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
)

func (k Keeper) UpdateListedClass(ctx sdk.Context, listing types.NftListing) {
	switch listing.State {
	case types.ListingState_LISTING:
		k.SetListedClass(ctx, listing)
	case types.ListingState_SUCCESSFUL_BID:
		// todo delete nftid from class
	}
}

func (k Keeper) SetListedClass(ctx sdk.Context, listing types.NftListing) {
	store := ctx.KVStore(k.storeKey)
	bzIdlist := store.Get(types.ClassKey(listing.ClassIdBytes()))
	if bzIdlist == nil {
		bz := k.cdc.MustMarshal(
			&types.ListedClass{
				ClassId: listing.NftId.ClassId,
				NftIds:  []string{listing.NftId.NftId},
			},
		)
		store.Set(types.ClassKey(listing.ClassIdBytes()), bz)
	} else {

		// todo delete dumplicate nftid
		class := types.ListedClass{}
		k.cdc.MustUnmarshal(bzIdlist, &class)
		class.NftIds = append(class.NftIds, listing.NftId.NftId)
		bz := k.cdc.MustMarshal(&class)
		store.Set(types.ClassKey(listing.ClassIdBytes()), bz)
	}
}

func (k Keeper) GetListedClassByClassIdBytes(ctx sdk.Context, classIdByte []byte) (types.ListedClass, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(classIdByte)
	if bz == nil {
		return types.ListedClass{}, types.ErrNftListingDoesNotExist
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
