package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	record, exists := k.GetRewardRecord(ctx, accAddr)
	if !exists {
		record.Address = accAddr.String()
		return &types.QueryAllRewardsResponse{RewardRecord: record}, nil
	}

	return &types.QueryAllRewardsResponse{RewardRecord: record}, nil
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

	allRewards, exists := k.GetRewardRecord(ctx, accAddr)
	if !exists {
		return &types.QueryRewardResponse{Reward: sdk.NewInt64Coin(req.Denom, 0)}, nil
	}

	exists, reward := allRewards.Rewards.Find(req.Denom)
	if !exists {
		return &types.QueryRewardResponse{Reward: sdk.NewInt64Coin(req.Denom, 0)}, nil
	}

	return &types.QueryRewardResponse{Reward: reward}, nil
}

func (k Keeper) RecipientAddressWithNftId(c context.Context, req *types.QueryRecipientAddressWithNftIdRequest) (*types.QueryRecipientAddressWithNftIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	ctx := sdk.UnwrapSDKContext(c)

	recipient, exists := k.GetRecipientByNftId(ctx, nftbackedloantypes.NftId{ClassId: req.ClassId, TokenId: req.TokenId})
	if !exists {
		return &types.QueryRecipientAddressWithNftIdResponse{Address: ""}, nil
	}

	return &types.QueryRecipientAddressWithNftIdResponse{Address: recipient}, nil
}
