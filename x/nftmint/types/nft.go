package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

func NewNFT(classID, nftID, nftUri string) nfttypes.NFT {
	return nfttypes.NFT{
		ClassId: classID,
		Id:      nftID,
		Uri:     nftUri,
	}
}

func NewNFTAttributes(classID, nftID string, minter sdk.AccAddress) NFTAttributes {
	return NFTAttributes{
		ClassId: classID,
		NftId:   nftID,
		Minter:  minter.Bytes(),
	}
}
