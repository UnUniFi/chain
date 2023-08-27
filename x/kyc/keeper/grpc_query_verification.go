package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/kyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerificationAll(c context.Context, req *types.QueryAllVerificationRequest) (*types.QueryAllVerificationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var verifications []types.Verification
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	verificationStore := prefix.NewStore(store, types.KeyPrefix(types.VerificationKeyPrefix))

	pageRes, err := query.Paginate(verificationStore, req.Pagination, func(key []byte, value []byte) error {
		var verification types.Verification
		if err := k.cdc.Unmarshal(value, &verification); err != nil {
			return err
		}

		verifications = append(verifications, verification)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVerificationResponse{Verification: verifications, Pagination: pageRes}, nil
}

func (k Keeper) Verification(c context.Context, req *types.QueryGetVerificationRequest) (*types.QueryGetVerificationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVerification(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVerificationResponse{Verification: val}, nil
}
