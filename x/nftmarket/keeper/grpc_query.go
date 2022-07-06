package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/nftmarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryParamsResponse{
		Params: k.GetParamSet(ctx),
	}, nil
}

func (k Keeper) NftListing(c context.Context, req *types.QueryNftListingRequest) (*types.QueryNftListingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	listing, err := k.GetNftListingByIdBytes(ctx, types.NftBytes(req.ClassId, req.NftId))
	if err != nil {
		return nil, err
	}

	return &types.QueryNftListingResponse{
		Listing: listing,
	}, nil
}

func (k Keeper) ListedNfts(c context.Context, req *types.QueryListedNftsRequest) (*types.QueryListedNftsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	listings := k.GetAllNftListings(ctx)
	return &types.QueryListedNftsResponse{
		Listings: listings,
	}, nil
}

func (k Keeper) Loans(c context.Context, req *types.QueryLoansRequest) (*types.QueryLoansResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryLoansResponse{
		Loans: k.GetAllDebts(ctx),
	}, nil
}

func (k Keeper) CDPsList(c context.Context, req *types.QueryCDPsListRequest) (*types.QueryCDPsListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.QueryCDPsListResponse{}, nil
}

func (k Keeper) NftBids(c context.Context, req *types.QueryNftBidsRequest) (*types.QueryNftBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	bids := k.GetBidsByNft(ctx, types.NftBytes(req.ClassId, req.NftId))
	return &types.QueryNftBidsResponse{
		Bids: bids,
	}, nil
}

func (k Keeper) BidderBids(c context.Context, req *types.QueryBidderBidsRequest) (*types.QueryBidderBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	bids := k.GetBidsByBidder(ctx, sdk.AccAddress(req.Bidder))
	return &types.QueryBidderBidsResponse{
		Bids: bids,
	}, nil
}

func (k Keeper) Rewards(c context.Context, req *types.QueryRewardsRequest) (*types.QueryRewardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.QueryRewardsResponse{}, nil
}
