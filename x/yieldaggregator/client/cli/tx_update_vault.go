package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func CmdTxUpdateVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-vault [id] [name] [description] [fee-collector]",
		Short: "update the vault info",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			name := args[1]
			description := args[2]
			feeCollector := args[3]

			msg := types.NewMsgUpdateVault(clientCtx.GetFromAddress().String(), id, name, description, feeCollector)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
