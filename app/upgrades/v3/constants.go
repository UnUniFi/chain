package v3

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"

	"github.com/UnUniFi/chain/app/upgrades"

	epochstypes "github.com/UnUniFi/chain/x/epochs/types"
	icacallbackstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/icacallbacks/types"
	interchainquerytypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/interchainquery/types"
	recordstypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
	stakeibctypes "github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/types"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"
)

const UpgradeName string = "v3"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{icacontrollertypes.SubModuleName, group.ModuleName, epochstypes.ModuleName, icacallbackstypes.ModuleName, interchainquerytypes.ModuleName, recordstypes.ModuleName, stakeibctypes.ModuleName, yieldaggregatortypes.ModuleName},
		Deleted: []string{},
	},
}
