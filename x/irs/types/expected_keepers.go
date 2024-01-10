package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin

	SendCoins(ctx sdk.Context, senderAddr sdk.AccAddress, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error

	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}

type RecordsKeeper interface {
	GetUserRedemptionRecordBySenderAndHostZone(ctx sdk.Context, sender sdk.AccAddress, zoneId string) math.Int
	VaultTransfer(ctx sdk.Context, vaultId uint64, contractAddr sdk.AccAddress, msg *ibctypes.MsgTransfer) error
	GetVaultPendingDeposit(ctx sdk.Context, vaultId uint64) math.Int
	IncreaseVaultPendingDeposit(ctx sdk.Context, vaultId uint64, amount math.Int)
	DecreaseVaultPendingDeposit(ctx sdk.Context, vaultId uint64, amount math.Int)
}
