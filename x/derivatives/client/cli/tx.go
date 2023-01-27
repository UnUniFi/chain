package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	ununifiType "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/derivatives/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		CmdMintLiquidityProviderToken(),
		CmdBurnLiquidityProviderToken(),
		CmdOpenPosition(),
		CmdClosePosition(),
		CmdReportLiquidationNeededPosition(),
	)

	return cmd
}

func CmdMintLiquidityProviderToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-lpt [amount]",
		Short: "mint liquidity provider token",
		Long: strings.TrimSpace(
			fmt.Sprintf(`mint liquidity provider token.
Example:
$ %s tx %s mint-lpt --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.MsgMintLiquidityProviderToken{
				Sender: ununifiType.StringAccAddress(sender),
				Amount: amount,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdBurnLiquidityProviderToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-lpt [amount]",
		Short: "burn liquidity provider token",
		Long: strings.TrimSpace(
			fmt.Sprintf(`burn liquidity provider token.
Example:
$ %s tx %s burn-lpt --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()
			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("invalid amount")
			}

			msg := types.MsgBurnLiquidityProviderToken{
				Sender: ununifiType.StringAccAddress(sender),
				Amount: amount,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdOpenPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "open-position",
		Short:                      fmt.Sprintf("%s open position subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		CmdOpenPerpetualFuturesPosition(),
		CmdOpenPerpetualOptionsPosition(),
	)

	return cmd
}

func CmdOpenPerpetualFuturesPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perpetual-futures",
		Short: "open perpetual futures position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`open perpetual futures position.
Example:
$ %s tx %s open-position perpetual-futures --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			msg := types.MsgOpenPosition{
				Sender: ununifiType.StringAccAddress(sender),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdOpenPerpetualOptionsPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perpetual-options",
		Short: "open perpetual options position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`open perpetual options position.
Example:
$ %s tx %s open-position perpetual-options --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			msg := types.MsgOpenPosition{
				Sender: ununifiType.StringAccAddress(sender),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdClosePosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-position",
		Short: "report liquidation needed position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`close position.
Example:
$ %s tx %s close-position --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			msg := types.MsgClosePosition{
				Sender: ununifiType.StringAccAddress(sender),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdReportLiquidationNeededPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-liquidation",
		Short: "report liquidation needed position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`report liquidation needed position.
Example:
$ %s tx %s report-liquidation --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			msg := types.MsgReportLiquidationNeededPosition{
				Sender: ununifiType.StringAccAddress(sender),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
