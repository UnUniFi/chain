package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/lcnem/jpyx/x/cdp/types"
)

func CmdCreateCdp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-cdp [collateral] [debt] [collateral-type]",
		Short: "Creates a new cdp",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new cdp, depositing some collateral and drawing some debt.
Example:
$ %s tx %s create-cdp 10ubtc 10jpyx ubtc-a --from myKeyName --chain-id jpyx-3-test
`, version.AppName, types.ModuleName)),
		Args:  cobra.ExactArgs(3),
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

			msg := types.NewMsgCreateCdp(clientCtx.GetFromAddress(), collateral, debt, collateralType)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
