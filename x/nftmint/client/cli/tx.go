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
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "nftmint transactions subcommands",
		Long:                       "Provides the most common nft minting applications, compatible with Ethereum's erc721 contract",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreateClass(),
		CmdMintNFT(),
	)

	return cmd
}

func CmdCreateClass() *cobra.Command {
	cmd := &cobra.Command{
		// TODO: write appropriate guide
		Use:   "create-class [class-name] [base-token-uri]] [token-supply-cap] [minting-permission] --from [sender]",
		Args:  cobra.ExactArgs(4),
		Short: "create class for minting NFTs",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Example:$ %s tx %s create-class <class-name> <base-token-uri> <token-supply-cap> <minting-permission: 0=OnlyOwner, 1=Anyone>"+
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
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsCreateClass)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		// TODO: write appropriate guide
		Use:   "mint-nft [class-id] [nft-id] [receiver] --from [sender]",
		Args:  cobra.ExactArgs(3),
		Short: "mint NFT under specific class by class-id",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Note: nft-id will be that nft-uri combined with base token uri of the class-id, like <base-token-uri><nft-id>"+
				"$ %s tx %s mint-nft <class-id> <nft-id> <receiver>"+
				"--from <sender> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>", version.AppName, types.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			recipient := args[2]
			recipientAddr, err := sdk.AccAddressFromBech32(recipient)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintNFT(
				sender,
				args[0],
				args[1],
				recipientAddr,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsCreateClass)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
