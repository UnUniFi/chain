package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:                  DefaultParams(),
		AssetManagementAccounts: []AssetManagementAccount{},
		AssetManagementTargets:  []AssetManagementTarget{},
		FarmingOrders:           []FarmingOrder{},
		FarmingUnits:            []FarmingUnit{},
		UserDeposits:            []UserDeposit{},
		DailyPercents:           []DailyPercent{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	return gs.Params.Validate()
}
