package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func ValidateListNftMsg(k Keeper, ctx sdk.Context, msg *types.MsgListNft) error {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	return CheckListNft(k, ctx, sender, msg.NftId, msg.BidDenom, msg.MinDepositRate)
}

func CheckListNft(k Keeper, ctx sdk.Context, sender sdk.AccAddress, nftId types.NftId, bidToken string, minimumDepositRate sdk.Dec) error {
	// check listing already exists
	_, err := k.GetNftListingByIdBytes(ctx, nftId.IdBytes())
	if err == nil {
		return err
	}

	// Check nft exists
	_, found := k.nftKeeper.GetNFT(ctx, nftId.ClassId, nftId.TokenId)
	if !found {
		return types.ErrNftDoesNotExists
	}

	// check ownership of nft
	owner := k.nftKeeper.GetOwner(ctx, nftId.ClassId, nftId.TokenId)
	if owner.String() != sender.String() {
		return types.ErrNotNftOwner
	}

	return nil
}
