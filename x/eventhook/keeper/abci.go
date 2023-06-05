package keeper

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/eventhook/types"
)

func searchAttribute(key string, value string, attributes []abci.EventAttribute) bool {
	for _, attribute := range attributes {
		if string(attribute.Key) == key && string(attribute.Value) == value {
			return true
		}
	}

	return false
}

func inspectEventForHook(event sdk.Event, hook types.Hook) bool {
	for _, attribute := range hook.EventAttributes {
		if !searchAttribute(attribute.Key, attribute.Value, event.Attributes) {
			return false
		}
	}
	return true
}

func (k Keeper) CallHook(ctx sdk.Context, event sdk.Event, hook types.Hook) {
	// TODO: call cosmwasm contract
}

func (k Keeper) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	eventHookMap := make(map[string][]types.Hook)

	for _, event := range ctx.EventManager().Events() {
		eventHookMap[string(event.Type)] = k.GetAllHook(ctx, event.Type)
	}

	for _, event := range ctx.EventManager().Events() {
		hooks := eventHookMap[string(event.Type)]
		for _, hook := range hooks {
			if inspectEventForHook(event, hook) {
				k.CallHook(ctx, event, hook)
			}
		}
	}

	return []abci.ValidatorUpdate{}
}
