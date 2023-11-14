package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/irs/types"
)

func CmdQueryDenomInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "denom-infos",
		Short: "list denom infos",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDenomInfosRequest{}

			res, err := queryClient.DenomInfos(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
