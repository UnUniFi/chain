package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) IntermediaryAccountInfo(c context.Context, req *types.QueryIntermediaryAccountInfoRequest) (*types.QueryIntermediaryAccountInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	info := k.GetIntermediaryAccountInfo(ctx)
	return &types.QueryIntermediaryAccountInfoResponse{Addrs: info.Addrs}, nil
}
