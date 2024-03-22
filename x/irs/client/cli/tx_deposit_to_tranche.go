package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func CmdTxDepositToTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-to-tranche [trancheId] [trancheType] [spendingToken] [requiredYt]",
		Short: "Deposit to tranche",
		Long: `Deposit to tranche
			ununifid tx irs deposit-to-tranche ununifi1unyuj8qnmygvzuex3dwmg9yzt9alhvyeat0uu0jedg2wj33efl5q5gcjfn FIXED_YIELD	1699977229
			ununifid tx irs deposit-to-tranche ununifi1unyuj8qnmygvzuex3dwmg9yzt9alhvyeat0uu0jedg2wj33efl5q5gcjfn LEVERAGED_VARIABLE_YIELD 1699977229
		`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			trancheId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			trancheTypeStr := args[1]
			token, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			requiredYt, ok := sdk.NewIntFromString(args[3])
			if !ok {
				requiredYt = sdk.ZeroInt()
			}

			msg := types.NewMsgDepositToTranche(
				clientCtx.GetFromAddress().String(),
				uint64(trancheId),
				types.TrancheType(types.TrancheType_value[trancheTypeStr]),
				token,
				requiredYt,
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
