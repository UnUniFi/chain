package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/UnUniFi/chain/x/yieldaggregator/client/cli"
)

var ProposalAddYieldFarmHandler = govclient.NewProposalHandler(cli.NewSubmitProposalAddYieldFarmTxCmd)
var ProposalUpdateYieldFarm = govclient.NewProposalHandler(cli.NewSubmitProposalUpdateYieldFarmTxCmd)
var ProposalStopYieldFarm = govclient.NewProposalHandler(cli.NewSubmitProposalStopYieldFarmTxCmd)
var ProposalRemoveYieldFarm = govclient.NewProposalHandler(cli.NewSubmitProposalRemoveYieldFarmTxCmd)
var ProposalAddYieldFarmTarget = govclient.NewProposalHandler(cli.NewSubmitProposalAddYieldFarmTargetTxCmd)
var ProposalUpdateYieldFarmTarget = govclient.NewProposalHandler(cli.NewSubmitProposalUpdateYieldFarmTargetTxCmd)
var ProposalStopYieldFarmTarget = govclient.NewProposalHandler(cli.NewSubmitProposalStopYieldFarmTargetTxCmd)
var ProposalRemoveYieldFarmTarget = govclient.NewProposalHandler(cli.NewSubmitProposalRemoveYieldFarmTargetTxCmd)
