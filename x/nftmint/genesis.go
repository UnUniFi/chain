package nftmint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftmint/keeper"
	"github.com/UnUniFi/chain/x/nftmint/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, nftKeeper types.NftKeeper, gs types.GenesisState) {
	k.SetParamSet(ctx, gs.Params)

	for _, classAttributes := range gs.ClassAttributesList {
		if err := InitClassRelatingData(ctx, k, nftKeeper, *classAttributes); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis export genesis state for nftmarket module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	classAttributesList := k.GetClassAttributesList(ctx)

	return types.GenesisState{
		Params:              k.GetParamSet(ctx),
		ClassAttributesList: classAttributesList,
	}
}

func InitClassRelatingData(ctx sdk.Context, k keeper.Keeper, nftKeeper types.NftKeeper, classAttributes types.ClassAttributes) error {
	class, exists := nftKeeper.GetClass(ctx, classAttributes.ClassId)
	if !exists {
		return sdkerrors.Wrap(nfttypes.ErrClassNotExists, classAttributes.ClassId)
	}

	params := k.GetParamSet(ctx)
	if err := types.ValidateCreateClass(
		params,
		class.Name, class.Symbol, classAttributes.BaseTokenUri, class.Description,
		classAttributes.MintingPermission,
		classAttributes.TokenSupplyCap,
	); err != nil {
		return err
	}

	if err := k.SetClassAttributes(ctx, types.NewClassAttributes(
		class.Id,
		classAttributes.Owner.AccAddress(),
		classAttributes.BaseTokenUri,
		classAttributes.MintingPermission,
		classAttributes.TokenSupplyCap,
	)); err != nil {
		return err
	}

	owningClassIdList := k.AddClassIDToOwningClassIdList(ctx, classAttributes.Owner.AccAddress(), class.Id)
	if err := k.SetOwningClassIdList(ctx, owningClassIdList); err != nil {
		return err
	}

	classNameIdList := k.AddClassNameIdList(ctx, class.Name, class.Id)
	if err := k.SetClassNameIdList(ctx, classNameIdList); err != nil {
		return err
	}

	return nil
}
