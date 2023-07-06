package keeper

import (
	"context"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
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
		return k.ListedNftsByOwner(c, acc)
	} else {
		listings := k.GetAllNftListings(ctx)
		res, err := k.GetNftListingDetails(ctx, listings)
		if err != nil {
			panic(err)
		}
		return &types.QueryListedNftsResponse{
			Listings: res,
		}, nil
	}

}

func (k Keeper) GetNftListingDetails(ctx sdk.Context, listings []types.NftListing) ([]types.NftListingDetail, error) {
	res := []types.NftListingDetail{}
	for _, v := range listings {
		nftInfo, found := k.nftKeeper.GetNFT(ctx, v.NftId.ClassId, v.NftId.NftId)
		if !found {
			return []types.NftListingDetail{}, types.ErrNotExistsNft
		}
		detail := types.NftListingDetail{
			Listing: v,
			NftInfo: types.NftInfo{
				Id:      nftInfo.GetId(),
				Uri:     nftInfo.GetUri(),
				UriHash: nftInfo.GetUriHash(),
			},
		}
		res = append(res, detail)
	}
	return res, nil
}

func (k Keeper) ListedNftsByOwner(c context.Context, address sdk.AccAddress) (*types.QueryListedNftsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	listings := k.GetListingsByOwner(ctx, address)
	res, err := k.GetNftListingDetails(ctx, listings)
	if err != nil {
		panic(err)
	}
	return &types.QueryListedNftsResponse{
		Listings: res,
	}, nil
}

// todo add pagination
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

	var nfts []types.NftInfo
	var pnfts []*types.NftInfo
	for i, v := range class.NftIds {
		if limit <= i {
			break
		}
		nftInfo, hasNft := k.nftKeeper.GetNFT(ctx, class.ClassId, v)
		if !hasNft {
			return nil, status.Error(codes.NotFound, "not found nft")
		}
		nfts = append(nfts, types.NftInfo{Id: nftInfo.Id, Uri: nftInfo.Uri, UriHash: nftInfo.UriHash})
	}

	for i := range nfts {
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

	// ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryLoansResponse{
		// todo impl
		Loans: []types.Loan{},
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
	bids := k.GetBidsByNft(ctx, nftId.IdBytes())
	// Change the order of bids to  descending order
	sort.SliceStable(bids, func(i, j int) bool {
		if bids[i].BidAmount.Amount.LT(bids[j].BidAmount.Amount) {
			return false
		}
		if bids[i].BidAmount.Amount.GT(bids[j].BidAmount.Amount) {
			return true
		}
		if bids[i].CreatedAt.After(bids[j].CreatedAt) {
			return true
		}
		return false
	})
	max, err := types.MaxBorrowAmount(bids, ctx.BlockTime())
	if err != nil {
		return nil, err
	}
	deposits := sdk.NewCoin(max.Denom, sdk.NewInt(0))
	loan := types.Loan{NftId: nftId, Loan: sdk.NewCoin(max.Denom, sdk.NewInt(0))}

	for _, v := range bids {
		deposits = deposits.Add(v.DepositAmount)
		loan.Loan = loan.Loan.Add(v.Borrow.Amount)
	}

	return &types.QueryLoanResponse{
		Loan:           loan,
		BorrowingLimit: max,
		TotalDeposit:   deposits,
	}, nil
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
		if v.Id.Bidder == req.Bidder {
			bidderBid = v
		}
	}
	if (bidderBid.Equal(types.NftBid{})) {
		return nil, status.Error(codes.InvalidArgument, "does not match bidder")
	}

	allPaid := listing.State >= types.ListingState_LIQUIDATION && bidderBid.BidAmount.Amount.Equal(bidderBid.DepositAmount.Amount)
	return &types.QueryPaymentStatusResponse{
		PaymentStatus: types.PaymentStatus{
			NftId:            listing.NftId,
			State:            listing.State,
			Bidder:           bidderBid.Id.Bidder,
			Amount:           bidderBid.BidAmount,
			AutomaticPayment: bidderBid.AutomaticPayment,
			PaidAmount:       bidderBid.DepositAmount.Amount,
			CreatedAt:        bidderBid.CreatedAt,
			AllPaid:          allPaid,
		},
	}, nil
}

func (k Keeper) Liquidation(c context.Context, req *types.QueryLiquidationRequest) (*types.QueryLiquidationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	listing, err := k.GetNftListingByIdBytes(ctx, types.NftBytes(req.ClassId, req.NftId))
	if err != nil {
		return nil, err
	}

	bids := types.NftBids(k.GetBidsByNft(ctx, listing.NftId.IdBytes()))
	bids = bids.SortLowerBiddingPeriod()
	liquidations := &types.Liquidations{}
	// after 1 hour
	afterAnHour := ctx.BlockTime().Add(time.Hour)

	for _, bid := range bids {
		if bid.Borrow.Amount.Amount.Equal(sdk.ZeroInt()) {
			continue
		}

		liq := types.Liquidation{
			Amount: sdk.NewCoin(listing.BidDenom, sdk.ZeroInt()),
		}
		liq.Amount = bid.LiquidationAmount(afterAnHour)
		liq.LiquidationDate = bid.ExpiryAt
		if liquidations.Liquidation == nil {
			liquidations.Liquidation = &liq
		} else {
			liquidations.NextLiquidation = append(liquidations.NextLiquidation, liq)
		}
	}

	return &types.QueryLiquidationResponse{
		Liquidations: liquidations,
	}, nil
}
