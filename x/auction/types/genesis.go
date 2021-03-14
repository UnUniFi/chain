package types

import "fmt"

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		AuctionList: []*Auction{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in auction
	auctionIdMap := make(map[string]bool)

	for _, elem := range gs.AuctionList {
		if _, ok := auctionIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for auction")
		}
		auctionIdMap[elem.Id] = true
	}

	return nil
}
