package cli

import (
	"context"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-vault",
		Short: "list all vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVaultRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VaultAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-vault [denom]",
		Short: "shows a vault",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			denom := args[0]

			params := &types.QueryGetVaultRequest{
				Denom: denom,
			}

			res, err := queryClient.Vault(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
