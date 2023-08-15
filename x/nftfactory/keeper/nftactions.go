package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftfactory/types"
)

func (k Keeper) mintTo(ctx sdk.Context, token nft.NFT, mintTo string) error {
	// verify that denom is an x/nftfactory denom
	_, _, err := types.DeconstructClassId(token.ClassId)
	if err != nil {
		return err
	}

	addr, err := sdk.AccAddressFromBech32(mintTo)
	if err != nil {
		return err
	}

	err = k.nftKeeper.Mint(ctx, token, addr)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) burnFrom(ctx sdk.Context, classId string, tokenId string) error {
	// verify that denom is an x/nftfactory denom
	_, _, err := types.DeconstructClassId(classId)
	if err != nil {
		return err
	}

	err = k.nftKeeper.Burn(ctx, classId, tokenId)
	if err != nil {
		return err
	}

	return nil
}
