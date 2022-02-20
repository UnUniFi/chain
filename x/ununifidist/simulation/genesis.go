package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/UnUniFi/chain/x/ununifidist/types"
)

const (
	SecondsPerYear = 31536000
	BaseAprPadding = "0.000000003022265980"
)

// RandomizedGenState generates a random GenesisState for kavadist module
func RandomizedGenState(simState *module.SimulationState) {
	params := genRandomParams(simState)
	if err := params.Validate(); err != nil {
		panic(err)
	}

	ununifidistGenesis := types.NewGenesisState(params, types.DefaultPreviousBlockTime, types.DefaultGovDenom)
	if err := ununifidistGenesis.Validate(); err != nil {
		panic(err)
	}

	bz, err := json.MarshalIndent(&ununifidistGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&ununifidistGenesis)
}

func genRandomParams(simState *module.SimulationState) types.Params {
	periods := genRandomPeriods(simState.Rand, simState.GenTimestamp)
	params := types.NewParams(true, periods)
	return params
}

func genRandomPeriods(r *rand.Rand, timestamp time.Time) []types.Period {
	var periods []types.Period
	numPeriods := simulation.RandIntBetween(r, 1, 10)
	periodStart := timestamp
	for i := 0; i < numPeriods; i++ {
		durationMultiplier := simulation.RandIntBetween(r, 1, 3)
		duration := time.Duration(int64(24*durationMultiplier)) * time.Hour
		periodEnd := periodStart.Add(duration)
		inflation := genRandomInflation(r)
		period := types.NewPeriod(periodStart, periodEnd, inflation)
		periods = append(periods, period)
		periodStart = periodEnd
	}
	return periods
}

func genRandomInflation(r *rand.Rand) sdk.Dec {
	aprPadding, _ := sdk.NewDecFromStr(BaseAprPadding)
	extraAprInflation := simulation.RandomDecAmount(r, sdk.MustNewDecFromStr("0.25"))
	for extraAprInflation.LT(aprPadding) {
		extraAprInflation = extraAprInflation.Add(aprPadding)
	}
	aprInflation := sdk.OneDec().Add(extraAprInflation)

	inflationSpr, err := aprInflation.ApproxRoot(uint64(SecondsPerYear))
	if err != nil {
		panic(fmt.Sprintf("error generating random inflation %v", err))
	}
	return inflationSpr
}

func genRandomActive(r *rand.Rand) bool {
	threshold := 50
	value := simulation.RandIntBetween(r, 1, 100)
	return value > threshold
}
