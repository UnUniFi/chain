package cli

import (
	"context"
	"fmt"
	"strconv"

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
		Use:   "show-cdp [id] [collateral-type]",
		Short: "shows a cdp",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("cdp-id '%s' not a valid uint", args[0])
			}

			params := &types.QueryGetCdpRequest{
				Id:             id,
				CollateralType: args[2],
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
