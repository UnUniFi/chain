package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateCdp{}, "cdp/CreateCdp", nil)
	cdc.RegisterConcrete(&MsgUpdateCdp{}, "cdp/UpdateCdp", nil)
	cdc.RegisterConcrete(&MsgDeleteCdp{}, "cdp/DeleteCdp", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCdp{},
		&MsgUpdateCdp{},
		&MsgDeleteCdp{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
