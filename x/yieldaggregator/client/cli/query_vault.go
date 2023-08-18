package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func CmdListVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-vault",
		Short: "list all vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVaultRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VaultAll(context.Background(), params)
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

func CmdShowVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-vault [id]",
		Short: "shows a vault",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetVaultRequest{
				Id: id,
			}

			res, err := queryClient.Vault(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdVaultAllByShareHolder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-vault-by-share-holder [holder]",
		Short: "List vaults by share holder",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVaultByShareHolderRequest{
				ShareHolder: args[0],
			}

			res, err := queryClient.VaultAllByShareHolder(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdVaultEstimatedMintAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-mint-amount [id] [deposit-amount]",
		Short: "Estimate mint amount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryEstimateMintAmountRequest{
				Id:            uint64(id),
				DepositAmount: args[1],
			}

			res, err := queryClient.EstimateMintAmount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdVaultEstimatedRedeemAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-redeem-amount [id] [burn-amount]",
		Short: "Estimate redeem amount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryEstimateRedeemAmountRequest{
				Id:         uint64(id),
				BurnAmount: args[1],
			}

			res, err := queryClient.EstimateRedeemAmount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
