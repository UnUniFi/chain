package cli

import (
	"context"

	"github.com/UnUniFi/chain/x/pricefeed/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-price [id]",
		Short: "shows a price",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetPriceRequest{
				MarketId: args[0],
			}

			res, err := queryClient.Price(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
