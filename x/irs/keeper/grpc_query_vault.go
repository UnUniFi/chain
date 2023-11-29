package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) Vault(c context.Context, req *types.QueryVaultRequest) (*types.QueryVaultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	vault, found := k.GetVault(ctx, req.StrategyContract)
	if !found {
		return nil, types.ErrVaultNotFound
	}

	return &types.QueryVaultResponse{
		Vault: vault,
	}, nil
}
