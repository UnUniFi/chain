package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) DepositToVault(ctx context.Context, msg *types.MsgDepositToVault) (*types.MsgDepositToVaultResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.DepositAndMintLPToken(sdkCtx, sender, msg.VaultId, msg.Amount.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositToVaultResponse{}, nil
}
