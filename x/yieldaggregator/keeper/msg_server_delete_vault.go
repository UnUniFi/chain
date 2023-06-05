package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
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

	// TODO: reenable positive fund check - just disabled to remove invalid vaults on testnet
	// // ensure no funds available on the vault
	// totalVaultAmount := k.VaultAmountTotal(ctx, vault)
	// if totalVaultAmount.IsPositive() {
	// 	return nil, types.ErrVaultHasPositiveBalance
	// }
	// TODO: add owner check

	// transfer deposit
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(vault.OwnerDeposit))
	if err != nil {
		return nil, err
	}

	k.Keeper.RemoveVault(ctx, msg.VaultId)
	return &types.MsgDeleteVaultResponse{}, nil
}
