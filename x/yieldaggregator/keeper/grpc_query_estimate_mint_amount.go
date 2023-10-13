package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
	// "github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimateMintAmount(c context.Context, req *types.QueryEstimateMintAmountRequest) (*types.QueryEstimateMintAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	vault, found := k.GetVault(ctx, req.Id)
	if !found {
		return nil, types.ErrInvalidVaultId
	}

	depositAmount, ok := sdk.NewIntFromString(req.DepositAmount)
	if !ok {
		return nil, types.ErrInvalidAmount
	}

	mintAmount := k.EstimateMintAmountInternal(ctx, vault.Id, sdk.NewCoin(vault.Denom, depositAmount))

	return &types.QueryEstimateMintAmountResponse{
		MintAmount: mintAmount,
	}, nil
}
