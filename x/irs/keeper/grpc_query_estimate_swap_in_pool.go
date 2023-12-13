package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateSwapInPool(c context.Context, req *types.QueryEstimateSwapInPoolRequest) (*types.QueryEstimateSwapInPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.PoolId)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	amount, err := k.SimulateSwapPoolTokens(ctx, tranche, req.Amount)
	if err != nil {
		return nil, err
	}

	return &types.QueryEstimateSwapInPoolResponse{
		Amount: amount,
	}, nil
}
