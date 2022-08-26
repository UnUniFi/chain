package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/UnUniFi/chain/x/decentralized-vault/types"
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

	cmd.AddCommand(
		CmdNftLocked(),
		CmdNftUnlocked(),
		CmdNftTransferRequest(),
		CmdRejectTransferRequest(),
		CmdTransferred(),
	)

	return cmd
}

func CmdNftLocked() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft-locked [receiver] [nft-id] [uri] [uri-hash]",
		Short: "locked nft on other network",
		Long: strings.TrimSpace(
			fmt.Sprintf(`locked nft on other network.
Example:
$ %s tx %s nft-locked ununifi1wgjh88unam4tuln0ju6l6q6cd08zk2vs87uytv a10 uri 888838  --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			receiver, err := sdk.AccAddressFromBech32(args[0])
			msg := types.NewMsgNftLocked(clientCtx.GetFromAddress(), receiver, args[1], args[2], args[3])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdNftUnlocked() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft-unlocked [target_address] [nft-id]",
		Short: "unlocked nft on other network",
		Long: strings.TrimSpace(
			fmt.Sprintf(`unlocked nft on other network.
Example:
$ %s tx %s nft-locked ununifi1wgjh88unam4tuln0ju6l6q6cd08zk2vs87uytv a10 --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			target, err := sdk.AccAddressFromBech32(args[0])
			msg := types.NewMsgNftUnlocked(clientCtx.GetFromAddress(), target, args[1])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdNftTransferRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-request [nft-id]",
		Short: "nft transfer request",
		Long: strings.TrimSpace(
			fmt.Sprintf(`nft transfer request.
Example:
$ %s tx %s transfer-request a10 --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNftTransferRequest(clientCtx.GetFromAddress(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRejectTransferRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reject-transfer [nft-id]",
		Short: "reject nft transfer request",
		Long: strings.TrimSpace(
			fmt.Sprintf(`reject nft transfer request.
Example:
$ %s tx %s reject-transfer a10 --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNftRejectTransfer(clientCtx.GetFromAddress(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdTransferred() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transferred [nft-id]",
		Short: "nft transferred",
		Long: strings.TrimSpace(
			fmt.Sprintf(`nft transferred.
Example:
$ %s tx %s transferred a10 --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNftTransferred(clientCtx.GetFromAddress(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
