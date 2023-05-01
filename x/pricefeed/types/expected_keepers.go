package types

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
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
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)

	// View
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetSupply(ctx sdk.Context, denom string) sdk.Coin

	// Send
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}
