package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/records/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) UserRedemptionRecordAll(c context.Context, req *types.QueryAllUserRedemptionRecordRequest) (*types.QueryAllUserRedemptionRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userRedemptionRecords []types.UserRedemptionRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userRedemptionRecordStore := prefix.NewStore(store, types.KeyPrefix(types.UserRedemptionRecordKey))

	pageRes, err := query.Paginate(userRedemptionRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var userRedemptionRecord types.UserRedemptionRecord
		if err := k.Cdc.Unmarshal(value, &userRedemptionRecord); err != nil {
			return err
		}

		userRedemptionRecords = append(userRedemptionRecords, userRedemptionRecord)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserRedemptionRecordResponse{UserRedemptionRecord: userRedemptionRecords, Pagination: pageRes}, nil
}

func (k Keeper) UserRedemptionRecord(c context.Context, req *types.QueryGetUserRedemptionRecordRequest) (*types.QueryGetUserRedemptionRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	userRedemptionRecord, found := k.GetUserRedemptionRecord(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetUserRedemptionRecordResponse{UserRedemptionRecord: userRedemptionRecord}, nil
}

func (k Keeper) UserRedemptionRecordForUser(c context.Context, req *types.QueryAllUserRedemptionRecordForUserRequest) (*types.QueryAllUserRedemptionRecordForUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// validate the address
	_, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, req.Address)
	}

	var userRedemptionRecords []types.UserRedemptionRecord

	ctx := sdk.UnwrapSDKContext(c)

	// limit loop to 50 records for performance
	var loopback uint64
	loopback = req.Limit
	if loopback > 50 {
		loopback = 50
	}
	var i uint64
	for i = 0; i < loopback; i++ {
		if i > req.Day {
			// we have reached the end of the records
			break
		}
		currentDay := req.Day - i
		// query the user redemption record for the current day
		userRedemptionRecord, found := k.GetUserRedemptionRecord(ctx, types.UserRedemptionRecordKeyFormatter(req.ChainId, currentDay, req.Address))
		if !found {
			continue
		}
		userRedemptionRecords = append(userRedemptionRecords, userRedemptionRecord)
	}

	return &types.QueryAllUserRedemptionRecordForUserResponse{UserRedemptionRecord: userRedemptionRecords}, nil
}

func (k Keeper) EpochUnbondingRecordAll(c context.Context, req *types.QueryAllEpochUnbondingRecordRequest) (*types.QueryAllEpochUnbondingRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var epochUnbondingRecords []types.EpochUnbondingRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	epochUnbondingRecordStore := prefix.NewStore(store, types.KeyPrefix(types.EpochUnbondingRecordKey))

	pageRes, err := query.Paginate(epochUnbondingRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var epochUnbondingRecord types.EpochUnbondingRecord
		if err := k.Cdc.Unmarshal(value, &epochUnbondingRecord); err != nil {
			return err
		}

		epochUnbondingRecords = append(epochUnbondingRecords, epochUnbondingRecord)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEpochUnbondingRecordResponse{EpochUnbondingRecord: epochUnbondingRecords, Pagination: pageRes}, nil
}

func (k Keeper) EpochUnbondingRecord(c context.Context, req *types.QueryGetEpochUnbondingRecordRequest) (*types.QueryGetEpochUnbondingRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	epochUnbondingRecord, found := k.GetEpochUnbondingRecord(ctx, req.EpochNumber)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEpochUnbondingRecordResponse{EpochUnbondingRecord: epochUnbondingRecord}, nil
}

func (k Keeper) DepositRecordAll(c context.Context, req *types.QueryAllDepositRecordRequest) (*types.QueryAllDepositRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var depositRecords []types.DepositRecord
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	depositRecordStore := prefix.NewStore(store, types.KeyPrefix(types.DepositRecordKey))

	pageRes, err := query.Paginate(depositRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var depositRecord types.DepositRecord
		if err := k.Cdc.Unmarshal(value, &depositRecord); err != nil {
			return err
		}

		depositRecords = append(depositRecords, depositRecord)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDepositRecordResponse{DepositRecord: depositRecords, Pagination: pageRes}, nil
}

func (k Keeper) DepositRecord(c context.Context, req *types.QueryGetDepositRecordRequest) (*types.QueryGetDepositRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	depositRecord, found := k.GetDepositRecord(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetDepositRecordResponse{DepositRecord: depositRecord}, nil
}
