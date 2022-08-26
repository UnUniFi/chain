package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgNftLocked{}

func NewMsgNftLocked(sender, receiver sdk.AccAddress, nftId, uri, uriHash string) MsgNftLocked {
	return MsgNftLocked{
		Sender:    sender.Bytes(),
		ToAddress: receiver.Bytes(),
		NftId:     nftId,
		Uri:       uri,
		UriHash:   uriHash,
	}
}

// Route return the message type used for routing the message.
func (msg MsgNftLocked) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgNftLocked) Type() string { return "wrapped_nft_locked" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgNftLocked) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgNftLocked) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgNftLocked) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgNftUnlocked{}

func NewMsgNftUnlocked(sender, target sdk.AccAddress, nftId string) MsgNftUnlocked {
	return MsgNftUnlocked{
		Sender:    sender.Bytes(),
		ToAddress: target.Bytes(),
		NftId:     nftId,
	}
}

// Route return the message type used for routing the message.
func (msg MsgNftUnlocked) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgNftUnlocked) Type() string { return "wrapped_nft_unlocked" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgNftUnlocked) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgNftUnlocked) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgNftUnlocked) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

var _ sdk.Msg = &MsgNftTransferRequest{}

func NewMsgNftTransferRequest(sender sdk.AccAddress, nftId string) MsgNftTransferRequest {
	return MsgNftTransferRequest{
		Sender: sender.Bytes(),
		NftId:  nftId,
	}
}

// Route return the message type used for routing the message.
func (msg MsgNftTransferRequest) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgNftTransferRequest) Type() string { return "wrapped_nft_transfer_request" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgNftTransferRequest) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgNftTransferRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgNftTransferRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

var _ sdk.Msg = &MsgNftRejectTransfer{}

func NewMsgNftRejectTransfer(sender sdk.AccAddress, nftId string) MsgNftRejectTransfer {
	return MsgNftRejectTransfer{
		Sender: sender.Bytes(),
		NftId:  nftId,
	}
}

// Route return the message type used for routing the message.
func (msg MsgNftRejectTransfer) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgNftRejectTransfer) Type() string { return "wrapped_nft_reject_transfer" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgNftRejectTransfer) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgNftRejectTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgNftRejectTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

var _ sdk.Msg = &MsgNftTransferred{}

func NewMsgNftTransferred(sender sdk.AccAddress, nftId string) MsgNftTransferred {
	return MsgNftTransferred{
		Sender: sender.Bytes(),
		NftId:  nftId,
	}
}

// Route return the message type used for routing the message.
func (msg MsgNftTransferred) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgNftTransferred) Type() string { return "wrapped_nft_transferred" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgNftTransferred) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgNftTransferred) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgNftTransferred) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}
