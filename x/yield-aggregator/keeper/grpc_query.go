package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Vaults(c context.Context, req *types.QueryVaultsRequest) (*types.QueryVaultsResponse, error) {
	panic("implement me")
}

func (k Keeper) Vault(c context.Context, req *types.QueryVaultRequest) (*types.QueryVaultResponse, error) {
	panic("implement me")
}

func (k Keeper) Strategies(c context.Context, req *types.QueryStrategiesRequest) (*types.QueryStrategiesResponse, error) {
	panic("implement me")
}
