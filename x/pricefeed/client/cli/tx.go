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
	"github.com/cosmos/cosmos-sdk/version"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/pricefeed/types"
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
		CmdPostPrice(),
	)

	return cmd
}

func CmdPostPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "postprice [market-id] [price] [expiry-seconds]",
		Short: "postprice ubtc:uusdc 24528.185864015486004064 2024-02-20T12:00:38Z",
		Long: strings.TrimSpace(
			fmt.Sprintf(`post price.
Example:
$ %s tx %s postprice uusdc:ubtc 24528.185864015486004064 60  --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			marketId := args[0]
			price := sdk.MustNewDecFromStr(args[1])
			now := time.Now()
			expirySec, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}
			expiry := now.Add(time.Second * time.Duration(expirySec))
			msg := types.NewMsgPostPrice(clientCtx.GetFromAddress().String(), marketId, price, expiry, sdk.NewCoin("uusdc", sdk.NewInt(1000))) // TODO: deposit
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
