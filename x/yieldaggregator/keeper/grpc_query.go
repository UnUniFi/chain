package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) AssetManagementAccount(c context.Context, req *types.QueryAssetManagementAccountRequest) (*types.QueryAssetManagementAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.QueryAssetManagementAccountResponse{}, nil
}

func (k Keeper) AllAssetManagementAccounts(c context.Context, req *types.QueryAllAssetManagementAccountsRequest) (*types.QueryAllAssetManagementAccountsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.QueryAllAssetManagementAccountsResponse{}, nil
}

func (k Keeper) UserInfo(c context.Context, req *types.QueryUserInfoRequest) (*types.QueryUserInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.QueryUserInfoResponse{}, nil
}

func (k Keeper) AllUserInfos(c context.Context, req *types.QueryAllUserInfosRequest) (*types.QueryAllUserInfosResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	return &types.QueryAllUserInfosResponse{}, nil
}
