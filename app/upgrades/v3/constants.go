package v3

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/group"

	"github.com/UnUniFi/chain/app/upgrades"

	epochstypes "github.com/UnUniFi/chain/x/epochs/types"
	icacallbackstypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/icacallbacks/types"
	interchainquerytypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/interchainquery/types"
	recordstypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/records/types"
	stakeibctypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/stakeibc/types"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
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
