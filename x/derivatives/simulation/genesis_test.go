package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/UnUniFi/chain/x/derivatives/simulation"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)
	var derivativesGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &derivativesGenesis)

	assert.Equal(t, derivativesGenesis.Params.PoolParams.QuoteTicker, "usd")
	assert.Len(t, derivativesGenesis.Params.PoolParams.AcceptedAssets, 2)
	assert.Equal(t, derivativesGenesis.Params.PoolParams.BaseLptMintFee, sdk.NewDecWithPrec(1, 2))
	assert.Equal(t, derivativesGenesis.Params.PoolParams.BaseLptRedeemFee, sdk.NewDecWithPrec(1, 2))
	assert.Equal(t, derivativesGenesis.Params.PoolParams.BorrowingFeeRatePerHour, sdk.NewDecWithPrec(1, 6))
	assert.Equal(t, derivativesGenesis.Params.PoolParams.ReportLiquidationRewardRate, sdk.NewDecWithPrec(1, 6))
	assert.Equal(t, derivativesGenesis.Params.PerpetualFutures.CommissionRate, sdk.NewDecWithPrec(1, 6))
	assert.Equal(t, derivativesGenesis.Params.PerpetualFutures.MarginMaintenanceRate, sdk.NewDecWithPrec(5, 1))
	assert.Equal(t, derivativesGenesis.Params.PerpetualFutures.ImaginaryFundingRateProportionalCoefficient, sdk.NewDecWithPrec(1, 4))
	assert.Len(t, derivativesGenesis.Params.PerpetualFutures.Markets, 2)
}

func TestRandomizedGenStateWithPanics(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	tests := []struct {
		simState module.SimulationState
		panicMsg string
	}{
		{
			module.SimulationState{}, "invalid memory address or nil pointer dereference",
		},
		{
			module.SimulationState{
				AppParams: make(simtypes.AppParams),
				Cdc:       cdc,
				Rand:      r,
			},
			"assignment to entry in nil map",
		},
	}

	for _, tt := range tests {
		require.Panicsf(t, func() { simulation.RandomizedGenState(&tt.simState) }, tt.panicMsg)
	}
}
