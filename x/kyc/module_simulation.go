package kyc

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"testchain/testutil/sample"
	kycsimulation "testchain/x/kyc/simulation"
	"testchain/x/kyc/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = kycsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateProvider = "op_weight_msg_provider"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProvider int = 100

	opWeightMsgUpdateProvider = "op_weight_msg_provider"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProvider int = 100

	opWeightMsgDeleteProvider = "op_weight_msg_provider"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProvider int = 100

	opWeightMsgCreateVerification = "op_weight_msg_verification"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVerification int = 100

	opWeightMsgUpdateVerification = "op_weight_msg_verification"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateVerification int = 100

	opWeightMsgDeleteVerification = "op_weight_msg_verification"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteVerification int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	kycGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ProviderList: []types.Provider{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ProviderCount: 2,
		VerificationList: []types.Verification{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&kycGenesis)
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

	var weightMsgCreateProvider int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProvider, &weightMsgCreateProvider, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProvider = defaultWeightMsgCreateProvider
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProvider,
		kycsimulation.SimulateMsgCreateProvider(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProvider int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateProvider, &weightMsgUpdateProvider, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProvider = defaultWeightMsgUpdateProvider
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProvider,
		kycsimulation.SimulateMsgUpdateProvider(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProvider int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteProvider, &weightMsgDeleteProvider, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProvider = defaultWeightMsgDeleteProvider
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProvider,
		kycsimulation.SimulateMsgDeleteProvider(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateVerification int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVerification, &weightMsgCreateVerification, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVerification = defaultWeightMsgCreateVerification
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVerification,
		kycsimulation.SimulateMsgCreateVerification(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateVerification int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateVerification, &weightMsgUpdateVerification, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateVerification = defaultWeightMsgUpdateVerification
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateVerification,
		kycsimulation.SimulateMsgUpdateVerification(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteVerification int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteVerification, &weightMsgDeleteVerification, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVerification = defaultWeightMsgDeleteVerification
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteVerification,
		kycsimulation.SimulateMsgDeleteVerification(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
