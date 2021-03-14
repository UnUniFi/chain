package keeper

import "github.com/lcnem/jpyx/x/cdp/types"

// SetHooks sets the cdp keeper hooks
func (k *Keeper) SetHooks(hooks types.CDPHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set validator hooks twice")
	}
	k.hooks = hooks
	return k
}
