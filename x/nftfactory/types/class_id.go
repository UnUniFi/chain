package types

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleDenomPrefix = "factory"
	MaxSubdenomLength = 512
	MaxHrpLength      = 16
	// MaxCreatorLength = 59 + MaxHrpLength
	MaxCreatorLength = 59 + MaxHrpLength
)

// GetClassId constructs a denom string for tokens created by tokenfactory
// based on an input creator address and a subdenom
// The denom constructed is factory/{creator}/{subdenom}
func GetClassId(creator, subclass string) (string, error) {
	if len(subclass) > MaxSubdenomLength {
		return "", ErrSubclassTooLong
	}
	if len(subclass) > MaxCreatorLength {
		return "", ErrCreatorTooLong
	}
	if strings.Contains(creator, "/") {
		return "", ErrInvalidCreator
	}
	denom := strings.Join([]string{ModuleDenomPrefix, creator, subclass}, "/")
	return denom, sdk.ValidateDenom(denom)
}

// DeconstructClassId takes a token denom string and verifies that it is a valid
// denom of the tokenfactory module, and is of the form `factory/{creator}/{subdenom}`
// If valid, it returns the creator address and subdenom
func DeconstructClassId(classId string) (creator, subclass string, err error) {
	err = sdk.ValidateDenom(classId)
	if err != nil {
		return "", "", err
	}

	strParts := strings.Split(classId, "/")
	if len(strParts) < 3 {
		return "", "", errorsmod.Wrapf(ErrInvalidClassId, "not enough parts of class id %s", classId)
	}

	if strParts[0] != ModuleDenomPrefix {
		return "", "", errorsmod.Wrapf(ErrInvalidClassId, "class id prefix is incorrect. Is: %s.  Should be: %s", strParts[0], ModuleDenomPrefix)
	}

	creator = strParts[1]
	_, err = sdk.AccAddressFromBech32(creator)
	if err != nil {
		return "", "", errorsmod.Wrapf(ErrInvalidClassId, "Invalid creator address (%s)", err)
	}

	// Handle the case where a denom has a slash in its subdenom. For example,
	// when we did the split, we'd turn factory/accaddr/atomderivative/sikka into ["factory", "accaddr", "atomderivative", "sikka"]
	// So we have to join [2:] with a "/" as the delimiter to get back the correct subdenom which should be "atomderivative/sikka"
	subclass = strings.Join(strParts[2:], "/")

	return creator, subclass, nil
}
