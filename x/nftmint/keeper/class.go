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

func (k Keeper) SetClassAttributes(ctx sdk.Context, classAttributes types.ClassAttributes) {
	bz := k.cdc.MustMarshal(&classAttributes)

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixClassAttributes))
	prefixStore.Set([]byte(classAttributes.ClassId), bz)
}

func (k Keeper) SetOwningClassList(ctx sdk.Context, owningClassList types.OwningClassList) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixOwningClassList)

	bz := k.cdc.MustMarshal(&owningClassList)
	owningClassListKey := types.OwningClassListKey(owningClassList.Owner)
	prefixStore.Set(owningClassListKey, bz)
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

func (k Keeper) GetOwningClassList(ctx sdk.Context, owner sdk.AccAddress) (types.OwningClassList, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixOwningClassList)

	var owningClassList types.OwningClassList
	bz := prefixStore.Get(owner.Bytes())
	if len(bz) == 0 {
		return types.OwningClassList{}, false
	}

	k.cdc.MustUnmarshal(bz, &owningClassList)
	return owningClassList, true
}

func (k Keeper) AddClassIDToOwningClassList(ctx sdk.Context, owner sdk.AccAddress, classID string) {
	owningClassList, exists := k.GetOwningClassList(ctx, owner)
	if !exists {
		owningClassList = types.NewOwningClassList(owner)
	}
	owningClassList.ClassId = append(owningClassList.ClassId, classID)
	k.SetOwningClassList(ctx, owningClassList)
}

func (k Keeper) DeleteClassIDInOwningClassList(ctx sdk.Context, owner sdk.AccAddress, classID string) error {
	owningClassList, exists := k.GetOwningClassList(ctx, owner)
	if !exists {
		return sdkerrors.Wrap(types.ErrOwningClassListNotExists, owner.String())
	}

	index := SliceIndex(owningClassList.ClassId, classID)
	if index == -1 {
		return sdkerrors.Wrap(types.ErrIndexNotFoundInOwningClassIDs, classID)
	}

	owningClassList.ClassId = RemoveIndex(owningClassList.ClassId, index)
	return nil
}
