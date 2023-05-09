package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	// ProposalTypeAddYieldFarm defines the type for a ProposalAddYieldFarm
	ProposalTypeAddYieldFarm          = "AddYieldFarm"
	ProposalTypeUpdateYieldFarm       = "UpdateYieldFarm"
	ProposalTypeStopYieldFarm         = "StopYieldFarm"
	ProposalTypeRemoveYieldFarm       = "RemoveYieldFarm"
	ProposalTypeAddYieldFarmTarget    = "AddYieldFarmTarget"
	ProposalTypeUpdateYieldFarmTarget = "UpdateYieldFarmTarget"
	ProposalTypeStopYieldFarmTarget   = "StopYieldFarmTarget"
	ProposalTypeRemoveYieldFarmTarget = "RemoveYieldFarmTarget"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeAddYieldFarm)
	govtypes.RegisterProposalType(ProposalTypeUpdateYieldFarm)
	govtypes.RegisterProposalType(ProposalTypeStopYieldFarm)
	govtypes.RegisterProposalType(ProposalTypeRemoveYieldFarm)
	govtypes.RegisterProposalType(ProposalTypeAddYieldFarmTarget)
	govtypes.RegisterProposalType(ProposalTypeUpdateYieldFarmTarget)
	govtypes.RegisterProposalType(ProposalTypeStopYieldFarmTarget)
	govtypes.RegisterProposalType(ProposalTypeRemoveYieldFarmTarget)
}

// Assert ProposalAddYieldFarm implements govtypes.Content at compile-time
var _ govtypes.Content = &ProposalAddYieldFarm{}

func NewProposalAddYieldFarm(title, description string, assetManagementAccount *AssetManagementAccount) *ProposalAddYieldFarm {
	return &ProposalAddYieldFarm{
		Title:       title,
		Description: description,
		Account:     assetManagementAccount,
	}
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p *ProposalAddYieldFarm) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (p *ProposalAddYieldFarm) ProposalType() string { return ProposalTypeAddYieldFarm }

// ValidateBasic validates the parameter change proposal
func (p *ProposalAddYieldFarm) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

var _ govtypes.Content = &ProposalUpdateYieldFarm{}

func NewProposalUpdateYieldFarm(title, description string, assetManagementAccount *AssetManagementAccount) *ProposalUpdateYieldFarm {
	return &ProposalUpdateYieldFarm{
		Title:       title,
		Description: description,
		Account:     assetManagementAccount,
	}
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p *ProposalUpdateYieldFarm) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (p *ProposalUpdateYieldFarm) ProposalType() string { return ProposalTypeUpdateYieldFarm }

// ValidateBasic validates the parameter change proposal
func (p *ProposalUpdateYieldFarm) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

var _ govtypes.Content = &ProposalStopYieldFarm{}

func NewProposalStopYieldFarm(title, description, assetManagementAccountId string) *ProposalStopYieldFarm {
	return &ProposalStopYieldFarm{
		Title:       title,
		Description: description,
		Id:          assetManagementAccountId,
	}
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p *ProposalStopYieldFarm) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (p *ProposalStopYieldFarm) ProposalType() string { return ProposalTypeStopYieldFarm }

// ValidateBasic validates the parameter change proposal
func (p *ProposalStopYieldFarm) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

var _ govtypes.Content = &ProposalRemoveYieldFarm{}

func NewProposalRemoveYieldFarm(title, description, assetManagementAccountId string) *ProposalRemoveYieldFarm {
	return &ProposalRemoveYieldFarm{
		Title:       title,
		Description: description,
		Id:          assetManagementAccountId,
	}
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p *ProposalRemoveYieldFarm) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (p *ProposalRemoveYieldFarm) ProposalType() string { return ProposalTypeRemoveYieldFarm }

// ValidateBasic validates the parameter change proposal
func (p *ProposalRemoveYieldFarm) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

var _ govtypes.Content = &ProposalAddYieldFarmTarget{}

func NewProposalAddYieldFarmTarget(title, description string, target *AssetManagementTarget) *ProposalAddYieldFarmTarget {
	return &ProposalAddYieldFarmTarget{
		Title:       title,
		Description: description,
		Target:      target,
	}
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p *ProposalAddYieldFarmTarget) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (p *ProposalAddYieldFarmTarget) ProposalType() string { return ProposalTypeAddYieldFarmTarget }

// ValidateBasic validates the parameter change proposal
func (p *ProposalAddYieldFarmTarget) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

var _ govtypes.Content = &ProposalUpdateYieldFarmTarget{}

func NewProposalUpdateYieldFarmTarget(title, description string, target *AssetManagementTarget) *ProposalUpdateYieldFarmTarget {
	return &ProposalUpdateYieldFarmTarget{
		Title:       title,
		Description: description,
		Target:      target,
	}
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p *ProposalUpdateYieldFarmTarget) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (p *ProposalUpdateYieldFarmTarget) ProposalType() string {
	return ProposalTypeUpdateYieldFarmTarget
}

// ValidateBasic validates the parameter change proposal
func (p *ProposalUpdateYieldFarmTarget) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

var _ govtypes.Content = &ProposalStopYieldFarmTarget{}

func NewProposalStopYieldFarmTarget(title, description, assetManagementAccountId, assetManagementTargetId string) *ProposalStopYieldFarmTarget {
	return &ProposalStopYieldFarmTarget{
		Title:                    title,
		Description:              description,
		AssetManagementAccountId: assetManagementAccountId,
		Id:                       assetManagementTargetId,
	}
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p *ProposalStopYieldFarmTarget) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (p *ProposalStopYieldFarmTarget) ProposalType() string {
	return ProposalTypeStopYieldFarmTarget
}

// ValidateBasic validates the parameter change proposal
func (p *ProposalStopYieldFarmTarget) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

var _ govtypes.Content = &ProposalRemoveYieldFarmTarget{}

func NewProposalRemoveYieldFarmTarget(title, description, assetManagementAccountId, assetManagementTargetId string) *ProposalRemoveYieldFarmTarget {
	return &ProposalRemoveYieldFarmTarget{
		Title:                    title,
		Description:              description,
		AssetManagementAccountId: assetManagementAccountId,
		Id:                       assetManagementTargetId,
	}
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p *ProposalRemoveYieldFarmTarget) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (p *ProposalRemoveYieldFarmTarget) ProposalType() string {
	return ProposalTypeRemoveYieldFarmTarget
}

// ValidateBasic validates the parameter change proposal
func (p *ProposalRemoveYieldFarmTarget) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}
