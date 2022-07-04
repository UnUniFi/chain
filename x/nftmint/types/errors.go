package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrClassIdInvalidLength          = sdkerrors.Register(ModuleName, 1, "class id length is invalid")
	ErrClassAttributesNotExists      = sdkerrors.Register(ModuleName, 2, "class attributes does not exist")
	ErrOwningClassIdListNotExists    = sdkerrors.Register(ModuleName, 3, "owning class list does not exist")
	ErrIndexNotFoundInOwningClassIDs = sdkerrors.Register(ModuleName, 4, "class id is not found in list")
	ErrInvalidMintingPermission      = sdkerrors.Register(ModuleName, 5, "invalid minting permission")
)
