package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgListNft{}

// todo: Implementation fields
// BidToken, MinBid, BidHook, ListingType
func NewMsgListNft(sender string, nftId NftId, bidDenom string, minimumDepositRate sdk.Dec, minBiddingPeriod time.Duration) MsgListNft {
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
var _ sdk.Msg = &MsgCancelListing{}

func NewMsgCancelNftListing(sender string, nftId NftId) MsgCancelListing {
	return MsgCancelListing{
		Sender: sender,
		NftId:  nftId,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgCancelListing) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgCancelListing) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgPlaceBid{}

// todo
func NewMsgPlaceBid(sender string, nftId NftId, price, deposit sdk.Coin,
	interestRate sdk.Dec, expiry time.Time, automaticPayment bool) MsgPlaceBid {
	return MsgPlaceBid{
		Sender:           sender,
		NftId:            nftId,
		Price:            price,
		AutomaticPayment: automaticPayment,
		Expiry:           expiry,
		InterestRate:     interestRate,
		Deposit:          deposit,
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

func NewMsgCancelBid(sender string, nftId NftId) MsgCancelBid {
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

func NewMsgSellingDecision(sender string, nftId NftId) MsgSellingDecision {
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

// // ensure Msg interface compliance at compile time
// var _ sdk.Msg = &MsgEndNftListing{}

// func NewMsgEndNftListing(sender string, nftId NftIdentifier) MsgEndNftListing {
// 	return MsgEndNftListing{
// 		Sender: sender,
// 		NftId:  nftId,
// 	}
// }

// // ValidateBasic does a simple validation check that doesn't require access to state.
// func (msg MsgEndNftListing) ValidateBasic() error {
// 	return nil
// }

// // GetSigners returns the addresses of signers that must sign.
// func (msg MsgEndNftListing) GetSigners() []sdk.AccAddress {
// 	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
// 	return []sdk.AccAddress{addr}
// }

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgPayRemainder{}

func NewMsgPayRemainder(sender string, nftId NftId) MsgPayRemainder {
	return MsgPayRemainder{
		Sender: sender,
		NftId:  nftId,
	}
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgPayRemainder) ValidateBasic() error {
	return nil
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgPayRemainder) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgBorrow{}

func NewMsgBorrow(sender string, nftId NftId, borrowBids []BorrowBid) MsgBorrow {
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

func NewMsgRepay(sender string, nftId NftId, repayBids []BorrowBid) MsgRepay {
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
