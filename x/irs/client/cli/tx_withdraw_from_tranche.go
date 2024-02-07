package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/irs/types"
)

func CmdTxWithdrawFromTranche() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-from-tranche [trancheId] [trancheType] [tokens] [requiredRedeemAmount]",
		Short: "Withdraw from tranche",
		Long: `Withdraw from tranche
			ununifid tx irs withdraw-from-tranche ununifi1unyuj8qnmygvzuex3dwmg9yzt9alhvyeat0uu0jedg2wj33efl5q5gcjfn FIXED_YIELD	100pt
			ununifid tx irs withdraw-from-tranche ununifi1unyuj8qnmygvzuex3dwmg9yzt9alhvyeat0uu0jedg2wj33efl5q5gcjfn LEVERAGED_VARIABLE_YIELD 100yt
		  ununifid tx irs withdraw-from-tranche ununifi1unyuj8qnmygvzuex3dwmg9yzt9alhvyeat0uu0jedg2wj33efl5q5gcjfn NORMAL_YIELD 100pt,100yt 50
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

			tokens, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			requiredRedeemAmount, ok := sdk.NewIntFromString(args[3])
			if !ok {
				requiredRedeemAmount = sdk.ZeroInt()
			}

			msg := types.NewMsgWithdrawFromTranche(
				clientCtx.GetFromAddress().String(),
				uint64(trancheId),
				types.TrancheType(types.TrancheType_value[trancheTypeStr]),
				tokens,
				requiredRedeemAmount,
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
