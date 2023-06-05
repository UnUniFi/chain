package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) MintNft(ctx sdk.Context, msg *types.MsgMintNft) error {
	classId := msg.ClassId
	nftId := msg.NftId
	_, exists := k.nftKeeper.GetNFT(ctx, classId, nftId)
	if exists {
		return nft.ErrNFTExists
	}

	_, hasId := k.nftKeeper.GetClass(ctx, classId)
	if !hasId {
		class := nft.Class{
			Id:          classId,
			Name:        classId,
			Symbol:      classId,
			Description: classId,
			Uri:         classId,
			UriHash:     classId,
		}
		k.nftKeeper.SaveClass(ctx, class)
	}

	expNFT := nft.NFT{
		ClassId: classId,
		Id:      nftId,
		Uri:     msg.NftUri,
		UriHash: msg.NftUriHash,
	}
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	err = k.nftKeeper.Mint(ctx, expNFT, sender)
	if err != nil {
		return err
	}

	return nil
}
