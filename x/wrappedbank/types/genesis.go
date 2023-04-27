package types

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{}
}

func (gs GenesisState) Validate() error {
	return nil
}

func NewGenesisState(params Params) GenesisState {
	return GenesisState{}
}
