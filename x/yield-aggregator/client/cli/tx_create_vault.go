package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CmdTxCreateVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vault [denom] [commission-rate] [fee] [deposit] [[strategy-id] [strategy-weight] ...]",
		Short: "create a new vault",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			denom := args[0]
			commissionRate, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}
			fee, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}
			strategyWeights := make([]types.StrategyWeight, 0)
			// TODO: append strategyWeights

			msg := types.NewMsgCreateVault(clientCtx.GetFromAddress().String(), denom, commissionRate, strategyWeights, fee, deposit)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
