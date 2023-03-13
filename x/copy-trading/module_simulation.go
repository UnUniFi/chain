package copy_trading

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/UnUniFi/chain/testutil/sample"
	copytradingsimulation "github.com/UnUniFi/chain/x/copy-trading/simulation"
	"github.com/UnUniFi/chain/x/copy-trading/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = copytradingsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateExemplaryTrader = "op_weight_msg_exemplary_trader"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateExemplaryTrader int = 100

	opWeightMsgUpdateExemplaryTrader = "op_weight_msg_exemplary_trader"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateExemplaryTrader int = 100

	opWeightMsgDeleteExemplaryTrader = "op_weight_msg_exemplary_trader"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteExemplaryTrader int = 100

	opWeightMsgCreateTracing = "op_weight_msg_tracing"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTracing int = 100

	opWeightMsgUpdateTracing = "op_weight_msg_tracing"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTracing int = 100

	opWeightMsgDeleteTracing = "op_weight_msg_tracing"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTracing int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	copytradingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ExemplaryTraderList: []types.ExemplaryTrader{
			{
				Address: sample.AccAddress(),
			},
			{
				Address: sample.AccAddress(),
			},
		},
		TracingList: []types.Tracing{
			{
				Address: sample.AccAddress(),
			},
			{
				Address: sample.AccAddress(),
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&copytradingGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateExemplaryTrader int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateExemplaryTrader, &weightMsgCreateExemplaryTrader, nil,
		func(_ *rand.Rand) {
			weightMsgCreateExemplaryTrader = defaultWeightMsgCreateExemplaryTrader
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateExemplaryTrader,
		copytradingsimulation.SimulateMsgCreateExemplaryTrader(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateExemplaryTrader int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateExemplaryTrader, &weightMsgUpdateExemplaryTrader, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateExemplaryTrader = defaultWeightMsgUpdateExemplaryTrader
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateExemplaryTrader,
		copytradingsimulation.SimulateMsgUpdateExemplaryTrader(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteExemplaryTrader int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteExemplaryTrader, &weightMsgDeleteExemplaryTrader, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteExemplaryTrader = defaultWeightMsgDeleteExemplaryTrader
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteExemplaryTrader,
		copytradingsimulation.SimulateMsgDeleteExemplaryTrader(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTracing int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTracing, &weightMsgCreateTracing, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTracing = defaultWeightMsgCreateTracing
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTracing,
		copytradingsimulation.SimulateMsgCreateTracing(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTracing int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTracing, &weightMsgDeleteTracing, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTracing = defaultWeightMsgDeleteTracing
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTracing,
		copytradingsimulation.SimulateMsgDeleteTracing(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
