package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateExemplaryTrader{}, "copy-trading/CreateExemplaryTrader", nil)
	cdc.RegisterConcrete(&MsgUpdateExemplaryTrader{}, "copy-trading/UpdateExemplaryTrader", nil)
	cdc.RegisterConcrete(&MsgDeleteExemplaryTrader{}, "copy-trading/DeleteExemplaryTrader", nil)
	cdc.RegisterConcrete(&MsgCreateTracing{}, "copy-trading/CreateTracing", nil)
	cdc.RegisterConcrete(&MsgDeleteTracing{}, "copy-trading/DeleteTracing", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateExemplaryTrader{},
		&MsgUpdateExemplaryTrader{},
		&MsgDeleteExemplaryTrader{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTracing{},
		&MsgDeleteTracing{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
