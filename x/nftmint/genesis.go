package nftmint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmint/keeper"
	"github.com/UnUniFi/chain/x/nftmint/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper types.AccountKeeper, gs types.GenesisState) {
	k.SetParamSet(ctx, gs.Params)
	for _, classAttributes := range gs.ClassAttributesList {
		k.SetClassAttributes(ctx, *classAttributes)
	}

	for _, classNameIdList := range gs.ClassNameIdLists {
		k.SetClassNameIdList(ctx, *classNameIdList)
	}

	for _, owningClassIdList := range gs.OwningClassIdLists {
		k.SetOwningClassIdList(ctx, *owningClassIdList)
	}
}

// ExportGenesis export genesis state for nftmarket module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	classAttributesList := k.GetClassAttributesList(ctx)
	owningClassIdLists := k.GetOwningClassIdLists(ctx)
	classNameIdLists := k.GetClassNameIdLists(ctx)

	return types.GenesisState{
		Params:              k.GetParamSet(ctx),
		ClassAttributesList: classAttributesList,
		OwningClassIdLists:  owningClassIdLists,
		ClassNameIdLists:    classNameIdLists,
	}
}
