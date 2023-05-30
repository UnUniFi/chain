// Logic implementation for the distribution of the ecosystem-incentive reward to stakers
package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

// AllocateTokens performs reward and fee distribution to all validators based
// on the F1 fee distribution specification.
func (k Keeper) AllocateTokensToStakers(ctx sdk.Context, amount sdk.Coin) error {
	// transfer collected fees to the fee_collector module account eventually for the distribution module account
	// NOTE: But, it's worth considering when to actually send tokens to the target module account in this case
	// Because in this style that we send tokens every time the hook method is called, there's the possibility that the sending numbers gets too high to affects the perfomance of the
	// node and app.
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeDistributionNameForStakers, sdk.NewCoins(amount))
	if err != nil {
		panic(err)
	}

	// emit Event for the distribution of the ecosystem-incentive reward to stakers
	_ = ctx.EventManager().EmitTypedEvent(&types.EventDistributionForStakers{
		DistributedAmount: amount,
		BlockHeight:       ctx.BlockHeight(),
	})
	return nil
}
