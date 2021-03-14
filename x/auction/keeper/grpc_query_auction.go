package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/lcnem/jpyx/x/auction/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AuctionAll(c context.Context, req *types.QueryAllAuctionRequest) (*types.QueryAllAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var auctions []*types.Auction
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	auctionStore := prefix.NewStore(store, types.KeyPrefix(types.AuctionKey))

	pageRes, err := query.Paginate(auctionStore, req.Pagination, func(key []byte, value []byte) error {
		var auction types.Auction
		if err := k.cdc.UnmarshalBinaryBare(value, &auction); err != nil {
			return err
		}

		auctions = append(auctions, &auction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAuctionResponse{Auction: auctions, Pagination: pageRes}, nil
}

func (k Keeper) Auction(c context.Context, req *types.QueryGetAuctionRequest) (*types.QueryGetAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var auction types.Auction
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.AuctionKey + req.Id)), &auction)

	return &types.QueryGetAuctionResponse{Auction: &auction}, nil
}
