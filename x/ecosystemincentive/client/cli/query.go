package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group ecosystemincentive queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "ecosystem-incentive",
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryRecordedIncentiveUnitId(),
		CmdQueryAllRewards(),
		CmdQueryIncentiveUnit(),
		CmdQueryIncentiveUnitIdsByAddr(),
	)

	return cmd
}

func CmdQueryRecordedIncentiveUnitId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recorded-incentive-unit-id [class-id] [nft-id]",
		Short: "shows incentive-unit-id recorded with the class and nft ids",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryRecordedIncentiveUnitIdRequest{
				ClassId: args[0],
				NftId:   args[1],
			}

			res, err := queryClient.RecordedIncentiveUnitId(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryAllRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-rewards [address]",
		Short: "shows all rewards that defined address have",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAllRewardsRequest{
				SubjectAddr: args[0],
			}

			res, err := queryClient.AllRewards(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryIncentiveUnit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "incentive-unit [incentive-unit-id]",
		Short: "shows incentive-unit data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryIncentiveUnitRequest{
				IncentiveUnitId: args[0],
			}

			res, err := queryClient.IncentiveUnit(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryIncentiveUnitIdsByAddr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "incentive-unit-ids-by-addr [address]",
		Short: "shows incentive-unit-ids to which the address belongs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryIncentiveUnitIdsByAddrRequest{
				Address: args[0],
			}

			res, err := queryClient.IncentiveUnitIdsByAddr(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
