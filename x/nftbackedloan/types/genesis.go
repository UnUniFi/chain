package types

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	params := DefaultParams()
	return &GenesisState{
		Params: params,
	}
}

func (gs GenesisState) Validate() error {
	return nil
}
