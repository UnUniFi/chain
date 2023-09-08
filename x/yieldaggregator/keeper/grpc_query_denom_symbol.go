package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) DenomSymbolMap(c context.Context, req *types.QueryDenomSymbolMapRequest) (*types.QueryDenomSymbolMapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryDenomSymbolMapResponse{
		Mappings: k.GetAllDenomSymbolMap(ctx),
	}, nil
}

// rpc RegisterDenomSymbolMap(MsgRegisterDenomSymbolMap) returns (MsgRegisterDenomSymbolMapResponse);
