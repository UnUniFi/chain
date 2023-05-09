package cli

import (
	"context"
	"fmt"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group derivatives queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdQueryLiquidityProviderTokenRealAPY(), CmdQueryLiquidityProviderTokenNominalAPY(), CmdQueryAddressPositions())
	cmd.AddCommand(CmdQueryAvailableAssetInPoolByDenom())
	cmd.AddCommand(CmdQueryAvailableAssetsInPool())

	return cmd
}

func CmdQueryLiquidityProviderTokenRealAPY() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lpt-real-apy [beforeHeight] [afterHeight]",
		Short: "shows the real Annual Percent Yield between beforeHeight and afterHeight",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.LiquidityProviderTokenRealAPY(context.Background(), &types.QueryLiquidityProviderTokenRealAPYRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryLiquidityProviderTokenNominalAPY() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lpt-nominal-apy [beforeHeight] [afterHeight]",
		Short: "shows the nominal Annual Percent Yield between beforeHeight and afterHeight",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.LiquidityProviderTokenNominalAPY(context.Background(), &types.QueryLiquidityProviderTokenNominalAPYRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryAddressPositions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "positions [address]",
		Short: "shows the positions owned by the designated address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			_, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.AddressPositions(context.Background(), &types.QueryAddressPositionsRequest{Address: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryAvailableAssetInPoolByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "available-asset [denom]",
		Short: "shows the available amount of the asset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AvailableAssetInPoolByDenom(
				context.Background(),
				&types.QueryAvailableAssetInPoolByDenomRequest{
					Denom: args[0],
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryAvailableAssetsInPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "available-assets",
		Short: "shows the available amount of all assets in the pool",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AvailableAssetsInPool(
				context.Background(),
				&types.QueryAvailableAssetsInPoolRequest{},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd

}
