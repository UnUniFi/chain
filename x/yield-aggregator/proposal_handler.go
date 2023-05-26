package yield_aggregator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/UnUniFi/chain/x/yield-aggregator/keeper"
	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

// NewYieldAggregatorProposalHandler creates a new governance Handler
func NewYieldAggregatorProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ProposalAddStrategy:
			return handleProposalAddStrategy(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func handleProposalAddStrategy(ctx sdk.Context, k keeper.Keeper, p *types.ProposalAddStrategy) error {
	k.AppendStrategy(ctx, p.Denom, types.Strategy{
		Denom:           p.Denom,
		ContractAddress: p.ContractAddress,
		Name:            p.Name,
		GitUrl:          p.GitUrl,
	})
	return nil
}
