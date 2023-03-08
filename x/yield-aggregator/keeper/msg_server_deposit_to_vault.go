package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DepositToVault(ctx context.Context, msg *types.MsgDepositToVault) (*types.MsgDepositToVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		// TODO
		return nil, nil
	}

	err = k.Keeper.MintLPToken(sdkCtx, sender, msg.VaultId, msg.Amount.Amount)
	if err != nil {
		// TODO
		return nil, nil
	}

	return &types.MsgDepositToVaultResponse{}, nil
}
