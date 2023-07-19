package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nft/types"
)

// AccountKeeper expected interface for the account keeper (noalias)
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
}

type BankKeeper interface {
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

type NftKeeper interface {
	GetClass(ctx sdk.Context, classID string) (nfttypes.Class, bool)
	SaveClass(ctx sdk.Context, class nfttypes.Class) error

	Mint(ctx sdk.Context, token nfttypes.NFT, receiver sdk.AccAddress) error
	Burn(ctx sdk.Context, classID string, nftID string) error

	Update(ctx sdk.Context, token nfttypes.NFT) error
	Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error

	GetNFT(ctx sdk.Context, classID, nftID string) (nfttypes.NFT, bool)
	GetNFTsOfClassByOwner(ctx sdk.Context, classID string, owner sdk.AccAddress) (nfts []nfttypes.NFT)
	GetNFTsOfClass(ctx sdk.Context, classID string) (nfts []nfttypes.NFT)
	GetOwner(ctx sdk.Context, classID string, nftID string) sdk.AccAddress
	GetBalance(ctx sdk.Context, classID string, owner sdk.AccAddress) uint64
	GetTotalSupply(ctx sdk.Context, classID string) uint64
	HasNFT(ctx sdk.Context, classID, id string) bool
	GetNftData(ctx sdk.Context, classId string, id string) (types.NftData, bool)
	SetNftData(ctx sdk.Context, classId string, id string, data types.NftData) error
}

type NftbackedloanHooks interface {
	AfterNftListed(ctx sdk.Context, nftIdentifier NftIdentifier, txMemo string)
	AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier NftIdentifier, fee sdk.Coin)
	AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier NftIdentifier)
}
