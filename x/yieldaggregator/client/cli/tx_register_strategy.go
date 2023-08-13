package cli

import (
	"fmt"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

const (
	FlagTitle           = "title"
	FlagDescription     = "description"
	FlagDenom           = "denom"
	FlagContractAddress = "contract-addr"
	FlagName            = "name"
	FlagGitUrl          = "git-url"
)

func FlagProposalAddStrategyTx() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagTitle, "", "title of the proposal")
	fs.String(FlagDescription, "", "description of the proposal")
	fs.String(govcli.FlagDeposit, "", "initial deposit on the proposal")
	fs.String(FlagDenom, "", "denom of the strategy")
	fs.String(FlagContractAddress, "", "contract address of the strategy")
	fs.String(FlagName, "", "name of the strategy")
	fs.String(FlagGitUrl, "", "git url of the strategy")

	return fs
}

// CmdTxRegisterStrategy returns a CLI command handler for creating
// a proposal to add new strategy
func CmdTxRegisterStrategy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-strategy",
		Args:  cobra.ExactArgs(0),
		Short: "Submit a proposal to register a strategy",
		Long:  fmt.Sprintf(`Submit a proposal to register a strategy.`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			title, err := cmd.Flags().GetString(FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(FlagDenom)
			if err != nil {
				return err
			}

			contractAddr, err := cmd.Flags().GetString(FlagContractAddress)
			if err != nil {
				return err
			}

			name, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}

			gitUrl, err := cmd.Flags().GetString(FlagGitUrl)
			if err != nil {
				return err
			}

			content := types.NewProposalAddStrategy(title, description, denom, contractAddr, name, gitUrl)

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govv1beta1.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagProposalAddStrategyTx())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
