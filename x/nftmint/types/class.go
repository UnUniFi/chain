package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

func NewClass(classID, name, symbol, description, classUri string) nfttypes.Class {
	return nfttypes.Class{
		Id:          classID,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Uri:         classUri,
	}
}

func NewClassAttributes(
	classID string,
	owner sdk.AccAddress,
	baseTokenUri string,
	mintingpermission MintingPermission,
	tokenSupplyCap uint64,
) ClassAttributes {
	return ClassAttributes{
		ClassId:           classID,
		Owner:             owner.Bytes(),
		BaseTokenUri:      baseTokenUri,
		MintingPermission: mintingpermission,
		TokenSupplyCap:    tokenSupplyCap,
	}
}

func NewOwningClassIdList(owner sdk.AccAddress) OwningClassIdList {
	var classIDList []string
	return OwningClassIdList{Owner: owner.Bytes(), ClassId: classIDList}
}
