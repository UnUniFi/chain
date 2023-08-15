package stakeibc

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/keeper"
	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/stakeibc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker of stakeibc module
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, bk types.BankKeeper, ak types.AccountKeeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// Iterate over all host zones and verify redemption rate
	for _, hz := range k.GetAllHostZone(ctx) {
		rrSafe, err := k.IsRedemptionRateWithinSafetyBounds(ctx, hz)
		if !rrSafe {
			panic(fmt.Sprintf("[INVARIANT BROKEN!!!] %s's RR is %s. ERR: %v", hz.GetChainId(), hz.RedemptionRate.String(), err.Error()))
		}
	}

	redemptions := k.RecordsKeeper.GetAllUserRedemptionRecord(ctx)
	for _, redemption := range redemptions {
		cacheCtx, _ := ctx.CacheContext()
		_, err := k.WithdrawUndelegatedTokensToChain(cacheCtx, &types.MsgClaimUndelegatedTokens{
			Creator:    redemption.Receiver,
			HostZoneId: redemption.HostZoneId,
			Epoch:      redemption.EpochNumber,
			Sender:     redemption.Receiver,
		})
		if err == nil {
			fmt.Println("Successful WithdrawUndelegatedTokensToChain", err, redemption)
			_, err := k.WithdrawUndelegatedTokensToChain(ctx, &types.MsgClaimUndelegatedTokens{
				Creator:    redemption.Receiver,
				HostZoneId: redemption.HostZoneId,
				Epoch:      redemption.EpochNumber,
				Sender:     redemption.Receiver,
			})
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println("ERROR WithdrawUndelegatedTokensToChain", err, redemption)
		}
	}
}
