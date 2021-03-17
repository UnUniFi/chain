package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/lcnem/jpyx/x/pricefeed/types"
	"github.com/spf13/cobra"
)

func CmdListOracle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-oracle [market-id]",
		Short: "list all oracle",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOracleRequest{
				MarketId:   args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.OracleAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
