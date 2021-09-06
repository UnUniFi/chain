package botanydist

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lcnem/jpyx/x/botanydist/keeper"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.MintPeriodInflation(ctx)
	if err != nil {
		panic(err)
	}
}
