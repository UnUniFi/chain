package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) DenomInfos(c context.Context, req *types.QueryDenomInfosRequest) (*types.QueryDenomInfosResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryDenomInfosResponse{
		Info: k.GetAllDenomInfo(ctx),
	}, nil
}
