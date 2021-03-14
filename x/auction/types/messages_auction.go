package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPlaceBid{}

func NewMsgPlaceBid(creator string) *MsgPlaceBid {
  return &MsgPlaceBid{
		Creator: creator,
	}
}

func (msg *MsgPlaceBid) Route() string {
  return RouterKey
}

func (msg *MsgPlaceBid) Type() string {
  return "CreateAuction"
}

func (msg *MsgPlaceBid) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgPlaceBid) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgPlaceBid) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateAuction{}
