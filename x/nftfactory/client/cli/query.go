package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/UnUniFi/chain/x/nftfactory/types"

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
		CmdQueryClassIdsByName(),
		CmdQueryNFTMinter(),
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
		Use:   "class-attributes [class-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the class attributes by class-id",
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
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryClassIdsByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-ids-by-owner [owner-address]",
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
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryClassIdsByName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-ids-by-name [class-name]",
		Args:  cobra.ExactArgs(1),
		Short: "Query classIDs which have the class name",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ClassIdsByName(
				context.Background(),
				&types.QueryClassIdsByNameRequest{ClassName: args[0]},
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

func CmdQueryNFTMinter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft-minter [class-id] [nft-id]",
		Args:  cobra.ExactArgs(2),
		Short: "Query nft minter with class and nft id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NFTMinter(
				context.Background(),
				&types.QueryNFTMinterRequest{
					ClassId: args[0],
					NftId:   args[1],
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
