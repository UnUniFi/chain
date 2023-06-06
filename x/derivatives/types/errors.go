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
	ErrMarginAssetNotValid      = sdkerrors.Register(ModuleName, 8, "margin asset is not valid")
	ErrNegativeMargin           = sdkerrors.Register(ModuleName, 9, "remaining margin must be positive")
	ErrInvalidLeverage          = sdkerrors.Register(ModuleName, 10, "invalid leverage")
	ErrInvalidPositionSize      = sdkerrors.Register(ModuleName, 11, "invalid position size")
	ErrInsufficientPoolFund     = sdkerrors.Register(ModuleName, 12, "insufficient pool fund")
	ErrInvalidPositionInstance  = sdkerrors.Register(ModuleName, 13, "invalid position instance")
	ErrNotImplemented           = sdkerrors.Register(ModuleName, 14, "not implemented")
)
