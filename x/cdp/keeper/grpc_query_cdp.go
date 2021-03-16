package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/lcnem/jpyx/x/cdp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CdpAll(c context.Context, req *types.QueryAllCdpRequest) (*types.QueryAllCdpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cdps []*types.CDP
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	cdpStore := prefix.NewStore(store, types.KeyPrefix(types.CdpKey))

	pageRes, err := query.Paginate(cdpStore, req.Pagination, func(key []byte, value []byte) error {
		var cdp types.CDP
		if err := k.cdc.UnmarshalBinaryBare(value, &cdp); err != nil {
			return err
		}

		cdps = append(cdps, &cdp)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCdpResponse{Cdp: cdps, Pagination: pageRes}, nil
}

func (k Keeper) Cdp(c context.Context, req *types.QueryGetCdpRequest) (*types.QueryGetCdpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cdp types.CDP
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CdpKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.CdpKey+req.Id)), &cdp)

	return &types.QueryGetCdpResponse{Cdp: &cdp}, nil
}
