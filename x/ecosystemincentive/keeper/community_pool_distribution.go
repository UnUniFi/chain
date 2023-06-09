package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
)

// AllocateTokensToCommunityPool performs reward and fee distribution to all validators based
// on the F1 fee distribution specification.
func (k Keeper) AllocateTokensToCommunityPool(ctx sdk.Context, rewardAmount sdk.Coin) error {
	// transfer collected fees to the fee_collector module account eventually for the distribution module account
	// NOTE: But, it's worth considering when to actually send tokens to the target module account in this case
	// Because in this style that we send tokens every time the hook method is called, there's the possibility that the sending numbers gets too high to affects the perfomance of the
	// node and app.
	eiModuleAcc := authtypes.NewModuleAddress(types.ModuleName)
	err := k.ckKeeper.FundCommunityPool(ctx, sdk.NewCoins(rewardAmount), eiModuleAcc)
	if err != nil {
		// TODO: we need better panic handling
		panic(err)
	}

	// emit Event for the distribution of the ecosystem-incentive reward to stakers
	_ = ctx.EventManager().EmitTypedEvent(&types.EventDistributionForStakers{
		DistributedAmount: rewardAmount,
		BlockHeight:       ctx.BlockHeight(),
	})
	return nil
}
