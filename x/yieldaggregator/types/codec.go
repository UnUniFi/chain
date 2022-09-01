package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgAddFarmingOrder{},
		&MsgDeleteFarmingOrder{},
		&MsgActivateFarmingOrder{},
		&MsgInactivateFarmingOrder{},
		&MsgExecuteFarmingOrders{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ProposalAddYieldFarm{},
		&ProposalUpdateYieldFarm{},
		&ProposalStopYieldFarm{},
		&ProposalRemoveYieldFarm{},
		&ProposalAddYieldFarmTarget{},
		&ProposalUpdateYieldFarmTarget{},
		&ProposalStopYieldFarmTarget{},
		&ProposalRemoveYieldFarmTarget{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
