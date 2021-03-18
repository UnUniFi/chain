package types

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/exported"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type AccountKeeper interface {
	// Return a new account with the next account number and the specified address. Does not save the new account to the store.
	NewAccountWithAddress(sdk.Context, sdk.AccAddress) authtypes.AccountI

	// Return a new account with the next account number. Does not save the new account to the store.
	NewAccount(sdk.Context, authtypes.AccountI) authtypes.AccountI

	// Retrieve an account from the store.
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI

	// Set an account in the store.
	SetAccount(sdk.Context, authtypes.AccountI)

	// Remove an account from the store.
	RemoveAccount(sdk.Context, authtypes.AccountI)

	// Iterate over all accounts, calling the provided function. Stop iteraiton when it returns false.
	IterateAccounts(sdk.Context, func(authtypes.AccountI) bool)

	// Fetch the public key of an account at a specified address
	GetPubKey(sdk.Context, sdk.AccAddress) (cryptotypes.PubKey, error)

	// Fetch the sequence of an account at a specified address.
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)

	// Fetch the next account number, and increment the internal counter.
	GetNextAccountNumber(sdk.Context) uint64

	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
}

type BankKeeper interface {
	// View
	ValidateBalance(ctx sdk.Context, addr sdk.AccAddress) error
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool

	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetAccountsBalances(ctx sdk.Context) []banktypes.Balance
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	IterateAccountBalances(ctx sdk.Context, addr sdk.AccAddress, cb func(coin sdk.Coin) (stop bool))
	IterateAllBalances(ctx sdk.Context, cb func(address sdk.AccAddress, coin sdk.Coin) (stop bool))

	// Send
	InputOutputCoins(ctx sdk.Context, inputs []banktypes.Input, outputs []banktypes.Output) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error

	GetParams(ctx sdk.Context) banktypes.Params
	SetParams(ctx sdk.Context, params banktypes.Params)

	SendEnabledCoin(ctx sdk.Context, coin sdk.Coin) bool
	SendEnabledCoins(ctx sdk.Context, coins ...sdk.Coin) error

	BlockedAddr(addr sdk.AccAddress) bool

	//
	InitGenesis(sdk.Context, *banktypes.GenesisState)
	ExportGenesis(sdk.Context) *banktypes.GenesisState

	GetSupply(ctx sdk.Context) exported.SupplyI
	SetSupply(ctx sdk.Context, supply exported.SupplyI)

	GetDenomMetaData(ctx sdk.Context, denom string) banktypes.Metadata
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
	IterateAllDenomMetaData(ctx sdk.Context, cb func(banktypes.Metadata) bool)

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	DelegateCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error

	DelegateCoins(ctx sdk.Context, delegatorAddr, moduleAccAddr sdk.AccAddress, amt sdk.Coins) error
	UndelegateCoins(ctx sdk.Context, moduleAccAddr, delegatorAddr sdk.AccAddress, amt sdk.Coins) error
	MarshalSupply(supplyI exported.SupplyI) ([]byte, error)
	UnmarshalSupply(bz []byte) (exported.SupplyI, error)
}
