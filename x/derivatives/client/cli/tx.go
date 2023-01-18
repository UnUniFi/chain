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
		CmdClaimLiquidityProviderRewards(),
		CmdOpenPerpetualFuturesPosition(),
		CmdClosePerpetualFuturesPosition(),
		CmdOpenPerpetualOptionsPosition(),
		CmdClosePerpetualOptionsPosition(),
	)

	return cmd
}

func CmdMintLiquidityProviderToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-lpt",
		Short: "mint liquidity provider token",
		Long: strings.TrimSpace(
			fmt.Sprintf(`mint liquidity provider token.
Example:
$ %s tx %s mint-lpt --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			msg := types.MsgMintLiquidityProviderToken{
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

func CmdBurnLiquidityProviderToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-lpt",
		Short: "burn liquidity provider token",
		Long: strings.TrimSpace(
			fmt.Sprintf(`burn liquidity provider token.
Example:
$ %s tx %s burn-lpt --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			msg := types.MsgBurnLiquidityProviderToken{
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

func CmdClaimLiquidityProviderRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-lp-rewards",
		Short: "claim liquidity provider rewards",
		Long: strings.TrimSpace(
			fmt.Sprintf(`claim liquidity provider rewards.
Example:
$ %s tx %s claim-lp-rewards --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			msg := types.MsgClaim{
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

func CmdOpenPerpetualFuturesPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open-position-perpetual-futures",
		Short: "open perpetual futures position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`open perpetual futures position.
Example:
$ %s tx %s open-position-perpetual-futures --from myKeyName --chain-id ununifi-x
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

func CmdClosePerpetualFuturesPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-position-perpetual-futures",
		Short: "close perpetual futures position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`close perpetual futures position.
Example:
$ %s tx %s close-position-perpetual-futures --from myKeyName --chain-id ununifi-x
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

func CmdOpenPerpetualOptionsPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open-position-perpetual-options",
		Short: "open perpetual options position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`open perpetual options position.
Example:
$ %s tx %s open-position-perpetual-options --from myKeyName --chain-id ununifi-x
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

func CmdClosePerpetualOptionsPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-position-perpetual-options",
		Short: "close perpetual options position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`close perpetual options position.
Example:
$ %s tx %s close-position-perpetual-options --from myKeyName --chain-id ununifi-x
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
