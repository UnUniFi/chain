package v1

const UpgradeName string = "upgrade_v1"

const TotalAmountValidator int64 = 4656862745096
const TotalAmountExceptValidator int64 = 304963298762
const FromAddressValidator string = "ununifi19srj7ga7t2pyflz7f50le5fv0wa9kuf7tmdtla"
const FromAddressAirdrop string = "ununifi1r500cehqg5u6fhsaysmhu4cnw5pz3lxcqhgaq7"

type ResultList struct {
	Validator                       []BankSendTarget `json:"validator"`
	AirdropCommunityRewardModerator []BankSendTarget `json:"airdropCommunityRewardModerator"`
}
type BankSendTarget struct {
	ToAddress     string `json:"toAddress,omitempty"`
	Denom         string `json:"denom,omitempty"`
	Amount        int64  `json:"amount,omitempty"`
	VestingStarts int64  `json:"vesting_starts,omitempty"`
	VestingEnds   int64  `json:"vesting_ends,omitempty"`
}
