package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateRedeemMaturedYt(c context.Context, req *types.QueryEstimateRedeemMaturedYtRequest) (*types.QueryEstimateRedeemMaturedYtResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	ytDenom := types.YtDenom(tranche)
	ytAmount, ok := sdk.NewIntFromString(req.YtAmount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	redeemAmount, err := k.CalculateRedeemYtAmount(ctx, tranche, sdk.NewCoin(ytDenom, ytAmount))
	if err != nil {
		return nil, err
	}
	return &types.QueryEstimateRedeemMaturedYtResponse{
		RedeemAmount: sdk.NewCoin(tranche.DepositDenom, redeemAmount),
	}, nil
}
