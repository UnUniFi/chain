package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	// "github.com/cosmos/cosmos-sdk/client/flags"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
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
		CmdRegister(),
	)
	return cmd
}

func CmdRegister() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register [flags]",
		Args:  cobra.ExactArgs(0),
		Short: "register incentive-unit to get ecosystem-incentive reward",
		Long:  "Example command: $ %s tx %s register --register-file [json-file-path]",
		Example: `Example of a json file to pass:
{
	"incentive-id": "incentive-unit1",
	"subject-addrs": [
		"ununifi17gs6kgph4657epky2ctl9sf66ucyua939nexgl",
		"ununifi1w9s3wpkh0kfk0t40m4lwjsx6h2v6gktsvfrgux"
	],
	"weights": [
		"0.50",
		"0.50"
	]
}
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			incentiveId, subjectAddrs, weights, err := BuildRegisterInputs(cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgRegister(
				clientCtx.GetFromAddress(),
				incentiveId,
				subjectAddrs,
				weights,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetRegister())
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func BuildRegisterInputs(fs *pflag.FlagSet) (string, []sdk.AccAddress, []sdk.Dec, error) {
	registerInputs, err := parseRegisterFlags(fs)
	if err != nil {
		return "", nil, nil, err
	}
	incentiveId := registerInputs.IncentiveId

	var subjectAddrs []sdk.AccAddress
	for _, addr := range registerInputs.SubjectAddrs {
		accAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return "", nil, nil, err
		}
		subjectAddrs = append(subjectAddrs, accAddr)
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
