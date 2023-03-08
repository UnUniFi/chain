package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransferVaultOwnership(ctx context.Context, msg *types.MsgTransferVaultOwnership) (*types.MsgTransferVaultOwnershipResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	vault, found := k.Keeper.GetVault(sdkCtx, msg.VaultId)
	if !found {
		// TODO
		return nil, nil
	}

	if vault.Owner != msg.Sender {
		// TODO
		return nil, nil
	}

	_, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		// TODO
		return nil, err
	}

	vault.Owner = msg.Recipient
	k.Keeper.SetVault(sdkCtx, vault)

	return &types.MsgTransferVaultOwnershipResponse{}, nil
}
