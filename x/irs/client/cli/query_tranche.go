package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/irs/types"
)

func CmdListStrategyTranches() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-strategy-tranches [strategy_contract]",
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

func CmdListAllTranches() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-tranches",
		Short: "list all tranches",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTranchesRequest{}

			res, err := queryClient.AllTranches(context.Background(), params)
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

func CmdShowTranchePtAPYs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tranche-pt-apys [id]",
		Short: "shows a tranche's PT APYs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			params := &types.QueryTranchePtAPYsRequest{
				Id: uint64(id),
			}

			res, err := queryClient.TranchePtAPYs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowTrancheYtAPYs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tranche-yt-apys [id]",
		Short: "shows a tranche's YT APYs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			params := &types.QueryTrancheYtAPYsRequest{
				Id: uint64(id),
			}

			res, err := queryClient.TrancheYtAPYs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowTranchePoolAPYs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-tranche-pool-apys [id]",
		Short: "shows a tranche's Pool APYs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			params := &types.QueryTranchePoolAPYsRequest{
				Id: uint64(id),
			}

			res, err := queryClient.TranchePoolAPYs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
