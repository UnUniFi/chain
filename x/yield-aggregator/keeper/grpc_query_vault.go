package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VaultAll(c context.Context, req *types.QueryAllVaultRequest) (*types.QueryAllVaultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vaults []types.Vault
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vaultStore := prefix.NewStore(store, types.KeyPrefix(types.VaultKey))

	pageRes, err := query.Paginate(vaultStore, req.Pagination, func(key []byte, value []byte) error {
		var vault types.Vault
		if err := k.cdc.Unmarshal(value, &vault); err != nil {
			return err
		}

		vaults = append(vaults, vault)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVaultResponse{Vaults: vaults, Pagination: pageRes}, nil
}

func (k Keeper) Vault(c context.Context, req *types.QueryGetVaultRequest) (*types.QueryGetVaultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	vault, found := k.GetVault(ctx, req.Denom)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVaultResponse{Vault: vault}, nil
}
