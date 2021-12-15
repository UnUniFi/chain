package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/pricefeed/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	return &types.QueryAllOracleResponse{Oracles: ununifitypes.StringAccAddresses(oracles)}, nil
}
