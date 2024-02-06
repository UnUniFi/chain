package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateRequiredDepositSwapToYt(c context.Context, req *types.QueryEstimateRequiredDepositSwapToYtRequest) (*types.QueryEstimateRequiredDepositSwapToYtResponse, error) {
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
	requiredDeposit, err := k.CalculateRequiredDepositSwapToYt(ctx, tranche, desiredAmount)
	if err != nil {
		return nil, err
	}
	return &types.QueryEstimateRequiredDepositSwapToYtResponse{
		RequiredDepositAmount: requiredDeposit,
	}, nil
}
