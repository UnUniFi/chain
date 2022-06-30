package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/nftmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryParamsResponse{
		Params: k.GetParamSet(ctx),
	}, nil
}

func (k Keeper) ClassAttributes(c context.Context, req *types.QueryClassAttributesRequest) (*types.QueryClassAttributesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	ctx := sdk.UnwrapSDKContext(c)
	classAttributes, found := k.GetClassAttributes(ctx, req.ClassId)
	if !found {
		return nil, sdkerrors.Wrap(nfttypes.ErrClassNotExists, "class which has that class id doesn't exist")
	}
	return &types.QueryClassAttributesResponse{
		ClassAttributes: &classAttributes,
	}, nil
}

func (k Keeper) NFTMinter(context.Context, *types.QueryNFTMinterRequest) (*types.QueryNFTMinterResponse, error) {
	return &types.QueryNFTMinterResponse{}, nil
}

func (k Keeper) ClassIdByName(context.Context, *types.QueryClassIdsByNameRequest) (*types.QueryClassIdsByNameResponse, error) {
	return &types.QueryClassIdsByNameResponse{}, nil
}

func (k Keeper) ClassIdsByOwner(c context.Context, req *types.QueryClassIdsByOwnerRequest) (*types.QueryClassIdsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	ctx := sdk.UnwrapSDKContext(c)
	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}
	owningClassIdList, found := k.GetOwningClassIdList(ctx, owner)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOwningClassIdListNotExists, "owner doesn't have any class")
	}

	return &types.QueryClassIdsByOwnerResponse{
		OwningClassIdList: &owningClassIdList,
	}, nil
}
