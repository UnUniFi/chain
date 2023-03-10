package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	stakeibckeeper "github.com/UnUniFi/chain/x/stakeibc/keeper"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	cdc            codec.BinaryCodec
	storeKey       storetypes.StoreKey
	paramstore     paramtypes.Subspace
	bankKeeper     types.BankKeeper
	wasmKeeper     wasmtypes.ContractOpsKeeper
	stakeibcKeeper stakeibckeeper.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	bk types.BankKeeper,
	wasmKeeper wasmtypes.ContractOpsKeeper,
	stakeibcKeeper stakeibckeeper.Keeper,
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
		stakeibcKeeper: stakeibcKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
