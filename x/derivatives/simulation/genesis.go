package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/UnUniFi/chain/x/derivatives/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func RandomGenesisBool(r *rand.Rand) bool {
	// 90% chance
	return r.Int63n(100) < 90
}

func RandomizedGenState(simState *module.SimulationState) {
	sdk.NewCoins()
	// numAccs := int64(len(simState.Accounts))

	bankGenesis := types.GenesisState{}

	paramsBytes, err := json.MarshalIndent(&bankGenesis.Params, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated bank parameters:\n%s\n", paramsBytes)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bankGenesis)
}
