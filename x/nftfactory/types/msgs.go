package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
)

var (
	_ sdk.Msg = &MsgCreateClass{}
	_ sdk.Msg = &MsgUpdateClass{}
	_ sdk.Msg = &MsgMintNFT{}
	_ sdk.Msg = &MsgBurnNFT{}
	_ sdk.Msg = &MsgChangeAdmin{}
)

func NewMsgCreateClass(
	sender string,
	subclass string,
	name string,
	symbol string,
	description string,
	uri string,
	uriHash string,
) *MsgCreateClass {
	return &MsgCreateClass{
		Sender:      sender,
		Subclass:    subclass,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Uri:         uri,
		UriHash:     uriHash,
	}
}

func (msg MsgCreateClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgCreateClass) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgUpdateClass(
	sender string,
	classId string,
	name string,
	symbol string,
	description string,
	uri string,
	uriHash string,
) *MsgUpdateClass {
	return &MsgUpdateClass{
		Sender:      sender,
		ClassId:     classId,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Uri:         uri,
		UriHash:     uriHash,
	}
}

func (msg MsgUpdateClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgUpdateClass) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgMintNFT(
	sender string,
	classID, tokenID string,
	recipient string,
) *MsgMintNFT {
	return &MsgMintNFT{
		Sender:    sender,
		ClassId:   classID,
		TokenId:   tokenID,
		Recipient: recipient,
	}
}

func (msg MsgMintNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	if err := ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrEmptyClassID, "Invalid class id (%s)", msg.ClassId)
	}

	if err := ValidateNFTID(msg.TokenId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrEmptyNFTID, "Invalid nft id (%s)", msg.TokenId)
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgBurnNFT(
	burner string,
	classId, tokenId string,
) *MsgBurnNFT {
	return &MsgBurnNFT{
		Sender:  burner,
		ClassId: classId,
		TokenId: tokenId,
	}
}

func (msg MsgBurnNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	if err := ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrEmptyClassID, "Invalid class id (%s)", msg.ClassId)
	}

	if err := ValidateNFTID(msg.TokenId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrEmptyNFTID, "Invalid nft id (%s)", msg.TokenId)
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgBurnNFT) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgChangeAdmin(
	sender string,
	classId string,
	newAdmin string,
) *MsgChangeAdmin {
	return &MsgChangeAdmin{
		Sender:   sender,
		ClassId:  classId,
		NewAdmin: newAdmin,
	}
}

func (msg MsgChangeAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgChangeAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
