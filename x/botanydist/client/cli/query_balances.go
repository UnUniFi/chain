package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/lcnem/jpyx/x/botanydist/types"
	"github.com/spf13/cobra"
)

func CmdListReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balances",
		Short: "show botanydist module account balances",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetBalancesRequest{}

			res, err := queryClient.Balances(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
