package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTitle                         = "title"
	FlagDescription                   = "description"
	FlagDeposit                       = "deposit"
	FlagAssetManagementAccountId      = "assetmanagement-account-id"
	FlagAssetManagementAccountName    = "assetmanagement-account-name"
	FlagEnabled                       = "enabled"
	FlagAssetManagementAccountAddress = "assetmanagement-account-address"
	FlagAssetManagementTargetId       = "assetmanagement-target-id"
	FlagUnbondingSeconds              = "unbonding-seconds"
	FlagAssetConditions               = "asset-conditions"
	FlagIntegrateType                 = "integration-type"
	FlagModName                       = "module-name"
	FlagExecuteOrders                 = "execute-orders"
	FlagFarmingOrderId                = "farming-order-id"
	FlagStrategyType                  = "strategy-type"
	FlagWhitelistedTargetIds          = "whitelisted-target-ids"
	FlagBlacklistedTargetIds          = "blacklisted-target-ids"
	FlagMaxUnbondingSeconds           = "max-unbonding-seconds"
	FlagOverallRatio                  = "overall-ratio"
	FlagMin                           = "min"
	FlagMax                           = "max"
	FlagDate                          = "date"
	FlagActive                        = "active"
)

func FlagProposalTx() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagTitle, "", "title of the proposal")
	fs.String(FlagDescription, "", "description of the proposal")
	fs.String(FlagDeposit, "", "initial deposit on the proposal")
	return fs
}

func FlagAddAssetManagementAccount() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagAssetManagementAccountId, "", "id of the asset management account")
	fs.String(FlagAssetManagementAccountName, "", "name of the asset management account")
	fs.Bool(FlagEnabled, true, "flag if account is enabled or not")
	return fs
}

func FlagStopAssetManagementAccount() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagAssetManagementAccountId, "", "id of the asset management account")
	return fs
}

func FlagAddAssetManagementTarget() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagAssetManagementAccountId, "", "id of the asset management account")
	fs.String(FlagAssetManagementAccountAddress, "", "address of the asset management target")
	fs.String(FlagAssetManagementTargetId, "", "id of the asset management target")
	fs.Uint64(FlagUnbondingSeconds, 0, "unbonding seconds")
	fs.String(FlagAssetConditions, "", "asset conditions string")
	fs.String(FlagIntegrateType, "", "integration type, GOLANG_MOD | COSMWASM")
	fs.String(FlagModName, "", "module name to invest on, stakeibc | yieldfarm")
	return fs
}

func FlagStopAssetManagementTarget() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagAssetManagementAccountId, "", "id of the asset management account")
	fs.String(FlagAssetManagementTargetId, "", "id of the asset management target")

	return fs
}

func FlagDepositCmd() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.Bool(FlagExecuteOrders, false, "required flag when execute orders as part of deposit")

	return fs
}

func FlagFarmingOrder() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagFarmingOrderId, "", "order id")
	fs.String(FlagStrategyType, "", "strategy type")
	fs.String(FlagWhitelistedTargetIds, "", "whitelisted target ids")
	fs.String(FlagBlacklistedTargetIds, "", "blacklisted target ids")
	fs.Uint64(FlagMaxUnbondingSeconds, 0, "maximum unbonding seconds on the order")
	fs.Uint32(FlagOverallRatio, 0, "ratio on fund split")
	fs.String(FlagMin, "", "minimum deposit amount")
	fs.String(FlagMax, "", "maximum deposit amount")
	fs.Uint64(FlagDate, 0, "order start timestamp")
	fs.Bool(FlagActive, false, "shows if order is active or not")

	return fs
}
