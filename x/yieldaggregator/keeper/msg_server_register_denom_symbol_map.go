package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) RegisterDenomSymbolMap(ctx context.Context, msg *types.MsgRegisterDenomSymbolMap) (*types.MsgRegisterDenomSymbolMapResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	for _, dsm := range msg.Mappings {
		k.SetDenomSymbolMap(sdkCtx, dsm.Key, dsm.Value)
	}

	return &types.MsgRegisterDenomSymbolMapResponse{}, nil
}
