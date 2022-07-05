package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: the validation against:
// Name
// BaseTokenUri, ClassUri
// TokenSupplyCap
// MintingPermission
// Symbol
// Description

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

func ValidateName(className string) error {
	return nil
}

func ValidateUri(uri string) error {
	return nil
}

func ValidateTokenSupplyCap(tokenSupplyCap uint64) error {
	return nil
}

func ValidateSymbol(classSymbol string) error {
	return nil
}

func ValidateDescription(classDescription string) error {
	return nil
}
