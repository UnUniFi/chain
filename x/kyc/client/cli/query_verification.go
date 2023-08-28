package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/kyc/types"
)

func CmdListVerification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-verification",
		Short: "list all verification",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVerificationRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VerificationAll(context.Background(), params)
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

func CmdShowVerification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-verification [address] [provider-id]",
		Short: "shows a verification",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			address := args[0]
			providerId, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryGetVerificationRequest{
				Customer:   address,
				ProviderId: uint64(providerId),
			}

			res, err := queryClient.Verification(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
