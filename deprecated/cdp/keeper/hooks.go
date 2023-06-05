package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/deprecated/cdp/types"
)

// Implements StakingHooks interface
var _ types.CdpHooks = Keeper{}

// AfterCdpCreated - call hook if registered
func (k Keeper) AfterCdpCreated(ctx sdk.Context, cdp types.Cdp) {
	if k.hooks != nil {
		k.hooks.AfterCdpCreated(ctx, cdp)
	}
}

// BeforeCdpModified - call hook if registered
func (k Keeper) BeforeCdpModified(ctx sdk.Context, cdp types.Cdp) {
	if k.hooks != nil {
		k.hooks.BeforeCdpModified(ctx, cdp)
	}
}
