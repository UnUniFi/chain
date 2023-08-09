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

func NewClassOwnership(
	classID string,
	owner sdk.AccAddress,
) ClassOwnership {
	return ClassOwnership{
		ClassId: classID,
		Owner:   owner.String(),
	}
}
