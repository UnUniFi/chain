package v2

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/nft"

	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"

	"github.com/UnUniFi/chain/app/upgrades"
)

const UpgradeName string = "v2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{nft.StoreKey, consensusparamtypes.StoreKey, crisistypes.StoreKey},
		Deleted: []string{},
	},
}
