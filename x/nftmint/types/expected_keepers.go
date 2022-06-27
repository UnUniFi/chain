package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

type AccountKeeper interface {
	GetSequence(ctx sdk.Context, addr sdk.AccAddress) (uint64, error)
}

type NftKeeper interface {
	NewClass(class nfttypes.Class)

	Mint(ctx sdk.Context, token nfttypes.NFT, receiver sdk.AccAddress) error
	Burn(ctx sdk.Context, classID string, nftID string) error
	Update(ctx sdk.Context, token nfttypes.NFT) error
	Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error

	GetClass(classid string) nfttypes.Class
	GetClasses() []nfttypes.Class

	GetNFT(ctx sdk.Context, classID, nftID string) (nfttypes.NFT, bool)
	GetNFTsOfClassByOwner(ctx sdk.Context, classID string, owner sdk.AccAddress) (nfts []nfttypes.NFT)
	GetNFTsOfClass(ctx sdk.Context, classID string) (nfts []nfttypes.NFT)

	GetOwner(ctx sdk.Context, classID string, nftID string) sdk.AccAddress
	GetBalance(ctx sdk.Context, classID string, owner sdk.AccAddress) uint64
	GetTotalSupply(ctx sdk.Context, classID string) uint64
}
