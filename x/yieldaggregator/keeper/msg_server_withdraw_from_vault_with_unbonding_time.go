package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) WithdrawFromVaultWithUnbondingTime(ctx context.Context, msg *types.MsgWithdrawFromVaultWithUnbondingTime) (*types.MsgWithdrawFromVaultWithUnbondingTimeResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.BurnLPTokenAndRedeem(sdkCtx, sender, msg.VaultId, msg.LpTokenAmount)
	if err != nil {
		return nil, err
	}

	// TODO: if bonded amount < withdraw with unbonding time - execute, otherwise fail

	return &types.MsgWithdrawFromVaultWithUnbondingTimeResponse{}, nil
}
