package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group decentralizedvault queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryTransferRequestedNft(),
		CmdQueryTransferRequestedNfts(),
	)

	return cmd
}

func CmdQueryTransferRequestedNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer_request [nft_id]",
		Short: "shows transfer request nft",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryTransferRequestedNftRequest{
				NftId: args[0],
			}

			res, err := queryClient.TransferRequestedNft(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryTransferRequestedNfts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer_requests",
		Short: "shows all transfer request nft",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			//todo add limit
			queryClient := types.NewQueryClient(clientCtx)

			limit, err := cmd.Flags().GetInt32(FlagNftTransferRequestLimit)
			if err != nil {
				return err
			}

			params := &types.QueryTransferRequestedNftsRequest{
				NftLimit: limit,
			}

			res, err := queryClient.TransferRequestedNfts(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().Int32(FlagNftTransferRequestLimit, 1, "nft transfer request limit")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
