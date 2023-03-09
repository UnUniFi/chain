package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/UnUniFi/chain/x/yieldaggregatorv1/client/cli"
)

var ProposalAddYieldFarmHandler = govclient.NewProposalHandler(cli.NewSubmitProposalAddYieldFarmTxCmd, nil)
var ProposalUpdateYieldFarm = govclient.NewProposalHandler(cli.NewSubmitProposalUpdateYieldFarmTxCmd, nil)
var ProposalStopYieldFarm = govclient.NewProposalHandler(cli.NewSubmitProposalStopYieldFarmTxCmd, nil)
var ProposalRemoveYieldFarm = govclient.NewProposalHandler(cli.NewSubmitProposalRemoveYieldFarmTxCmd, nil)
var ProposalAddYieldFarmTarget = govclient.NewProposalHandler(cli.NewSubmitProposalAddYieldFarmTargetTxCmd, nil)
var ProposalUpdateYieldFarmTarget = govclient.NewProposalHandler(cli.NewSubmitProposalUpdateYieldFarmTargetTxCmd, nil)
var ProposalStopYieldFarmTarget = govclient.NewProposalHandler(cli.NewSubmitProposalStopYieldFarmTargetTxCmd, nil)
var ProposalRemoveYieldFarmTarget = govclient.NewProposalHandler(cli.NewSubmitProposalRemoveYieldFarmTargetTxCmd, nil)
