package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/yieldaggregator module sentinel errors
var (
	ErrAssetManagementAccountAlreadyExists = sdkerrors.Register(ModuleName, 2, "asset management account already exists")
	ErrAssetManagementAccountDoesNotExists = sdkerrors.Register(ModuleName, 3, "asset management account does not exist")
)
