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

	"github.com/UnUniFi/chain/x/nftfactory/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "nftfactory transactions subcommands",
		Long:                       "Provides the most common nft minting applications, compatible with Ethereum's erc721 contract",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreateClass(),
		CmdUpdateClass(),
		CmdMintNFT(),
		CmdBurnNFT(),
		CmdChangeAdmin(),
	)

	return cmd
}

func CmdCreateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-class [subclass] --from [sender]",
		Args:  cobra.ExactArgs(1),
		Short: "create class for minting NFTs",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Example:$ %s tx %s create-class <subclass>"+
				"--from <sender> "+
				"--name <name> "+
				"--symbol <symbol> "+
				"--description <description> "+
				"--class-uri <uri> "+
				"--class-uri-hash <uri-hash> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>", version.AppName, types.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subClass := args[0]

			className, err := cmd.Flags().GetString(FlagName)
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
			classUri, err := cmd.Flags().GetString(FlagUri)
			if err != nil {
				return err
			}
			classUriHash, err := cmd.Flags().GetString(FlagUriHash)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateClass(
				clientCtx.GetFromAddress().String(),
				subClass,
				className,
				classSymbol,
				classDescription,
				classUri,
				classUriHash,
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

func CmdUpdateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-class [class-id] --from [sender]",
		Args:  cobra.ExactArgs(1),
		Short: "update class",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Example:$ %s tx %s create-class <class-id> <name>"+
				"--from <sender> "+
				"--name <name> "+
				"--symbol <symbol> "+
				"--description <description> "+
				"--class-uri <uri> "+
				"--class-uri-hash <uri-hash> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>", version.AppName, types.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			classId := args[0]

			className, err := cmd.Flags().GetString(FlagName)
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
			classUri, err := cmd.Flags().GetString(FlagUri)
			if err != nil {
				return err
			}
			classUriHash, err := cmd.Flags().GetString(FlagUriHash)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateClass(
				clientCtx.GetFromAddress().String(),
				classId,
				className,
				classSymbol,
				classDescription,
				classUri,
				classUriHash,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsUpdateClass)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-nft [class-id] [token-id] [receiver] --from [sender]",
		Args:  cobra.ExactArgs(3),
		Short: "mint NFT under specific class by class-id",
		Long: strings.TrimSpace(fmt.Sprintf(
			"$ %s tx %s mint-nft <class-id> <token-id> <receiver>"+
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

			classId := args[0]
			tokenId := args[1]
			recipient := args[2]
			recipientAddr, err := sdk.AccAddressFromBech32(recipient)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintNFT(
				sender.String(),
				classId,
				tokenId,
				recipientAddr.String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdBurnNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-nft [class-id] [token-id] --from [sender]",
		Args:  cobra.ExactArgs(2),
		Short: "burn specified NFT",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			classId := args[0]
			tokenId := args[1]

			msg := types.NewMsgBurnNFT(
				sender.String(),
				classId,
				tokenId,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdChangeAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change-admin [class-id] [new-admin-address] --from [sender]",
		Args:  cobra.ExactArgs(2),
		Short: "Changes the admin address for a factory-created class. Must have admin authority to do so.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()
			classId := args[0]
			newAdmin := args[1]

			msg := types.NewMsgChangeAdmin(
				sender.String(),
				classId,
				newAdmin,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
