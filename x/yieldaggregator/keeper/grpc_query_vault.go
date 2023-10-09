package keeper

import (
	"context"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
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

	var vaultContainers []types.VaultContainer
	for _, vault := range vaults {
		vaultContainers = append(vaultContainers, types.VaultContainer{
			Vault:                vault,
			TotalBondedAmount:    k.VaultAmountInStrategies(ctx, vault),
			TotalUnbondingAmount: k.VaultUnbondingAmountInStrategies(ctx, vault),
			WithdrawReserve:      k.VaultWithdrawalAmount(ctx, vault),
		})
	}

	return &types.QueryAllVaultResponse{Vaults: vaultContainers, Pagination: pageRes}, nil
}

func (k Keeper) Vault(c context.Context, req *types.QueryGetVaultRequest) (*types.QueryGetVaultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	vault, found := k.GetVault(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	strategies := []types.Strategy{}
	for _, strategyWeight := range vault.StrategyWeights {
		strategy, found := k.GetStrategy(ctx, strategyWeight.Denom, strategyWeight.StrategyId)
		if !found {
			continue
		}
		strategies = append(strategies, strategy)
	}

	return &types.QueryGetVaultResponse{
		Vault:                vault,
		Strategies:           strategies,
		TotalBondedAmount:    k.VaultAmountInStrategies(ctx, vault),
		TotalUnbondingAmount: k.VaultUnbondingAmountInStrategies(ctx, vault),
		WithdrawReserve:      k.VaultWithdrawalAmount(ctx, vault),
	}, nil
}

func (k Keeper) VaultAllByShareHolder(c context.Context, req *types.QueryAllVaultByShareHolderRequest) (*types.QueryAllVaultByShareHolderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	addr, err := sdk.AccAddressFromBech32(req.ShareHolder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	balances := k.bankKeeper.GetAllBalances(ctx, addr)

	var vaultContainers []types.VaultContainer
	for _, coin := range balances {
		denomParts := strings.Split(coin.Denom, "/")

		if len(denomParts) != 3 {
			continue
		}
		if denomParts[0] != types.ModuleName || denomParts[1] != "vaults" {
			continue
		}
		vaultId, err := strconv.Atoi(denomParts[2])
		if err != nil {
			continue
		}
		vault, found := k.GetVault(ctx, uint64(vaultId))
		if !found {
			continue
		}

		vaultContainers = append(vaultContainers, types.VaultContainer{
			Vault:                vault,
			TotalBondedAmount:    k.VaultAmountInStrategies(ctx, vault),
			TotalUnbondingAmount: k.VaultUnbondingAmountInStrategies(ctx, vault),
			WithdrawReserve:      k.VaultWithdrawalAmount(ctx, vault),
		})
	}

	return &types.QueryAllVaultByShareHolderResponse{Vaults: vaultContainers}, nil
}
