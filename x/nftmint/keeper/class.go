package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmint/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	PrefixClassId    = "ununifi/"
	LenHashByteToHex = 32 - 20
)

// Create class id on UnUniFi using addr sequence and addr byte
func createClassId(num uint64, addr sdk.Address) string {
	sequenceByte := UintToByte(num)
	addrByte := addr.Bytes()
	idByte := append(addrByte, sequenceByte...)

	idHash := sha256.Sum256(idByte)
	idString := hex.EncodeToString(idHash[LenHashByteToHex:])
	classID := PrefixClassId + strings.ToUpper(idString)

	return classID
}

func (k Keeper) SetClassAttributes(ctx sdk.Context, classAttributes types.ClassAttributes) {
	bz := k.cdc.MustMarshal(&classAttributes)

	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixClassAttributes))
	prefixStore.Set([]byte(classAttributes.ClassId), bz)
}

func (k Keeper) SetOwningClassList(ctx sdk.Context, owningClassIdList types.OwningClassIdList) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixOwningClassIdList)

	bz := k.cdc.MustMarshal(&owningClassIdList)
	owningClassIdListKey := types.OwningClassIdListKey(owningClassIdList.Owner.AccAddress())
	prefixStore.Set(owningClassIdListKey, bz)
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

func (k Keeper) AddClassIDToOwningClassIdList(ctx sdk.Context, owner sdk.AccAddress, classID string) {
	owningClassIdList, exists := k.GetOwningClassIdList(ctx, owner)
	if !exists {
		owningClassIdList = types.NewOwningClassIdList(owner)
	}
	owningClassIdList.ClassId = append(owningClassIdList.ClassId, classID)
	k.SetOwningClassList(ctx, owningClassIdList)
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
