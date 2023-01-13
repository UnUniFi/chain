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
	balances := k.GetVaultAmount(ctx, key)
	afterBalances := k.subAmount(balances, amt)
	if !afterBalances.IsValid() {
		// To Do use crisis
		fmt.Println(afterBalances)
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}

func (k Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, key []byte, senderModule, recipientModule string, amt sdk.Coins) error {
	balances := k.GetVaultAmount(ctx, key)
	afterBalances := k.subAmount(balances, amt)
	if !afterBalances.IsValid() {
		// To Do use crisis
		fmt.Println(afterBalances)
	}
	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt)
}

func (k Keeper) SendCoinsFromAccountToModule(ctx sdk.Context, key []byte, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	newAmount := k.CheckVaultAndAddAmount(ctx, key, amt)
	k.SetVaultAmount(ctx, key, newAmount)
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}

// function for controlling sdk.Coins
func (k Keeper) CheckVaultAndAddAmount(ctx sdk.Context, key []byte, amt sdk.Coins) sdk.Coins {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(key)
	if bz != nil {
		var p sdk.Coins
		err := json.Unmarshal(bz, &p)
		if err != nil {
			panic(err)
		}
		k.addAmount(p, amt)
	}
	return amt
}

func (k Keeper) addAmount(balance sdk.Coins, amt sdk.Coins) sdk.Coins {
	for _, coin := range amt {
		balance = balance.Add(coin)
	}
	return balance
}

func (k Keeper) subAmount(balance sdk.Coins, amt sdk.Coins) sdk.Coins {
	for _, coin := range amt {
		balance = balance.Sub(coin)
	}
	return balance
}

func (k Keeper) GetVaultAmount(ctx sdk.Context, key []byte) sdk.Coins {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(key)
	var amt sdk.Coins
	err := json.Unmarshal(bz, &amt)
	if err != nil {
		panic(err)
	}
	return amt
}

func (k Keeper) SetVaultAmount(ctx sdk.Context, key []byte, amt sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	bz, err := amt.MarshalJSON()
	if err != nil {
		panic(err)
	}
	store.Set(key, bz)
}
