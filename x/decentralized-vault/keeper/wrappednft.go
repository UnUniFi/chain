package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

func (k Keeper) MintWrappedNft(ctx sdk.Context, msg *types.MsgNftLocked) error {
	//todo: check trust sender address
	nftId := msg.NftId
	_, exists := k.nftKeeper.GetNFT(ctx, types.WrappedClassId, nftId)
	if exists {
		return nft.ErrNFTExists
	}

	_, hasId := k.nftKeeper.GetClass(ctx, types.WrappedClassId)
	if !hasId {
		class := nft.Class{
			Id:          types.WrappedClassId,
			Name:        types.WrappedClassName,
			Symbol:      types.WrappedClassSymbol,
			Description: types.WrappedClassDescription,
		}
		k.nftKeeper.SaveClass(ctx, class)
	}

	expNFT := nft.NFT{
		ClassId: types.WrappedClassId,
		Id:      nftId,
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
	}
	err := k.nftKeeper.Mint(ctx, expNFT, msg.ToAddress.AccAddress())

	return err
}

func (k Keeper) BurnWrappedNft(ctx sdk.Context, msg *types.MsgNftUnlocked) error {
	//todo: check trust sender address
	_, exists := k.nftKeeper.GetNFT(ctx, types.WrappedClassId, msg.NftId)
	if !exists {
		return nft.ErrNFTNotExists
	}

	// todo: check nft market
	// todo: check nft owner

	err := k.nftKeeper.Burn(ctx, types.WrappedClassId, msg.NftId)

	return err
}

func (k Keeper) DepositWrappedNft(ctx sdk.Context, msg *types.MsgNftTransferRequest) error {
	//todo: check trust sender address
	_, exists := k.nftKeeper.GetNFT(ctx, types.WrappedClassId, msg.NftId)
	if !exists {
		return nft.ErrNFTNotExists
	}

	// todo: check nft market
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	owner := k.nftKeeper.GetOwner(ctx, types.WrappedClassId, msg.NftId)
	fmt.Println(owner.String())
	fmt.Println(msg.Sender.AccAddress().String())
	if owner.String() != msg.Sender.AccAddress().String() {
		return types.ErrNotNftOwner
	}

	transferRequest := types.TransferRequest{
		NftId:      msg.NftId,
		Owner:      owner.String(),
		EthAddress: msg.EthAddress,
	}
	k.SetTransferRequest(ctx, transferRequest)
	err := k.nftKeeper.Transfer(ctx, types.WrappedClassId, msg.NftId, moduleAddr)
	return err
}

func (k Keeper) WithdrawWrappedNft(ctx sdk.Context, msg *types.MsgNftRejectTransfer) error {
	//todo: check trust sender address
	_, exists := k.nftKeeper.GetNFT(ctx, types.WrappedClassId, msg.NftId)
	if !exists {
		return nft.ErrNFTNotExists
	}

	// todo: check nft market

	req, err := k.GetTransferRequestByIdBytes(ctx, []byte(msg.NftId))
	// todo: error check
	owner, err := sdk.AccAddressFromBech32(req.Owner)
	err = k.nftKeeper.Transfer(ctx, types.WrappedClassId, msg.NftId, owner)
	return err
}

func (k Keeper) SetTransferRequest(ctx sdk.Context, transferRequest types.TransferRequest) {
	nftIdBytes := transferRequest.IdBytes()
	bz := k.cdc.MustMarshal(&transferRequest)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.TransferRequestKey(nftIdBytes), bz)
}

func (k Keeper) GetTransferRequestByIdBytes(ctx sdk.Context, nftIdBytes []byte) (types.TransferRequest, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TransferRequestKey(nftIdBytes))
	if bz == nil {
		return types.TransferRequest{}, types.ErrTransferRequestDoesNotExists
	}
	transferRequest := types.TransferRequest{}
	k.cdc.MustUnmarshal(bz, &transferRequest)
	return transferRequest, nil
}

func (k Keeper) NftTransferred(ctx sdk.Context, msg *types.MsgNftTransferred) error {
	//todo: check trust sender address
	_, exists := k.nftKeeper.GetNFT(ctx, types.WrappedClassId, msg.NftId)
	if !exists {
		return nft.ErrNFTNotExists
	}

	// todo: check nft market
	// todo: check nft owner

	err := k.nftKeeper.Burn(ctx, types.WrappedClassId, msg.NftId)

	return err
}
