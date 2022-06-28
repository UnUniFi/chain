package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/nftmint/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	nftTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "nftmint transactions subcommands",
		Long:                       "Provides the most common nft logic for upper-level applications, compatible with Ethereum's erc721 contract",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	nftTxCmd.AddCommand(
		NewCmdCreateClass(),
	)

	return nftTxCmd
}

func NewCmdCreateClass() *cobra.Command {
	cmd := &cobra.Command{
		// TODO: write appropriate guide
		Use:   "create-class [class-name] [nft-id] [receiver] --from [sender]",
		Args:  cobra.ExactArgs(4),
		Short: "create class for minting NFTs",
		Long: strings.TrimSpace(fmt.Sprintf(
			"$ %s tx %s send <class-name> <base-token-uri> <token-supply-cap> <minting-permission>"+
				"--from <sender> "+
				"--symbol <class-symbol> "+
				"--description <class-description> "+
				"--class-uri <class-uri> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>", version.AppName, types.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenSupplyCap, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}
			mintingPermission, err := strconv.ParseInt(args[3], 10, 32)
			if err != nil {
				return err
			}
			classSymbol, err := cmd.Flags().GetString(FlagSymbol)
			if err != nil {
				return err
			}
			classDescription, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}
			classUri, err := cmd.Flags().GetString(FlagClassUri)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateClass(
				clientCtx.GetFromAddress(),
				args[0],
				args[1],
				tokenSupplyCap,
				types.MintingPermission(mintingPermission),
				classSymbol,
				classDescription,
				classUri,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().AddFlagSet(FsCreateClass)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
