package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k Keeper) BeginBlocker(ctx sdk.Context) {
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}
	vaults := k.GetAllVault(ctx)
	for _, vault := range vaults {
		// register new tranches per cycle
		if int64(vault.Cycle+vault.LastTrancheTime) < ctx.BlockTime().Unix() {
			info := k.GetStrategyDepositInfo(ctx, vault.StrategyContract)
			k.SetTranchePool(ctx, types.TranchePool{
				Id:               k.GetLastTrancheId(ctx) + 1,
				StrategyContract: vault.StrategyContract,
				Denom:            info.Denom,
				DepositDenom:     info.DepositDenom,
				StartTime:        uint64(ctx.BlockTime().Unix()),
				Maturity:         vault.MaxMaturity,
				SwapFee:          params.TradeFeeRate,
				ExitFee:          sdk.ZeroDec(),
				TotalShares:      sdk.Coin{},
				PoolAssets:       sdk.Coins{},
			})
		}
	}
}
