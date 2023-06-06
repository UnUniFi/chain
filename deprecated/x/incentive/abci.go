package incentive

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/deprecated/x/incentive/keeper"
)

// BeginBlocker runs at the start of every block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)
	for _, rp := range params.CdpMintingRewardPeriods {
		err := k.AccumulateCdpMintingRewards(ctx, rp)
		if err != nil {
			panic(err)
		}
	}
}
