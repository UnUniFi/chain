package v3

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/group"

	"github.com/UnUniFi/chain/app/upgrades"
)

const UpgradeName string = "v3"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{group.ModuleName},
		Deleted: []string{},
	},
}
