package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

// GetParamSet returns token params from the global param store
func (k Keeper) GetParamSet(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSpace.GetParamSet(ctx, &p)
	return p
}

// SetParamSet sets token params to the global param store
func (k Keeper) SetParamSet(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetNetworks returns the networks from params
func (k Keeper) GetNetworks(ctx sdk.Context) []types.Network {
	return k.GetParamSet(ctx).Networks
}

// GetOracles returns the oracles
func (k Keeper) GetOracles(ctx sdk.Context, networkId string) []sdk.AccAddress {
	for _, m := range k.GetNetworks(ctx) {
		if networkId == m.NetworkId {
			return ununifitypes.AccAddresses(m.Oracles)
		}
	}
	return []sdk.AccAddress{}
}
