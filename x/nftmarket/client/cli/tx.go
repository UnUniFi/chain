package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/UnUniFi/chain/x/nftmarket/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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

	cmd.AddCommand(CmdCreateListing())

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
