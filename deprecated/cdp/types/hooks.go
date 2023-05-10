package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// MultiCdpHooks combine multiple cdp hooks, all hook functions are run in array sequence
type MultiCdpHooks []CdpHooks

// NewMultiCdpHooks returns a new MultiCdpHooks
func NewMultiCdpHooks(hooks ...CdpHooks) MultiCdpHooks {
	return hooks
}

// BeforeCdpModified runs before a cdp is modified
func (h MultiCdpHooks) BeforeCdpModified(ctx sdk.Context, cdp Cdp) {
	for i := range h {
		h[i].BeforeCdpModified(ctx, cdp)
	}
}

// AfterCdpCreated runs before a cdp is created
func (h MultiCdpHooks) AfterCdpCreated(ctx sdk.Context, cdp Cdp) {
	for i := range h {
		h[i].AfterCdpCreated(ctx, cdp)
	}
}
