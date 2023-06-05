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

	"github.com/UnUniFi/chain/deprecated/cdp/types"
)

func CmdCreateCdp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-cdp [collateral] [debt] [collateral-type]",
		Short: "Creates a new cdp",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new cdp, depositing some collateral and drawing some debt.
Example:
$ %s tx %s create-cdp 10ubtc 10jpu ubtc-a --from myKeyName --chain-id ununifi-5-test
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
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

func CmdDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [owner] [collateral] [collateral-type]",
		Short: "Deposit",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Deposit.
Example:
$ %s tx %s deposit [owner-address] 10ubtc ubtc-a --from myKeyName --chain-id ununifi-5-test
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			collateral, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			collateralType := args[2]

			msg := types.NewMsgDeposit(owner, clientCtx.GetFromAddress(), collateral, collateralType)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [owner] [collateral] [collateral-type]",
		Short: "Withdraw",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Withdraw.
Example:
$ %s tx %s withdraw [owner-address] 10ubtc ubtc-a --from myKeyName --chain-id ununifi-5-test
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			collateral, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			collateralType := args[2]

			msg := types.NewMsgWithdraw(owner, clientCtx.GetFromAddress(), collateral, collateralType)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDrawDebt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "draw-debt [collateral-type] [principal]",
		Short: "Draw debt",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Draw debt.
Example:
$ %s tx %s draw-debt ubtc-a [principal] --from myKeyName --chain-id ununifi-5-test
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			collateralType := args[0]

			principal, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDrawDebt(clientCtx.GetFromAddress(), collateralType, principal)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRepayDebt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repay-debt [collateral-type] [payment]",
		Short: "Repay debt",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Repay debt.
Example:
$ %s tx %s repay-debt ubtc-a [payment] --from myKeyName --chain-id ununifi-5-test
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			collateralType := args[0]

			payment, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgRepayDebt(clientCtx.GetFromAddress(), collateralType, payment)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdLiquidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidate [borrower] [collateral-type]",
		Short: "Liquidate",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Liquidate.
Example:
$ %s tx %s liquidate [borrower] ubtc-a --from myKeyName --chain-id ununifi-5-test
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			borrower, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			collateralType := args[1]

			msg := types.NewMsgLiquidate(clientCtx.GetFromAddress(), borrower, collateralType)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
