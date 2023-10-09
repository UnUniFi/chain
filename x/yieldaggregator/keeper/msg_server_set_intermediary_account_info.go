package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) SetIntermediaryAccountInfo(ctx context.Context, msg *types.MsgSetIntermediaryAccountInfo) (*types.MsgSetIntermediaryAccountInfoResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	k.Keeper.SetIntermediaryAccountInfo(sdkCtx, msg.Addrs)
	return &types.MsgSetIntermediaryAccountInfoResponse{}, nil
}
