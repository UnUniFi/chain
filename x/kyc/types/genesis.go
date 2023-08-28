package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ProviderList:     []Provider{},
		VerificationList: []Verification{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in provider
	providerIdMap := make(map[uint64]bool)
	providerCount := gs.GetProviderCount()
	for _, elem := range gs.ProviderList {
		if _, ok := providerIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for provider")
		}
		if elem.Id >= providerCount {
			return fmt.Errorf("provider id should be lower or equal than the last id")
		}
		providerIdMap[elem.Id] = true
	}
	// Check for duplicated index in verification
	verificationIndexMap := make(map[string]struct{})

	for _, elem := range gs.VerificationList {
		index := string(VerificationKey(elem.Address, elem.ProviderId))
		if _, ok := verificationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for verification")
		}
		verificationIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
