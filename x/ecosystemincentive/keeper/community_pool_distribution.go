package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

// AllocateTokensToCommunityPool performs reward and fee distribution to the community pool
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

func (k Keeper) GetCommunityPoolRewardRate(ctx sdk.Context) sdk.Dec {
	params := k.GetParams(ctx)
	rewardParams := params.RewardParams

	for _, rewardParam := range rewardParams {
		if rewardParam.ModuleName == nftbackedloantypes.ModuleName {
			for _, rewardRate := range rewardParam.RewardRate {
				if rewardRate.RewardType == types.RewardType_COMMUNITY_POOL {
					return rewardRate.Rate
				}
			}
		}
	}

	// if target param wasn't found somehow, return zero dec
	return sdk.ZeroDec()
}
