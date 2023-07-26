package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
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
	sender string,
	name, baseTokenUri string,
	tokenSupplyCap uint64,
	mintingPermission MintingPermission,
	symbol, description, classUri string,
) *MsgCreateClass {
	return &MsgCreateClass{
		Sender:            sender,
		Name:              name,
		BaseTokenUri:      baseTokenUri,
		TokenSupplyCap:    tokenSupplyCap,
		MintingPermission: mintingPermission,
		Symbol:            symbol,
		Description:       description,
		ClassUri:          classUri,
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

func NewMsgMintNFT(
	sender string,
	classID, nftID string,
	recipient string,
) *MsgMintNFT {
	return &MsgMintNFT{
		Sender:    sender,
		ClassId:   classID,
		NftId:     nftID,
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

	if err := ValidateNFTID(msg.NftId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrEmptyNFTID, "Invalid nft id (%s)", msg.NftId)
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgSendClassOwnership(sender string, classID string, recipient string) *MsgSendClassOwnership {
	return &MsgSendClassOwnership{
		Sender:    sender,
		ClassId:   classID,
		Recipient: recipient,
	}
}

func (msg MsgSendClassOwnership) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	if err := ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrEmptyClassID, "Invalid class id (%s)", msg.ClassId)
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgSendClassOwnership) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgUpdateBaseTokenUri(sender string, classID, baseTokenUri string) *MsgUpdateBaseTokenUri {
	return &MsgUpdateBaseTokenUri{
		Sender:       sender,
		ClassId:      classID,
		BaseTokenUri: baseTokenUri,
	}
}

func (msg MsgUpdateBaseTokenUri) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	if err := ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrEmptyClassID, "Invalid class id (%s)", msg.ClassId)
	}
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgUpdateBaseTokenUri) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgUpdateTokenSupplyCap(sender string, classID string, tokenSupplyCap uint64) *MsgUpdateTokenSupplyCap {
	return &MsgUpdateTokenSupplyCap{
		Sender:         sender,
		ClassId:        classID,
		TokenSupplyCap: tokenSupplyCap,
	}
}

func (msg MsgUpdateTokenSupplyCap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address is not valid")
	}

	if err := ValidateClassID(msg.ClassId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrEmptyClassID, "Invalid class id (%s)", msg.ClassId)
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgUpdateTokenSupplyCap) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

func NewMsgBurnNFT(
	burner string,
	classID, nftID string,
) *MsgBurnNFT {
	return &MsgBurnNFT{
		Sender:  burner,
		ClassId: classID,
		NftId:   nftID,
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

	if err := ValidateNFTID(msg.NftId); err != nil {
		return sdkerrors.Wrapf(nfttypes.ErrEmptyNFTID, "Invalid nft id (%s)", msg.NftId)
	}

	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgBurnNFT) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
