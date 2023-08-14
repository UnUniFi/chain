package v3_2

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	buildertypes "github.com/skip-mev/pob/x/builder/types"

	"github.com/UnUniFi/chain/app/upgrades"
	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"
)

const UpgradeName string = "v3_2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{nftfactorytypes.ModuleName, buildertypes.ModuleName},
		Deleted: []string{},
	},
}
