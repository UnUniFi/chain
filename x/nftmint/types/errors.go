package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrClassIdInvalidLength          = sdkerrors.Register(ModuleName, 1, "class id length is invalid")
	ErrClassAttributesNotExists      = sdkerrors.Register(ModuleName, 2, "class attributes does not exist")
	ErrOwningClassIdListNotExists    = sdkerrors.Register(ModuleName, 3, "owning class list does not exist")
	ErrIndexNotFoundInOwningClassIDs = sdkerrors.Register(ModuleName, 4, "class id is not found in list")
	ErrInvalidMintingPermission      = sdkerrors.Register(ModuleName, 5, "invalid minting permission for the class")
	ErrClassNameIdListNotExists      = sdkerrors.Register(ModuleName, 6, "class name id list with this class name does not exists")
	ErrClassNameInvalidLength        = sdkerrors.Register(ModuleName, 7, "invalid class name length on UnUniFi")
	ErrUriInvalidLength              = sdkerrors.Register(ModuleName, 8, "invalid uri length on UnUniFi")
	ErrInvalidTokenSupplyCap         = sdkerrors.Register(ModuleName, 9, "invalid token supply cap on UnUniFi")
	ErrClassSymbolInvalidLength      = sdkerrors.Register(ModuleName, 10, "invalid class symbol length on UnUniFi")
	ErrClassDescriptionInvalidLength = sdkerrors.Register(ModuleName, 11, "invalid class description length on UnUniFi")
	ErrTokenSupplyBelow              = sdkerrors.Register(ModuleName, 12, "updating token supply cap number is over the number of the current supplied token")
	ErrNftAttributesNotExists        = sdkerrors.Register(ModuleName, 13, "nft attributes does not exist")
)
