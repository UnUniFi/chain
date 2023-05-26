package types

import (
	"errors"
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreateCdp{}
	_ sdk.Msg = &MsgDeposit{}
	_ sdk.Msg = &MsgWithdraw{}
	_ sdk.Msg = &MsgDrawDebt{}
	_ sdk.Msg = &MsgRepayDebt{}
	_ sdk.Msg = &MsgLiquidate{}
)

// NewMsgCreateCdp returns a new MsgPlaceBid.
func NewMsgCreateCdp(sender sdk.AccAddress, collateral sdk.Coin, principal sdk.Coin, collateralType string) MsgCreateCdp {
	return MsgCreateCdp{
		Sender:         sender.Bytes(),
		Collateral:     collateral,
		Principal:      principal,
		CollateralType: collateralType,
	}
}

// Route return the message type used for routing the message.
func (msg MsgCreateCdp) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgCreateCdp) Type() string { return "create_cdp" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgCreateCdp) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if msg.Collateral.IsZero() || !msg.Collateral.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "collateral amount %s", msg.Collateral)
	}
	if msg.Principal.IsZero() || !msg.Principal.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "principal amount %s", msg.Principal)
	}
	if strings.TrimSpace(msg.CollateralType) == "" {
		return fmt.Errorf("collateral type cannot be empty")
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgCreateCdp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgCreateCdp) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// NewMsgDeposit returns a new MsgDeposit
func NewMsgDeposit(owner sdk.AccAddress, depositor sdk.AccAddress, collateral sdk.Coin, collateralType string) MsgDeposit {
	return MsgDeposit{
		Owner:          owner.Bytes(),
		Depositor:      depositor.Bytes(),
		Collateral:     collateral,
		CollateralType: collateralType,
	}
}

// Route return the message type used for routing the message.
func (msg MsgDeposit) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgDeposit) Type() string { return "deposit_cdp" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgDeposit) ValidateBasic() error {
	if msg.Owner.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner address cannot be empty")
	}
	if msg.Depositor.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if !msg.Collateral.IsValid() || msg.Collateral.IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "collateral amount %s", msg.Collateral)
	}
	if strings.TrimSpace(msg.CollateralType) == "" {
		return fmt.Errorf("collateral type cannot be empty")
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Depositor.AccAddress()}
}

// NewMsgWithdraw returns a new MsgDeposit
func NewMsgWithdraw(owner sdk.AccAddress, depositor sdk.AccAddress, collateral sdk.Coin, collateralType string) MsgWithdraw {
	return MsgWithdraw{
		Owner:          owner.Bytes(),
		Depositor:      depositor.Bytes(),
		Collateral:     collateral,
		CollateralType: collateralType,
	}
}

// Route return the message type used for routing the message.
func (msg MsgWithdraw) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgWithdraw) Type() string { return "withdraw_cdp" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgWithdraw) ValidateBasic() error {
	if msg.Owner.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner address cannot be empty")
	}
	if msg.Depositor.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if !msg.Collateral.IsValid() || msg.Collateral.IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "collateral amount %s", msg.Collateral)
	}
	if strings.TrimSpace(msg.CollateralType) == "" {
		return fmt.Errorf("collateral type cannot be empty")
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Depositor.AccAddress()}
}

// NewMsgDrawDebt returns a new MsgDrawDebt
func NewMsgDrawDebt(sender sdk.AccAddress, collateralType string, principal sdk.Coin) MsgDrawDebt {
	return MsgDrawDebt{
		Sender:         sender.Bytes(),
		CollateralType: collateralType,
		Principal:      principal,
	}
}

// Route return the message type used for routing the message.
func (msg MsgDrawDebt) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgDrawDebt) Type() string { return "draw_cdp" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgDrawDebt) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if strings.TrimSpace(msg.CollateralType) == "" {
		return errors.New("cdp collateral type cannot be blank")
	}
	if msg.Principal.IsZero() || !msg.Principal.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "principal amount %s", msg.Principal)
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgDrawDebt) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgDrawDebt) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// NewMsgRepayDebt returns a new MsgRepayDebt
func NewMsgRepayDebt(sender sdk.AccAddress, collateralType string, payment sdk.Coin) MsgRepayDebt {
	return MsgRepayDebt{
		Sender:         sender.Bytes(),
		CollateralType: collateralType,
		Payment:        payment,
	}
}

// Route return the message type used for routing the message.
func (msg MsgRepayDebt) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgRepayDebt) Type() string { return "repay_cdp" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgRepayDebt) ValidateBasic() error {
	if msg.Sender.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if strings.TrimSpace(msg.CollateralType) == "" {
		return errors.New("cdp collateral type cannot be blank")
	}
	if msg.Payment.IsZero() || !msg.Payment.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "payment amount %s", msg.Payment)
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgRepayDebt) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgRepayDebt) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender.AccAddress()}
}

// NewMsgLiquidate returns a new MsgLiquidate
func NewMsgLiquidate(keeper, borrower sdk.AccAddress, ctype string) MsgLiquidate {
	return MsgLiquidate{
		Keeper:         keeper.Bytes(),
		Borrower:       borrower.Bytes(),
		CollateralType: ctype,
	}
}

// Route return the message type used for routing the message.
func (msg MsgLiquidate) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgLiquidate) Type() string { return "liquidate" }

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgLiquidate) ValidateBasic() error {
	if msg.Keeper.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "keeper address cannot be empty")
	}
	if msg.Borrower.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "borrower address cannot be empty")
	}
	if strings.TrimSpace(msg.CollateralType) == "" {
		return sdkerrors.Wrap(ErrInvalidCollateral, "collateral type cannot be empty")
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgLiquidate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgLiquidate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Keeper.AccAddress()}
}
