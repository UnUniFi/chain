package v1_beta4

import (
	"github.com/UnUniFi/chain/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const UpgradeName string = "v1-beta.4"

const TotalAmountValidator int64 = 9271174601858

// const TotalAmountEcocsytemDevelopment int64 = 408999999969
const TotalAmountEcocsytemDevelopment int64 = 402207556569 // 408999999969 - 6792443400 (TotalAmount excluding Vanya)
const TotalAmountMarketing int64 = 906880000000
const TotalAmountAdvisors int64 = 2100000000000
const TotalAmountTransferredValidator int64 = 2205862352941 // validator -> validator (In preparation for the processing of money transfers to validators, when combining accounts into a single)
const TotalDelegationAmountValidator int64 = 2050518544482  // Ecocsytem Development -> validator
const FundAmountValidator int64 = 20000000
const FromAddressValidator string = "ununifi1q6jfv5un5cc7lh26njttg0tje0jevt93shy9zv"
const FromAddressEcocsytemDevelopment string = "ununifi1pa29ejcfrylh69pvntrx3va9xej69tnx7re567"
const FromAddressMarketing string = "ununifi1khe6yv4zswaergkrzv0dmq3afcda5fx4jjmf07"
const FromAddressAdvisors string = "ununifi1y8430kzjeudf8x0zyvcdgqlzcnwt3zqzedayt9"
const Denom string = "uguu"

type ResultList struct {
	Validator            []BankSendTarget        `json:"validator"`
	LendValidator        []BankSendTarget        `json:"lendValidator"`
	EcocsytemDevelopment []BankSendTarget        `json:"ecocsytemDevelopment"`
	Marketing            []BankSendTarget        `json:"marketing"`
	Advisors             []BankSendTarget        `json:"advisors"`
	Others               []SpecialBankSendTarget `json:"others"`
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
