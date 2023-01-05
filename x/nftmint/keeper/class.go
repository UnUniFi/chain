package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/nftmint/types"
)

const (
	PrefixClassId    = "ununifi-"
	LenHashByteToHex = 32 - 20
)

// CreateClass does validate the contents of MsgCreateClass and operate whole flow for CreateClass message
func (k Keeper) CreateClass(ctx sdk.Context, classID string, msg *types.MsgCreateClass) error {
	exists := k.nftKeeper.HasClass(ctx, classID)
	if exists {
		return sdkerrors.Wrap(nfttypes.ErrClassExists, classID)
	}

	params := k.GetParamSet(ctx)
	err := types.ValidateCreateClass(
		params,
		msg.Name, msg.Symbol, msg.BaseTokenUri, msg.Description,
		msg.MintingPermission,
		msg.TokenSupplyCap,
	)
	if err != nil {
		return err
	}

	if err := k.nftKeeper.SaveClass(ctx, types.NewClass(classID, msg.Name, msg.Symbol, msg.Description, msg.ClassUri)); err != nil {
		return err
	}

	if err = k.SetClassAttributes(ctx, types.NewClassAttributes(classID, msg.Sender.AccAddress(), msg.BaseTokenUri, msg.MintingPermission, msg.TokenSupplyCap)); err != nil {
		return err
	}

	owningClassIdList := k.AddClassIDToOwningClassIdList(ctx, msg.Sender.AccAddress(), classID)
	if err = k.SetOwningClassIdList(ctx, owningClassIdList); err != nil {
		return err
	}

	classNameIdList := k.AddClassNameIdList(ctx, msg.Name, classID)
	if err = k.SetClassNameIdList(ctx, classNameIdList); err != nil {
		return err
	}

	return nil
}

// Create class id on UnUniFi using addr sequence and addr byte
func CreateClassId(num uint64, addr sdk.Address) string {
	sequenceByte := UintToByte(num)
	addrByte := addr.Bytes()
	idByte := append(addrByte, sequenceByte...)

	idHash := sha256.Sum256(idByte)
	idString := hex.EncodeToString(idHash[LenHashByteToHex:])
	classID := PrefixClassId + strings.ToUpper(idString)

	return classID
}

// SendClassOwnership does validate the contents of MsgSendClassOwnership and operate whole flow for SendClassOwnership message
func (k Keeper) SendClassOwnership(ctx sdk.Context, msg *types.MsgSendClassOwnership) error {
	if !k.nftKeeper.HasClass(ctx, msg.ClassId) {
		return sdkerrors.Wrap(nfttypes.ErrClassNotExists, msg.ClassId)
	}

	classAttirbutes, exists := k.GetClassAttributes(ctx, msg.ClassId)
	if !exists {
		return sdkerrors.Wrap(types.ErrClassAttributesNotExists, msg.ClassId)
	}

	if !msg.Sender.AccAddress().Equals(classAttirbutes.Owner.AccAddress()) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not the owner of the class", msg.Sender.AccAddress().String())
	}

	classAttirbutes.Owner = msg.Recipient
	if err := k.SetClassAttributes(ctx, classAttirbutes); err != nil {
		return err
	}

	return nil
}

// UpdateTokenSupplyCap does validate the contents of MsgUpdateTokenSupplyCap and operate whole flow for UpdateTokenSupplyCap message
func (k Keeper) UpdateTokenSupplyCap(ctx sdk.Context, msg *types.MsgUpdateTokenSupplyCap) error {
	classAttributes, exists := k.GetClassAttributes(ctx, msg.ClassId)
	if !exists {
		return sdkerrors.Wrap(types.ErrClassAttributesNotExists, msg.ClassId)
	}

	if err := k.IsUpgradable(ctx, msg.Sender.AccAddress(), classAttributes); err != nil {
		return err
	}

	params := k.GetParamSet(ctx)
	if err := types.ValidateTokenSupplyCap(params.MaxNFTSupplyCap, msg.TokenSupplyCap); err != nil {
		return err
	}
	currentSupply := k.nftKeeper.GetTotalSupply(ctx, msg.ClassId)
	if err := types.ValidateTokenSupply(currentSupply, msg.TokenSupplyCap); err != nil {
		return err
	}

	classAttributes.TokenSupplyCap = msg.TokenSupplyCap
	if err := k.SetClassAttributes(ctx, classAttributes); err != nil {
		return err
	}

	return nil
}

// UpdateBaseTokenUri does validate the contents of MsgUpdateBaseTokenUri and operate whole flow for UpdateBaseTokenUri message
func (k Keeper) UpdateBaseTokenUri(ctx sdk.Context, msg *types.MsgUpdateBaseTokenUri) error {
	classAttributes, exists := k.GetClassAttributes(ctx, msg.ClassId)
	if !exists {
		return sdkerrors.Wrap(types.ErrClassAttributesNotExists, msg.ClassId)
	}

	if err := k.IsUpgradable(ctx, msg.Sender.AccAddress(), classAttributes); err != nil {
		return err
	}

	params := k.GetParamSet(ctx)
	if err := types.ValidateUri(params.MinUriLen, params.MaxUriLen, msg.BaseTokenUri); err != nil {
		return err
	}

	classAttributes.BaseTokenUri = msg.BaseTokenUri
	if err := k.SetClassAttributes(ctx, classAttributes); err != nil {
		return err
	}

	if err := k.UpdateNFTUri(ctx, classAttributes.ClassId, classAttributes.BaseTokenUri); err != nil {
		return err
	}

	return nil
}

// check if update relating messages are permitted
func (k Keeper) IsUpgradable(ctx sdk.Context, sender sdk.AccAddress, classAttributes types.ClassAttributes) error {
	if exists := k.nftKeeper.HasClass(ctx, classAttributes.ClassId); !exists {
		return sdkerrors.Wrap(nfttypes.ErrClassNotExists, classAttributes.ClassId)
	}

	if !sender.Equals(classAttributes.Owner.AccAddress()) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not the owner of the class", sender.String())
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

func (k Keeper) GetClassAttributesList(ctx sdk.Context) (classAttributesList []*types.ClassAttributes) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixClassAttributes)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var classAttributes types.ClassAttributes
		k.cdc.MustUnmarshal(iterator.Value(), &classAttributes)
		classAttributesList = append(classAttributesList, &classAttributes)
	}

	return
}

func (k Keeper) GetOwningClassIdLists(ctx sdk.Context) (owningClassIdLists []*types.OwningClassIdList) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixOwningClassIdList)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var owningClassIdList types.OwningClassIdList
		k.cdc.MustUnmarshal(iterator.Value(), &owningClassIdList)
		owningClassIdLists = append(owningClassIdLists, &owningClassIdList)
	}

	return
}

func (k Keeper) GetClassNameIdLists(ctx sdk.Context) (classNameIdLists []*types.ClassNameIdList) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixClassNameIdList)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var classNameIdList types.ClassNameIdList
		k.cdc.MustUnmarshal(iterator.Value(), &classNameIdList)
		classNameIdLists = append(classNameIdLists, &classNameIdList)
	}

	return
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
