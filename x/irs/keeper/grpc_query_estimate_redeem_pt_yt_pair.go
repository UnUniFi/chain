package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateRedeemPtYtPair(c context.Context, req *types.QueryEstimateRedeemPtYtPairRequest) (*types.QueryEstimateRedeemPtYtPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.PoolId)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	redeemAmount, ok := sdk.NewIntFromString(req.YtAmount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	ptAmount, err := k.CalculateRedeemRequiredPtAmount(ctx, tranche, redeemAmount)
	if err != nil {
		return nil, err
	}

	ptDenom := types.PtDenom(tranche)
	return &types.QueryEstimateRedeemPtYtPairResponse{
		PtAmount: sdk.NewCoin(ptDenom, ptAmount),
	}, nil
}
