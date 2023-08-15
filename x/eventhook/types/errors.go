package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/eventhook module sentinel errors
var (
	ErrHookNotFound = sdkerrors.Register(ModuleName, 1, "hook not found")
)
