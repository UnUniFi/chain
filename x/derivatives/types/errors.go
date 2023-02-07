package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/derivatives module sentinel errors
var (
	ErrInvalidRedeemAmount      = sdkerrors.Register(ModuleName, 1, "redeem amount exceeds user balance")
	ErrNoLiquidityProviderToken = sdkerrors.Register(ModuleName, 2, "no liquidity provider token")
)
