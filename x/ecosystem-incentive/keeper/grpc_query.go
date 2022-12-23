package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

// AllRewards returns the RewardStore defined by subject address
func (k Keeper) AllRewards(c context.Context, req *types.QueryAllRewardsRequest) (*types.QueryAllRewardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	ctx := sdk.UnwrapSDKContext(c)

	accAddr, err := sdk.AccAddressFromBech32(req.SubjectAddr)
	if err != nil {
		return nil, err
	}

	allRewards, exists := k.GetRewardStore(ctx, accAddr)
	if !exists {
		return nil, types.ErrAddressNotHaveReward
	}

	return &types.QueryAllRewardsResponse{Rewards: allRewards}, nil
}

func (k Keeper) IncentiveUnit(c context.Context, req *types.QueryIncentiveUnitRequest) (*types.QueryIncentiveUnitResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	ctx := sdk.UnwrapSDKContext(c)

	incentiveUnit, exists := k.GetIncentiveUnit(ctx, req.IncentiveUnitId)
	if !exists {
		return nil, types.ErrNotRegisteredIncentiveUnitId
	}

	return &types.QueryIncentiveUnitResponse{IncentiveUnit: &incentiveUnit}, nil
}

func (k Keeper) Reward(c context.Context, req *types.QueryRewardRequest) (*types.QueryRewardResponse, error) {

	return &types.QueryRewardResponse{}, nil
}

func (k Keeper) RecordedIncentiveUnitId(c context.Context, req *types.QueryRecordedIncentiveUnitIdRequest) (*types.QueryRecordedIncentiveUnitIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	ctx := sdk.UnwrapSDKContext(c)

	nftIdentifier := nftmarkettypes.NftIdentifier{
		ClassId: req.ClassId,
		NftId:   req.NftId,
	}

	incentiveUnitid, exists := k.GetIncentiveUnitIdByNftId(ctx, nftIdentifier)
	if !exists {
		return nil, sdkerrors.Wrapf(types.ErrIncentiveUnitIdByNftIdDoesntExist, "class id: %s\nnft id: %s", req.ClassId, req.NftId)
	}
	return &types.QueryRecordedIncentiveUnitIdResponse{IncentiveUnitId: incentiveUnitid}, nil
}

func (k Keeper) IncentiveUnitIdsByAddr(c context.Context, req *types.QueryIncentiveUnitIdsByAddrRequest) (*types.QueryIncentiveUnitIdsByAddrResponse, error) {
	return &types.QueryIncentiveUnitIdsByAddrResponse{}, nil
}
