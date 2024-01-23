package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) EstimateMintPtYtPair(c context.Context, req *types.QueryEstimateMintPtYtPairRequest) (*types.QueryEstimateMintPtYtPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	tranche, found := k.GetTranchePool(ctx, req.Id)
	if !found {
		return nil, types.ErrTrancheNotFound
	}
	if req.Denom != tranche.Denom {
		return nil, types.ErrInvalidDepositDenom
	}
	depositAmount, ok := sdk.NewIntFromString(req.Amount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}
	mintAmount, err := k.CalculateMintPtAmount(ctx, tranche, sdk.NewCoin(req.Denom, depositAmount))
	if err != nil {
		return nil, err
	}

	ptDenom := types.PtDenom(tranche)
	ytDenom := types.YtDenom(tranche)
	return &types.QueryEstimateMintPtYtPairResponse{
		PtAmount: sdk.NewCoin(ptDenom, mintAmount),
		YtAmount: sdk.NewCoin(ytDenom, depositAmount),
	}, nil
}
