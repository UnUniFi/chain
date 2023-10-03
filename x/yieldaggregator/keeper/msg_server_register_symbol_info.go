package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) RegisterSymbolInfo(ctx context.Context, msg *types.MsgSymbolInfos) (*types.MsgSymbolInfosResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	for _, dsm := range msg.Info {
		k.SetSymbolInfo(sdkCtx, dsm)
	}

	return &types.MsgSymbolInfosResponse{}, nil
}
