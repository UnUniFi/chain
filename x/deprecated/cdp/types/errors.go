package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cdp module sentinel errors

var (
	// ErrCdpAlreadyExists error for duplicate cdps
	ErrCdpAlreadyExists = sdkerrors.Register(ModuleName, 2, "cdp already exists")
	// ErrInvalidCollateralLength error for invalid collateral input length
	ErrInvalidCollateralLength = sdkerrors.Register(ModuleName, 3, "only one collateral type per cdp")
	// ErrCollateralNotSupported error for unsupported collateral
	ErrCollateralNotSupported = sdkerrors.Register(ModuleName, 4, "collateral not supported")
	// ErrDebtNotSupported error for unsupported debt
	ErrDebtNotSupported = sdkerrors.Register(ModuleName, 5, "debt not supported")
	// ErrExceedsDebtLimit error for attempted draws that exceed debt limit
	ErrExceedsDebtLimit = sdkerrors.Register(ModuleName, 6, "proposed debt increase would exceed debt limit")
	// ErrInvalidCollateralRatio error for attempted draws that are below liquidation ratio
	ErrInvalidCollateralRatio = sdkerrors.Register(ModuleName, 7, "proposed collateral ratio is below liquidation ratio")
	// ErrCdpNotFound error cdp not found
	ErrCdpNotFound = sdkerrors.Register(ModuleName, 8, "cdp not found")
	// ErrDepositNotFound error for deposit not found
	ErrDepositNotFound = sdkerrors.Register(ModuleName, 9, "deposit not found")
	// ErrInvalidDeposit error for invalid deposit
	ErrInvalidDeposit = sdkerrors.Register(ModuleName, 10, "invalid deposit")
	// ErrInvalidPayment error for invalid payment
	ErrInvalidPayment = sdkerrors.Register(ModuleName, 11, "invalid payment")
	//ErrDepositNotAvailable error for withdrawing deposits in liquidation
	ErrDepositNotAvailable = sdkerrors.Register(ModuleName, 12, "deposit in liquidation")
	// ErrInvalidWithdrawAmount error for invalid withdrawal amount
	ErrInvalidWithdrawAmount = sdkerrors.Register(ModuleName, 13, "withdrawal amount exceeds deposit")
	//ErrCdpNotAvailable error for depositing to a Cdp in liquidation
	ErrCdpNotAvailable = sdkerrors.Register(ModuleName, 14, "cannot modify cdp in liquidation")
	// ErrBelowDebtFloor error for creating a cdp with debt below the minimum
	ErrBelowDebtFloor = sdkerrors.Register(ModuleName, 15, "proposed cdp debt is below minimum")
	// ErrLoadingAugmentedCdp error loading augmented cdp
	ErrLoadingAugmentedCdp = sdkerrors.Register(ModuleName, 16, "augmented cdp could not be loaded from cdp")
	// ErrInvalidDebtRequest error for invalid principal input length
	ErrInvalidDebtRequest = sdkerrors.Register(ModuleName, 17, "only one principal type per cdp")
	// ErrDenomPrefixNotFound error for denom prefix not found
	ErrDenomPrefixNotFound = sdkerrors.Register(ModuleName, 18, "denom prefix not found")
	// ErrPricefeedDown error for when a price for the input denom is not found
	ErrPricefeedDown = sdkerrors.Register(ModuleName, 19, "no price found for collateral")
	// ErrInvalidCollateral error for when the input collateral denom does not match the expected collateral denom
	ErrInvalidCollateral = sdkerrors.Register(ModuleName, 20, "invalid collateral for input collateral type")
	// ErrAccountNotFound error for when no account is found for an input address
	ErrAccountNotFound = sdkerrors.Register(ModuleName, 21, "account not found")
	// ErrInsufficientBalance error for when an account does not have enough funds
	ErrInsufficientBalance = sdkerrors.Register(ModuleName, 22, "insufficient balance")
	// ErrNotLiquidatable error for when an cdp is not liquidatable
	ErrNotLiquidatable = sdkerrors.Register(ModuleName, 23, "cdp collateral ratio not below liquidation ratio")
)
