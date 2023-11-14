package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/irs/types"
)

func CmdTxWithdrawFromTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-from-tranche [strategy_contract] [tranche_type] [maturity]",
		Short: "Withdraw from tranche",
		Long: `Withdraw from tranche
			ununifid tx irs withdraw-from-tranche ununifi1unyuj8qnmygvzuex3dwmg9yzt9alhvyeat0uu0jedg2wj33efl5q5gcjfn FIXED_YIELD	1699977229
			ununifid tx irs withdraw-from-tranche ununifi1unyuj8qnmygvzuex3dwmg9yzt9alhvyeat0uu0jedg2wj33efl5q5gcjfn LEVERAGED_VARIABLE_YIELD 1699977229
		`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			strategyContract := args[0]
			trancheTypeStr := args[1]
			trancheMaturity, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositToTranche(
				clientCtx.GetFromAddress().String(),
				strategyContract,
				types.TrancheType(types.TrancheType_value[trancheTypeStr]),
				uint64(trancheMaturity),
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
