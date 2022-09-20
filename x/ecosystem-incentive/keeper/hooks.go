package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

type Hooks struct {
	k Keeper
}

var _ nftmarkettypes.NftmarketHooks = Hooks{}

// Hooks create new ecosystem-incentive hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// ------------------- Nftmarket Module Hooks -------------------
func (h Hooks) AfterNftListed(ctx sdk.Context, nftIdentifier []byte, txMemo string) {

}

func (h Hooks) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier []byte, fee sdk.Coin) {

}

func (h Hooks) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier []byte) {

}
