package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/pricefeed/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Markets(c context.Context, req *types.QueryMarketsRequest) (*types.QueryMarketsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var post types.Market
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PostKey+req.Id)), &post)

	return &types.QueryMarketsResponse{Post: &post}, nil
}

func (k Keeper) Oracle(c context.Context, req *types.QueryOracleRequest) (*types.QueryOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var post types.Market
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PostKey+req.Id)), &post)

	return &types.QueryOracleResponse{Post: &post}, nil
}

func (k Keeper) RawPrice(c context.Context, req *types.QueryRawPriceRequest) (*types.QueryRawPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var post types.Market
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PostKey+req.Id)), &post)

	return &types.QueryRawPriceResponse{Post: &post}, nil
}
