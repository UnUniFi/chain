package v1

const UpgradeName = "upgrade_v1"

type ResultList struct {
	Response []BankSendTarget `json:"response"`
}
type BankSendTarget struct {
	FromAddress string `json:"fromAddress,omitempty"`
	ToAddress   string `json:"toAddress,omitempty"`
	Amount      string `json:"amount,omitempty"`
	Denom       int64  `json:"denom,omitempty"`
}
