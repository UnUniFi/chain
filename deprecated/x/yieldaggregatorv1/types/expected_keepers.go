package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	yieldfarmtypes "github.com/UnUniFi/chain/deprecated/x/yieldfarm/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

type YieldFarmKeeper interface {
	SetFarmerInfo(ctx sdk.Context, obj yieldfarmtypes.FarmerInfo)
	DeleteFarmerInfo(ctx sdk.Context, addr sdk.AccAddress)
	GetFarmerInfo(ctx sdk.Context, addr sdk.AccAddress) yieldfarmtypes.FarmerInfo
	GetAllFarmerInfos(ctx sdk.Context) []yieldfarmtypes.FarmerInfo
	Deposit(ctx sdk.Context, user sdk.AccAddress, coins sdk.Coins) error
	Withdraw(ctx sdk.Context, user sdk.AccAddress, coins sdk.Coins) error
	ClaimRewards(ctx sdk.Context, user sdk.AccAddress) sdk.Coins
}
