package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/derivatives module sentinel errors
var (
	ErrInvalidRedeemAmount              = sdkerrors.Register(ModuleName, 1, "redeem amount exceeds user balance")
	ErrNoLiquidityProviderToken         = sdkerrors.Register(ModuleName, 2, "no liquidity provider token")
	ErrInvalidCoins                     = sdkerrors.Register(ModuleName, 3, "invalid coins")
	ErrZeroLpTokenPrice                 = sdkerrors.Register(ModuleName, 4, "zero lp token price")
	ErrorMarginNotEnough                = sdkerrors.Register(ModuleName, 5, "margin is not enough")
	ErrorInvalidPositionParams          = sdkerrors.Register(ModuleName, 6, "invalid param for position")
	ErrInsufficientAssetBalance         = sdkerrors.Register(ModuleName, 7, "insufficient asset balance")
	ErrMarginAssetNotValid              = sdkerrors.Register(ModuleName, 8, "margin asset is not valid")
	ErrNegativeMargin                   = sdkerrors.Register(ModuleName, 9, "remaining margin must be positive")
	ErrInvalidLeverage                  = sdkerrors.Register(ModuleName, 10, "invalid leverage")
	ErrInvalidPositionSize              = sdkerrors.Register(ModuleName, 11, "invalid position size")
	ErrInsufficientPoolFund             = sdkerrors.Register(ModuleName, 12, "insufficient pool fund")
	ErrInvalidPositionInstance          = sdkerrors.Register(ModuleName, 13, "invalid position instance")
	ErrNotImplemented                   = sdkerrors.Register(ModuleName, 14, "not implemented")
	ErrInsufficientAmount               = sdkerrors.Register(ModuleName, 15, "insufficient amount")
	ErrPositionDoesNotExist             = sdkerrors.Register(ModuleName, 16, "position does not exist")
	ErrLiquidationNotNeeded             = sdkerrors.Register(ModuleName, 17, "liquidation is not needed")
	ErrLiquidationNeeded                = sdkerrors.Register(ModuleName, 18, "liquidation is needed")
	ErrUnauthorized                     = sdkerrors.Register(ModuleName, 19, "unauthorized")
	ErrTooMuchMarginToWithdraw          = sdkerrors.Register(ModuleName, 20, "too much margin to withdraw")
	ErrInsufficientAvailablePoolBalance = sdkerrors.Register(ModuleName, 21, "insufficient available pool balance")
	ErrPositionNFTNotFound              = sdkerrors.Register(ModuleName, 22, "position NFT not found")
	ErrNotPositionNFTOwner              = sdkerrors.Register(ModuleName, 23, "not position NFT owner")
	ErrPositionNFTSendDisabled          = sdkerrors.Register(ModuleName, 24, "position NFT send disabled")
	ErrNoPendingPaymentManager          = sdkerrors.Register(ModuleName, 25, "no pending payment manager")
)
