package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/cdp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var params types.Params
	ctx := sdk.UnwrapSDKContext(c)

	params = k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: &params}, nil
}
