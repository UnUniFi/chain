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
		CmdQueryRewards(),
		CmdQueryRecipientAddressWithNftId(),
	)
	return cmd
}

func CmdQueryRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rewards [address] [denom]",
		Short: "shows ecosystem reward by address & denom",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			var denom string
			if len(args) > 1 {
				denom = args[1]
			} else {
				denom = ""
			}

			req := &types.QueryEcosystemRewardsRequest{
				Address: args[0],
				Denom:   denom,
			}

			res, err := queryClient.EcosystemRewards(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryRecipientAddressWithNftId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recipient-with-nft [class_id] [token_id]",
		Short: "shows recipient address by nft id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryRecipientAddressWithNftIdRequest{
				ClassId: args[0],
				TokenId: args[1],
			}

			res, err := queryClient.RecipientAddressWithNftId(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
