package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgListNft{}

// todo: Implementation fields
// BidToken, MinBid, BidHook, ListingType
func NewMsgListNft(sender sdk.AccAddress, nftId NftIdentifier) MsgListNft {
	return MsgListNft{
		Sender:        sender.Bytes(),
		NftId:         nftId,
		BidToken:      "uguu",
		MinBid:        sdk.NewInt(1),
		BidActiveRank: 1,
	}
}

// Route return the message type used for routing the message.
func (msg MsgListNft) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgListNft) Type() string { return "list_nft" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgListNft) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgListNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgListNft) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgCancelNftListing{}

func NewMsgCancelNftListing(sender sdk.AccAddress) MsgCancelNftListing {
	return MsgCancelNftListing{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgCancelNftListing) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgCancelNftListing) Type() string { return "cancel_nft_listing" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgCancelNftListing) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgCancelNftListing) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgCancelNftListing) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgExpandListingPeriod{}

func NewMsgExpandListingPeriod(sender sdk.AccAddress) MsgExpandListingPeriod {
	return MsgExpandListingPeriod{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgExpandListingPeriod) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgExpandListingPeriod) Type() string { return "expand_listing_period" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgExpandListingPeriod) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgExpandListingPeriod) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgExpandListingPeriod) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgPlaceBid{}

func NewMsgPlaceBid(sender sdk.AccAddress) MsgPlaceBid {
	return MsgPlaceBid{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgPlaceBid) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgPlaceBid) Type() string { return "place_bid" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgPlaceBid) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgPlaceBid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgPlaceBid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgCancelBid{}

func NewMsgCancelBid(sender sdk.AccAddress) MsgCancelBid {
	return MsgCancelBid{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgCancelBid) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgCancelBid) Type() string { return "cancel_bid" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgCancelBid) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgCancelBid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgCancelBid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgEndNftListing{}

func NewMsgEndNftListing(sender sdk.AccAddress) MsgEndNftListing {
	return MsgEndNftListing{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgEndNftListing) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgEndNftListing) Type() string { return "end_nft_listing" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgEndNftListing) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgEndNftListing) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgEndNftListing) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgPayFullBid{}

func NewMsgPayFullBid(sender sdk.AccAddress) MsgPayFullBid {
	return MsgPayFullBid{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgPayFullBid) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgPayFullBid) Type() string { return "pay_full_bid" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgPayFullBid) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgPayFullBid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgPayFullBid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgBorrow{}

func NewMsgBorrow(sender sdk.AccAddress) MsgBorrow {
	return MsgBorrow{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgBorrow) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgBorrow) Type() string { return "borrow" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgBorrow) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgBorrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgBorrow) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgRepay{}

func NewMsgRepay(sender sdk.AccAddress) MsgRepay {
	return MsgRepay{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgRepay) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgRepay) Type() string { return "repay" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgRepay) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgRepay) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgRepay) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgMintStableCoin{}

func NewMsgMintStableCoin(sender sdk.AccAddress) MsgMintStableCoin {
	return MsgMintStableCoin{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgMintStableCoin) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgMintStableCoin) Type() string { return "mint_stable_coin" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgMintStableCoin) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgMintStableCoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgMintStableCoin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgMintStableCoin{}

func NewMsgBurnStableCoin(sender sdk.AccAddress) MsgBurnStableCoin {
	return MsgBurnStableCoin{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgBurnStableCoin) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgBurnStableCoin) Type() string { return "burn_stable_coin" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgBurnStableCoin) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgBurnStableCoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgBurnStableCoin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgMintStableCoin{}

func NewMsgLiquidate(sender sdk.AccAddress) MsgLiquidate {
	return MsgLiquidate{
		Sender: sender.Bytes(),
	}
}

// Route return the message type used for routing the message.
func (msg MsgLiquidate) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgLiquidate) Type() string { return "liquidate" }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgLiquidate) ValidateBasic() error {
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgLiquidate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgLiquidate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}
