package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/UnUniFi/chain/x/wrappedbank/types"
)

type Keeper struct {
	cdc        codec.Codec
	txCfg      client.TxConfig
	storeKey   storetypes.StoreKey
	memKey     storetypes.StoreKey
	paramSpace paramtypes.Subspace
	bankKeeper types.BankKeeper
}

func NewKeeper(cdc codec.Codec, txCfg client.TxConfig, storeKey,
	memKey storetypes.StoreKey, paramSpace paramtypes.Subspace,
	bankKeeper types.BankKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		txCfg:      txCfg,
		storeKey:   storeKey,
		memKey:     memKey,
		paramSpace: paramSpace,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
