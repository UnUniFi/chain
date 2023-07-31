package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"
	pftypes "github.com/UnUniFi/chain/x/pricefeed/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

// AccountKeeper expected interface for the account keeper (noalias)
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
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

type NftKeeper interface {
	SaveClass(ctx sdk.Context, class nfttypes.Class) error

	Mint(ctx sdk.Context, token nfttypes.NFT, receiver sdk.AccAddress) error
	Burn(ctx sdk.Context, classID string, nftID string) error
	Update(ctx sdk.Context, token nfttypes.NFT) error
	Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error

	GetClass(ctx sdk.Context, classID string) (nfttypes.Class, bool)
	GetClasses(ctx sdk.Context) (classes []*nfttypes.Class)
	HasClass(ctx sdk.Context, classId string) bool

	GetNFT(ctx sdk.Context, classID, nftID string) (nfttypes.NFT, bool)
	GetNFTsOfClassByOwner(ctx sdk.Context, classID string, owner sdk.AccAddress) (nfts []nfttypes.NFT)
	GetNFTsOfClass(ctx sdk.Context, classID string) (nfts []nfttypes.NFT)
	HasNFT(ctx sdk.Context, classID, nftID string) bool

	GetOwner(ctx sdk.Context, classID string, nftID string) sdk.AccAddress
	GetBalance(ctx sdk.Context, classID string, owner sdk.AccAddress) uint64
	GetTotalSupply(ctx sdk.Context, classID string) uint64
}

type NftfactoryKeeper interface {
	MintNFT(ctx sdk.Context, msg *nftfactorytypes.MsgMintNFT) error
	BurnNFT(ctx sdk.Context, msg *nftfactorytypes.MsgBurnNFT) error
}
