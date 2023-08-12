package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) DeleteVault(goCtx context.Context, msg *types.MsgDeleteVault) (*types.MsgDeleteVaultResponse, error) {
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	vault, found := k.Keeper.GetVault(ctx, msg.VaultId)
	if !found {
		return nil, types.ErrInvalidVaultId
	}

	// TODO: force all strategies to unstake for each holders

	if vault.Owner != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	// transfer deposit
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(vault.OwnerDeposit))
	if err != nil {
		return nil, err
	}

	k.Keeper.RemoveVault(ctx, msg.VaultId)
	return &types.MsgDeleteVaultResponse{}, nil
}
