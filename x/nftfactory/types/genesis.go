package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:  DefaultParams(),
		Classes: []GenesisClass{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	err := gs.Params.Validate()
	if err != nil {
		return err
	}

	seenDenoms := map[string]bool{}

	for _, class := range gs.GetClasses() {
		if seenDenoms[class.GetClassId()] {
			return sdkerrors.Wrapf(ErrInvalidGenesis, "duplicate class id: %s", class.GetClassId())
		}
		seenDenoms[class.GetClassId()] = true

		_, _, err := DeconstructClassId(class.GetClassId())
		if err != nil {
			return err
		}

		if class.AuthorityMetadata.Admin != "" {
			_, err = sdk.AccAddressFromBech32(class.AuthorityMetadata.Admin)
			if err != nil {
				return sdkerrors.Wrapf(ErrInvalidAuthorityMetadata, "Invalid admin address (%s)", err)
			}
		}
	}

	return nil
}
