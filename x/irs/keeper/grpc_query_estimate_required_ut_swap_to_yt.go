package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateRequiredUtSwapToYt(c context.Context, req *types.QueryEstimateRequiredUtSwapToYtRequest) (*types.QueryEstimateRequiredUtSwapToYtResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	desiredAmount, ok := sdk.NewIntFromString(req.DesiredYtAmount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	requiredUt, err := k.CalculateRequiredDepositSwapToYt(ctx, tranche, desiredAmount)
	if err != nil {
		return nil, err
	}
	return &types.QueryEstimateRequiredUtSwapToYtResponse{
		RequiredUtAmount: requiredUt,
	}, nil
}
