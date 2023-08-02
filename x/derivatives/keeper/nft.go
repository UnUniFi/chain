package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetPositionNFTSendDisabled(ctx sdk.Context, positionId string) (bool, error) {
	data, found := k.nftKeeper.GetNftData(ctx, types.PositionNFTClassId, positionId)
	if !found {
		return false, types.ErrPositionNFTNotFound
	}
	return data.SendDisabled, nil
}

func (k Keeper) GetPositionNFTOwner(ctx sdk.Context, positionId string) sdk.AccAddress {
	owner := k.nftKeeper.GetOwner(ctx, types.PositionNFTClassId, positionId)
	return owner
}

func (k Keeper) GetAddressNFTPositions(ctx sdk.Context, address sdk.AccAddress) []types.Position {
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

func (k Keeper) ClosePositionNFT(ctx sdk.Context, positionId string) error {
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	err := k.nftKeeper.Transfer(ctx, types.PositionNFTClassId, positionId, moduleAddr)
	if err != nil {
		return err
	}
	err = k.BurnPositionNFT(ctx, positionId)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) BurnPositionNFT(ctx sdk.Context, positionId string) error {
	moduleAddr := k.GetModuleAddress()
	msgBurnNFT := nftfactorytypes.MsgBurnNFT{
		Sender:  moduleAddr.String(),
		ClassId: types.PositionNFTClassId,
		NftId:   positionId,
	}
	err := k.nftfactoryKeeper.BurnNFT(ctx, &msgBurnNFT)
	if err != nil {
		return err
	}
	return nil
}
