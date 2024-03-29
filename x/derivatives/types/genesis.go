package types

import (
	"fmt"
	// this line is used by starport scaffolding # genesis/types/import
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	positionIdMap := make(map[string]bool)
	for _, elem := range gs.Positions {
		if _, ok := positionIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for position")
		}
		if _, err := sdk.AccAddressFromBech32(elem.OpenerAddress); err != nil {
			return fmt.Errorf("invalid address")
		}
		positionIdMap[elem.Id] = true
	}
	return gs.Params.Validate()
}
