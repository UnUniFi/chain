package v4_0_1

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/UnUniFi/chain/app/upgrades"
)

const UpgradeName string = "v4"

type ResultList struct {
	Campaign []BankSendTarget `json:"campaign"`
}

type BankSendTarget struct {
	Number        int64  `json:"number,omitempty"`
	ToAddress     string `json:"toAddress,omitempty"`
	Denom         string `json:"denom,omitempty"`
	Amount        int64  `json:"amount,omitempty"`
	VestingStarts int64  `json:"vesting_starts,omitempty"`
	VestingEnds   int64  `json:"vesting_ends,omitempty"`
}

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
