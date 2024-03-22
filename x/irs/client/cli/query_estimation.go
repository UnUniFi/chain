package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/irs/types"
)

func CmdEstimateSwapInPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-swap-in-pool [id] [amount]",
		Short: "estimate token swap result in the pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			token, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryEstimateSwapInPoolRequest{
				Id:     uint64(id),
				Amount: token.Amount.String(),
				Denom:  token.Denom,
			}

			res, err := queryClient.EstimateSwapInPool(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEstimateMintPtYtPair() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-mint-pt-yt-pair [id] [amount]",
		Short: "estimate mint PT & YT result by depositing token",
		Long: `Example:
    ununifid query irs estimate-mint-pt-yt-pair 1 1000000uguu
    `,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			token, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryEstimateMintPtYtPairRequest{
				Id:     uint64(id),
				Amount: token.Amount.String(),
				Denom:  token.Denom,
			}

			res, err := queryClient.EstimateMintPtYtPair(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEstimateRedeemPtYtPair() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-redeem-pt-yt-pair [id] [amount]",
		Short: "estimate redeem pt-yt-pair result by each input amount",
		Long: `Example:
		ununifid query irs estimate-redeem-pt-yt-pair 1 1000000irs/tranche/1/pt
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			token, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryEstimateRedeemPtYtPairRequest{
				Id:     uint64(id),
				Denom:  token.Denom,
				Amount: token.Amount.String(),
			}

			res, err := queryClient.EstimateRedeemPtYtPair(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEstimateSwapToYt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-swap-to-yt [id] [amount]",
		Short: "estimate swap to YT result by underlying token amount",
		Long: `Example:
		ununifid query irs estimate-swap-to-yt 1 1000000uguu
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			token, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryEstimateSwapToYtRequest{
				Id:     uint64(id),
				Denom:  token.Denom,
				Amount: token.Amount.String(),
			}

			res, err := queryClient.EstimateSwapToYt(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEstimateRequiredDepositSwapToYt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-required-deposit-swap-to-yt [id] [desired-yt-amount]",
		Short: "estimate require Deposit-Token to swap to YT by desired YT amount",
		Long: `Example:
		ununifid query irs estimate-required-deposit-swap-to-yt 1 1000000
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			desiredYt, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("error parsing amount")
			}

			params := &types.QueryEstimateRequiredDepositSwapToYtRequest{
				Id:              uint64(id),
				DesiredYtAmount: desiredYt.String(),
			}

			res, err := queryClient.EstimateRequiredDepositSwapToYt(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEstimateRedeemMaturedYt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-redeem-matured-yt [id] [yt-amount]",
		Short: "estimate redeem matured YT to deposited token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			desiredYt, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("error parsing amount")
			}

			params := &types.QueryEstimateRedeemMaturedYtRequest{
				Id:       uint64(id),
				YtAmount: desiredYt.String(),
			}

			res, err := queryClient.EstimateRedeemMaturedYt(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEstimateMintLiquidityPoolToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-mint-liquidity-pool-token [id] [lp-amount]",
		Short: "estimate mint liquidity pool token result by each input amount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			token, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryEstimateMintLiquidityPoolTokenRequest{
				Id:     uint64(id),
				Denom:  token.Denom,
				Amount: token.Amount.String(),
			}

			res, err := queryClient.EstimateMintLiquidityPoolToken(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEstimateRedeemLiquidityPoolToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-redeem-liquidity-pool-token [id] [pool-token-amount]",
		Short: "estimate redeem liquidity pool token by  pool-token amount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("error parsing amount")
			}

			params := &types.QueryEstimateRedeemLiquidityPoolTokenRequest{
				Id:     uint64(id),
				Amount: amount.String(),
			}

			res, err := queryClient.EstimateRedeemLiquidityPoolToken(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
