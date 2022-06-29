package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// nftmint message types
const (
	TypeMsgCreateClass = "create-class"
	TypeMsgMintNFT     = "mint-nft"
)

var (
	_ sdk.Msg = &MsgCreateClass{}
	_ sdk.Msg = &MsgMintNFT{}
)

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

// TODO: Impl validate func
func (msg MsgCreateClass) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	// TODO: the validation against:
	// Name
	// BaseTokenUri
	// TokenSupplyCap
	// MintingPermission
	// Symbol
	// Description
	// ClassUri

	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgCreateClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgCreateClass) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

func NewMsgMintNFT(
	sender sdk.AccAddress,
	classID, nftID string,
	recipient sdk.AccAddress,
) *MsgMintNFT {
	return &MsgMintNFT{
		Sender:    sender.Bytes(),
		ClassId:   classID,
		NftId:     nftID,
		Recipient: recipient.Bytes(),
	}
}

func (msg MsgMintNFT) Route() string { return RouterKey }

func (msg MsgMintNFT) Type() string { return TypeMsgMintNFT }

// TODO: Impl validate func
func (msg MsgMintNFT) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	// TODO: the validation against:
	// class-id validation
	// if class-id exists
	// nft-id validation
	//

	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}
