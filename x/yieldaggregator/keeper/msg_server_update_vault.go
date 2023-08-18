package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) UpdateVault(goCtx context.Context, msg *types.MsgUpdateVault) (*types.MsgUpdateVaultResponse, error) {
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	vault, found := k.GetVault(ctx, msg.Id)
	if !found {
		return nil, types.ErrVaultNotFound
	}

	vault.Name = msg.Name
	vault.Description = msg.Description
	vault.FeeCollectorAddress = msg.FeeCollectorAddress
	k.SetVault(ctx, vault)

	return &types.MsgUpdateVaultResponse{}, nil
}
