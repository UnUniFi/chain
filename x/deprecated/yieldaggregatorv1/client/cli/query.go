package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/yieldaggregatorv1/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group yieldaggregator queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryAssetManagementAccount(),
		CmdQueryAllAssetManagementAccounts(),
		CmdQueryUserInfo(),
		CmdQueryAllFarmingUnits(),
		CmdQueryDailyRewardPercents(),
	)

	return cmd
}

func CmdQueryAssetManagementAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "asset-management-account [id]",
		Short: "queries asset management account details by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AssetManagementAccount(context.Background(), &types.QueryAssetManagementAccountRequest{
				Id: args[0],
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

func CmdQueryAllAssetManagementAccounts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-asset-management-accounts",
		Short: "queries asset management account details by id",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllAssetManagementAccounts(context.Background(), &types.QueryAllAssetManagementAccountsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryUserInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-info [user]",
		Short: "query user information by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.UserInfo(context.Background(), &types.QueryUserInfoRequest{
				Address: args[0],
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

func CmdQueryAllFarmingUnits() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-farming-units",
		Short: "query all farming units",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllFarmingUnits(context.Background(), &types.QueryAllFarmingUnitsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryDailyRewardPercents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "daily-reward-percents",
		Short: "query all daily reward percents",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.DailyRewardPercents(context.Background(), &types.QueryDailyRewardPercentsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
