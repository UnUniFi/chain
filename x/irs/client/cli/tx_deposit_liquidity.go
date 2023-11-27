package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func CmdTxDepositLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-liquidity [tranche-id] [share-out-amount] [token-in-maxs]",
		Short: "deposit liquidity",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			trancheId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			shareOutAmount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("error parsing amount")
			}

			tokenInMaxs, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositLiquidity(clientCtx.GetFromAddress().String(), uint64(trancheId), shareOutAmount, tokenInMaxs)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
