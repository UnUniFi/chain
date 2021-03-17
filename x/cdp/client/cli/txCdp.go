package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/cdp/types"
)

func CmdCreateCdp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-cdp [collateral] [debt] [collateral-type]",
		Short: "Creates a new cdp",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			collateral, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			debt, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			collateralType := args[2]

			msg := types.NewMsgCreateCDP(clientCtx.GetFromAddress(), collateral, debt, collateralType)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
