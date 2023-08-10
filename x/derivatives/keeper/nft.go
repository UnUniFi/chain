package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	nft "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) GetPositionNFTSendDisabled(ctx sdk.Context, positionId string) (bool, error) {
	futureDisabled, err := k.GetFuturePositionNFTSendDisabled(ctx, positionId)
	if err == nil {
		return futureDisabled, nil
	}
	optionDisabled, err := k.GetOptionPositionNFTSendDisabled(ctx, positionId)
	if err != nil {
		return false, err
	}
	return optionDisabled, nil
}

func (k Keeper) GetFuturePositionNFTSendDisabled(ctx sdk.Context, positionId string) (bool, error) {
	data, found := k.nftKeeper.GetNftData(ctx, types.PerpFuturePositionNFTClassId, positionId)
	if !found {
		return false, types.ErrPositionNFTNotFound
	}
	return data.SendDisabled, nil
}

func (k Keeper) GetOptionPositionNFTSendDisabled(ctx sdk.Context, positionId string) (bool, error) {
	data, found := k.nftKeeper.GetNftData(ctx, types.PerpOptionPositionNFTClassId, positionId)
	if !found {
		return false, types.ErrPositionNFTNotFound
	}
	return data.SendDisabled, nil
}

func (k Keeper) GetPositionNFTOwner(ctx sdk.Context, positionId string) sdk.AccAddress {
	futureOwner := k.GetFuturePositionNFTOwner(ctx, positionId)
	if futureOwner != nil {
		return futureOwner
	}
	optionOwner := k.GetOptionPositionNFTOwner(ctx, positionId)
	if optionOwner != nil {
		return optionOwner
	}
	return nil
}

func (k Keeper) GetFuturePositionNFTOwner(ctx sdk.Context, positionId string) sdk.AccAddress {
	owner := k.nftKeeper.GetOwner(ctx, types.PerpFuturePositionNFTClassId, positionId)
	return owner
}

func (k Keeper) GetOptionPositionNFTOwner(ctx sdk.Context, positionId string) sdk.AccAddress {
	owner := k.nftKeeper.GetOwner(ctx, types.PerpOptionPositionNFTClassId, positionId)
	return owner
}

func (k Keeper) GetAddressNFTPositions(ctx sdk.Context, address sdk.AccAddress) []types.Position {
	futurePositions := k.GetAddressNFTFuturePositions(ctx, address)
	optionPositions := k.GetAddressNFTOptionPositions(ctx, address)
	positions := append(futurePositions, optionPositions...)
	return positions
}

func (k Keeper) GetAddressNFTFuturePositions(ctx sdk.Context, address sdk.AccAddress) []types.Position {
	nfts := k.nftKeeper.GetNFTsOfClassByOwner(ctx, types.PerpFuturePositionNFTClassId, address)
	positions := []types.Position{}
	for _, nft := range nfts {
		position := k.GetPositionWithId(ctx, nft.Id)
		if position == nil {
			continue
		}
		positions = append(positions, *position)
	}
	return positions
}

func (k Keeper) GetAddressNFTOptionPositions(ctx sdk.Context, address sdk.AccAddress) []types.Position {
	nfts := k.nftKeeper.GetNFTsOfClassByOwner(ctx, types.PerpOptionPositionNFTClassId, address)
	positions := []types.Position{}
	for _, nft := range nfts {
		position := k.GetPositionWithId(ctx, nft.Id)
		if position == nil {
			continue
		}
		positions = append(positions, *position)
	}
	return positions
}

func (k Keeper) MintFuturePositionNFT(ctx sdk.Context, position types.Position) error {
	receiver, err := sdk.AccAddressFromBech32(position.OpenerAddress)
	err = k.nftKeeper.Mint(ctx, nft.NFT{
		ClassId: types.PerpFuturePositionNFTClassId,
		Id:      position.Id,
		Uri:     "",
	},
		receiver)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) CloseFuturePositionNFT(ctx sdk.Context, positionId string) error {
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	err := k.nftKeeper.Transfer(ctx, types.PerpFuturePositionNFTClassId, positionId, moduleAddr)
	if err != nil {
		return err
	}
	err = k.BurnFuturePositionNFT(ctx, positionId)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) BurnFuturePositionNFT(ctx sdk.Context, positionId string) error {
	err := k.nftKeeper.Burn(ctx, types.PerpFuturePositionNFTClassId, positionId)
	if err != nil {
		return err
	}
	return nil
}
