package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/UnUniFi/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// TypeMsgPostPrice type of PostPrice msg
	TypeMsgPostPrice = "post_price"

	// MaxExpiry defines the max expiry time defined as UNIX time (9999-12-31 23:59:59 +0000 UTC)
	MaxExpiry = 253402300799
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgPostPrice{}

// NewMsgPostPrice creates a new post price msg
func NewMsgPostPrice(
	from sdk.AccAddress,
	assetCode string,
	price sdk.Dec,
	expiry time.Time,
	deposit sdk.Coin,
) MsgPostPrice {
	return MsgPostPrice{
		From:     from.Bytes(),
		MarketId: assetCode,
		Price:    types.NewDecFromSDKDec(price),
		Expiry:   expiry,
		Deposit:  deposit,
	}
}

// Route Implements Msg.
func (msg MsgPostPrice) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgPostPrice) Type() string { return TypeMsgPostPrice }

// GetSignBytes Implements Msg.
func (msg MsgPostPrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgPostPrice) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From.AccAddress()}
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgPostPrice) ValidateBasic() error {
	if msg.From.AccAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if strings.TrimSpace(msg.MarketId) == "" {
		return errors.New("market id cannot be blank")
	}
	if msg.Price.IsNegative() {
		return fmt.Errorf("price cannot be negative: %s", msg.Price.String())
	}
	if msg.Expiry.Unix() <= 0 {
		return errors.New("must set an expiration time")
	}
	return nil
}
