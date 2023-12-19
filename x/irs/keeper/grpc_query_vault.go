package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) VaultByContract(c context.Context, req *types.QueryVaultByContractRequest) (*types.QueryVaultByContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	vault, found := k.GetVault(ctx, req.StrategyContract)
	if !found {
		return nil, types.ErrVaultNotFound
	}

	return &types.QueryVaultByContractResponse{
		Vault: vault,
	}, nil
}
