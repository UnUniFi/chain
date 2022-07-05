package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftmint/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	PrefixClassId    = "ununifi/"
	LenHashByteToHex = 32 - 20
)

func (k Keeper) CreateClass(ctx sdk.Context, classID string, msg *types.MsgCreateClass) error {
	exists := k.nftKeeper.HasClass(ctx, classID)
	if exists {
		return sdkerrors.Wrap(nfttypes.ErrClassExists, classID)
	}

	err := k.nftKeeper.SaveClass(
		ctx,
		types.NewClass(classID, msg.Name, msg.Symbol, msg.Description, msg.ClassUri),
	)
	if err != nil {
		return err
	}

	err = k.SetClassAttributes(
		ctx,
		types.NewClassAttributes(classID, msg.Sender.AccAddress(), msg.BaseTokenUri, msg.MintingPermission, msg.TokenSupplyCap),
	)
	if err != nil {
		return err
	}

	owningClassIdList := k.AddClassIDToOwningClassIdList(ctx, msg.Sender.AccAddress(), classID)
	err = k.SetOwningClassIdList(ctx, owningClassIdList)
	if err != nil {
		return err
	}

	classNameIdList := k.AddClassNameIdList(ctx, msg.Name, classID)
	err = k.SetClassNameIdList(ctx, classNameIdList)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) SetClassAttributes(ctx sdk.Context, classAttributes types.ClassAttributes) error {
	bz, err := k.cdc.Marshal(&classAttributes)
	if err != nil {
		return sdkerrors.Wrap(err, "Marshal nftmint.ClassAttributes failed")
	}
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixClassAttributes))
	prefixStore.Set([]byte(classAttributes.ClassId), bz)
	return nil
}

func (k Keeper) SetOwningClassIdList(ctx sdk.Context, owningClassIdList types.OwningClassIdList) error {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixOwningClassIdList)

	bz, err := k.cdc.Marshal(&owningClassIdList)
	if err != nil {
		return sdkerrors.Wrap(err, "Marshal nftmint.OwningClassIdList failed")
	}
	owningClassIdListKey := types.OwningClassIdListKey(owningClassIdList.Owner.AccAddress())
	prefixStore.Set(owningClassIdListKey, bz)
	return nil
}

func (k Keeper) SetClassNameIdList(ctx sdk.Context, classNameIdList types.ClassNameIdList) error {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixClassNameIdList)

	bz, err := k.cdc.Marshal(&classNameIdList)
	if err != nil {
		return sdkerrors.Wrap(err, "Marshal nftmint.ClassNameIdList failed")
	}
	prefixStore.Set([]byte(classNameIdList.ClassName), bz)
	return nil
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

func (k Keeper) GetOwningClassIdList(ctx sdk.Context, owner sdk.AccAddress) (types.OwningClassIdList, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixOwningClassIdList)

	var owningClassIdList types.OwningClassIdList
	bz := prefixStore.Get(owner.Bytes())
	if len(bz) == 0 {
		return types.OwningClassIdList{}, false
	}
	k.cdc.MustUnmarshal(bz, &owningClassIdList)
	return owningClassIdList, true
}

func (k Keeper) GetClassNameIdList(ctx sdk.Context, className string) (types.ClassNameIdList, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixClassNameIdList)

	var classNameIdList types.ClassNameIdList
	bz := prefixStore.Get([]byte(className))
	if len(bz) == 0 {
		return types.ClassNameIdList{}, false
	}
	k.cdc.MustUnmarshal(bz, &classNameIdList)
	return classNameIdList, true
}

func (k Keeper) AddClassIDToOwningClassIdList(ctx sdk.Context, owner sdk.AccAddress, classID string) types.OwningClassIdList {
	owningClassIdList, exists := k.GetOwningClassIdList(ctx, owner)
	if !exists {
		owningClassIdList = types.NewOwningClassIdList(owner)
	}
	owningClassIdList.ClassId = append(owningClassIdList.ClassId, classID)
	return owningClassIdList
}

func (k Keeper) AddClassNameIdList(ctx sdk.Context, className string, classID string) types.ClassNameIdList {
	classNameIdList, exists := k.GetClassNameIdList(ctx, className)
	if !exists {
		classNameIdList = types.NewClassNameIdList(className)
	}
	classNameIdList.ClassId = append(classNameIdList.ClassId, classID)
	return classNameIdList
}

func (k Keeper) DeleteClassIDInOwningClassList(ctx sdk.Context, owner sdk.AccAddress, classID string) error {
	owningClassIdList, exists := k.GetOwningClassIdList(ctx, owner)
	if !exists {
		return sdkerrors.Wrap(types.ErrOwningClassIdListNotExists, owner.String())
	}

	index := SliceIndex(owningClassIdList.ClassId, classID)
	if index == -1 {
		return sdkerrors.Wrap(types.ErrIndexNotFoundInOwningClassIDs, classID)
	}

	owningClassIdList.ClassId = RemoveIndex(owningClassIdList.ClassId, index)
	return nil
}
