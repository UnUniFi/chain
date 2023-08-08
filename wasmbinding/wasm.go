package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	interchainquerykeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/keeper"
	recordskeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/keeper"
)

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	icqKeeper *interchainquerykeeper.Keeper,
	recordsKeeper *recordskeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin()

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, icqKeeper, recordsKeeper),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
