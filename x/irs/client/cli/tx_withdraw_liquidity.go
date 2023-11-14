package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func CmdTxWithdrawLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-liquidity [strategy_contract] [share-amount]",
		Short: "withdraw liquidity",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			strategyContract := args[0]

			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("error parsing amount")
			}

			msg := types.NewMsgWithdrawLiquidity(clientCtx.GetFromAddress().String(), strategyContract, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
