package v2

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	"github.com/cosmos/cosmos-sdk/x/nft"

	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"

	"github.com/UnUniFi/chain/app/upgrades"

	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"

	auctiontypes "github.com/UnUniFi/chain/deprecated/x/auction/types"
	cdptypes "github.com/UnUniFi/chain/deprecated/x/cdp/types"
	incentivetypes "github.com/UnUniFi/chain/deprecated/x/incentive/types"
	ununifidisttypes "github.com/UnUniFi/chain/deprecated/x/ununifidist/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"
)

const UpgradeName string = "v2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{group.ModuleName, nft.ModuleName, consensusparamtypes.ModuleName, crisistypes.ModuleName, ibcfeetypes.ModuleName, icahosttypes.SubModuleName},
		Deleted: []string{auctiontypes.ModuleName, cdptypes.ModuleName, incentivetypes.ModuleName, ununifidisttypes.ModuleName, pricefeedtypes.ModuleName},
	},
}
