package keeper

import (
	"context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
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

	var auctions []*codectypes.Any
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	auctionStore := prefix.NewStore(store, types.KeyPrefix(types.AuctionKey))

	pageRes, err := query.Paginate(auctionStore, req.Pagination, func(key []byte, value []byte) error {
		var auction codectypes.Any
		if err := k.cdc.UnmarshalBinaryBare(value, &auction); err != nil {
			return err
		}

		auctions = append(auctions, &auction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAuctionResponse{Auctions: auctions, Pagination: pageRes}, nil
}

func (k Keeper) Auction(c context.Context, req *types.QueryGetAuctionRequest) (*types.QueryGetAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	auction, err := k.GetAuction(ctx, req.Id)

	if err {
		return nil, status.Error(codes.NotFound, "not found")
	}
	auctionAny, _ := codectypes.NewAnyWithValue(auction)

	return &types.QueryGetAuctionResponse{Auction: auctionAny}, nil
}
