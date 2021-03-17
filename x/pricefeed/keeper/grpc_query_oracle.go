package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	jpyx "github.com/lcnem/jpyx/types"
	"github.com/lcnem/jpyx/x/pricefeed/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OracleAll(c context.Context, req *types.QueryAllOracleRequest) (*types.QueryAllOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var oracles []sdk.AccAddress
	ctx := sdk.UnwrapSDKContext(c)

	oracles, err := k.GetOracles(ctx, req.MarketId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryAllOracleResponse{Oracles: jpyx.StringAccAddresses(oracles)}, nil
}
