package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}

// PricefeedKeeper defines the expected interface for the pricefeed  (noalias)
type PricefeedKeeper interface {
	GetCurrentPrice(sdk.Context, string) (pftypes.CurrentPrice, error)
	GetParams(sdk.Context) pftypes.Params
	GetTicker(ctx sdk.Context, denom string) (string, error)
	GetMarketId(ctx sdk.Context, lhsTicker string, rhsTicker string) string
	GetMarketIdFromDenom(ctx sdk.Context, lhsDenom string, rhsDenom string) (string, error)
	// These are used for testing TODO replace mockApp with keeper in tests to remove these
	SetParams(sdk.Context, pftypes.Params)
}
