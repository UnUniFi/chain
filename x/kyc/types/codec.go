package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateProvider{}, "kyc/CreateProvider", nil)
	cdc.RegisterConcrete(&MsgUpdateProvider{}, "kyc/UpdateProvider", nil)
	cdc.RegisterConcrete(&MsgDeleteProvider{}, "kyc/DeleteProvider", nil)
	cdc.RegisterConcrete(&MsgCreateVerification{}, "kyc/CreateVerification", nil)
	cdc.RegisterConcrete(&MsgUpdateVerification{}, "kyc/UpdateVerification", nil)
	cdc.RegisterConcrete(&MsgDeleteVerification{}, "kyc/DeleteVerification", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProvider{},
		&MsgUpdateProvider{},
		&MsgDeleteProvider{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVerification{},
		&MsgUpdateVerification{},
		&MsgDeleteVerification{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
