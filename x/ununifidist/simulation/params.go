package simulation

import (
	"fmt"
	"math/rand"

	"github.com/UnUniFi/chain/x/ununifidist/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	active := genRandomActive(r)
	periods := genRandomPeriods(r, simtypes.RandTimestamp(r))
	if err := types.NewParams(active, periods).Validate(); err != nil {
		panic(err)
	}

	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyActive),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%t", active)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyPeriods),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%v", periods)
			},
		),
	}
}
