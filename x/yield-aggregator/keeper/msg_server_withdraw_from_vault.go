package keeper

import (
	"context"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawFromVault(ctx context.Context, msg *types.MsgWithdrawFromVault) (*types.MsgWithdrawFromVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.BurnLPToken(sdkCtx, sender, msg.VaultId, msg.LpTokenAmount)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawFromVaultResponse{}, nil
}
