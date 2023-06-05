package keeper

import (
	"context"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/deprecated/nftmarketv1/types"
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
	if req.Owner != "" {
		acc, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid request. address wrong")
		}
		return k.ListedNftsByOwner(ctx, acc)
	} else {
		listings := k.GetAllNftListings(ctx)
		return &types.QueryListedNftsResponse{
			Listings: listings,
		}, nil
	}

}

func (k Keeper) ListedNftsByOwner(c context.Context, address sdk.AccAddress) (*types.QueryListedNftsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	listings := k.GetListingsByOwner(ctx, address)
	return &types.QueryListedNftsResponse{
		Listings: listings,
	}, nil
}

// todo add pagenation
func (k Keeper) ListedClasses(c context.Context, req *types.QueryListedClassesRequest) (*types.QueryListedClassesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	classes, err := k.GetListedClasses(ctx)
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
	class, err := k.GetListedClassByClassIdBytes(ctx, types.ClassIdKey(classId))
	if err != nil {
		return nil, err
	}
	classInfo, hasClass := k.nftKeeper.GetClass(ctx, class.ClassId)
	if !hasClass {
		return nil, status.Error(codes.NotFound, "not found class")
	}

	var nfts []types.ListedNft
	var pnfts []*types.ListedNft
	for i, v := range class.NftIds {
		if limit <= i {
			break
		}
		nftInfo, hasNft := k.nftKeeper.GetNFT(ctx, class.ClassId, v)
		if !hasNft {
			return nil, status.Error(codes.NotFound, "not found nft")
		}
		nfts = append(nfts, types.ListedNft{Id: nftInfo.Id, Uri: nftInfo.Uri, UriHash: nftInfo.UriHash})
	}

	for i, _ := range nfts {
		pnfts = append(pnfts, &nfts[i])
	}

	return &types.QueryListedClassResponse{
		ClassId:     class.ClassId,
		Name:        classInfo.Name,
		Description: classInfo.Description,
		Symbol:      classInfo.Symbol,
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
	return &types.QueryLoansResponse{
		Loans: k.GetAllDebts(ctx),
	}, nil
}

func (k Keeper) Loan(c context.Context, req *types.QueryLoanRequest) (*types.QueryLoanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	nftId := types.NftIdentifier{
		ClassId: req.ClassId,
		NftId:   req.NftId,
	}
	ctx := sdk.UnwrapSDKContext(c)
	nft, err := k.GetNftListingByIdBytes(ctx, nftId.IdBytes())
	if err != nil {
		return &types.QueryLoanResponse{
			Loan:           types.Loan{},
			BorrowingLimit: sdk.ZeroInt(),
		}, nil
	}
	bids := k.GetBidsByNft(ctx, nftId.IdBytes())
	// Change the order of bids to  descending order
	sort.SliceStable(bids, func(i, j int) bool {
		if bids[i].Amount.Amount.LT(bids[j].Amount.Amount) {
			return false
		}
		if bids[i].Amount.Amount.GT(bids[j].Amount.Amount) {
			return true
		}
		if bids[i].BidTime.After(bids[j].BidTime) {
			return true
		}
		return false
	})
	max := sdk.ZeroInt()
	for i, v := range bids {
		if i+1 > int(nft.BidActiveRank) {
			break
		}
		max = max.Add(v.PaidAmount)
	}

	return &types.QueryLoanResponse{
		Loan:           k.GetDebtByNft(ctx, nftId.IdBytes()),
		BorrowingLimit: max,
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
	bids := k.GetBidsByBidder(ctx, sdk.AccAddress(sdk.MustAccAddressFromBech32(req.Bidder)))
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

func (k Keeper) PaymentStatus(c context.Context, req *types.QueryPaymentStatusRequest) (*types.QueryPaymentStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	nft := types.NftIdentifier{
		ClassId: req.ClassId,
		NftId:   req.NftId,
	}
	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
	if err != nil {
		return &types.QueryPaymentStatusResponse{}, err
	}
	bids := k.GetBidsByNft(ctx, nft.IdBytes())
	if len(bids) == 0 {
		return nil, status.Error(codes.InvalidArgument, "not existing bidder")
	}

	var bidderBid types.NftBid
	for _, v := range bids {
		if v.Bidder == req.Bidder {
			bidderBid = v
		}
	}
	if (bidderBid == types.NftBid{}) {
		return nil, status.Error(codes.InvalidArgument, "does not match bidder")
	}

	allPaid := listing.State >= types.ListingState_END_LISTING && bidderBid.Amount.Amount.Equal(bidderBid.PaidAmount)
	return &types.QueryPaymentStatusResponse{
		PaymentStatus: types.PaymentStatus{
			NftId:            listing.NftId,
			State:            listing.State,
			Bidder:           bidderBid.Bidder,
			Amount:           bidderBid.Amount,
			AutomaticPayment: bidderBid.AutomaticPayment,
			PaidAmount:       bidderBid.PaidAmount,
			BidTime:          bidderBid.BidTime,
			AllPaid:          allPaid,
		},
	}, nil
}
