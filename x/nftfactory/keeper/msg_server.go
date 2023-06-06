package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftfactory/types"
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

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	seq, err := k.keeper.accountKeeper.GetSequence(ctx, sender)
	if err != nil {
		return nil, err
	}

	classID := CreateClassId(seq, sender)
	err = k.keeper.CreateClass(ctx, classID, msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventCreateClass{
		Owner:             msg.Sender,
		ClassId:           classID,
		BaseTokenUri:      msg.BaseTokenUri,
		TokenSupplyCap:    strconv.FormatUint(msg.TokenSupplyCap, 10),
		MintingPermission: msg.MintingPermission,
	})

	return &types.MsgCreateClassResponse{}, nil
}

func (k msgServer) SendClassOwnership(c context.Context, msg *types.MsgSendClassOwnership) (*types.MsgSendClassOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if err := k.keeper.SendClassOwnership(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventSendClassOwnership{
		Sender:   msg.Sender,
		ClassId:  msg.ClassId,
		Receiver: msg.Recipient,
	})
	return &types.MsgSendClassOwnershipResponse{}, nil
}

func (k msgServer) UpdateBaseTokenUri(c context.Context, msg *types.MsgUpdateBaseTokenUri) (*types.MsgUpdateBaseTokenUriResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if err := k.keeper.UpdateBaseTokenUri(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventUpdateBaseTokenUri{
		Owner:        msg.Sender,
		ClassId:      msg.ClassId,
		BaseTokenUri: msg.BaseTokenUri,
	})
	return &types.MsgUpdateBaseTokenUriResponse{}, nil
}

func (k msgServer) UpdateTokenSupplyCap(c context.Context, msg *types.MsgUpdateTokenSupplyCap) (*types.MsgUpdateTokenSupplyCapResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if err := k.keeper.UpdateTokenSupplyCap(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventUpdateTokenSupplyCap{
		Owner:          msg.Sender,
		ClassId:        msg.ClassId,
		TokenSupplyCap: strconv.FormatUint(msg.TokenSupplyCap, 10),
	})
	return &types.MsgUpdateTokenSupplyCapResponse{}, nil
}

func (k msgServer) MintNFT(c context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if err := k.keeper.MintNFT(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventMintNFT{
		ClassId: msg.ClassId,
		NftId:   msg.NftId,
		Owner:   msg.Recipient,
		Minter:  msg.Sender,
	})
	return &types.MsgMintNFTResponse{}, nil
}

func (k msgServer) BurnNFT(c context.Context, msg *types.MsgBurnNFT) (*types.MsgBurnNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if err := k.keeper.BurnNFT(ctx, msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(&types.EventBurnNFT{
		Burner:  msg.Sender,
		ClassId: msg.ClassId,
		NftId:   msg.NftId,
	})
	return &types.MsgBurnNFTResponse{}, nil
}
