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
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	redeemAmount, ok := sdk.NewIntFromString(req.DesiredUtAmount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	pt, yt, err := k.CalculateRedeemRequiredPtAndYtAmount(ctx, tranche, redeemAmount)
	if err != nil {
		return nil, err
	}

	return &types.QueryEstimateRedeemPtYtPairResponse{
		PtAmount: pt,
		YtAmount: yt,
	}, nil
}
