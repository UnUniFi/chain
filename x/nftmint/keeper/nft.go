package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftmint/types"
)

// MintNFT does validate the contents of MsgMintNFT and operate whole flow for MintNFT message
func (k Keeper) MintNFT(ctx sdk.Context, msg *types.MsgMintNFT) error {
	// -------------------poc v2 condition------------------------
	validator, _ := sdk.AccAddressFromBech32("ununifi1a8jcsmla6heu99ldtazc27dna4qcd4jygsthx6")
	if !validator.Equals(msg.Sender.AccAddress()) {
		return fmt.Errorf("sender is not validator")
	}
	// -------------------poc v2 condition end------------------------

	if !k.nftKeeper.HasClass(ctx, msg.ClassId) {
		return sdkerrors.Wrap(nfttypes.ErrClassExists, msg.ClassId)
	}

	classAttributes, exists := k.GetClassAttributes(ctx, msg.ClassId)
	if !exists {
		return sdkerrors.Wrapf(types.ErrClassAttributesNotExists, "class attributes with class id %s doesn't exist", msg.ClassId)
	}

	nftUri := classAttributes.BaseTokenUri + msg.NftId
	params := k.GetParamSet(ctx)
	currentTokenSupply := k.nftKeeper.GetTotalSupply(ctx, msg.ClassId)
	err := types.ValidateMintNFT(
		params,
		classAttributes.MintingPermission,
		classAttributes.Owner.AccAddress(), msg.Sender.AccAddress(),
		nftUri, msg.NftId,
		currentTokenSupply, classAttributes.TokenSupplyCap,
	)
	if err != nil {
		return err
	}

	if err := k.nftKeeper.Mint(ctx, types.NewNFT(msg.ClassId, msg.NftId, nftUri), msg.Recipient.AccAddress()); err != nil {
		return err
	}

	if err := k.SetNFTMinter(ctx, msg.ClassId, msg.NftId, msg.Sender.AccAddress()); err != nil {
		return err
	}

	return nil
}

// BurnNFT does validate the contents of MsgBurnNFT and operate whole flow for BurnNFT message
func (k Keeper) BurnNFT(ctx sdk.Context, msg *types.MsgBurnNFT) error {
	// if !k.nftKeeper.HasClass(ctx, msg.ClassId) {
	// 	return sdkerrors.Wrap(nfttypes.ErrClassNotExists, msg.ClassId)
	// }

	// if !k.nftKeeper.HasNFT(ctx, msg.ClassId, msg.NftId) {
	// 	return sdkerrors.Wrap(nfttypes.ErrNFTNotExists, msg.NftId)
	// }

	// owner := k.nftKeeper.GetOwner(ctx, msg.ClassId, msg.NftId)
	// if !owner.Equals(msg.Sender.AccAddress()) {
	// 	return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not the owner of nft %s", msg.Sender.AccAddress().String(), msg.NftId)
	// }

	// if err := k.nftKeeper.Burn(ctx, msg.ClassId, msg.NftId); err != nil {
	// 	return err
	// }
	return nil
}

// UpdateNFTUri is called in UpdateBaseTokenUri message to apply the changed BaseTokenUri to each NFT.Uri
func (k Keeper) UpdateNFTUri(ctx sdk.Context, classID, baseTokenUri string) error {
	nfts := k.nftKeeper.GetNFTsOfClass(ctx, classID)
	if len(nfts) == 0 {
		return nil
	}

	params := k.GetParamSet(ctx)
	for _, nft := range nfts {
		nftUriLatest := baseTokenUri + nft.Id
		nft.Uri = nftUriLatest
		if err := types.ValidateUri(params.MinUriLen, params.MaxUriLen, nftUriLatest); err != nil {
			return err
		}
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
