package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/derivatives module sentinel errors
var (
	ErrInvalidRedeemAmount      = sdkerrors.Register(ModuleName, 1, "redeem amount exceeds user balance")
	ErrNoLiquidityProviderToken = sdkerrors.Register(ModuleName, 2, "no liquidity provider token")
	ErrInvalidCoins             = sdkerrors.Register(ModuleName, 3, "invalid coins")
	ErrZeroLpTokenPrice         = sdkerrors.Register(ModuleName, 4, "zero lp token price")
	ErrorMarginNotEnough        = sdkerrors.Register(ModuleName, 5, "margin is not enough")
	ErrorInvalidPositionParams  = sdkerrors.Register(ModuleName, 6, "invalid param for position")
	ErrInsufficientAssetBalance = sdkerrors.Register(ModuleName, 7, "insufficient asset balance")
)
