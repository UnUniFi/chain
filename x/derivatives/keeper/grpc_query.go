package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ClaimableLiquidityProviderRewards(c context.Context, req *types.QueryClaimableLiquidityProviderRewardsRequest) (*types.QueryClaimableLiquidityProviderRewardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	amount, err := nil, nil // TODO
	if err != nil {
		return nil, err
	}

	return &types.QueryClaimableLiquidityProviderRewardsResponse{}, nil
}

func (k Keeper) Positions(c context.Context, req *types.QueryPositionsRequest) (*types.QueryPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	positions := k.GetUserPositions(ctx, sdk.AccAddress(req.Address))

	return &types.QueryPositionsResponse{
		Positions: positions,
	}, nil
}
