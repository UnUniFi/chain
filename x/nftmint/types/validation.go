package types

// TODO: the validation against:
// Name
// BaseTokenUri, ClassUri
// TokenSupplyCap
// MintingPermission
// Symbol
// Description

func ValidateName(className string) error {
	return nil
}

func ValidateUri(uri string) error {
	return nil
}

func ValidateTokenSupplyCap(tokenSupplyCap uint64) error {
	return nil
}

func ValidateMintingPermission(MintingPermission MintingPermission) error {
	return nil
}

func ValidateSymbol(classSymbol string) error {
	return nil
}

func ValidateDescription(classDescription string) error {
	return nil
}
