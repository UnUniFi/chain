package types

import (
	"bytes"
	"fmt"
	"time"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	params := DefaultParams()
	return &GenesisState{
		Params:               params,
		CdpAccumulationTimes: GenesisAccumulationTimes{},
		CdpMintingClaims:     DefaultCdpClaims,
		Denoms:               DefaultGenesisDenoms(),
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
	if err := GenesisAccumulationTimes(gs.CdpAccumulationTimes).Validate(); err != nil {
		return err
	}
	if err := gs.Denoms.Validate(); err != nil {
		return err
	}

	return CdpMintingClaims(gs.CdpMintingClaims).Validate()
}

// NewGenesisState returns a new genesis state
func NewGenesisState(params Params, jpyxAccumTimes GenesisAccumulationTimes, c CdpMintingClaims, denoms *GenesisDenoms) GenesisState {
	return GenesisState{
		Params:               params,
		CdpAccumulationTimes: jpyxAccumTimes,
		CdpMintingClaims:     c,
		Denoms:               denoms,
	}
}

// Equal checks whether two gov GenesisState structs are equivalent
func (gs GenesisState) Equal(gs2 GenesisState) bool {
	b1 := ModuleCdc.MustMarshalBinaryBare(&gs)
	b2 := ModuleCdc.MustMarshalBinaryBare(&gs2)
	return bytes.Equal(b1, b2)
}

// IsEmpty returns true if a GenesisState is empty
func (gs GenesisState) IsEmpty() bool {
	return gs.Equal(GenesisState{})
}

// NewGenesisAccumulationTime returns a new GenesisAccumulationTime
func NewGenesisAccumulationTime(ctype string, prevTime time.Time) GenesisAccumulationTime {
	return GenesisAccumulationTime{
		CollateralType:           ctype,
		PreviousAccumulationTime: prevTime,
	}
}

// GenesisAccumulationTimes slice of GenesisAccumulationTime
type GenesisAccumulationTimes []GenesisAccumulationTime

// Validate performs validation of GenesisAccumulationTimes
func (gats GenesisAccumulationTimes) Validate() error {
	for _, gat := range gats {
		if err := gat.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate performs validation of GenesisAccumulationTime
func (gat GenesisAccumulationTime) Validate() error {
	if len(gat.CollateralType) == 0 {
		return fmt.Errorf("genesis accumulation time's collateral type must be defined")
	}
	return nil
}

func DefaultGenesisDenoms() *GenesisDenoms {
	return &GenesisDenoms{
		PrincipalDenom:        DefaultPrincipalDenom,
		CdpMintingRewardDenom: DefaultCDPMintingRewardDenom,
	}
}

func (denoms *GenesisDenoms) Validate() error {
	if len(denoms.PrincipalDenom) == 0 {
		return fmt.Errorf("GovDenom is nil or empty")
	}
	if len(denoms.CdpMintingRewardDenom) == 0 {
		return fmt.Errorf("GovDenom is nil or empty")
	}

	return nil
}
