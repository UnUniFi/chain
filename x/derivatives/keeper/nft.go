package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetNftOwner(ctx sdk.Context, positionId string) sdk.AccAddress {
	owner := k.nftKeeper.GetOwner(ctx, types.PositionNFTClassId, positionId)
	return owner
}

func (k Keeper) GetAddressNftPositions(ctx sdk.Context, address sdk.AccAddress) []types.Position {
	nfts := k.nftKeeper.GetNFTsOfClassByOwner(ctx, types.PositionNFTClassId, address)
	positions := []types.Position{}
	for _, nft := range nfts {
		position := k.GetPositionWithId(ctx, nft.Id)
		positions = append(positions, *position)
	}
	return positions
}

func (k Keeper) MintPositionNFT(ctx sdk.Context, position types.Position) error {
	moduleAddr := k.GetModuleAddress()
	msgMintNFT := nftfactorytypes.MsgMintNFT{
		Sender:    moduleAddr.String(),
		ClassId:   types.PositionNFTClassId,
		NftId:     position.Id,
		Recipient: position.Address,
	}
	err := k.nftfactoryKeeper.MintNFT(ctx, &msgMintNFT)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) BurnPositionNFT(ctx sdk.Context, position types.Position) error {
	moduleAddr := k.GetModuleAddress()
	msgBurnNFT := nftfactorytypes.MsgBurnNFT{
		Sender:  moduleAddr.String(),
		ClassId: types.PositionNFTClassId,
		NftId:   position.Id,
	}
	err := k.nftfactoryKeeper.BurnNFT(ctx, &msgBurnNFT)
	if err != nil {
		return err
	}
	return nil
}
