package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/irs/types"
)

func CmdListTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-tranches [strategy_contract]",
		Short: "list all tranches by strategy contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTranchesRequest{
				StrategyContract: args[0],
			}

			res, err := queryClient.Tranches(context.Background(), params)
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

func CmdShowTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tranche [id]",
		Short: "shows a tranche",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			params := &types.QueryTrancheRequest{
				Id: uint64(id),
			}

			res, err := queryClient.Tranche(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
