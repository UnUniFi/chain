package v1

const UpgradeName string = "upgrade_v1"

const TOTAL_AMOUNT_VALIDATOR int64 = 100001        // 4656862745096
const TOTAL_AMOUNT_EXCEPT_VALIDATOR int64 = 100005 // 294215472734

type ResultList struct {
	Validator                       []BankSendTarget `json:"validator"`
	AirdropCommunityRewardModerator []BankSendTarget `json:"airdropCommunityRewardModerator"`
}
type BankSendTarget struct {
	FromAddress   string `json:"fromAddress,omitempty"`
	ToAddress     string `json:"toAddress,omitempty"`
	Amount        string `json:"amount,omitempty"`
	Denom         int64  `json:"denom,omitempty"`
	VestingStarts int64  `json:"vesting_starts,omitempty"`
	VestingEnds   int64  `json:"vesting_ends,omitempty"`
}
