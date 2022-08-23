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
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		NewSubmitProposalAddYieldFarmTxCmd(),
		NewSubmitProposalUpdateYieldFarmTxCmd(),
		NewSubmitProposalStopYieldFarmTxCmd(),
		NewSubmitProposalRemoveYieldFarmTxCmd(),
		NewSubmitProposalAddYieldFarmTargetTxCmd(),
		NewSubmitProposalUpdateYieldFarmTargetTxCmd(),
		NewSubmitProposalStopYieldFarmTargetTxCmd(),
	)

	return cmd
}

// NewSubmitProposalAddYieldFarmTxCmd returns a CLI command handler for creating
// a proposal to add new yield farm
func NewSubmitProposalAddYieldFarmTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-add-yieldfarm",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to add yield farm",
		Long:  fmt.Sprintf(`Submit a proposal to add yield farm.`),
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

			assetManagementAccId, err := cmd.Flags().GetString(FlagAssetManagementAccountId)
			if err != nil {
				return err
			}

			assetManagementAccName, err := cmd.Flags().GetString(FlagAssetManagementAccountName)
			if err != nil {
				return err
			}

			enabled, err := cmd.Flags().GetBool(FlagEnabled)
			if err != nil {
				return err
			}

			content := types.NewProposalAddYieldFarm(title, description, &types.AssetManagementAccount{
				Id:      assetManagementAccId,
				Name:    assetManagementAccName,
				Enabled: enabled,
			})

			depositStr, err := cmd.Flags().GetString(FlagDeposit)
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

	cmd.Flags().AddFlagSet(FlagProposalTx())
	cmd.Flags().AddFlagSet(FlagAddAssetManagementAccount())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// ProposalUpdateYieldFarm returns a CLI command handler for creating
// a proposal to update a yield farm
func NewSubmitProposalUpdateYieldFarmTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-update-yieldfarm",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to update a yield farm",
		Long:  fmt.Sprintf(`Submit a proposal to update a yield farm.`),
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

			assetManagementAccId, err := cmd.Flags().GetString(FlagAssetManagementAccountId)
			if err != nil {
				return err
			}

			assetManagementAccName, err := cmd.Flags().GetString(FlagAssetManagementAccountName)
			if err != nil {
				return err
			}

			enabled, err := cmd.Flags().GetBool(FlagEnabled)
			if err != nil {
				return err
			}

			content := types.NewProposalUpdateYieldFarm(title, description, &types.AssetManagementAccount{
				Id:      assetManagementAccId,
				Name:    assetManagementAccName,
				Enabled: enabled,
			})
			depositStr, err := cmd.Flags().GetString(FlagDeposit)
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

	cmd.Flags().AddFlagSet(FlagProposalTx())
	cmd.Flags().AddFlagSet(FlagAddAssetManagementAccount())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// ProposalStopYieldFarm returns a CLI command handler for creating
// a proposal to stop a yield farm
func NewSubmitProposalStopYieldFarmTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-stop-yieldfarm",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to stop a yield farm",
		Long:  fmt.Sprintf(`Submit a proposal to stop a yield farm.`),
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

			assetManagementAccId, err := cmd.Flags().GetString(FlagAssetManagementAccountId)
			if err != nil {
				return err
			}

			content := types.NewProposalStopYieldFarm(title, description, assetManagementAccId)
			depositStr, err := cmd.Flags().GetString(FlagDeposit)
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

	cmd.Flags().AddFlagSet(FlagProposalTx())
	cmd.Flags().AddFlagSet(FlagStopAssetManagementAccount())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// ProposalRemoveYieldFarm returns a CLI command handler for creating
// a proposal to remove a yield farm
func NewSubmitProposalRemoveYieldFarmTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-remove-yieldfarm",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to remove a yield farm",
		Long:  fmt.Sprintf(`Submit a proposal to remove a yield farm.`),
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

			assetManagementAccId, err := cmd.Flags().GetString(FlagAssetManagementAccountId)
			if err != nil {
				return err
			}

			content := types.NewProposalRemoveYieldFarm(title, description, assetManagementAccId)
			depositStr, err := cmd.Flags().GetString(FlagDeposit)
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

	cmd.Flags().AddFlagSet(FlagProposalTx())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// ProposalAddYieldFarmTarget returns a CLI command handler for creating
// a proposal to add a yield farm target
func NewSubmitProposalAddYieldFarmTargetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-add-yieldfarmtarget",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to add a yield farm target",
		Long:  fmt.Sprintf(`Submit a proposal to add a yield farm target.`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			target, err := getAssetManagementTargetFromFlags(cmd)
			if err != nil {
				return err
			}

			content := types.NewProposalAddYieldFarmTarget(title, description, target)

			msg, err := govv1beta1.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagProposalTx())
	cmd.Flags().AddFlagSet(FlagAddAssetManagementTarget())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// ProposalUpdateYieldFarmTarget returns a CLI command handler for creating
