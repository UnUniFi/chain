package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

type Keeper struct {
	cdc           codec.Codec
	storeKey      storetypes.StoreKey
	memKey        storetypes.StoreKey
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	nftKeeper     types.NftKeeper
	authority     string
	hooks         types.NftbackedloanHooks
}

func NewKeeper(
	cdc codec.Codec, storeKey,
	memKey storetypes.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	nftKeeper types.NftKeeper,
	authority string,
) Keeper {

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		nftKeeper:     nftKeeper,
		authority:     authority,
		hooks:         nil,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetHooks sets the nftbackedloan keeper hooks
func (k *Keeper) SetHooks(hooks types.NftbackedloanHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set validator hooks twice")
	}
	k.hooks = hooks
	return k
}
