package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftfactory/types"
)

// InitGenesis initializes the tokenfactory module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, genClass := range genState.GetClasses() {
		creator, _, err := types.DeconstructDenom(genClass.GetClassId())
		if err != nil {
			panic(err)
		}
		err = k.createClassAfterValidation(ctx, creator, genClass.GetClassId())
		if err != nil {
			panic(err)
		}
		err = k.setAuthorityMetadata(ctx, genClass.GetClassId(), genClass.GetAuthorityMetadata())
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the tokenfactory module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genClasses := []types.GenesisClass{}
	iterator := k.GetAllDenomsIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		class := string(iterator.Value())

		authorityMetadata, err := k.GetAuthorityMetadata(ctx, class)
		if err != nil {
			panic(err)
		}

		genClasses = append(genClasses, types.GenesisClass{
			ClassId:           class,
			AuthorityMetadata: authorityMetadata,
		})
	}

	return &types.GenesisState{
		Classes: genClasses,
		Params:  k.GetParams(ctx),
	}
}
