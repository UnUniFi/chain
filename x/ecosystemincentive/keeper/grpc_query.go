package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
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

	accAddr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	allRewards, exists := k.GetRewardStore(ctx, accAddr)
	if !exists {
		return nil, types.ErrAddressNotHaveReward
	}

	return &types.QueryAllRewardsResponse{Rewards: allRewards}, nil
}

func (k Keeper) Reward(c context.Context, req *types.QueryRewardRequest) (*types.QueryRewardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	ctx := sdk.UnwrapSDKContext(c)

	accAddr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return &types.QueryRewardResponse{Reward: sdk.Coin{}}, err
	}

	allRewards, exists := k.GetRewardStore(ctx, accAddr)
	if !exists {
		return &types.QueryRewardResponse{Reward: sdk.Coin{}}, types.ErrAddressNotHaveReward
	}

	exists, reward := allRewards.Rewards.Find(req.Denom)
	if !exists {
		return &types.QueryRewardResponse{Reward: sdk.Coin{}}, types.ErrDenomRewardNotExists
	}

	return &types.QueryRewardResponse{Reward: reward}, nil
}

func (k Keeper) RecipientContainer(c context.Context, req *types.QueryRecipientContainerRequest) (*types.QueryRecipientContainerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	ctx := sdk.UnwrapSDKContext(c)

	recipientContainer, exists := k.GetRecipientContainer(ctx, req.Id)
	if !exists {
		return nil, types.ErrNotRegisteredRecipientContainerId
	}

	return &types.QueryRecipientContainerResponse{RecipientContainer: &recipientContainer}, nil
}

func (k Keeper) RecordedRecipientContainerId(c context.Context, req *types.QueryRecordedRecipientContainerIdRequest) (*types.QueryRecordedRecipientContainerIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	ctx := sdk.UnwrapSDKContext(c)

	nftIdentifier := nftbackedloantypes.NftIdentifier{
		ClassId: req.ClassId,
		NftId:   req.NftId,
	}

	recipientContainerid, exists := k.GetRecipientContainerIdByNftId(ctx, nftIdentifier)
	if !exists {
		return nil, sdkerrors.Wrapf(types.ErrRecipientContainerIdByNftIdDoesntExist, "class id: %s\nnft id: %s", req.ClassId, req.NftId)
	}
	return &types.QueryRecordedRecipientContainerIdResponse{RecipientContainerId: recipientContainerid}, nil
}

func (k Keeper) BelongingRecipientContainerIdsByAddr(c context.Context, req *types.QueryBelongingRecipientContainerIdsByAddrRequest) (*types.QueryBelongingRecipientContainerIdsByAddrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	ctx := sdk.UnwrapSDKContext(c)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return &types.QueryBelongingRecipientContainerIdsByAddrResponse{}, err
	}
	recipientContainerIdsByAddr := k.GetRecipientContainerIdsByAddr(ctx, addr)

	if len(recipientContainerIdsByAddr.RecipientContainerIds) == 0 {
		return &types.QueryBelongingRecipientContainerIdsByAddrResponse{}, err
	}

	return &types.QueryBelongingRecipientContainerIdsByAddrResponse{RecipientContainerIds: recipientContainerIdsByAddr.RecipientContainerIds}, nil
}
