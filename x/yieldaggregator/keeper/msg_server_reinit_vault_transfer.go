package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) ReinitVaultTransfer(ctx context.Context, msg *types.MsgReinitVaultTransfer) (*types.MsgReinitVaultTransferResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	vault, found := k.Keeper.GetVault(sdkCtx, msg.VaultId)
	if !found {
		return nil, types.ErrVaultNotFound
	}

	strategy, found := k.Keeper.GetStrategy(sdkCtx, msg.StrategyDenom, msg.StrategyId)
	if !found {
		return nil, types.ErrStrategyNotFound
	}

	err := k.Keeper.ExecuteVaultTransfer(sdkCtx, vault, strategy, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgReinitVaultTransferResponse{}, nil
}
