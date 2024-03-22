package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group yieldaggregator queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdListVault(),
		CmdListStrategy(),
		CmdShowStrategy(),
		CmdShowVault(),
		CmdVaultAllByShareHolder(),
		CmdVaultEstimatedMintAmount(),
		CmdVaultEstimatedRedeemAmount(),
		CmdQuerySymbolInfo(),
		CmdQueryDenomInfo(),
		CmdQueryIntermediaryAccounts(),
		CmdVaultAddress(),
	)

	return cmd
}
