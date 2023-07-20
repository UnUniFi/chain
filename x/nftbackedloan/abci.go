package nftbackedloan

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/keeper"
)

// EndBlocker updates the current pricefeed
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	// process liquidation of expired bids
	k.ProcessLiquidateExpiredBids(ctx)

	// handle full payment period endings
	k.HandleFullPaymentsPeriodEndings(ctx)

	// deliver successful bids
	k.DeliverSuccessfulBids(ctx)
}
