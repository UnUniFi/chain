package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftfactory/types"

	"github.com/cosmos/cosmos-sdk/x/nft"
)

type msgServer struct {
	keeper Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateClass(c context.Context, msg *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	classId, err := k.keeper.CreateClass(ctx, msg.Sender, msg.Subclass)
	if err != nil {
		return nil, err
	}

	class := nft.Class{
		Id:          classId,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Description: msg.Description,
		Uri:         msg.Uri,
		UriHash:     msg.UriHash,
	}
	err = k.keeper.nftKeeper.SaveClass(ctx, class)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventCreateClass{
		Sender:  msg.Sender,
		ClassId: classId,
	})

	return &types.MsgCreateClassResponse{}, nil
}

func (k msgServer) UpdateClass(c context.Context, msg *types.MsgUpdateClass) (*types.MsgUpdateClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	authorityMetadata, err := k.keeper.GetAuthorityMetadata(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}

	if msg.Sender != authorityMetadata.GetAdmin() {
		return nil, types.ErrUnauthorized
	}

	k.keeper.nftKeeper.SaveClass(ctx, nft.Class{
		Id:          msg.ClassId,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Description: msg.Description,
		Uri:         msg.Uri,
		UriHash:     msg.UriHash,
	})

	ctx.EventManager().EmitTypedEvent(&types.EventUpdateClass{
		Sender:  msg.Sender,
		ClassId: msg.ClassId,
	})

	return &types.MsgUpdateClassResponse{}, nil
}

func (k msgServer) MintNFT(c context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	_, classExists := k.keeper.nftKeeper.GetClass(ctx, msg.ClassId)
	if !classExists {
		return nil, types.ErrClassDoesNotExist.Wrapf("class id: %s", msg.ClassId)
	}

	authorityMetadata, err := k.keeper.GetAuthorityMetadata(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}

	if msg.Sender != authorityMetadata.GetAdmin() {
		return nil, types.ErrUnauthorized
	}

	err = k.keeper.mintTo(ctx, nft.NFT{
		ClassId: msg.ClassId,
		Id:      msg.TokenId,
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
	}, msg.Sender)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventMintNFT{
		Sender:    msg.Sender,
		ClassId:   msg.ClassId,
		TokenId:   msg.TokenId,
		Recipient: msg.Recipient,
	})
	return &types.MsgMintNFTResponse{}, nil
}

func (k msgServer) BurnNFT(c context.Context, msg *types.MsgBurnNFT) (*types.MsgBurnNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	authorityMetadata, err := k.keeper.GetAuthorityMetadata(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}

	if msg.Sender != authorityMetadata.GetAdmin() {
		return nil, types.ErrUnauthorized
	}

	err = k.keeper.burnFrom(ctx, msg.ClassId, msg.TokenId)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventBurnNFT{
		Sender:  msg.Sender,
		ClassId: msg.ClassId,
		TokenId: msg.TokenId,
	})
	return &types.MsgBurnNFTResponse{}, nil
}

func (k msgServer) ChangeAdmin(c context.Context, msg *types.MsgChangeAdmin) (*types.MsgChangeAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	authorityMetadata, err := k.keeper.GetAuthorityMetadata(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}

	if msg.Sender != authorityMetadata.GetAdmin() {
		return nil, types.ErrUnauthorized
	}

	err = k.keeper.setAdmin(ctx, msg.ClassId, msg.NewAdmin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventChangeAdmin{
		Admin:    msg.Sender,
		ClassId:  msg.ClassId,
		NewAdmin: msg.NewAdmin,
	})
	return &types.MsgChangeAdminResponse{}, nil
}
