package kyc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/kyc/keeper"
	"github.com/UnUniFi/chain/x/kyc/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the provider
	for _, elem := range genState.ProviderList {
		k.SetProvider(ctx, elem)
	}

	// Set provider count
	k.SetProviderCount(ctx, genState.ProviderCount)
	// Set all the verification
	for _, elem := range genState.VerificationList {
		k.SetVerification(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ProviderList = k.GetAllProvider(ctx)
	genesis.ProviderCount = k.GetProviderCount(ctx)
	genesis.VerificationList = k.GetAllVerification(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
