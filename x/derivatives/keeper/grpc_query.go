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

func (k Keeper) AllOpeningPositions(c context.Context, req *types.QueryAllOpeningPositionsRequest) (*types.QueryAllOpeningPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	positions := k.GetAllOpenedPositions(ctx)

	return &types.QueryAllOpeningPositionsResponse{
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
	positions := k.GetAddressOpenedPositions(ctx, req.Address.AccAddress())

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

func (k Keeper) ClosedPosition(c context.Context, req *types.QueryClosedPositionRequest) (*types.QueryClosedPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	position := k.GetClosedPosition(ctx, req.Id)

	return &types.QueryClosedPositionResponse{
		Position: position,
	}, nil
}

func (k Keeper) PerpetualFutures(c context.Context, req *types.QueryPerpetualFuturesRequest) (*types.QueryPerpetualFuturesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	metricsQuoteTicker := ""
	volume24Hours := sdk.NewDec(0)
	fees24Hours := sdk.NewDec(0)
	longPositions := sdk.NewDec(0)
	shortPositions := sdk.NewDec(0)

	return &types.QueryPerpetualFuturesResponse{
		MetricsQuoteTicker: metricsQuoteTicker,
		Volume_24Hours:     &volume24Hours,
		Fees_24Hours:       &fees24Hours,
		LongPositions:      &longPositions,
		ShortPositions:     &shortPositions,
	}, nil
}

func (k Keeper) PerpetualFuturesPair(c context.Context, req *types.QueryPerpetualFuturesPairRequest) (*types.QueryPerpetualFuturesPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	price := sdk.NewDec(0)
	metricsQuoteTicker := ""
	volume24Hours := sdk.NewDec(0)
	fees24Hours := sdk.NewDec(0)
	longPositions := sdk.NewDec(0)
	shortPositions := sdk.NewDec(0)

	return &types.QueryPerpetualFuturesPairResponse{
		Price:              &price,
		MetricsQuoteTicker: metricsQuoteTicker,
		Volume_24Hours:     &volume24Hours,
		Fees_24Hours:       &fees24Hours,
		LongPositions:      &longPositions,
		ShortPositions:     &shortPositions,
	}, nil
}

func (k Keeper) PerpetualOptions(c context.Context, req *types.QueryPerpetualOptionsRequest) (*types.QueryPerpetualOptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic

	return &types.QueryPerpetualOptionsResponse{}, nil
}

func (k Keeper) PerpetualOptionsPair(c context.Context, req *types.QueryPerpetualOptionsPairRequest) (*types.QueryPerpetualOptionsPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic

	return &types.QueryPerpetualOptionsPairResponse{}, nil
}

func (k Keeper) Pool(c context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	metricsQuoteTicker := ""
	poolMarketCap := k.GetPoolMarketCap(ctx)
	volume24Hours := sdk.NewDec(0)
	fees24Hours := sdk.NewDec(0)

	return &types.QueryPoolResponse{
		MetricsQuoteTicker: metricsQuoteTicker,
		PoolMarketCap:      &poolMarketCap,
		Volume_24Hours:     &volume24Hours,
		Fees_24Hours:       &fees24Hours,
	}, nil
}
