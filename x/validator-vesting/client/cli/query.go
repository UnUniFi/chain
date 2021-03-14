package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/lcnem/jpyx/x/validator-vesting/types"
)

// GetQueryCmd returns the cli query commands for the jsmndist module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	valVestingQueryCmd := &cobra.Command{
		Use:   types.QueryPath,
		Short: "Querying commands for the validator vesting module",
	}

	valVestingQueryCmd.AddCommand(flags.GetCommands(
		queryCirculatingSupply(queryRoute, cdc),
		queryTotalSupply(queryRoute, cdc),
		queryCirculatingSupplyHARD(queryRoute, cdc),
		queryCirculatingSupplyJPYX(queryRoute, cdc),
		queryTotalSupplyHARD(queryRoute, cdc),
		queryTotalSupplyJPYX(queryRoute, cdc),
	)...)

	return valVestingQueryCmd

}

func queryCirculatingSupply(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "circulating-supply",
		Short: "Get circulating supply",
		Long:  "Get the current circulating supply of kava tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryCirculatingSupply), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cdc.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

func queryTotalSupply(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "total-supply",
		Short: "Get total supply",
		Long:  "Get the current total supply of kava tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryTotalSupply), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cdc.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

func queryCirculatingSupplyHARD(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "circulating-supply-hard",
		Short: "Get HARD circulating supply",
		Long:  "Get the current circulating supply of HARD tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryCirculatingSupplyHARD), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cdc.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

func queryCirculatingSupplyJPYX(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "circulating-supply-jpyx",
		Short: "Get JPYX circulating supply",
		Long:  "Get the current circulating supply of JPYX tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryCirculatingSupplyJPYX), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cdc.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

func queryTotalSupplyHARD(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "total-supply-hard",
		Short: "Get HARD total supply",
		Long:  "Get the current total supply of HARD tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryTotalSupplyHARD), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cdc.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintOutput(out)
		},
	}
}

func queryTotalSupplyJPYX(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "total-supply-jpyx",
		Short: "Get JPYX total supply",
		Long:  "Get the current total supply of JPYX tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryTotalSupplyJPYX), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cdc.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintOutput(out)
		},
	}
}
