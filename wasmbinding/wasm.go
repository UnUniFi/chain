package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	yieldaggregatorkeeper "github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	interchainquerykeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/keeper"
	recordskeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/keeper"
)

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	icqKeeper *interchainquerykeeper.Keeper,
	recordsKeeper *recordskeeper.Keeper,
	yieldaggregatorKeeper *yieldaggregatorkeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin()

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, icqKeeper, recordsKeeper, yieldaggregatorKeeper),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
