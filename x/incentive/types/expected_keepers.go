package types

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cdptypes "github.com/lcnem/jpyx/x/cdp/types"
)

// AccountKeeper expected interface for the account keeper (noalias)
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

// SupplyKeeper defines the expected supply keeper for module accounts  (noalias)
type BankKeeper interface {
	InitGenesis(sdk.Context, *banktypes.GenesisState)
	ExportGenesis(sdk.Context) *banktypes.GenesisState

	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	GetTotalSupply(ctx sdk.Context) sdk.Coins
	IterateTotalSupply(ctx sdk.Context, cb func(sdk.Coin) bool)

	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
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
}

// StakingKeeper expected interface for the auction keeper (noalias)
type StakingKeeper interface {
}

// CdpKeeper defines the expected interface for the pricefeed  (noalias)
type CdpKeeper interface {
	GetInterestFactor(ctx sdk.Context, collateralType string) (sdk.Dec, bool)
	GetTotalPrincipal(ctx sdk.Context, collateralType string, principalDenom string) (total sdk.Int)
	GetCdpByOwnerAndCollateralType(ctx sdk.Context, owner sdk.AccAddress, collateralType string) (cdptypes.CDP, bool)
	GetCollateral(ctx sdk.Context, collateralType string) (cdptypes.CollateralParam, bool)
}
