package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteVault(goCtx context.Context, msg *types.MsgDeleteVault) (*types.MsgDeleteVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	vault, found := k.Keeper.GetVault(ctx, msg.VaultId)
	if !found {
		return nil, types.ErrInvalidVaultId
	}

	// ensure no funds available on the vault
	totalVaultAmount := k.VaultAmountTotal(ctx, vault)
	if totalVaultAmount.IsPositive() {
		return nil, types.ErrVaultHasPositiveBalance
	}

	// transfer deposit
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(vault.OwnerDeposit))
	if err != nil {
		return nil, err
	}

	k.Keeper.RemoveVault(ctx, msg.VaultId)
	return &types.MsgDeleteVaultResponse{}, nil
}
