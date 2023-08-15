package v3_2_1

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7/types"
	buildertypes "github.com/skip-mev/pob/x/builder/types"

	"github.com/UnUniFi/chain/app/upgrades"
	nftfactorytypes "github.com/UnUniFi/chain/x/nftfactory/types"
)

const UpgradeName string = "v3_2_1"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{nftfactorytypes.StoreKey, buildertypes.StoreKey, ibchookstypes.StoreKey},
		Deleted: []string{},
	},
}
