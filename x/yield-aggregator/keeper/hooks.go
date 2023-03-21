package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	epochstypes "github.com/UnUniFi/chain/x/epochs/types"
)

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	// every epoch
	epochIdentifier := epochInfo.Identifier

	// process unbondings
	if epochIdentifier == "day" {
		vaults := k.GetAllVault(ctx)
		for _, vault := range vaults {
			totalAmount := k.VaultAmountTotal(ctx, vault)
			reserve := k.VaultWithdrawalAmount(ctx, vault)
			unbonding := k.VaultUnbondingAmountInStrategies(ctx, vault)

			targetUnbonded := totalAmount.ToDec().Mul(vault.WithdrawReserveRate).RoundInt()
			if targetUnbonded.LT(reserve.Add(unbonding)) {
				continue
			}
			amountToUnbond := targetUnbonded.Sub(reserve.Add(unbonding))
			for _, strategyWeight := range vault.StrategyWeights {
				strategy, found := k.GetStrategy(ctx, vault.Denom, strategyWeight.StrategyId)
				if !found {
					continue
				}
				strategyAmount := amountToUnbond.ToDec().Mul(strategyWeight.Weight).RoundInt()
				cacheCtx, _ := ctx.CacheContext()
				err := k.UnstakeFromStrategy(cacheCtx, vault, strategy, strategyAmount)
				if err != nil {
					fmt.Println("Epoch unstaking error", err.Error())
				} else {
					err = k.UnstakeFromStrategy(ctx, vault, strategy, strategyAmount)
					if err != nil {
						panic(fmt.Sprintln("Epoch unstaking error", err))
					}
				}
			}
		}
	}
}

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
}

// Hooks wrapper struct for incentives keeper
type Hooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// epochs hooks
func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	h.k.BeforeEpochStart(ctx, epochInfo)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	h.k.AfterEpochEnd(ctx, epochInfo)
}