// a proposal to update a yield farm target
func NewSubmitProposalUpdateYieldFarmTargetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-update-yieldfarmtarget",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to update a yield farm target",
		Long:  fmt.Sprintf(`Submit a proposal to update a yield farm target.`),
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

			target, err := getAssetManagementTargetFromFlags(cmd)
			if err != nil {
				return err
			}

			content := types.NewProposalUpdateYieldFarmTarget(title, description, target)
			depositStr, err := cmd.Flags().GetString(FlagDeposit)
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

	cmd.Flags().AddFlagSet(FlagProposalTx())
	cmd.Flags().AddFlagSet(FlagAddAssetManagementTarget())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getAssetManagementTargetFromFlags(cmd *cobra.Command) (*types.AssetManagementTarget, error) {
	assetManagementAccId, err := cmd.Flags().GetString(FlagAssetManagementAccountId)
	if err != nil {
		return nil, err
	}

	assetManagementAccAddress, err := cmd.Flags().GetString(FlagAssetManagementAccountAddress)
	if err != nil {
		return nil, err
	}

	assetManagementTargetId, err := cmd.Flags().GetString(FlagAssetManagementTargetId)
	if err != nil {
		return nil, err
	}

	unbondingSeconds, err := cmd.Flags().GetUint64(FlagUnbondingSeconds)
	if err != nil {
		return nil, err
	}

	assetConditionsStr, err := cmd.Flags().GetString(FlagAssetConditions)
	if err != nil {
		return nil, err
	}

	assetConditions := []types.AssetCondition{}
	assetConditionStrArr := strings.Split(assetConditionsStr, ",")
	for _, assetConditionStr := range assetConditionStrArr {
		split := strings.Split(assetConditionStr, ":")
		if len(split) != 2 {
			return nil, fmt.Errorf("invalid asset condition expression")
		}
		ratio, err := strconv.Atoi(split[2])
		if err != nil {
			return nil, err
		}
		assetConditions = append(assetConditions, types.AssetCondition{
			Denom: split[0],
			Min:   split[1],
			Ratio: uint32(ratio),
		})
	}

	integrateType, err := cmd.Flags().GetString(FlagIntegrateType)
	if err != nil {
		return nil, err
	}

	return &types.AssetManagementTarget{
		Id:                       assetManagementTargetId,
		AssetManagementAccountId: assetManagementAccId,
		AccountAddress:           assetManagementAccAddress,
		AssetConditions:          assetConditions,
		UnbondingTime:            time.Second * time.Duration(unbondingSeconds),
		IntegrateInfo: types.IntegrateInfo{
			Type:              types.IntegrateType(types.IntegrateType_value[integrateType]),
			ContractIbcPortId: "",
			ModName:           "",
		},
	}, nil
}

// ProposalStopYieldFarmTarget returns a CLI command handler for creating
// a proposal to stop a yield farm target
func NewSubmitProposalStopYieldFarmTargetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-stop-yieldfarmtarget",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to stop a yield farm target",
		Long:  fmt.Sprintf(`Submit a proposal to stop a yield farm target.`),
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

			assetManagementAccId, err := cmd.Flags().GetString(FlagAssetManagementAccountId)
			if err != nil {
				return err
			}

			assetManagementTargetId, err := cmd.Flags().GetString(FlagAssetManagementTargetId)
			if err != nil {
				return err
			}

			content := types.NewProposalStopYieldFarmTarget(title, description, assetManagementAccId, assetManagementTargetId)
			depositStr, err := cmd.Flags().GetString(FlagDeposit)
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

	cmd.Flags().AddFlagSet(FlagProposalTx())
	cmd.Flags().AddFlagSet(FlagStopAssetManagementTarget())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// ProposalRemoveYieldFarmTarget returns a CLI command handler for creating
// a proposal to remove a yield farm target
func NewSubmitProposalRemoveYieldFarmTargetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal-remove-yieldfarmtarget",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to remove a yield farm target",
		Long:  fmt.Sprintf(`Submit a proposal to remove a yield farm target.`),
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

			assetManagementAccId, err := cmd.Flags().GetString(FlagAssetManagementAccountId)
			if err != nil {
				return err
			}

			assetManagementTargetId, err := cmd.Flags().GetString(FlagAssetManagementTargetId)
			if err != nil {
				return err
			}

			content := types.NewProposalRemoveYieldFarmTarget(title, description, assetManagementAccId, assetManagementTargetId)
			depositStr, err := cmd.Flags().GetString(FlagDeposit)
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

	cmd.Flags().AddFlagSet(FlagProposalTx())
	cmd.Flags().AddFlagSet(FlagStopAssetManagementTarget())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
