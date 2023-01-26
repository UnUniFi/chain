package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) LiquidityProviderTokenRealAPY(c context.Context, req *types.QueryLiquidityProviderTokenRealAPYRequest) (*types.QueryLiquidityProviderTokenRealAPYResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	rate := k.GetLPNominalYieldRate(ctx, req.BeforeHeight, req.AfterHeight)
	annualized := k.AnnualizeYieldRate(ctx, rate, req.BeforeHeight, req.AfterHeight)

	return &types.QueryLiquidityProviderTokenRealAPYResponse{Apy: &annualized}, nil
}

func (k Keeper) LiquidityProviderTokenNominalAPY(c context.Context, req *types.QueryLiquidityProviderTokenNominalAPYRequest) (*types.QueryLiquidityProviderTokenNominalAPYResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	rate := k.GetLPNominalYieldRate(ctx, req.BeforeHeight, req.AfterHeight)
	annualized := k.AnnualizeYieldRate(ctx, rate, req.BeforeHeight, req.AfterHeight)

	return &types.QueryLiquidityProviderTokenNominalAPYResponse{Apy: &annualized}, nil
}

func (k Keeper) AllPositions(c context.Context, req *types.QueryAllPositionsRequest) (*types.QueryAllPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	positions := k.GetAllPositions(ctx)

	return &types.QueryAllPositionsResponse{
		Positions: positions,
	}, nil
}

func (k Keeper) AddressOpeningPositions(c context.Context, req *types.QueryAddressOpeningPositionsRequest) (*types.QueryAddressOpeningPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// If address is empty, error should be emitted. It is better to prepare another query for all positions
	if req.Address.AccAddress().String() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}
	positions := k.GetAddressOpeningPositions(ctx, req.Address.AccAddress())

	return &types.QueryAddressOpeningPositionsResponse{
		Positions: positions,
	}, nil
}

func (k Keeper) AddressClosedPositions(c context.Context, req *types.QueryAddressClosedPositionsRequest) (*types.QueryAddressClosedPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// If address is empty, error should be emitted. It is better to prepare another query for all positions
	if req.Address.AccAddress().String() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}
	positions := k.GetAddressClosedPositions(ctx, req.Address.AccAddress())

	return &types.QueryAddressClosedPositionsResponse{
		Positions: positions,
	}, nil
}
