package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "ecosystem-incentive",
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdWithdrawAllRewards(),
		CmdWithdrawReward(),
	)
	return cmd
}

func BuildRegisterInputs(fs *pflag.FlagSet) (string, []string, []sdk.Dec, error) {
	registerInputs, err := parseRegisterFlags(fs)
	if err != nil {
		return "", nil, nil, err
	}
	incentiveId := registerInputs.IncentiveUnitId

	var subjectAddrs []string
	for _, addr := range registerInputs.SubjectAddrs {
		subjectAddrs = append(subjectAddrs, addr)
	}

	var weights []sdk.Dec
	for _, weight := range registerInputs.Weights {
		weightDec, err := sdk.NewDecFromStr(weight)
		if err != nil {
			return "", nil, nil, err
		}
		weights = append(weights, weightDec)
	}

	return incentiveId, subjectAddrs, weights, nil
}

func CmdWithdrawAllRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-all-rewards [flags]",
		Args:  cobra.ExactArgs(0),
		Short: "withdraw all accumulated rewards in ecosystem-incentive for tx-sender",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawAllRewards(
				clientCtx.GetFromAddress().String(),
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

func CmdWithdrawReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-reward [denom] [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "withdraw accumulated reward for the specific denom in ecosystem-incentive for tx-sender",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawReward(
				clientCtx.GetFromAddress().String(),
				args[0],
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
