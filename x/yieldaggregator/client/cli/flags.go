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
	return fs
}

func FlagStopAssetManagementTarget() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagAssetManagementAccountId, "", "id of the asset management account")
	fs.String(FlagAssetManagementTargetId, "", "id of the asset management target")

	return fs
}
