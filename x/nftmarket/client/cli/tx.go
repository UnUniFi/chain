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
		CmdMintNft(),
		CmdCreateListing(),
		CmdCancelNftListing(),
		CmdExpandListingPeriod(),
		CmdSellingDecision(),
		CmdCreatePlaceBid(),
		CmdCancelBid(),
		CmdPayFullBid(),
		CmdEndListing(),
		CmdBorrow(),
		CmdRepay(),
		CmdMintStableCoin(),
		CmdBurnStableCoin(),
		CmdLiquidate(),
	)

	return cmd
}

func CmdMintNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [class-id] [nft-id] [uri] [uri-hash]",
		Short: "Mint an nft",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Mint an nft.
Example:
$ %s tx %s mint a10 a10 uri 888838  --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintNft(clientCtx.GetFromAddress(), args[0], args[1], args[2], args[3])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

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

func CmdSellingDecision() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "selling_decision [class-id] [nft-id]",
		Short: "broadcast selling decision message",
		Long: strings.TrimSpace(
			fmt.Sprintf(`broadcast selling decision message.
Example:
$ %s tx %s selling_decision 1 1 --from myKeyName --chain-id ununifi-x
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

			msg := types.NewMsgSellingDecision(clientCtx.GetFromAddress(), nftIde)

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
		Use:   "borrow [class-id] [nft-id] [amount]",
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

func CmdRepay() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repay [class-id] [nft-id] [amount]",
		Short: "repay loan on nft",
		Long: strings.TrimSpace(
			fmt.Sprintf(`repay loan on nft.
Example:
$ %s tx %s repay 1 1 100uguu --from myKeyName --chain-id ununifi-x
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

			msg := types.NewMsgRepay(clientCtx.GetFromAddress(), nftIde, borrowCoin)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdMintStableCoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint_stablecoin [class-id] [nft-id] [amount]",
		Short: "mint stablecoin",
		Long: strings.TrimSpace(
			fmt.Sprintf(`mint stablecoin.
Example:
$ %s tx %s mint_stablecoin 1 1 100usd --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// classId := args[0]
			// nftId := args[1]
			// nftIde := types.NftIdentifier{
			// 	ClassId: classId,
			// 	NftId:   nftId,
			// }

			// borrowCoin, err := sdk.ParseCoinNormalized(args[2])
			// if err != nil {
			// 	return err
			// }

			msg := types.NewMsgMintStableCoin(clientCtx.GetFromAddress())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdBurnStableCoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn_stablecoin [class-id] [nft-id] [amount]",
		Short: "burn stablecoin",
		Long: strings.TrimSpace(
			fmt.Sprintf(`burn stablecoin.
Example:
$ %s tx %s burn_stablecoin 1 1 100usd --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// classId := args[0]
			// nftId := args[1]
			// nftIde := types.NftIdentifier{
			// 	ClassId: classId,
			// 	NftId:   nftId,
			// }

			// borrowCoin, err := sdk.ParseCoinNormalized(args[2])
			// if err != nil {
			// 	return err
			// }

			msg := types.NewMsgBurnStableCoin(clientCtx.GetFromAddress())

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
		Use:   "liquidate [class-id] [nft-id]",
		Short: "liquidate",
		Long: strings.TrimSpace(
			fmt.Sprintf(`liquidate.
Example:
$ %s tx %s liquidate 1 1 --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// classId := args[0]
			// nftId := args[1]
			// nftIde := types.NftIdentifier{
			// 	ClassId: classId,
			// 	NftId:   nftId,
			// }

			// borrowCoin, err := sdk.ParseCoinNormalized(args[2])
			// if err != nil {
			// 	return err
			// }

			msg := types.NewMsgLiquidate(clientCtx.GetFromAddress())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdCancelNftListing() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel_listing [class-id] [nft-id]",
		Short: "Cancel nft listing",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Cancel nft listing.
Example:
$ %s tx %s cancel_listing 1 1 --from myKeyName --chain-id ununifi-x
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

			msg := types.NewMsgCancelNftListing(clientCtx.GetFromAddress(), nftIde)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdExpandListingPeriod() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expand_nft_listing [class-id] [nft-id]",
		Short: "Expand nft listing",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Expand nft listing.
Example:
$ %s tx %s expand_nft_listing 1 1 --from myKeyName --chain-id ununifi-x
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

			msg := types.NewMsgExpandListingPeriod(clientCtx.GetFromAddress(), nftIde)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCancelBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancelbid [class-id] [nft-id]",
		Short: "Cancel bid on nft",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Cancel bid on nft.
Example:
$ %s tx %s cancelbid 1 1 --from myKeyName --chain-id ununifi-x
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

			msg := types.NewMsgCancelBid(clientCtx.GetFromAddress(), nftIde)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdPayFullBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pay_fullbid [class-id] [nft-id]",
		Short: "Pay full bid on nft",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Pay full bid on nft.
Example:
$ %s tx %s pay_fullbid 1 1 --from myKeyName --chain-id ununifi-x
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

			msg := types.NewMsgPayFullBid(clientCtx.GetFromAddress(), nftIde)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
