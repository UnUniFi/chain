package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/UnUniFi/chain/x/ununifidist/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc           codec.Codec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		paramSpace    paramtypes.Subspace
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(cdc codec.Codec, storeKey, memKey sdk.StoreKey, paramSpace paramtypes.Subspace, accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramSpace:    paramSpace,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetPreviousBlockTime get the blocktime for the previous block
func (k Keeper) GetPreviousBlockTime(ctx sdk.Context) (blockTime time.Time, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.KeyPrefix(types.PreviousBlockTimeKey))
	if b == nil {
		return time.Time{}, false
	}
	blockTime.UnmarshalBinary(b)

	return blockTime, true
}

// SetPreviousBlockTime set the time of the previous block
func (k Keeper) SetPreviousBlockTime(ctx sdk.Context, blockTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	b, _ := blockTime.MarshalBinary()
	store.Set(types.KeyPrefix(types.PreviousBlockTimeKey), b)
}

func (k Keeper) GetGovDenom(ctx sdk.Context) (govDenom string, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.KeyPrefix(types.GovDenomKey))
	govDenom = string(b)

	return govDenom, true
}

func (k Keeper) SetGovDenom(ctx sdk.Context, govDenom string) {
	store := ctx.KVStore(k.storeKey)
	b := []byte(govDenom)
	store.Set(types.KeyPrefix(types.GovDenomKey), b)
}
