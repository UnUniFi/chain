package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/pricefeed/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PriceAll(c context.Context, req *types.QueryAllPriceRequest) (*types.QueryAllPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var prices types.CurrentPrices
	ctx := sdk.UnwrapSDKContext(c)

	prices = k.GetCurrentPrices(ctx)

	return &types.QueryAllPriceResponse{Prices: prices}, nil
}

func (k Keeper) Price(c context.Context, req *types.QueryGetPriceRequest) (*types.QueryGetPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var price types.CurrentPrice
	ctx := sdk.UnwrapSDKContext(c)

	price, err := k.GetCurrentPrice(ctx, req.MarketId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPriceResponse{Price: price}, nil
}
