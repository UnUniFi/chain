package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/x/jsmndist/types"
)

// GetQueryCmd returns the cli query commands for the jsmndist module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	jsmndistQueryCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Querying commands for the jsmndist module",
	}

	jsmndistQueryCmd.AddCommand(flags.GetCommands(
		queryParamsCmd(queryRoute, cdc),
		queryBalanceCmd(queryRoute, cdc),
	)...)

	return jsmndistQueryCmd

}

func queryParamsCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "get the jsmndist module parameters",
		Long:  "Get the current global jsmndist module parameters.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query
			route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryGetParams)
			res, height, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var params types.Params
			if err := cdc.UnmarshalJSON(res, &params); err != nil {
				return fmt.Errorf("failed to unmarshal params: %w", err)
			}
			return cliCtx.PrintOutput(params)
		},
	}
}

func queryBalanceCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "balance",
		Short: "get the jsmndist module balance",
		Long:  "Get the current jsmndist module account balance.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryGetBalance)
			res, height, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			var coins sdk.Coins
			if err := cdc.UnmarshalJSON(res, &coins); err != nil {
				return fmt.Errorf("failed to unmarshal coin balance: %w", err)
			}
			return cliCtx.PrintOutput(coins)
		},
	}
}
