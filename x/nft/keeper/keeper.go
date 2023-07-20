package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft/keeper"

	"github.com/UnUniFi/chain/x/nft/types"
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

	var nftDataI types.NftDataI
	if err := k.cdc.UnpackAny(token.Data, nftDataI); err != nil {
		return types.NftData{}, false
	}

	switch nftData := nftDataI.(type) {
	case *types.NftData:
		return *nftData, true
	default:
		// if the type is not *types.NftData, return an empty NftData
		return types.NftData{}, true
	}
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
