package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group ecosystemincentive queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "ecosystem-incentive",
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryAllRewards(),
	)
	return cmd
}

func CmdQueryAllRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-rewards [address]",
		Short: "shows all rewards that defined address have",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAllRewardsRequest{
				Address: args[0],
			}

			res, err := queryClient.AllRewards(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
