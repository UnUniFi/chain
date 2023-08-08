package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/eventhook/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey

		wasmKeeper wasmtypes.ContractOpsKeeper
		authority  string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	wasmKeeper wasmtypes.ContractOpsKeeper,
	authority string,
) *Keeper {

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		wasmKeeper: wasmKeeper,
		authority:  authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
