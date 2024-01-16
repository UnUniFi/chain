package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateSwapUtToYt(c context.Context, req *types.QueryEstimateSwapUtToYtRequest) (*types.QueryEstimateSwapUtToYtResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	tokenInAmount, ok := sdk.NewIntFromString(req.Amount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	yt, err := k.CalculateSwapUtToYt(ctx, tranche, sdk.NewCoin(req.Denom, tokenInAmount))
	if err != nil {
		return nil, err
	}
	return &types.QueryEstimateSwapUtToYtResponse{
		YtAmount: yt,
	}, nil
}
