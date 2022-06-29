package keeper

import (
	"github.com/UnUniFi/chain/x/nftmint/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SaveNFTAttributes(ctx sdk.Context, nftAttributes types.NFTAttributes) error {
	// TODO: save in kvstore
	bz := k.cdc.MustMarshal(&nftAttributes)
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixNFTAttributes))

	nftAttributesKey := types.NFTAttributesKey(nftAttributes.ClassId, nftAttributes.NftId)
	prefixStore.Set(nftAttributesKey, bz)
	return nil
}

func (k Keeper) GetNFTAttributes(ctx sdk.Context, classID, nftID string) (types.NFTAttributes, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixNFTAttributes))
	nftAttributesKey := types.NFTAttributesKey(classID, nftID)

	var nftAttributes types.NFTAttributes
	bz := prefixStore.Get(nftAttributesKey)
	if len(bz) == 0 {
		return types.NFTAttributes{}, false
	}

	k.cdc.MustUnmarshal(bz, &nftAttributes)
	return nftAttributes, true
}
