package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func CmdTxCreateVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vault [symbol] [name] [description] [commission-rate] [withdraw-reserve-rate] [fee] [deposit] [strategy-weights] [fee-collector]",
		Short: "create a new vault",
		Long: `create a new vault
			ununifid tx yieldaggregator create-vault uguu 0.001 1000uguu 1000000uguu ibc/XXXD:1:0.1,ibc/XXXD:2:0.9
		`,
		Args: cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			symbol := args[0]
			name := args[1]
			description := args[2]
			commissionRate, err := sdk.NewDecFromStr(args[3])
			if err != nil {
				return err
			}
			withdrawReserveRate, err := sdk.NewDecFromStr(args[4])
			if err != nil {
				return err
			}
			fee, err := sdk.ParseCoinNormalized(args[5])
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinNormalized(args[6])
			if err != nil {
				return err
			}

			strategyWeightStrs := strings.Split(args[7], ",")

			strategyWeights := make([]types.StrategyWeight, 0)
			for _, strategyWeightStr := range strategyWeightStrs {
				split := strings.Split(strategyWeightStr, ":")
				if len(split) != 3 {
					return fmt.Errorf("invalid strategy weight: %s", strategyWeightStr)
				}
				strategyDenom := split[0]
				strategyId, err := strconv.Atoi(split[1])
				if err != nil {
					return err
				}
				weight, err := sdk.NewDecFromStr(split[2])
				if err != nil {
					return err
				}

				strategyWeights = append(strategyWeights, types.StrategyWeight{
					Denom:      strategyDenom,
					StrategyId: uint64(strategyId),
					Weight:     weight,
				})
			}

			msg := types.NewMsgCreateVault(
				clientCtx.GetFromAddress().String(),
				symbol,
				name,
				description,
				commissionRate,
				withdrawReserveRate,
				strategyWeights,
				fee,
				deposit,
				args[8],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
