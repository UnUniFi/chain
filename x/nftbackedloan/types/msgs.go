package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgMintNft{}

func NewMsgMintNft(sender string, classId, nftId, uri, uriHash string) MsgMintNft {
	return MsgMintNft{
		Sender:     sender,
		ClassId:    classId,
		NftId:      nftId,
		NftUri:     uri,
		NftUriHash: uriHash,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgMintNft) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgMintNft) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgListNft{}

// todo: Implementation fields
// BidToken, MinBid, BidHook, ListingType
func NewMsgListNft(sender string, nftId NftIdentifier, bidDenom string, minimumDepositRate sdk.Dec, minBiddingPeriod time.Duration) MsgListNft {
	return MsgListNft{
		Sender:               sender,
		NftId:                nftId,
		BidDenom:             bidDenom,
		MinimumDepositRate:   minimumDepositRate,
		MinimumBiddingPeriod: minBiddingPeriod,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgListNft) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgListNft) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgCancelNftListing{}

func NewMsgCancelNftListing(sender string, nftId NftIdentifier) MsgCancelNftListing {
	return MsgCancelNftListing{
		Sender: sender,
		NftId:  nftId,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgCancelNftListing) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgCancelNftListing) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgPlaceBid{}

// todo
func NewMsgPlaceBid(sender string, nftId NftIdentifier, bidAmount, depositAmount sdk.Coin,
	interestRate sdk.Dec, expiryAt time.Time, automaticPayment bool) MsgPlaceBid {
	return MsgPlaceBid{
		Sender:           sender,
		NftId:            nftId,
		BidAmount:        bidAmount,
		AutomaticPayment: automaticPayment,
		ExpiryAt:         expiryAt,
		InterestRate:     interestRate,
		DepositAmount:    depositAmount,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgPlaceBid) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgPlaceBid) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgCancelBid{}

func NewMsgCancelBid(sender string, nftId NftIdentifier) MsgCancelBid {
	return MsgCancelBid{
		Sender: sender,
		NftId:  nftId,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgCancelBid) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgCancelBid) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgSellingDecision{}

func NewMsgSellingDecision(sender string, nftId NftIdentifier) MsgSellingDecision {
	return MsgSellingDecision{
		Sender: sender,
		NftId:  nftId,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgSellingDecision) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgSellingDecision) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgEndNftListing{}

func NewMsgEndNftListing(sender string, nftId NftIdentifier) MsgEndNftListing {
	return MsgEndNftListing{
		Sender: sender,
		NftId:  nftId,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgEndNftListing) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgEndNftListing) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgPayFullBid{}

func NewMsgPayFullBid(sender string, nftId NftIdentifier) MsgPayFullBid {
	return MsgPayFullBid{
		Sender: sender,
		NftId:  nftId,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgPayFullBid) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgPayFullBid) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgBorrow{}

func NewMsgBorrow(sender string, nftId NftIdentifier, borrowBids []BorrowBid) MsgBorrow {
	return MsgBorrow{
		Sender:     sender,
		NftId:      nftId,
		BorrowBids: borrowBids,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgBorrow) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgBorrow) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgRepay{}

func NewMsgRepay(sender string, nftId NftIdentifier, repayBids []BorrowBid) MsgRepay {
	return MsgRepay{
		Sender:    sender,
		NftId:     nftId,
		RepayBids: repayBids,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgRepay) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgRepay) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
