ackage cli

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

	"github.com/UnUniFi/chain/x/nftmarket/types"
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
		CmdSellingDecision(),
		CmdCreatePlaceBid(),
		CmdCancelBid(),
		CmdPayFullBid(),
		CmdEndListing(),
		CmdBorrow(),
		CmdRepay(),
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

			minDepositRate, err := cmd.Flags().GetString(FlagMinimumDepositRate)
			if err != nil {
				return err
			}
			automaticRef, err := cmd.Flags().GetBool(FlagAutomaticRefinancing)
			if err != nil {
				return err
			}

			bidToken, err := cmd.Flags().GetString(FlagBidToken)
			if err != nil {
				return err
			}

			msg := types.NewMsgListNft(clientCtx.GetFromAddress(), nftIde, bidToken, minDepositRate, automaticRef)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(FlagBidToken, "uguu", "bid token")
	cmd.Flags().String(FlagMinimumDepositRate, "0.1", "min minimum deposit rate")
	cmd.Flags().BoolP(FlagAutomaticRefinancing, "r", false, "automatic refinancing")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCreatePlaceBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "placebid [class-id] [nft-id] [bid-amount] [deposit-amount] [deposit-interest-rate] [bidding_hour_time]",
		Short: "Creates a new place bid",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Creates a new place bid.
Example:
$ %s tx %s placebid 1 1 100uguu --automatic-payment --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			classIdArg := 0
			nftIdArg := 1
			bidArgs := 2
			depositArg := 3
			depositRateArg := 4
			bidEndArg := 5

			classId := args[classIdArg]
			nftId := args[nftIdArg]
			nftIde := types.NftIdentifier{
				ClassId: classId,
				NftId:   nftId,
			}
			bidCoin, err := sdk.ParseCoinNormalized(args[bidArgs])
			if err != nil {
				return err
			}

			depositCoin, err := sdk.ParseCoinNormalized(args[depositArg])
			if err != nil {
				return err
			}

			depositInterestRate, err := sdk.NewDecFromStr(args[depositRateArg])
			if err != nil {
				return err
			}

			bidding_duration_hour, err := strconv.Atoi(args[bidEndArg])
			if err != nil {
				return err
			}

			automaticPayment, err := cmd.Flags().GetBool(FlagAutomaticPayment)
			if err != nil {
				return err
			}
			now := time.Now()
			// todo fix me
			// bid_end_at := now.Add(time.Hour * time.Duration(bidding_duration_hour))
			bid_end_at := now.Add(time.Second * time.Duration(bidding_duration_hour))

			msg := types.NewMsgPlaceBid(clientCtx.GetFromAddress(), nftIde, bidCoin, depositCoin, depositInterestRate.String(), bid_end_at, automaticPayment)

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
