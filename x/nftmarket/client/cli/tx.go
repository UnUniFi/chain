package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/nftmarket/types"
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
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreateListing(),
		CmdCreatePlaceBid(),
		CmdEndListing(),
		CmdBorrow(),
	)

	return cmd
}

// todo: Implementation fields
// BidToken, MinBid, BidHook, ListingType
func CmdCreateListing() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listing [class-id] [nft-id]",
		Short: "Creates a new listing",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new listing, depositing nft.
Example:
$ %s tx %s listing 1 1 --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			classId := args[0]
			nftId := args[1]
			nftIde := types.NftIdentifier{
				ClassId: classId,
				NftId:   nftId,
			}

			msg := types.NewMsgListNft(clientCtx.GetFromAddress(), nftIde)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCreatePlaceBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "placebid [class-id] [nft-id] [amount]",
		Short: "Creates a new place bid",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Creates a new place bid.
Example:
$ %s tx %s placebid 1 1 100uguu --automatic-payment --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			classId := args[0]
			nftId := args[1]
			nftIde := types.NftIdentifier{
				ClassId: classId,
				NftId:   nftId,
			}
			bidCoin, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}
			automaticPayment, err := cmd.Flags().GetBool(FlagAutomaticPayment)
			if err != nil {
				return err
			}

			msg := types.NewMsgPlaceBid(clientCtx.GetFromAddress(), nftIde, bidCoin, automaticPayment)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().BoolP(FlagAutomaticPayment, "p", false, "automation payment")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdEndListing() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "endlisting [class-id] [nft-id]",
		Short: "end listing",
		Long: strings.TrimSpace(
			fmt.Sprintf(`end listing.
Example:
$ %s tx %s endlisting 1 1 --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			classId := args[0]
			nftId := args[1]
			nftIde := types.NftIdentifier{
				ClassId: classId,
				NftId:   nftId,
			}

			msg := types.NewMsgEndNftListing(clientCtx.GetFromAddress(), nftIde)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdBorrow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrow [class-id] [nft-id]",
		Short: "borrow denom",
		Long: strings.TrimSpace(
			fmt.Sprintf(`borrow denom.
Example:
$ %s tx %s borrow 1 1 100uguu --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			classId := args[0]
			nftId := args[1]
			nftIde := types.NftIdentifier{
				ClassId: classId,
				NftId:   nftId,
			}

			borrowCoin, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgBorrow(clientCtx.GetFromAddress(), nftIde, borrowCoin)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
