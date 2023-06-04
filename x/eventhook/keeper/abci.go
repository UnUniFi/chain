package keeper

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func searchAttribute(key string, value string, attributes []abci.EventAttribute) bool {
	for _, attribute := range attributes {
		if string(attribute.Key) == key && string(attribute.Value) == value {
			return true
		}
	}

	return false
}

type Hook struct {
	EventAttributes []struct {
		Key   string
		Value string
	}
}

func inspectEventForHook(event sdk.Event, hook Hook) bool {
	for _, attribute := range hook.EventAttributes {
		// search attribute.Key in event.Attributes
		if !searchAttribute(attribute.Key, attribute.Value, event.Attributes) {
			return false
		}
	}
	return true
}

func (k Keeper) CallHook(ctx sdk.Context, event sdk.Event, hook Hook) {

}

func (k Keeper) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	for _, event := range ctx.EventManager().Events() {
		hooks := []Hook{} // event.Type -> []Hook
		for _, hook := range hooks {
			if inspectEventForHook(event, hook) {
				k.CallHook(ctx, event, hook)
			}
		}
	}
}
