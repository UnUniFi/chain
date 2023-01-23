package derivatives

import (
	"github.com/UnUniFi/chain/x/derivatives/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	// TODO: impose imaginary funding rate
	// TODO: cutting losses
}

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// TODO: save PoolMarketCap
}
