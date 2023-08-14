package types

// DONTCOVER

import (
	fmt "fmt"

	errorsmod "cosmossdk.io/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrParsingParams            = errorsmod.Register(ModuleName, 1, "failed to marshal or unmarshal module params")
	ErrClassExists              = errorsmod.Register(ModuleName, 2, "attempting to create a denom that already exists (has bank metadata)")
	ErrUnauthorized             = errorsmod.Register(ModuleName, 3, "unauthorized account")
	ErrInvalidClassId           = errorsmod.Register(ModuleName, 4, "invalid class id")
	ErrInvalidCreator           = errorsmod.Register(ModuleName, 5, "invalid creator")
	ErrInvalidAuthorityMetadata = errorsmod.Register(ModuleName, 6, "invalid authority metadata")
	ErrInvalidGenesis           = errorsmod.Register(ModuleName, 7, "invalid genesis")
	ErrSubclassTooLong          = errorsmod.Register(ModuleName, 8, fmt.Sprintf("subclass too long, max length is %d bytes", MaxSubdenomLength))
	ErrCreatorTooLong           = errorsmod.Register(ModuleName, 9, fmt.Sprintf("creator too long, max length is %d bytes", MaxCreatorLength))
	ErrClassDoesNotExist        = errorsmod.Register(ModuleName, 10, "class does not exist")
	ErrUnableToCharge           = errorsmod.Register(ModuleName, 11, "unable to charge for denom creation")
)
