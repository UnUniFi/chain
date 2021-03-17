package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/jsmndist/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RewardAll(c context.Context, req *types.QueryAllRewardRequest) (*types.QueryAllRewardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rewards string
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.RewardKey)), &rewards)

	return &types.QueryAllRewardResponse{Params: &params}, nil
}
