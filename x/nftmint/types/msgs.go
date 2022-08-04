package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

// nftmint message types
const (
	TypeMsgCreateClass          = "create-class"
	TypeMsgMintNFT              = "mint-nft"
	TypeMsgSendClassOwnership   = "send-class"
	TypeMsgUpdateBaseTokenUri   = "update-base-token-uri"
	TypeMsgUpdateTokenSupplyCap = "update-token-supply-cap"
	TypeMsgBurnNFT              = "burn-nft"
)

var (
	_ sdk.Msg = &MsgCreateClass{}
	_ sdk.Msg = &MsgMintNFT{}
	_ sdk.Msg = &MsgSendClassOwnership{}
	_ sdk.Msg = &MsgUpdateBaseTokenUri{}
	_ sdk.Msg = &MsgUpdateTokenSupplyCap{}
	_ sdk.Msg = &MsgBurnNFT{}
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

func (msg MsgCreateClass) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}

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

func (msg MsgMintNFT) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}

	if err := nfttypes.ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrInvalidClassID, "Invalid class id (%s)", msg.ClassId)
	}

	if err := nfttypes.ValidateNFTID(msg.NftId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrInvalidID, "Invalid nft id (%s)", msg.NftId)
	}

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

func NewMsgSendClassOwnership(sender sdk.AccAddress, classID string, recipient sdk.AccAddress) *MsgSendClassOwnership {
	return &MsgSendClassOwnership{
		Sender:    sender.Bytes(),
		ClassId:   classID,
		Recipient: recipient.Bytes(),
	}
}

func (msg MsgSendClassOwnership) Route() string { return RouterKey }

func (msg MsgSendClassOwnership) Type() string { return TypeMsgSendClassOwnership }

func (msg MsgSendClassOwnership) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}

	if err := nfttypes.ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrInvalidClassID, "Invalid class id (%s)", msg.ClassId)
	}

	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgSendClassOwnership) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgSendClassOwnership) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

func NewMsgUpdateBaseTokenUri(sender sdk.AccAddress, classID, baseTokenUri string) *MsgUpdateBaseTokenUri {
	return &MsgUpdateBaseTokenUri{
		Sender:       sender.Bytes(),
		ClassId:      classID,
		BaseTokenUri: baseTokenUri,
	}
}

func (msg MsgUpdateBaseTokenUri) Route() string { return RouterKey }

func (msg MsgUpdateBaseTokenUri) Type() string { return TypeMsgUpdateBaseTokenUri }

func (msg MsgUpdateBaseTokenUri) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}

	if err := nfttypes.ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrInvalidClassID, "Invalid class id (%s)", msg.ClassId)
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgUpdateBaseTokenUri) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgUpdateBaseTokenUri) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

func NewMsgUpdateTokenSupplyCap(sender sdk.AccAddress, classID string, tokenSupplyCap uint64) *MsgUpdateTokenSupplyCap {
	return &MsgUpdateTokenSupplyCap{
		Sender:         sender.Bytes(),
		ClassId:        classID,
		TokenSupplyCap: tokenSupplyCap,
	}
}

func (msg MsgUpdateTokenSupplyCap) Route() string { return RouterKey }

func (msg MsgUpdateTokenSupplyCap) Type() string { return TypeMsgUpdateTokenSupplyCap }

func (msg MsgUpdateTokenSupplyCap) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}

	if err := nfttypes.ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrInvalidClassID, "Invalid class id (%s)", msg.ClassId)
	}

	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgUpdateTokenSupplyCap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgUpdateTokenSupplyCap) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

func NewMsgBurnNFT(
	burner sdk.AccAddress,
	classID, nftID string,
) *MsgBurnNFT {
	return &MsgBurnNFT{
		Sender:  burner.Bytes(),
		ClassId: classID,
		NftId:   nftID,
	}
}

func (msg MsgBurnNFT) Route() string { return RouterKey }

func (msg MsgBurnNFT) Type() string { return TypeMsgBurnNFT }

func (msg MsgBurnNFT) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}

	if err := nfttypes.ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrInvalidClassID, "Invalid class id (%s)", msg.ClassId)
	}

	if err := nfttypes.ValidateNFTID(msg.NftId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrInvalidID, "Invalid nft id (%s)", msg.NftId)
	}

	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgBurnNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgBurnNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}
