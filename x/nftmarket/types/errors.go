package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNftListingAlreadyExists = sdkerrors.Register(ModuleName, 1, "nft listing already exist")
	ErrNftListingDoesNotExist  = sdkerrors.Register(ModuleName, 2, "nft listing does not exist")
	ErrBidDoesNotExists        = sdkerrors.Register(ModuleName, 3, "nft bid does not exist")
	ErrNotSupportedBidToken    = sdkerrors.Register(ModuleName, 4, "not supported bid token")
	ErrNftDoesNotExists        = sdkerrors.Register(ModuleName, 5, "specified nft does not exist")
	ErrNotNftOwner             = sdkerrors.Register(ModuleName, 6, "not the owner of nft")
	ErrNotNftListingOwner      = sdkerrors.Register(ModuleName, 7, "not the owner of nft listing")
	ErrNftBidAlreadyExists     = sdkerrors.Register(ModuleName, 8, "bid already exists on the nft")
	ErrNftBidDoesNotExists     = sdkerrors.Register(ModuleName, 9, "bid does not exists on the nft")
)
