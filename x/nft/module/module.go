package module

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/cosmos/cosmos-sdk/x/nft"
	nftmodule "github.com/cosmos/cosmos-sdk/x/nft/module"

	newkeeper "github.com/UnUniFi/chain/x/nft/keeper"
	"github.com/UnUniFi/chain/x/nft/types"
)

// AppModuleBasic defines the basic application module used by the nft module.
type AppModuleBasic struct {
	nftmodule.AppModuleBasic
}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	a.AppModuleBasic.RegisterInterfaces(reg)
	types.RegisterInterfaces(reg)
}

// AppModule implements the sdk.AppModule interface
type AppModule struct {
	nftmodule.AppModule
	keeper newkeeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(am nftmodule.AppModule, k newkeeper.Keeper) AppModule {
	return AppModule{
		AppModule: am,
		keeper:    k,
	}
}

// RegisterServices registers a gRPC query service to respond to the
// module-specific gRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	nft.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	nft.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}
