package keeper

import (
	"github.com/UnUniFi/chain/x/nftmint/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

func (k Keeper) MintNFT(ctx sdk.Context, msg *types.MsgMintNFT) error {
	exists := k.nftKeeper.HasClass(ctx, msg.ClassId)
	if !exists {
		return sdkerrors.Wrap(nfttypes.ErrClassExists, msg.ClassId)
	}

	classAttributes, exists := k.GetClassAttributes(ctx, msg.ClassId)
	if !exists {
		return sdkerrors.Wrapf(types.ErrClassAttributesNotExists, "class attributes with class id %s doesn't exist", msg.ClassId)
	}
	// TODO: validate minting permission from ClassAttributes
	err := types.ValidateMintingPermission(classAttributes, msg.Sender.AccAddress())
	if err != nil {
		return err
	}

	nftUri := classAttributes.BaseTokenUri + msg.NftId
	// TODO: validate uri
	// err := types.ValidateUri(nftUri)

	// TODO: validate token supply cap

	err = k.nftKeeper.Mint(ctx, types.NewNFT(msg.ClassId, msg.NftId, nftUri), msg.Recipient.AccAddress())
	if err != nil {
		return err
	}

	err = k.SetNFTMinter(ctx, msg.ClassId, msg.NftId, msg.Sender.AccAddress())
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) BurnNFT(ctx sdk.Context, msg *types.MsgBurnNFT) error {
	if !k.nftKeeper.HasClass(ctx, msg.ClassId) {
		return sdkerrors.Wrap(nfttypes.ErrClassNotExists, msg.ClassId)
	}

	if !k.nftKeeper.HasNFT(ctx, msg.ClassId, msg.NftId) {
		return sdkerrors.Wrap(nfttypes.ErrNFTNotExists, msg.NftId)
	}

	owner := k.nftKeeper.GetOwner(ctx, msg.ClassId, msg.NftId)
	if !owner.Equals(msg.Sender.AccAddress()) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not the owner of nft %s", msg.Sender.AccAddress().String(), msg.NftId)
	}

	err := k.nftKeeper.Burn(ctx, msg.ClassId, msg.NftId)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) UpdateNFTUri(ctx sdk.Context, classID, baseTokenUri string) error {
	nfts := k.nftKeeper.GetNFTsOfClass(ctx, classID)
	if len(nfts) == 0 {
		return nil
	}

	for _, nft := range nfts {
		nftUriLatest := baseTokenUri + nft.Id
		nft.Uri = nftUriLatest
		// TODO: uri len validation
		if err := k.nftKeeper.Update(ctx, nft); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) SetNFTMinter(ctx sdk.Context, classID, nftID string, minter sdk.AccAddress) error {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixNFTMinter))

	prefixStore.Set(types.NFTMinterKey(classID, nftID), minter.Bytes())
	return nil
}

func (k Keeper) GetNFTMinter(ctx sdk.Context, classID, nftID string) (sdk.AccAddress, bool) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, []byte(types.KeyPrefixNFTMinter))

	bz := prefixStore.Get(types.NFTMinterKey(classID, nftID))
	if len(bz) == 0 {
		return nil, false
	}

	minter := sdk.AccAddress(bz)
	return minter, true
}
