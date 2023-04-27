package keeper

import (
	"context"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k Keeper) Balance(ctx context.Context, req *banktypes.QueryBalanceRequest) (*banktypes.QueryBalanceResponse, error) {
	return k.bankKeeper.Balance(ctx, req)
}
func (k Keeper) AllBalances(ctx context.Context, req *banktypes.QueryAllBalancesRequest) (*banktypes.QueryAllBalancesResponse, error) {
	return k.bankKeeper.AllBalances(ctx, req)
}
func (k Keeper) TotalSupply(ctx context.Context, req *banktypes.QueryTotalSupplyRequest) (*banktypes.QueryTotalSupplyResponse, error) {
	return k.bankKeeper.TotalSupply(ctx, req)
}

func (k Keeper) SupplyOf(ctx context.Context, req *banktypes.QuerySupplyOfRequest) (*banktypes.QuerySupplyOfResponse, error) {
	return k.bankKeeper.SupplyOf(ctx, req)
}
func (k Keeper) DenomMetadata(ctx context.Context, req *banktypes.QueryDenomMetadataRequest) (*banktypes.QueryDenomMetadataResponse, error) {
	return k.bankKeeper.DenomMetadata(ctx, req)
}
func (k Keeper) DenomsMetadata(ctx context.Context, req *banktypes.QueryDenomsMetadataRequest) (*banktypes.QueryDenomsMetadataResponse, error) {
	return k.bankKeeper.DenomsMetadata(ctx, req)
}
func (k Keeper) DenomOwners(c context.Context, req *banktypes.QueryDenomOwnersRequest) (*banktypes.QueryDenomOwnersResponse, error) {
	return k.bankKeeper.DenomOwners(c, req)
}
