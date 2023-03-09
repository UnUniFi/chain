package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

func (k msgServer) WithdrawFromVault(ctx context.Context, msg *types.MsgWithdrawFromVault) (*types.MsgWithdrawFromVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.BurnLPTokenAndRedeem(sdkCtx, sender, msg.VaultId, msg.LpTokenAmount)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawFromVaultResponse{}, nil
}
