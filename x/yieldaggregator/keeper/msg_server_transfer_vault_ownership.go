package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) TransferVaultOwnership(ctx context.Context, msg *types.MsgTransferVaultOwnership) (*types.MsgTransferVaultOwnershipResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	vault, found := k.Keeper.GetVault(sdkCtx, msg.VaultId)
	if !found {
		return nil, types.ErrInvalidVaultId
	}

	if vault.Owner != msg.Sender {
		return nil, types.ErrNotVaultOwner
	}

	_, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	vault.Owner = msg.Recipient
	k.Keeper.SetVault(sdkCtx, vault)

	return &types.MsgTransferVaultOwnershipResponse{}, nil
}
