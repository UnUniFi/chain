package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/UnUniFi/chain/x/vault/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		txCfg         client.TxConfig
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramSpace    paramtypes.Subspace
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	txCfg client.TxConfig, storeKey,
	memKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		txCfg:         txCfg,
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

// func (k Keeper) SendCoins(ctx sdk.Context, key []byte, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
// 	return nil
// }
func (k Keeper) SendCoinsFromModuleToAccount(ctx sdk.Context, key []byte, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(key)
	var p sdk.Coins
	err := json.Unmarshal(bz, &p)
	if err != nil {
		panic(err)
	}
	// calc
	// k.bankKeeper
	return nil
}
func (k Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, key []byte, senderModule, recipientModule string, amt sdk.Coins) error {
	return nil
}
func (k Keeper) SendCoinsFromAccountToModule(ctx sdk.Context, key []byte, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := amt.MarshalJSON()
	if err != nil {
		panic(err)
	}
	store.Set(key, bz)
	// k.bankKeeper
	return nil
}
