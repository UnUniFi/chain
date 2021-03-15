package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/lcnem/jpyx/x/cdp/types"
	"github.com/spf13/cobra"
)

func CmdListCdp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-cdp",
		Short: "list all cdp",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCdpRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CdpAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowCdp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-cdp [id]",
		Short: "shows a cdp",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCdpRequest{
				Id: args[0],
			}

			res, err := queryClient.Cdp(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
