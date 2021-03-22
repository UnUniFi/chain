package types

import (
	"bytes"
	fmt "fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DefaultParams(),
		Cdps{},
		Deposits{},
		DefaultCdpStartingID,
		DefaultDebtDenom,
		DefaultGovDenom,
		GenesisAccumulationTimes{},
		GenesisTotalPrincipals{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in cdp
	// cdpIdMap := make(map[string]bool)

	// for _, elem := range gs.CdpList {
	// 	if _, ok := cdpIdMap[elem.Id]; ok {
	// 		return fmt.Errorf("duplicated id for cdp")
	// 	}
	// 	cdpIdMap[elem.Id] = true
	// }

	// return nil

	if err := gs.Params.Validate(); err != nil {
		return err
	}

	if err := Cdps(gs.Cdps).Validate(); err != nil {
		return err
	}

	if err := Deposits(gs.Deposits).Validate(); err != nil {
		return err
	}

	if err := GenesisAccumulationTimes(gs.PreviousAccumulationTimes).Validate(); err != nil {
		return err
	}

	if err := GenesisTotalPrincipals(gs.TotalPrincipals).Validate(); err != nil {
		return err
	}

	if err := sdk.ValidateDenom(gs.DebtDenom); err != nil {
		return fmt.Errorf(fmt.Sprintf("debt denom invalid: %v", err))
	}

	if err := sdk.ValidateDenom(gs.GovDenom); err != nil {
		return fmt.Errorf(fmt.Sprintf("gov denom invalid: %v", err))
	}

	return nil
}

func NewGenesisState(params Params, cdps Cdps, deposits Deposits, startingCdpID uint64,
	debtDenom, govDenom string, prevAccumTimes GenesisAccumulationTimes,
	totalPrincipals GenesisTotalPrincipals) GenesisState {
	return GenesisState{
		Params:                    params,
		Cdps:                      cdps,
		Deposits:                  deposits,
		StartingCdpId:             startingCdpID,
		DebtDenom:                 debtDenom,
		GovDenom:                  govDenom,
		PreviousAccumulationTimes: prevAccumTimes,
		TotalPrincipals:           totalPrincipals,
	}
}

func validateSavingsRateDistributed(i interface{}) error {
	savingsRateDist, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if savingsRateDist.IsNegative() {
		return fmt.Errorf("savings rate distributed should not be negative: %s", savingsRateDist)
	}

	return nil
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

// NewGenesisTotalPrincipal returns a new GenesisTotalPrincipal
func NewGenesisTotalPrincipal(ctype string, principal sdk.Int) GenesisTotalPrincipal {
	return GenesisTotalPrincipal{
		CollateralType: ctype,
		TotalPrincipal: principal,
	}
}

// GenesisTotalPrincipals slice of GenesisTotalPrincipal
type GenesisTotalPrincipals []GenesisTotalPrincipal

// Validate performs validation of GenesisTotalPrincipal
func (gtp GenesisTotalPrincipal) Validate() error {
	if gtp.TotalPrincipal.IsNegative() {
		return fmt.Errorf("total principal should be positive, is %s for %s", gtp.TotalPrincipal, gtp.CollateralType)
	}
	return nil
}

// Validate performs validation of GenesisTotalPrincipals
func (gtps GenesisTotalPrincipals) Validate() error {
	for _, gtp := range gtps {
		if err := gtp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// NewGenesisAccumulationTime returns a new GenesisAccumulationTime
func NewGenesisAccumulationTime(ctype string, prevTime time.Time, factor sdk.Dec) GenesisAccumulationTime {
	return GenesisAccumulationTime{
		CollateralType:           ctype,
		PreviousAccumulationTime: prevTime,
		InterestFactor:           factor,
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
	if gat.InterestFactor.LT(sdk.OneDec()) {
		return fmt.Errorf("interest factor should be ≥ 1.0, is %s for %s", gat.InterestFactor, gat.CollateralType)
	}
	return nil
}
