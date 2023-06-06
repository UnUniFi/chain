package ununifidist

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/deprecated/x/ununifidist/keeper"
	"github.com/UnUniFi/chain/deprecated/x/ununifidist/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, gs types.GenesisState) {
	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	k.SetParams(ctx, gs.Params)

	// only set the previous block time if it's different than default
	if !gs.PreviousBlockTime.Equal(types.DefaultPreviousBlockTime) {
		k.SetPreviousBlockTime(ctx, gs.PreviousBlockTime)
	}

	// check if the module account exists
	moduleAcc := accountKeeper.GetModuleAccount(ctx, types.UnunifidistMacc)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.UnunifidistMacc))
	}

	k.SetGovDenom(ctx, gs.GovDenom)
}

// ExportGenesis export genesis state for cdp module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	params := k.GetParams(ctx)
	previousBlockTime, found := k.GetPreviousBlockTime(ctx)
	if !found {
		previousBlockTime = types.DefaultPreviousBlockTime
	}
	govDenom, found := k.GetGovDenom(ctx)
	if !found {
		govDenom = types.DefaultGovDenom
	}
	return types.NewGenesisState(params, previousBlockTime, govDenom)
}
