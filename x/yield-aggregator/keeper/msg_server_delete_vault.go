package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteVault(ctx context.Context, msg *types.MsgDeleteVault) (*types.MsgDeleteVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	vault, found := k.Keeper.GetVault(sdkCtx, msg.VaultId)
	if !found {
		// TODO
		return nil, nil
	}

	// transfer deposit
	err = k.bankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, sender, sdk.NewCoins(vault.OwnerDeposit))
	if err != nil {
		return nil, err
	}

	k.Keeper.RemoveVault(sdkCtx, msg.VaultId)

	return &types.MsgDeleteVaultResponse{}, nil
}
