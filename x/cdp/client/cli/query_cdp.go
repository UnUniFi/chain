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
		Use:   "show-cdp [owner] [collateral-type]",
		Short: "shows a cdp",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCdpRequest{
				Owner:          args[0],
				CollateralType: args[1],
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

func CmdListAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-account",
		Short: "list all account",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAccountRequest{}

			res, err := queryClient.AccountAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-deposit [owner] [collateral-type]",
		Short: "list all deposit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDepositRequest{
				Owner:          args[0],
				CollateralType: args[1],
			}

			res, err := queryClient.DepositAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
