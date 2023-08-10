package types

// DONTCOVER

import (
	fmt "fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrClassExists              = sdkerrors.Register(ModuleName, 2, "attempting to create a denom that already exists (has bank metadata)")
	ErrUnauthorized             = sdkerrors.Register(ModuleName, 3, "unauthorized account")
	ErrInvalidClassId           = sdkerrors.Register(ModuleName, 4, "invalid class id")
	ErrInvalidCreator           = sdkerrors.Register(ModuleName, 5, "invalid creator")
	ErrInvalidAuthorityMetadata = sdkerrors.Register(ModuleName, 6, "invalid authority metadata")
	ErrInvalidGenesis           = sdkerrors.Register(ModuleName, 7, "invalid genesis")
	ErrSubclassTooLong          = sdkerrors.Register(ModuleName, 8, fmt.Sprintf("subclass too long, max length is %d bytes", MaxSubdenomLength))
	ErrCreatorTooLong           = sdkerrors.Register(ModuleName, 9, fmt.Sprintf("creator too long, max length is %d bytes", MaxCreatorLength))
	ErrDenomDoesNotExist        = sdkerrors.Register(ModuleName, 10, "denom does not exist")
	ErrUnableToCharge           = sdkerrors.Register(ModuleName, 11, "unable to charge for denom creation")
)
