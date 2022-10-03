package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

type Hooks struct {
	k Keeper
}

var _ nftmarkettypes.NftmarketHooks = Hooks{}

// Hooks create new ecosystem-incentive hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// ------------------- Nftmarket Module Hooks -------------------

func (h Hooks) AfterNftListed(ctx sdk.Context, nftIdentifier []byte, txMemo []byte) {
	_, err := types.ParseMemo(txMemo)

	// return immediately if memo data cannot be decoded properly
	// this doesn't mean MsgListNft fail. It succeeds anyway.
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (h Hooks) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier []byte, fee sdk.Coin) {
	h.k.AccumulateReward(ctx, nftIdentifier, fee)
}

func (h Hooks) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier []byte) {
}
