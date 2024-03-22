package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
	yieldaggregatorkeeper "github.com/UnUniFi/chain/x/yieldaggregator/keeper"
)

type Keeper struct {
	cdc                   codec.BinaryCodec
	storeKey              storetypes.StoreKey
	bankKeeper            types.BankKeeper
	wasmKeeper            wasmtypes.ContractOpsKeeper
	wasmReader            wasmkeeper.Keeper
	recordsKeeper         types.RecordsKeeper
	YieldaggregatorKeeper yieldaggregatorkeeper.Keeper

	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bk types.BankKeeper,
	wasmKeeper wasmtypes.ContractOpsKeeper,
	wasmReader wasmkeeper.Keeper,
	recordsKeeper types.RecordsKeeper,
	yieldaggregatorKeeper yieldaggregatorkeeper.Keeper,
	authority string,
) Keeper {

	return Keeper{
		cdc:                   cdc,
		storeKey:              storeKey,
		bankKeeper:            bk,
		wasmKeeper:            wasmKeeper,
		wasmReader:            wasmReader,
		recordsKeeper:         recordsKeeper,
		YieldaggregatorKeeper: yieldaggregatorKeeper,
		authority:             authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
