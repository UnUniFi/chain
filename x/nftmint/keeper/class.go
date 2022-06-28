package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmint/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

// TODO: method to create class_id
func (k Keeper) CreateClassId(ctx sdk.Context, creator sdk.AccAddress) (string, error) {
	sequence, err := k.accountKeeper.GetSequence(ctx, creator)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d", sequence)

	// TODO: create random string from accaddress and sequence
	// q is bech32 or hex as encoding format
	classID := "initial"

	exists := k.nftKeeper.HasClass(ctx, classID)
	if exists {
		return "", sdkerrors.Wrap(nfttypes.ErrClassExists, classID)
	}

	return classID, nil
}

func (k Keeper) SaveClassAttributes(ctx sdk.Context, classAttributes types.ClassAttributes) {
	bz := k.cdc.MustMarshal(&classAttributes)

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixClassAttributes))
	prefixStore.Set([]byte(classAttributes.ClassId), bz)
}

func (k Keeper) GetClassAttributes(ctx sdk.Context, classID string) (types.ClassAttributes, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixClassAttributes)

	bz := prefixStore.Get([]byte(classID))
	if len(bz) == 0 {
		return types.ClassAttributes{}, false
	}
	var classAttributes types.ClassAttributes
	k.cdc.MustUnmarshal(bz, &classAttributes)
	return classAttributes, true
}
