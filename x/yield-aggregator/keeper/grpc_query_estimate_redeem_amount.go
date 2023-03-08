package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	// "github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimateRedeemAmount(c context.Context, req *types.QueryEstimateRedeemAmountRequest) (*types.QueryEstimateRedeemAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	vault, found := k.GetVault(ctx, req.Id)
	if !found {
		// TODO
		return nil, nil
	}
	redeemAmount := k.EstimateRedeemAmountInternal(ctx, vault.Denom, vault.Id, req.BurnAmount)

	return &types.QueryEstimateRedeemAmountResponse{
		RedeemAmount: redeemAmount,
	}, nil
}
