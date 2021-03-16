package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/lcnem/jpyx/x/pricefeed/types"
	"github.com/spf13/cobra"
)

func CmdShowOracle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-oracle [id]",
		Short: "shows a oracle",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetOracleRequest{
				Id: args[0],
			}

			res, err := queryClient.Oracle(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
