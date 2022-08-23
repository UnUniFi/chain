package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/decentralized-vault module sentinel errors
var (
	ErrSample                       = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrNotNftOwner                  = sdkerrors.Register(ModuleName, 1, "not the owner of nft")
	ErrTransferRequestDoesNotExists = sdkerrors.Register(ModuleName, 2, "transfer request does not exist")
)
