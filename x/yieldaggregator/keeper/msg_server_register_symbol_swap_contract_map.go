package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k msgServer) RegisterSymbolSwapContractMap(ctx context.Context, msg *types.MsgSymbolSwapContractMap) (*types.MsgSymbolSwapContractMapResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	for _, dsm := range msg.Mappings {
		k.SetSymbolSwapContractMap(sdkCtx, dsm.Key, dsm.Value)
	}

	return &types.MsgSymbolSwapContractMapResponse{}, nil
}
