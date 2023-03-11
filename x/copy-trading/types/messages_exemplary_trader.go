package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateExemplaryTrader = "create_exemplary_trader"
	TypeMsgUpdateExemplaryTrader = "update_exemplary_trader"
	TypeMsgDeleteExemplaryTrader = "delete_exemplary_trader"
)

var _ sdk.Msg = &MsgCreateExemplaryTrader{}

func NewMsgCreateExemplaryTrader(
    creator string,
    index string,
    
) *MsgCreateExemplaryTrader {
  return &MsgCreateExemplaryTrader{
		Creator : creator,
		Index: index,
		
	}
}

func (msg *MsgCreateExemplaryTrader) Route() string {
  return RouterKey
}

func (msg *MsgCreateExemplaryTrader) Type() string {
  return TypeMsgCreateExemplaryTrader
}

func (msg *MsgCreateExemplaryTrader) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateExemplaryTrader) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateExemplaryTrader) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateExemplaryTrader{}

func NewMsgUpdateExemplaryTrader(
    creator string,
    index string,
    
) *MsgUpdateExemplaryTrader {
  return &MsgUpdateExemplaryTrader{
		Creator: creator,
        Index: index,
        
	}
}

func (msg *MsgUpdateExemplaryTrader) Route() string {
  return RouterKey
}

func (msg *MsgUpdateExemplaryTrader) Type() string {
  return TypeMsgUpdateExemplaryTrader
}

func (msg *MsgUpdateExemplaryTrader) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateExemplaryTrader) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateExemplaryTrader) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgDeleteExemplaryTrader{}

func NewMsgDeleteExemplaryTrader(
    creator string,
    index string,
    
) *MsgDeleteExemplaryTrader {
  return &MsgDeleteExemplaryTrader{
		Creator: creator,
		Index: index,
        
	}
}
func (msg *MsgDeleteExemplaryTrader) Route() string {
  return RouterKey
}

func (msg *MsgDeleteExemplaryTrader) Type() string {
  return TypeMsgDeleteExemplaryTrader
}

func (msg *MsgDeleteExemplaryTrader) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteExemplaryTrader) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteExemplaryTrader) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}