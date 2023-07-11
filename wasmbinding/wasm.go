package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	interchainquerykeeper "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/interchainquery/keeper"
)

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	icqKeeper *interchainquerykeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin()

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, icqKeeper),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
