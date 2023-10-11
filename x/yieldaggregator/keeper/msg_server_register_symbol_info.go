package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) RegisterSymbolInfos(ctx context.Context, msg *types.MsgRegisterSymbolInfos) (*types.MsgRegisterSymbolInfosResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	for _, dsm := range msg.Info {
		k.SetSymbolInfo(sdkCtx, dsm)
	}

	return &types.MsgRegisterSymbolInfosResponse{}, nil
}
