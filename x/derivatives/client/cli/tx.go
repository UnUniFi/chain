package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	codecType "github.com/cosmos/cosmos-sdk/codec/types"

	ununifiType "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/derivatives/types"
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
		CmdReportLiquidation(),
		CmdReportLevyPeriod(),
	)

	return cmd
}

func CmdMintLiquidityProviderToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-to-pool [amount]",
		Short: "deposit to pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`deposit to pool.
Example:
$ %s tx %s deposit-to-pool --from myKeyName --chain-id ununifi-x
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

			msg := types.MsgDepositToPool{
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
		Use:   "withdraw-from-pool [amount] [redeem-denom]",
		Short: "withdraw from pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`withdraw from pool.
Example:
$ %s tx %s withdraw-from-pool 1 ubtc --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(2),
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
			denom := args[1]

			msg := types.NewMsgWithdrawFromPool(sender, amount, denom)

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
		// CmdOpenPerpetualOptionsPosition(),
	)

	return cmd
}

func CmdOpenPerpetualFuturesPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perpetual-futures [margin] [base-denom] [quote-denom] [long/short] [position-size] [leverage]",
		Short: "open perpetual futures position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`open perpetual futures position.
Example:
$ %s tx %s open-position perpetual-futures --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress()
			margin, err := sdk.ParseCoinNormalized(args[0])
			baseDenom := args[1]
			quoteDenom := args[2]
			if err != nil {
				return err
			}

			// todo: use args (or flags) to create position instance
			var positionType types.PositionType
			switch args[3] {
			case "long":
				positionType = types.PositionType_LONG
			case "short":
				positionType = types.PositionType_SHORT
			default:
				return fmt.Errorf("invalid position type")
			}

			positionSize, err := sdk.NewDecFromStr(args[4])
			if err != nil {
				return err
			}
			// str to uint32 for levarage from args[5]
			leverage, err := strconv.ParseUint(args[5], 10, 32)
			if err != nil {
				return err
			}

			positionInstVal := types.PerpetualFuturesPositionInstance{
				PositionType: positionType,
				Size_:        positionSize,
				Leverage:     uint32(leverage),
			}

			// positionInstBz, err := positionInstVal.Marshal()
			// if err != nil {
			// 	return err
			// }
			// positionInstI, err :=
			piAny, err := codecType.NewAnyWithValue(&positionInstVal)
			if err != nil {
				return err
			}

			// positionInstance := codecType.Any{
			// 	TypeUrl: "/ununifi.derivatives.PerpetualFuturesPositionInstance",
			// 	Value:   positionInstBz,
			// }

			market := types.Market{
				BaseDenom:  baseDenom,
				QuoteDenom: quoteDenom,
			}
			msg := types.NewMsgOpenPosition(sender, margin, market, *piAny)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// func CmdOpenPerpetualOptionsPosition() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "perpetual-options [margin] [base-denom] [quote-denom]",
// 		Short: "open perpetual options position",
// 		Long: strings.TrimSpace(
// 			fmt.Sprintf(`open perpetual options position.
// Example:
// $ %s tx %s open-position perpetual-options --from myKeyName --chain-id ununifi-x
// `, version.AppName, types.ModuleName)),
// 		Args: cobra.ExactArgs(3),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			return fmt.Errorf("not implemented")
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

func CmdClosePosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-position",
		Short: "report liquidation needed position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`close position.
Example:
$ %s tx %s close-position --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()
			potisionId := args[0]

			msg := types.NewMsgClosePosition(sender, potisionId)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdReportLiquidation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-liquidation [position-id] [reward-recipient]",
		Short: "report liquidation needed position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`report liquidation needed position.
Example:
$ %s tx %s report-liquidation --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()
			recipient, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.MsgReportLiquidation{
				Sender:          ununifiType.StringAccAddress(sender),
				PositionId:      args[0],
				RewardRecipient: ununifiType.StringAccAddress(recipient),
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

func CmdReportLevyPeriod() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-levy-period [position-id] [reward-recipient]",
		Short: "report levy needed position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`report levy needed position.
Example:
$ %s tx %s report-levy-period --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()
			recipient, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.MsgReportLevyPeriod{
				Sender:          ununifiType.StringAccAddress(sender),
				PositionId:      args[0],
				RewardRecipient: ununifiType.StringAccAddress(recipient),
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
