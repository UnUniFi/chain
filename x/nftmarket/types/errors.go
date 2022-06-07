package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNftListingDoesNotExist = sdkerrors.Register(ModuleName, 2, "nft listing does not exist")
	ErrBidDoesNotExists       = sdkerrors.Register(ModuleName, 3, "nft bid does not exist")
)
