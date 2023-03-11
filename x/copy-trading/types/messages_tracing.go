package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTracing = "create_tracing"
	TypeMsgUpdateTracing = "update_tracing"
	TypeMsgDeleteTracing = "delete_tracing"
)

var _ sdk.Msg = &MsgCreateTracing{}

func NewMsgCreateTracing(
    creator string,
    index string,
    
) *MsgCreateTracing {
  return &MsgCreateTracing{
		Creator : creator,
		Index: index,
		
	}
}

func (msg *MsgCreateTracing) Route() string {
  return RouterKey
}

func (msg *MsgCreateTracing) Type() string {
  return TypeMsgCreateTracing
}

func (msg *MsgCreateTracing) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTracing) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTracing) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateTracing{}

func NewMsgUpdateTracing(
    creator string,
    index string,
    
) *MsgUpdateTracing {
  return &MsgUpdateTracing{
		Creator: creator,
        Index: index,
        
	}
}

func (msg *MsgUpdateTracing) Route() string {
  return RouterKey
}

func (msg *MsgUpdateTracing) Type() string {
  return TypeMsgUpdateTracing
}

func (msg *MsgUpdateTracing) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTracing) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTracing) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgDeleteTracing{}

func NewMsgDeleteTracing(
    creator string,
    index string,
    
) *MsgDeleteTracing {
  return &MsgDeleteTracing{
		Creator: creator,
		Index: index,
        
	}
}
func (msg *MsgDeleteTracing) Route() string {
  return RouterKey
}

func (msg *MsgDeleteTracing) Type() string {
  return TypeMsgDeleteTracing
}

func (msg *MsgDeleteTracing) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTracing) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTracing) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}