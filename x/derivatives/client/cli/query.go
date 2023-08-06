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
	cmd.AddCommand(CmdQueryPool())
	cmd.AddCommand(CmdQueryLiquidityProviderTokenRealAPY())
	cmd.AddCommand(CmdQueryLiquidityProviderTokenNominalAPY())
	cmd.AddCommand(CmdQueryPerpetualFutures())
	cmd.AddCommand(CmdQueryPerpetualFuturesMarket())
	cmd.AddCommand(CmdQueryPerpetualOptions())
	cmd.AddCommand(CmdQueryAllPositions())
	cmd.AddCommand(CmdQueryPosition())
	cmd.AddCommand(CmdQueryPerpetualFuturesPositionSize())
	cmd.AddCommand(CmdQueryAddressPositions())
	cmd.AddCommand(CmdQueryDLPTokenRate())
	cmd.AddCommand(CmdQueryEstimateDLPTokenAmount())
	cmd.AddCommand(CmdQueryEstimateRedeemTokenAmount())
	cmd.AddCommand(CmdQueryAvailableAssetInPoolByDenom())
	cmd.AddCommand(CmdQueryAvailableAssetsInPool())
	cmd.AddCommand(CmdQueryAllPendingPaymentPositions())
	cmd.AddCommand(CmdQueryPendingPaymentPosition())

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

func CmdQueryPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool",
		Short: "shows the pool information",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Pool(context.Background(), &types.QueryPoolRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryPerpetualFutures() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perpetual-futures",
		Short: "shows the perpetual futures information",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.PerpetualFutures(context.Background(), &types.QueryPerpetualFuturesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryPerpetualFuturesMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perpetual-futures-market [base-denom] [quote-denom]",
		Short: "shows the perpetual futures market information",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			_, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.PerpetualFuturesMarket(
				context.Background(),
				&types.QueryPerpetualFuturesMarketRequest{
					BaseDenom:  args[0],
					QuoteDenom: args[1],
				})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryPerpetualOptions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perpetual-options",
		Short: "shows the perpetual options information",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.PerpetualOptions(context.Background(), &types.QueryPerpetualOptionsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryAllPositions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-positions",
		Short: "shows all positions",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.AllPositions(context.Background(), &types.QueryAllPositionsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "position [position-id]",
		Short: "shows position of the specified position id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Position(context.Background(), &types.QueryPositionRequest{PositionId: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryAllPendingPaymentPositions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending-payment-positions",
		Short: "shows all pending payment positions",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.AllPendingPaymentPositions(context.Background(), &types.QueryAllPendingPaymentPositionsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryPendingPaymentPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending-payment-position [position-id]",
		Short: "shows pending payment position of the specified position id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PendingPaymentPosition(context.Background(), &types.QueryPendingPaymentPositionRequest{PositionId: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryPerpetualFuturesPositionSize() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perpetual-futures-position-size [position-type] [address]",
		Short: "shows the perpetual futures position size of the specified address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var positionType types.PositionType
			switch args[3] {
			case "long":
				positionType = types.PositionType_LONG
			case "short":
				positionType = types.PositionType_SHORT
			default:
				return fmt.Errorf("invalid position type")
			}

			address, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PerpetualFuturesPositionSize(
				context.Background(),
				&types.QueryPerpetualFuturesPositionSizeRequest{
					PositionType: positionType,
					Address:      address.String(),
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

func CmdQueryDLPTokenRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delp-token-rate",
		Short: "shows the delp token rate",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.DLPTokenRates(
				context.Background(),
				&types.QueryDLPTokenRateRequest{},
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

func CmdQueryEstimateDLPTokenAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "estimate-delp-token-amount [mint-denom] [amount]",
		Long: "shows the estimated delp token amount for the specified amount of the asset",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.EstimateDLPTokenAmount(
				context.Background(),
				&types.QueryEstimateDLPTokenAmountRequest{
					MintDenom: args[0],
					Amount:    args[1],
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

func CmdQueryEstimateRedeemTokenAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "estimate-redeem-token-amount [redeem-denom] [amount]",
		Long: "shows the estimated redeem token amount for the specified amount of the asset",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.EstimateRedeemTokenAmount(
				context.Background(),
				&types.QueryEstimateRedeemTokenAmountRequest{
					RedeemDenom: args[0],
					LptAmount:   args[1],
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
