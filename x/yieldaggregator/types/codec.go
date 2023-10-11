package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDepositToVault{},
		&MsgWithdrawFromVault{},
		&MsgWithdrawFromVaultWithUnbondingTime{},
		&MsgCreateVault{},
		&MsgTransferVaultOwnership{},
		&MsgUpdateParams{},
		&MsgRegisterStrategy{},
		&MsgDeleteVault{},
		&MsgUpdateStrategy{},
		&MsgUpdateVault{},
		&MsgRegisterDenomInfos{},
		&MsgRegisterSymbolInfos{},
		&MsgSetIntermediaryAccountInfo{},
	)

	// Deprecated: Just for backward compatibility of query proposals
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ProposalAddStrategy{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
