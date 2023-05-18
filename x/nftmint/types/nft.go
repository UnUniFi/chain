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

func ValidateClassID(classID string) error {
	if len(classID) == 0 {
		return nfttypes.ErrEmptyClassID
	}
	return nil
}

func ValidateNFTID(nftID string) error {
	if len(nftID) == 0 {
		return nfttypes.ErrEmptyNFTID
	}
	return nil
}
