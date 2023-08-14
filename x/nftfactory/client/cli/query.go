package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/UnUniFi/chain/x/nftfactory/types"

	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group tokenfactory queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetParams(),
		GetCmdClassAuthorityMetadata(),
		GetCmdDenomsFromCreator(),
	)

	return cmd
}

// GetParams returns the params for the module
func GetParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params [flags]",
		Short: "Get the params for the x/nftfactory module",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdClassAuthorityMetadata returns the authority metadata for a queried denom
func GetCmdClassAuthorityMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-authority-metadata [class-id] [flags]",
		Short: "Get the authority metadata for a specific class id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			denom := strings.Split(args[0], "/")

			if len(denom) != 3 {
				return fmt.Errorf("invalid denom format, expected format: factory/[creator]/[subdenom]")
			}

			res, err := queryClient.ClassAuthorityMetadata(cmd.Context(), &types.QueryClassAuthorityMetadataRequest{
				Creator:  denom[1],
				Subclass: denom[2],
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

// GetCmdDenomsFromCreator a command to get a list of all tokens created by a specific creator address
func GetCmdDenomsFromCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "classes-from-creator [creator address] [flags]",
		Short: "Returns a list of all tokens created by a specific creator address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ClassesFromCreator(cmd.Context(), &types.QueryClassesFromCreatorRequest{
				Creator: args[0],
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
