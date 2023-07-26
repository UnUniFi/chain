package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/pricefeed/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RawPriceAll(c context.Context, req *types.QueryAllRawPriceRequest) (*types.QueryAllRawPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	prices := k.GetRawPrices(ctx, req.MarketId)

	return &types.QueryAllRawPriceResponse{Prices: prices}, nil
}
