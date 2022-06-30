package cli

import (
	// "context"
	"context"
	"fmt"

	"github.com/UnUniFi/chain/x/nftmint/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {

	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryClassAttributes(),
		CmdQueryClassIdsByOwner(),
	)

	return cmd
}

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows params",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryParamsRequest{}

			res, err := queryClient.Params(context.Background(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryClassAttributes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-owner",
		Args:  cobra.ExactArgs(1),
		Short: "Query the owner of class",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ClassAttributes(
				context.Background(),
				&types.QueryClassAttributesRequest{ClassId: args[0]},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res.ClassAttributes)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryClassIdsByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-ids-by-owner",
		Args:  cobra.ExactArgs(1),
		Short: "Query classIDs owned by the owner address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ClassIdsByOwner(
				context.Background(),
				&types.QueryClassIdsByOwnerRequest{Owner: args[0]},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res.OwningClassIdList)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
