package keeper

import (
	"github.com/UnUniFi/chain/x/nft/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft/keeper"
)

type Keeper struct {
	keeper.Keeper
	cdc codec.BinaryCodec
}

func NewKeeper(k keeper.Keeper, cdc codec.BinaryCodec) Keeper {
	return Keeper{
		Keeper: k,
		cdc:    cdc,
	}
}

func (k Keeper) GetNftData(ctx sdk.Context, classId string, id string) (types.NftData, bool) {
	token, found := k.Keeper.GetNFT(ctx, classId, id)
	if !found {
		return types.NftData{}, false
	}

	var nftData types.NftData
	if err := k.cdc.UnpackAny(token.Data, &nftData); err != nil {
		return types.NftData{}, false
	}

	return nftData, true
}

func (k Keeper) SetNftData(ctx sdk.Context, classId string, id string, data types.NftData) error {
	token, found := k.Keeper.GetNFT(ctx, classId, id)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown NFT %s", id)
	}

	dataAny, err := codectypes.NewAnyWithValue(&data)
	if err != nil {
		return err
	}

	token.Data = dataAny
	err = k.Keeper.Update(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
