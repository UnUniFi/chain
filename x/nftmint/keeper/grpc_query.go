package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/nftmint/types"

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

	// if err := nfttypes.ValidateClassID(req.ClassId); err != nil {
	// 	return nil, err
	// }

	ctx := sdk.UnwrapSDKContext(c)
	classAttributes, found := k.GetClassAttributes(ctx, req.ClassId)
	if !found {
		return nil, sdkerrors.Wrap(nfttypes.ErrClassNotExists, "class which has that class id doesn't exist")
	}
	return &types.QueryClassAttributesResponse{
		ClassAttributes: &classAttributes,
	}, nil
}

func (k Keeper) NFTMinter(c context.Context, req *types.QueryNFTMinterRequest) (*types.QueryNFTMinterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	// if err := nfttypes.ValidateClassID(req.ClassId); err != nil {
	// 	return nil, err
	// }
	// if err := nfttypes.ValidateNFTID(req.NftId); err != nil {
	// 	return nil, err
	// }

	ctx := sdk.UnwrapSDKContext(c)
	nftMinter, exists := k.GetNFTMinter(ctx, req.ClassId, req.NftId)
	if !exists {
		return nil, sdkerrors.Wrapf(types.ErrNftAttributesNotExists, "NftAttributes with this %s class and %s nft id doesn't exist", req.ClassId, req.NftId)
	}

	return &types.QueryNFTMinterResponse{
		Minter: nftMinter.String(),
	}, nil
}

func (k Keeper) ClassIdsByName(c context.Context, req *types.QueryClassIdsByNameRequest) (*types.QueryClassIdsByNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParamSet(ctx)
	if err := types.ValidateClassName(params.MinClassNameLen, params.MaxClassNameLen, req.ClassName); err != nil {
		return nil, err
	}

	classNameIdList, exists := k.GetClassNameIdList(ctx, req.ClassName)
	if !exists {
		return nil, sdkerrors.Wrap(types.ErrClassNameIdListNotExists, req.ClassName)
	}
	return &types.QueryClassIdsByNameResponse{
		ClassNameIdList: &classNameIdList,
	}, nil
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
