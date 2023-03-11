package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ExemplaryTraderList: []ExemplaryTrader{},
		TracingList:         []Tracing{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in exemplaryTrader
	exemplaryTraderIndexMap := make(map[string]struct{})

	for _, elem := range gs.ExemplaryTraderList {
		index := string(ExemplaryTraderKey(elem.Address))
		if _, ok := exemplaryTraderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for exemplaryTrader")
		}
		exemplaryTraderIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in tracing
	tracingIndexMap := make(map[string]struct{})

	for _, elem := range gs.TracingList {
		index := string(TracingKey(elem.Address))
		if _, ok := tracingIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tracing")
		}
		tracingIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
