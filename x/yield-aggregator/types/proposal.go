package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	// ProposalTypeAddStrategy defines the type for a ProposalAddStrategy
	ProposalTypeAddStrategy = "AddStrategy"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeAddStrategy)
}

// Assert ProposalAddStrategy implements govtypes.Content at compile-time
var _ govtypes.Content = &ProposalAddStrategy{}

func NewProposalAddStrategy(title, description string, denom, contractAddr, name string) *ProposalAddStrategy {
	return &ProposalAddStrategy{
		Title:           title,
		Description:     description,
		Denom:           denom,
		ContractAddress: contractAddr,
		Name:            name,
	}
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p *ProposalAddStrategy) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a parameter change proposal.
func (p *ProposalAddStrategy) ProposalType() string { return ProposalTypeAddStrategy }

// ValidateBasic validates the parameter change proposal
func (p *ProposalAddStrategy) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}
