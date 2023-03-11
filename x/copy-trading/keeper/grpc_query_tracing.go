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

func (k Keeper) TracingAll(c context.Context, req *types.QueryAllTracingRequest) (*types.QueryAllTracingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tracings []types.Tracing
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tracingStore := prefix.NewStore(store, types.KeyPrefix(types.TracingKeyPrefix))

	pageRes, err := query.Paginate(tracingStore, req.Pagination, func(key []byte, value []byte) error {
		var tracing types.Tracing
		if err := k.cdc.Unmarshal(value, &tracing); err != nil {
			return err
		}

		tracings = append(tracings, tracing)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTracingResponse{Tracing: tracings, Pagination: pageRes}, nil
}

func (k Keeper) Tracing(c context.Context, req *types.QueryGetTracingRequest) (*types.QueryGetTracingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTracing(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTracingResponse{Tracing: val}, nil
}

func (k Keeper) ExemplaryTraderTracing(c context.Context, req *types.QueryGetExemplaryTraderTracingRequest) (*types.QueryGetExemplaryTraderTracingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tracings []types.Tracing
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	// TODO: AddressTracingKeyPrefix
	tracingStore := prefix.NewStore(store, types.KeyPrefix(types.TracingKeyPrefix))

	pageRes, err := query.Paginate(tracingStore, req.Pagination, func(key []byte, value []byte) error {
		var tracing types.Tracing
		if err := k.cdc.Unmarshal(value, &tracing); err != nil {
			return err
		}

		tracings = append(tracings, tracing)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetExemplaryTraderTracingResponse{Tracing: tracings, Pagination: pageRes}, nil
}
