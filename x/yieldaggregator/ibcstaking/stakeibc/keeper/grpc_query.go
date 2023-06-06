package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/stakeibc/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Validators(c context.Context, req *types.QueryGetValidatorsRequest) (*types.QueryGetValidatorsResponse, error) {
	if req == nil || req.ChainId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	hostZone, found := k.GetHostZone(ctx, req.ChainId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetValidatorsResponse{Validators: hostZone.Validators}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) ModuleAddress(goCtx context.Context, req *types.QueryModuleAddressRequest) (*types.QueryModuleAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	addr := k.accountKeeper.GetModuleAccount(ctx, req.Name).GetAddress().String()

	return &types.QueryModuleAddressResponse{Addr: addr}, nil
}

// InterchainAccountFromAddress implements the Query/InterchainAccountFromAddress gRPC method
func (k Keeper) InterchainAccountFromAddress(goCtx context.Context, req *types.QueryInterchainAccountFromAddressRequest) (*types.QueryInterchainAccountFromAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	portID, err := icatypes.NewControllerPortID(req.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}

	addr, found := k.ICAControllerKeeper.GetInterchainAccountAddress(ctx, req.ConnectionId, portID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "no account found for portID %s", portID)
	}

	return types.NewQueryInterchainAccountResponse(addr), nil
}

func (k Keeper) ICAAccount(c context.Context, req *types.QueryGetICAAccountRequest) (*types.QueryGetICAAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetICAAccount(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetICAAccountResponse{ICAAccount: val}, nil
}

func (k Keeper) HostZoneAll(c context.Context, req *types.QueryAllHostZoneRequest) (*types.QueryAllHostZoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hostZones []types.HostZone
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	hostZoneStore := prefix.NewStore(store, types.KeyPrefix(types.HostZoneKey))

	pageRes, err := query.Paginate(hostZoneStore, req.Pagination, func(key []byte, value []byte) error {
		var hostZone types.HostZone
		if err := k.cdc.Unmarshal(value, &hostZone); err != nil {
			return err
		}

		hostZones = append(hostZones, hostZone)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHostZoneResponse{HostZone: hostZones, Pagination: pageRes}, nil
}

func (k Keeper) HostZone(c context.Context, req *types.QueryGetHostZoneRequest) (*types.QueryGetHostZoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	hostZone, found := k.GetHostZone(ctx, req.ChainId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHostZoneResponse{HostZone: hostZone}, nil
}

func (k Keeper) EpochTrackerAll(c context.Context, req *types.QueryAllEpochTrackerRequest) (*types.QueryAllEpochTrackerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var epochTrackers []types.EpochTracker
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	epochTrackerStore := prefix.NewStore(store, types.KeyPrefix(types.EpochTrackerKeyPrefix))

	pageRes, err := query.Paginate(epochTrackerStore, req.Pagination, func(key []byte, value []byte) error {
		var epochTracker types.EpochTracker
		if err := k.cdc.Unmarshal(value, &epochTracker); err != nil {
			return err
		}

		epochTrackers = append(epochTrackers, epochTracker)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEpochTrackerResponse{EpochTracker: epochTrackers, Pagination: pageRes}, nil
}

func (k Keeper) EpochTracker(c context.Context, req *types.QueryGetEpochTrackerRequest) (*types.QueryGetEpochTrackerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEpochTracker(
		ctx,
		req.EpochIdentifier,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEpochTrackerResponse{EpochTracker: val}, nil
}
