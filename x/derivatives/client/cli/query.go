package cli

import (
	"context"
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifiTypes "github.com/UnUniFi/chain/types"
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
	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdQueryClaimableLiquidityProviderRewards(), CmdQueryPositions())

	return cmd
}

func CmdQueryClaimableLiquidityProviderRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claimable-rewards [address]",
		Short: "shows the claimable liquidity provider rewards of the designated address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.ClaimableLiquidityProviderRewards(context.Background(), &types.QueryClaimableLiquidityProviderRewardsRequest{Address: (*ununifiTypes.StringAccAddress)(&address)})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryPositions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "positions [address]",
		Short: "shows the positions owned by the designated address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.Positions(context.Background(), &types.QueryPositionsRequest{Address: (*ununifiTypes.StringAccAddress)(&address)})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
