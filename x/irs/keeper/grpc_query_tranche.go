package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) Tranches(c context.Context, req *types.QueryTranchesRequest) (*types.QueryTranchesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryTranchesResponse{
		Tranches: k.GetTranchesByStrategy(ctx, req.StrategyContract),
	}, nil
}

func (k Keeper) Tranche(c context.Context, req *types.QueryTrancheRequest) (*types.QueryTrancheResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}

	return &types.QueryTrancheResponse{
		Tranche: tranche,
	}, nil
}
