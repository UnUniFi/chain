package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdTxDepositToVault(),
		CmdTxWithdrawFromVault(),
		CmdTxWithdrawFromVaultWithUnbondingTime(),
		CmdTxCreateVault(),
		CmdTxTransferVaultOwnership(),
		CmdTxDeleteVault(),
	)

	return cmd
}
