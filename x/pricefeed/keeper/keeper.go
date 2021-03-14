package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/lcnem/jpyx/x/pricefeed/types"
)

type (
	Keeper struct {
		cdc        codec.Marshaler
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramSpace paramtypes.Subspace
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey, paramSpace paramtypes.Subspace) *Keeper {
	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramSpace: paramSpace,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
