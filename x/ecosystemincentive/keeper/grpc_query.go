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

func (k Keeper) EcosystemRewards(c context.Context, req *types.QueryEcosystemRewardsRequest) (*types.QueryEcosystemRewardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	ctx := sdk.UnwrapSDKContext(c)

	accAddr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return &types.QueryEcosystemRewardsResponse{Rewards: sdk.Coins{}}, err
	}

	allRewards, exist := k.GetRewardRecord(ctx, accAddr)
	if !exist {
		return &types.QueryEcosystemRewardsResponse{Rewards: allRewards.Rewards}, nil
	}

	if req.Denom == "" {
		return &types.QueryEcosystemRewardsResponse{Rewards: allRewards.Rewards}, nil
	}

	exist, reward := allRewards.Rewards.Find(req.Denom)
	if !exist {
		return &types.QueryEcosystemRewardsResponse{Rewards: sdk.Coins{sdk.NewInt64Coin(req.Denom, 0)}}, nil
	}

	return &types.QueryEcosystemRewardsResponse{Rewards: sdk.Coins{reward}}, nil
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
