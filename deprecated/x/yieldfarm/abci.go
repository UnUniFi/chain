package yieldfarm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/deprecated/x/yieldfarm/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)
	farmers := k.GetAllFarmerInfos(ctx)
	for _, info := range farmers {
		rewards := sdk.Coins{}
		for _, coin := range info.Amount {
			rewards = rewards.Add(sdk.NewCoin(coin.Denom, coin.Amount.Mul(sdk.NewInt(int64(params.DailyReward))).Quo(sdk.NewInt(100))))
		}
		addr, err := sdk.AccAddressFromBech32(info.Account)
		if err != nil {
			continue
		}
		k.AllocateRewards(ctx, addr, rewards)
	}
}
