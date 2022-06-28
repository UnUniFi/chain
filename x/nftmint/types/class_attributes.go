package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewClassAttributes(
	classID string,
	owner sdk.AccAddress,
	baseTokenUri string,
	mintingpermission MintingPermission,
	tokenSupplyCap uint64,
) ClassAttributes {
	return ClassAttributes{
		ClassId:           classID,
		Owner:             owner.String(),
		BaseTokenUri:      baseTokenUri,
		MintingPermission: mintingpermission,
		TokenSupplyCap:    tokenSupplyCap,
	}
}
