package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

func ValidateCreateClass(
	params Params,
	name, symbol, uri, description string,
	mintingPermission MintingPermission,
	tokenSupplyCap uint64,
) error {
	if err := ValidateClassName(params.MinClassNameLen, params.MaxClassNameLen, name); err != nil {
		return err
	}
	if err := ValidateSymbol(params.MaxSymbolLen, symbol); err != nil {
		return err
	}
	if err := ValidateUri(params.MinUriLen, params.MaxUriLen, uri); err != nil {
		return err
	}
	if err := ValidateDescription(params.MaxDescriptionLen, description); err != nil {
		return err
	}
	if err := ValidateMintingPermission(mintingPermission, nil, nil); err != nil {
		return err
	}
	if err := ValidateTokenSupplyCap(params.MaxNFTSupplyCap, tokenSupplyCap); err != nil {
		return err
	}
	return nil
}

func ValidateMintNFT(
	params Params,
	mintingPermission MintingPermission,
	owner, minter sdk.AccAddress,
	uri, nftID string,
	currentTokenSupply, tokenSupplyCap uint64,
) error {
	if err := ValidateMintingPermission(mintingPermission, owner, minter); err != nil {
		return err
	}
	if err := ValidateUri(params.MinUriLen, params.MaxUriLen, uri); err != nil {
		return err
	}
	if err := ValidateTokenSupply(currentTokenSupply, tokenSupplyCap); err != nil {
		return err
	}
	if err := nfttypes.ValidateNFTID(nftID); err != nil {
		return err
	}
	return nil
}

func ValidateClassAttributes(classAttributes ClassAttributes, params Params) error {
	if err := ValidateUri(params.MinUriLen, params.MaxUriLen, classAttributes.BaseTokenUri); err != nil {
		return err
	}

	if err := ValidateTokenSupplyCap(params.MaxNFTSupplyCap, classAttributes.TokenSupplyCap); err != nil {
		return err
	}

	return nil
}

func ValidateMintingPermission(mintingPermission MintingPermission, owner, minter sdk.AccAddress) error {
	switch mintingPermission {
	case 0:
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
	if len == 0 {
		return nil
	}

	if len < int(minLen) || len > int(maxLen) {
		return sdkerrors.Wrap(ErrUriInvalidLength, uri)
	}
	return nil
}

func ValidateTokenSupplyCap(maxCap, tokenSupplyCap uint64) error {
	if tokenSupplyCap > maxCap {
		strTokenSupplyCap := strconv.FormatUint(tokenSupplyCap, 10)
		return sdkerrors.Wrap(ErrInvalidTokenSupplyCap, strTokenSupplyCap)
	}
	return nil
}

func ValidateTokenSupply(tokenSupply, tokenSupplyCap uint64) error {
	if tokenSupply >= tokenSupplyCap {
		strTokenSupply := strconv.FormatUint(tokenSupply, 10)
		return sdkerrors.Wrap(ErrTokenSupplyCapOver, strTokenSupply)
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
	if len == 0 {
		return nil
	}
	if len > int(maxLen) {
		return sdkerrors.Wrap(ErrClassDescriptionInvalidLength, classDescription)
	}
	return nil
}
