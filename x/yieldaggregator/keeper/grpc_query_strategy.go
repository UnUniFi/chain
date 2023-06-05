package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) StrategyAll(c context.Context, req *types.QueryAllStrategyRequest) (*types.QueryAllStrategyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var strategies []types.Strategy
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	strategyStore := prefix.NewStore(store, types.KeyPrefixStrategy(req.Denom))

	pageRes, err := query.Paginate(strategyStore, req.Pagination, func(key []byte, value []byte) error {
		var strategy types.Strategy
		if err := k.cdc.Unmarshal(value, &strategy); err != nil {
			return err
		}

		strategies = append(strategies, strategy)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStrategyResponse{Strategies: strategies, Pagination: pageRes}, nil
}

func (k Keeper) Strategy(c context.Context, req *types.QueryGetStrategyRequest) (*types.QueryGetStrategyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	strategy, found := k.GetStrategy(ctx, req.Denom, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetStrategyResponse{Strategy: strategy}, nil
}
