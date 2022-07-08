package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func ValidateMintingPermission(classAttributes ClassAttributes, minter sdk.AccAddress) error {
	switch classAttributes.MintingPermission {
	case 0:
		owner := classAttributes.Owner.AccAddress()
		if !owner.Equals(minter) {
			return ErrInvalidMintingPermission
		}
		return nil
	case 1:
		return nil
	default:
		return ErrInvalidMintingPermission
	}
}

// Measure length of the string as byte length

func ValidateClassName(minLen, maxLen uint64, className string) error {
	len := len(className)
	if len < int(minLen) || len > int(maxLen) {
		return sdkerrors.Wrap(ErrClassNameInvalidLength, className)
	}

	return nil
}

func ValidateUri(minLen, maxLen uint64, uri string) error {
	len := len(uri)
	if len < int(minLen) || len > int(maxLen) {
		return sdkerrors.Wrap(ErrUriInvalidLength, uri)
	}

	return nil
}

func ValidateTokenSupplyCap(maxCap uint64, tokenSupplyCap uint64) error {
	if tokenSupplyCap > maxCap {
		strTokenSupplyCap := strconv.FormatUint(tokenSupplyCap, 10)
		return sdkerrors.Wrap(ErrInvalidTokenSupplyCap, strTokenSupplyCap)
	}
	return nil
}

func ValidateSymbol(maxLen uint64, classSymbol string) error {
	len := len(classSymbol)
	if len > int(maxLen) {
		return sdkerrors.Wrap(ErrClassSymbolInvalidLength, classSymbol)
	}
	return nil
}

func ValidateDescription(maxLen uint64, classDescription string) error {
	len := len(classDescription)
	if len > int(maxLen) {
		return sdkerrors.Wrap(ErrClassDescriptionInvalidLength, classDescription)
	}
	return nil
}
