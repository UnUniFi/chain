package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/lcnem/jpyx/x/pricefeed/types"
	"github.com/spf13/cobra"
)

func CmdShowRawPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-rawprice [id]",
		Short: "shows a rawprice",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetRawPriceRequest{
				Id: args[0],
			}

			res, err := queryClient.RawPrice(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
