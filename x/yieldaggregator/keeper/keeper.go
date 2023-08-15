package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stakeibckeeper "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/keeper"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

type Keeper struct {
	cdc            codec.BinaryCodec
	storeKey       storetypes.StoreKey
	bankKeeper     types.BankKeeper
	wasmKeeper     wasmtypes.ContractOpsKeeper
	wasmReader     wasmkeeper.Keeper
	stakeibcKeeper stakeibckeeper.Keeper
	recordsKeeper  types.RecordsKeeper

	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bk types.BankKeeper,
	wasmKeeper wasmtypes.ContractOpsKeeper,
	wasmReader wasmkeeper.Keeper,
	stakeibcKeeper stakeibckeeper.Keeper,
	recordsKeeper types.RecordsKeeper,
	authority string,
) Keeper {

	return Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		bankKeeper:     bk,
		wasmKeeper:     wasmKeeper,
		wasmReader:     wasmReader,
		stakeibcKeeper: stakeibcKeeper,
		recordsKeeper:  recordsKeeper,
		authority:      authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
