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

// todo add pagenation
func (k Keeper) ListedClasses(c context.Context, req *types.QueryListedClassesRequest) (*types.QueryListedClassesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	classes, err := k.GetClasses(ctx)
	if err != nil {
		return nil, err
	}

	var limit int
	if int(req.NftLimit) > 0 {
		limit = int(req.NftLimit)
	} else {
		limit = 1
	}
	var listedClasses []*types.QueryListedClassResponse
	for _, v := range classes {
		listedClass, err := k.GetListedClass(ctx, v.ClassId, limit)
		if err != nil {
			return nil, status.Error(codes.NotFound, "not found nft")
		}
		listedClasses = append(listedClasses, listedClass)
	}

	return &types.QueryListedClassesResponse{
		Classes: listedClasses,
	}, nil
}

func (k Keeper) ListedClass(c context.Context, req *types.QueryListedClassRequest) (*types.QueryListedClassResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var limit int
	if int(req.NftLimit) > 0 {
		limit = int(req.NftLimit)
	} else {
		limit = 1
	}
	listedClass, err := k.GetListedClass(ctx, req.ClassId, limit)
	if err != nil {
		return nil, err
	}

	return listedClass, nil
}

func (k Keeper) GetListedClass(ctx sdk.Context, classId string, limit int) (*types.QueryListedClassResponse, error) {
	class, err := k.GetClass(ctx, types.ClassIdKey(classId))
	if err != nil {
		return nil, err
	}
	classInfo, hasClass := k.nftKeeper.GetClass(ctx, class.ClassId)
	if !hasClass {
		return nil, status.Error(codes.NotFound, "not found class")
	}

	var nfts []types.Nft
	var pnfts []*types.Nft
	for i, v := range class.NftIds {
		if limit <= i {
			break
		}
		nftInfo, hasNft := k.nftKeeper.GetNFT(ctx, class.ClassId, v)
		if !hasNft {
			return nil, status.Error(codes.NotFound, "not found nft")
		}
		nfts = append(nfts, types.Nft{Id: nftInfo.Id, Uri: nftInfo.Uri, UriHash: nftInfo.UriHash})
	}

	for i, _ := range nfts {
		pnfts = append(pnfts, &nfts[i])
	}

	return &types.QueryListedClassResponse{
		Class:       class.ClassId,
		Name:        classInfo.Name,
		Description: classInfo.Description,
		Sybol:       classInfo.Symbol,
		Uri:         classInfo.Uri,
		Urihash:     classInfo.UriHash,
		NftCount:    uint64(len(class.NftIds)),
		Nfts:        pnfts,
	}, nil
}

func (k Keeper) Loans(c context.Context, req *types.QueryLoansRequest) (*types.QueryLoansResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.QueryLoansResponse{}, nil
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
