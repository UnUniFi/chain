package v2_beta1

import (
	"github.com/UnUniFi/chain/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const UpgradeName string = "v2-beta.1"

const TotalAmountValidator int64 = 9949510918111
const TotalAmountExceptValidator int64 = 304963298762
const TotalForfeitAmount int64 = 30839520258
const FundAmountValidator int64 = 20000000
const FromAddressValidator string = "ununifi19srj7ga7t2pyflz7f50le5fv0wa9kuf7tmdtla"
const FromAddressAirdrop string = "ununifi1r500cehqg5u6fhsaysmhu4cnw5pz3lxcqhgaq7"
const ToAddressAirdropForfeit string = "ununifi1r500cehqg5u6fhsaysmhu4cnw5pz3lxcqhgaq7"
const Denom string = "uguu"

type ResultList struct {
	Validator                       []BankSendTarget        `json:"validator"`
	LendValidator                   []BankSendTarget        `json:"lendValidator"`
	AirdropCommunityRewardModerator []BankSendTarget        `json:"airdropCommunityRewardModerator"`
	AirdropForfeit                  []string                `json:"airdropForfeit"`
	Others                          []SpecialBankSendTarget `json:"others"`
}

type BankSendTarget struct {
	ToAddress     string `json:"toAddress,omitempty"`
	Denom         string `json:"denom,omitempty"`
	Amount        int64  `json:"amount,omitempty"`
	VestingStarts int64  `json:"vesting_starts,omitempty"`
	VestingEnds   int64  `json:"vesting_ends,omitempty"`
}

type SpecialBankSendTarget struct {
	FromAddress    string         `json:"fromAddress,omitempty"`
	BankSendTarget BankSendTarget `json:"bankSendTarget,omitempty"`
}

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}
