package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/decentralized-vault/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

func (k Keeper) NftLocked(ctx sdk.Context, msg *types.MsgNftLocked) error {
	isTrustworthySender, err := k.IsTrustedSender(ctx, msg.Sender.AccAddress())
	if !isTrustworthySender {
		return err
	}

	return k.MintWrappedNft(ctx, msg.NftId, msg.Uri, msg.UriHash, msg.ToAddress.AccAddress())
}

func (k Keeper) NftUnlocked(ctx sdk.Context, msg *types.MsgNftUnlocked) error {
	isTrustworthySender, err := k.IsTrustedSender(ctx, msg.Sender.AccAddress())
	if !isTrustworthySender {
		return err
	}

	if k.IsListedNft(ctx, msg.NftId) {
		return types.ErrInUseNft
	}

	if k.hasNft(ctx, msg.NftId, msg.ToAddress.AccAddress()) {
		return types.ErrNotNftOwner
	}

	return k.BurnWrappedNft(ctx, msg.NftId)
}

func (k Keeper) NftTransferRequest(ctx sdk.Context, msg *types.MsgNftTransferRequest) error {
	return k.DepositWrappedNft(ctx, msg)
}

func (k Keeper) NftRejectTransfer(ctx sdk.Context, msg *types.MsgNftRejectTransfer) error {
	return k.WithdrawWrappedNft(ctx, msg)
}

func (k Keeper) NftTransferred(ctx sdk.Context, msg *types.MsgNftTransferred) error {
	isTrustworthySender, err := k.IsTrustedSender(ctx, msg.Sender.AccAddress())
	if !isTrustworthySender {
		return err
	}

	if k.IsListedNft(ctx, msg.NftId) {
		return types.ErrInUseNft
	}

	if k.hasNft(ctx, msg.NftId, k.accountKeeper.GetModuleAddress(types.ModuleName)) {
		return types.ErrNotDepositedNft
	}

	return k.BurnWrappedNft(ctx, msg.NftId)
}

func (k Keeper) IsTrustedSender(ctx sdk.Context, senderAddress sdk.AccAddress) (bool, error) {
	oracles := k.GetOracles(ctx, "Ethereum")
	if len(oracles) == 0 {
		return false, types.ErrOracleDoesNotRegister
	}
	for _, oracle := range oracles {
		if senderAddress.Equals(oracle) {
			return true, nil
		}
	}
	return false, types.ErrOracleDoesNotMatch
}

func (k Keeper) MintWrappedNft(ctx sdk.Context, nftId, uri, uriHash string, receiver sdk.AccAddress) error {
	var err error
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
		err = k.nftKeeper.SaveClass(ctx, class)
		if err != nil {
			return err
		}
	}

	expNFT := nft.NFT{
		ClassId: types.WrappedClassId,
		Id:      nftId,
		Uri:     uri,
		UriHash: uriHash,
	}
	err = k.nftKeeper.Mint(ctx, expNFT, receiver)

	return err
}

func (k Keeper) BurnWrappedNft(ctx sdk.Context, nftId string) error {
	_, exists := k.nftKeeper.GetNFT(ctx, types.WrappedClassId, nftId)
	if !exists {
		return nft.ErrNFTNotExists
	}

	err := k.nftKeeper.Burn(ctx, types.WrappedClassId, nftId)

	return err
}

func (k Keeper) DepositWrappedNft(ctx sdk.Context, msg *types.MsgNftTransferRequest) error {
	_, exists := k.nftKeeper.GetNFT(ctx, types.WrappedClassId, msg.NftId)
	if !exists {
		return nft.ErrNFTNotExists
	}

	if k.IsListedNft(ctx, msg.NftId) {
		return types.ErrInUseNft
	}

	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	owner := k.nftKeeper.GetOwner(ctx, types.WrappedClassId, msg.NftId)
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
	isTrustworthySender, err := k.IsTrustedSender(ctx, msg.Sender.AccAddress())
	if !isTrustworthySender {
		return err
	}

	_, exists := k.nftKeeper.GetNFT(ctx, types.WrappedClassId, msg.NftId)
	if !exists {
		return nft.ErrNFTNotExists
	}

	if k.IsListedNft(ctx, msg.NftId) {
		return types.ErrInUseNft
	}
	if k.hasNft(ctx, msg.NftId, k.accountKeeper.GetModuleAddress(types.ModuleName)) {
		return types.ErrNotDepositedNft
	}

	req, err := k.GetTransferRequestByIdBytes(ctx, []byte(msg.NftId))
	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return err
	}

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

func (k Keeper) IsListedNft(ctx sdk.Context, nftId string) bool {
	nftIde := nftmarkettypes.NftIdentifier{
		ClassId: types.WrappedClassId,
		NftId:   nftId,
	}
	lisitngNft, err := k.nftmarketKeeper.GetNftListingByIdBytes(ctx, nftIde.IdBytes())
	if (lisitngNft != nftmarkettypes.NftListing{} || err != nil) {
		return false
	} else {
		return true
	}
}

func (k Keeper) hasNft(ctx sdk.Context, nftId string, address sdk.AccAddress) bool {
	owner := k.nftKeeper.GetOwner(ctx, types.WrappedClassId, nftId)
	if owner.String() != address.String() {
		return false
	} else {
		return true
	}
}
