package yieldaggregator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/UnUniFi/chain/x/yieldaggregatorv1/keeper"

	"github.com/UnUniFi/chain/x/yieldaggregatorv1/types"
)

// NewYieldAggregatorProposalHandler creates a new governance Handler
func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ProposalAddYieldFarm:
			return handleProposalAddYieldFarm(ctx, k, c)
		case *types.ProposalUpdateYieldFarm:
			return handleProposalUpdateYieldFarm(ctx, k, c)
		case *types.ProposalStopYieldFarm:
			return handleProposalStopYieldFarm(ctx, k, c)
		case *types.ProposalRemoveYieldFarm:
			return handleProposalRemoveYieldFarm(ctx, k, c)
		case *types.ProposalAddYieldFarmTarget:
			return handleProposalAddYieldFarmTarget(ctx, k, c)
		case *types.ProposalUpdateYieldFarmTarget:
			return handleProposalUpdateYieldFarmTarget(ctx, k, c)
		case *types.ProposalStopYieldFarmTarget:
			return handleProposalStopYieldFarmTarget(ctx, k, c)
		case *types.ProposalRemoveYieldFarmTarget:
			return handleProposalRemoveYieldFarmTarget(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func handleProposalAddYieldFarm(ctx sdk.Context, k keeper.Keeper, p *types.ProposalAddYieldFarm) error {
	return k.AddAssetManagementAccount(ctx, p.Account.Id, p.Account.Name)
}

func handleProposalUpdateYieldFarm(ctx sdk.Context, k keeper.Keeper, p *types.ProposalUpdateYieldFarm) error {
	return k.UpdateAssetManagementAccount(ctx, *p.Account)
}

func handleProposalStopYieldFarm(ctx sdk.Context, k keeper.Keeper, p *types.ProposalStopYieldFarm) error {
	acc := k.GetAssetManagementAccount(ctx, p.Id)
	if acc.Id == "" {
		return types.ErrAssetManagementAccountDoesNotExists
	}
	acc.Enabled = false
	k.SetAssetManagementAccount(ctx, acc)
	return nil
}

func handleProposalRemoveYieldFarm(ctx sdk.Context, k keeper.Keeper, p *types.ProposalRemoveYieldFarm) error {
	acc := k.GetAssetManagementAccount(ctx, p.Id)
	if acc.Id == "" {
		return types.ErrAssetManagementAccountDoesNotExists
	}
	k.DeleteAssetManagementAccount(ctx, p.Id)
	k.DeleteAssetManagementTargetsOfAccount(ctx, p.Id)
	return nil
}

func handleProposalAddYieldFarmTarget(ctx sdk.Context, k keeper.Keeper, p *types.ProposalAddYieldFarmTarget) error {
	k.SetAssetManagementTarget(ctx, *p.Target)
	return nil
}

func handleProposalUpdateYieldFarmTarget(ctx sdk.Context, k keeper.Keeper, p *types.ProposalUpdateYieldFarmTarget) error {
	// TODO: should automatically withdraw all the funds from the target
	target := k.GetAssetManagementTarget(ctx, p.Target.AssetManagementAccountId, p.Target.Id)
	if target.Id == "" {
		return types.ErrNoAssetManagementTargetExists
	}
	k.SetAssetManagementTarget(ctx, *p.Target)
	return nil
}

func handleProposalStopYieldFarmTarget(ctx sdk.Context, k keeper.Keeper, p *types.ProposalStopYieldFarmTarget) error {
	// TODO: should automatically withdraw all the funds from the target

	target := k.GetAssetManagementTarget(ctx, p.AssetManagementAccountId, p.Id)
	if target.Id == "" {
		return types.ErrNoAssetManagementTargetExists
	}
	target.Enabled = false
	k.SetAssetManagementTarget(ctx, target)
	return nil
}

func handleProposalRemoveYieldFarmTarget(ctx sdk.Context, k keeper.Keeper, p *types.ProposalRemoveYieldFarmTarget) error {
	// TODO: should ensure that all the tokens are unbonded via stop proposal

	target := k.GetAssetManagementTarget(ctx, p.AssetManagementAccountId, p.Id)
	if target.Id == "" {
		return types.ErrNoAssetManagementTargetExists
	}
	k.DeleteAssetManagementTarget(ctx, p.AssetManagementAccountId, p.Id)
	return nil
}
