package types

import (
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

func NewNFT(classID, nftID, nftUri string) nfttypes.NFT {
	return nfttypes.NFT{
		ClassId: classID,
		Id:      nftID,
		Uri:     nftUri,
	}
}
