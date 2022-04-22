package types

import (
	"bytes"
	fmt "fmt"
	time "time"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:            DefaultParams(),
		PreviousBlockTime: DefaultPreviousBlockTime,
		GovDenom:          DefaultGovDenom,
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	// return nil
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	if gs.PreviousBlockTime.Equal(time.Time{}) {
		return fmt.Errorf("previous block time not set")
	}
	if len(gs.GovDenom) == 0 {
		return fmt.Errorf("GovDenom is nil or empty")
	}
	return nil
}

// NewGenesisState returns a new genesis state
func NewGenesisState(params Params, previousBlockTime time.Time, govDenom string) GenesisState {
	return GenesisState{
		Params:            params,
		PreviousBlockTime: previousBlockTime,
		GovDenom:          govDenom,
	}
}

// Equal checks whether two gov GenesisState structs are equivalent
func (gs GenesisState) Equal(gs2 GenesisState) bool {
	b1 := ModuleCdc.MustMarshal(&gs)
	b2 := ModuleCdc.MustMarshal(&gs2)
	return bytes.Equal(b1, b2)
}

// IsEmpty returns true if a GenesisState is empty
func (gs GenesisState) IsEmpty() bool {
	return gs.Equal(GenesisState{})
}
