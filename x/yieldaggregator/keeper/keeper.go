package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	stakeibckeeper "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/stakeibc/keeper"
	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

type Keeper struct {
	cdc            codec.BinaryCodec
	storeKey       storetypes.StoreKey
	paramstore     paramtypes.Subspace
	bankKeeper     types.BankKeeper
	wasmKeeper     wasmtypes.ContractOpsKeeper
	wasmReader     wasmkeeper.Keeper
	stakeibcKeeper stakeibckeeper.Keeper
	recordsKeeper  types.RecordsKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	bk types.BankKeeper,
	wasmKeeper wasmtypes.ContractOpsKeeper,
	wasmReader wasmkeeper.Keeper,
	stakeibcKeeper stakeibckeeper.Keeper,
	recordsKeeper types.RecordsKeeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		paramstore:     paramSpace,
		bankKeeper:     bk,
		wasmKeeper:     wasmKeeper,
		wasmReader:     wasmReader,
		stakeibcKeeper: stakeibcKeeper,
		recordsKeeper:  recordsKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
