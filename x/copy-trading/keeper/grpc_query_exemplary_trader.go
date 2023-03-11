package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/copy-trading/types"
)

func (k Keeper) ExemplaryTraderAll(c context.Context, req *types.QueryAllExemplaryTraderRequest) (*types.QueryAllExemplaryTraderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var exemplaryTraders []types.ExemplaryTrader
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	exemplaryTraderStore := prefix.NewStore(store, types.KeyPrefix(types.ExemplaryTraderKeyPrefix))

	pageRes, err := query.Paginate(exemplaryTraderStore, req.Pagination, func(key []byte, value []byte) error {
		var exemplaryTrader types.ExemplaryTrader
		if err := k.cdc.Unmarshal(value, &exemplaryTrader); err != nil {
			return err
		}

		exemplaryTraders = append(exemplaryTraders, exemplaryTrader)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExemplaryTraderResponse{ExemplaryTrader: exemplaryTraders, Pagination: pageRes}, nil
}

func (k Keeper) ExemplaryTrader(c context.Context, req *types.QueryGetExemplaryTraderRequest) (*types.QueryGetExemplaryTraderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetExemplaryTrader(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetExemplaryTraderResponse{ExemplaryTrader: val}, nil
}
