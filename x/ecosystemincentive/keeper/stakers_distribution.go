// Logic implementation for the distribution of the ecosystem-incentive reward to stakers
package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

// AllocateTokensToStakers performs reward and fee distribution to all validators based
// on the F1 fee distribution specification.
func (k Keeper) AllocateTokensToStakers(ctx sdk.Context, rewardAmount sdk.Coin) error {
	// transfer collected fees to the fee_collector module account eventually for the distribution module account
	// NOTE: But, it's worth considering when to actually send tokens to the target module account in this case
	// Because in this style that we send tokens every time the hook method is called, there's the possibility that the sending numbers gets too high to affects the perfomance of the
	// node and app.
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeDistributionNameForStakers, sdk.NewCoins(rewardAmount))
	if err != nil {
		panic(err)
	}

	// emit Event for the distribution of the ecosystem-incentive reward to stakers
	_ = ctx.EventManager().EmitTypedEvent(&types.EventDistributionForStakers{
		DistributedAmount: rewardAmount,
		BlockHeight:       ctx.BlockHeight(),
	})
	return nil
}

func (k Keeper) GetStakersRewardRate(ctx sdk.Context) sdk.Dec {
	params := k.GetParams(ctx)
	rewardParams := params.RewardParams

	for _, rewardParam := range rewardParams {
		if rewardParam.ModuleName == nftbackedloantypes.ModuleName {
			for _, rewardRate := range rewardParam.RewardRate {
				if rewardRate.RewardType == types.RewardType_STAKERS {
					return rewardRate.Rate
				}
			}
		}
	}

	// if target param wasn't found somehow, return zero dec
	return sdk.ZeroDec()
}
