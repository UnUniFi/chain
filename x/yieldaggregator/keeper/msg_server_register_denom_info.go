package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) RegisterDenomInfos(ctx context.Context, msg *types.MsgRegisterDenomInfos) (*types.MsgRegisterDenomInfosResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	for _, dsm := range msg.Info {
		k.SetDenomInfo(sdkCtx, dsm)
	}

	return &types.MsgRegisterDenomInfosResponse{}, nil
}
