package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// nftmint message types
const (
	TypeMsgCreateClass = "create-class"
	TypeMsgMintNFT     = "mint-nft"
)

var _ sdk.Msg = &MsgCreateClass{}

func NewMsgCreateClass(
	sender sdk.AccAddress,
	name, baseTokenUri string,
	tokenSupplyCap uint64,
	mintingPermission MintingPermission,
	symbol, description, classUri string,
) *MsgCreateClass {
	return &MsgCreateClass{
		Sender:            sender.Bytes(),
		Name:              name,
		BaseTokenUri:      baseTokenUri,
		TokenSupplyCap:    tokenSupplyCap,
		MintingPermission: mintingPermission,
		Symbol:            symbol,
		Description:       description,
		ClassUri:          classUri,
	}
}

func (msg MsgCreateClass) Route() string { return RouterKey }

func (msg MsgCreateClass) Type() string { return TypeMsgCreateClass }

// TODO: impl validate func

// GetSigners implements Msg
func (m MsgCreateClass) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(string(m.Sender))
	return []sdk.AccAddress{signer}
}

var _ sdk.Msg = &MsgMintNFT{}

func NewMsgMintNFT(
	sender sdk.AccAddress,
	classId string,
	recipient sdk.AccAddress,
) *MsgMintNFT {
	return &MsgMintNFT{
		Sender:    sender.Bytes(),
		ClassId:   classId,
		Recipient: recipient.Bytes(),
	}
}

func (msg MsgMintNFT) Route() string { return RouterKey }

func (msg MsgMintNFT) Type() string { return TypeMsgMintNFT }

// TODO: impl validate func

// GetSigners implements Msg
func (m MsgMintNFT) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(string(m.Sender))
	return []sdk.AccAddress{signer}
}
