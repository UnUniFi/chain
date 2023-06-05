package nfttransfer

import (
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/bianjieai/nft-transfer"
	"github.com/bianjieai/nft-transfer/types"

	newkeeper "github.com/UnUniFi/chain/x/nft-transfer/keeper"
)

// AppModule represents the AppModule for this module
type AppModule struct {
	nfttransfer.AppModule
	keeper newkeeper.Keeper
}

// NewAppModule creates a new nft-transfer module
func NewAppModule(am nfttransfer.AppModule, k newkeeper.Keeper) AppModule {
	return AppModule{
		AppModule: am,
		keeper:    k,
	}
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}
