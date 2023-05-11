package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/UnUniFi/chain/x/deprecated/incentive/types"
)

func CmdClaimCdpMintingReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-cdp-minting-reward [multiplierName]",
		Short: "Claims cdp minting reward",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			multiplierName := args[0]

			msg := types.NewMsgClaimCdpMintingReward(clientCtx.GetFromAddress(), multiplierName)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
