package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftfactory/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) ClassAuthorityMetadata(ctx context.Context, req *types.QueryClassAuthorityMetadataRequest) (*types.QueryClassAuthorityMetadataResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	class_id := fmt.Sprintf("factory/%s/%s", req.GetCreator(), req.GetSubclass())

	authorityMetadata, err := k.GetAuthorityMetadata(sdkCtx, class_id)
	if err != nil {
		return nil, err
	}

	return &types.QueryClassAuthorityMetadataResponse{AuthorityMetadata: authorityMetadata}, nil
}

func (k Keeper) ClassesFromCreator(ctx context.Context, req *types.QueryClassesFromCreatorRequest) (*types.QueryClassesFromCreatorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	classes := k.getDenomsFromCreator(sdkCtx, req.GetCreator())
	return &types.QueryClassesFromCreatorResponse{Classes: classes}, nil
}
