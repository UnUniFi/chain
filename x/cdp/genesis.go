package cdp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/cdp/keeper"
	"github.com/lcnem/jpyx/x/cdp/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the cdp
	for _, elem := range genState.CdpList {
		k.SetCdp(ctx, *elem)
	}

	// Set cdp count
	k.SetCdpCount(ctx, int64(len(genState.CdpList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all cdp
	cdpList := k.GetAllCdp(ctx)
	for _, elem := range cdpList {
		elem := elem
		genesis.CdpList = append(genesis.CdpList, &elem)
	}

	return genesis
}
