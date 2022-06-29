package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/nftmint/types"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(context.Context, *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	return &types.QueryParamsResponse{}, nil
}

func (k Keeper) ClassOwner(context.Context, *types.QueryClassOwnerRequest) (*types.QueryClassOwnerResponse, error) {
	return &types.QueryClassOwnerResponse{}, nil
}

func (k Keeper) NFTMinter(context.Context, *types.QueryNFTMinterRequest) (*types.QueryNFTMinterResponse, error) {
	return &types.QueryNFTMinterResponse{}, nil
}

func (k Keeper) ClassIdByName(context.Context, *types.QueryClassIdsByNameRequest) (*types.QueryClassIdsByNameResponse, error) {
	return &types.QueryClassIdsByNameResponse{}, nil
}

func (k Keeper) ClassBaseTokenUri(context.Context, *types.QueryClassBaseTokenUriRequest) (*types.QueryClassBaseTokenUriResponse, error) {
	return &types.QueryClassBaseTokenUriResponse{}, nil
}

func (k Keeper) ClassTokenSupplyCap(context.Context, *types.QueryClassTokenSupplyCapRequest) (*types.QueryClassTokenSupplyCapResponse, error) {
	return &types.QueryClassTokenSupplyCapResponse{}, nil
}

func (k Keeper) ClassIdsByOwner(context.Context, *types.QueryClassIdsByOwnerRequest) (*types.QueryClassIdsByOwnerResponse, error) {
	return &types.QueryClassIdsByOwnerResponse{}, nil
}
