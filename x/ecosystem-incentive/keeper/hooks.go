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

func (h Hooks) AfterNftListed(ctx sdk.Context, nftIdentifier nftmarkettypes.NftIdentifier, txMemo string) {
	_, err := types.ParseMemo([]byte(txMemo))

	// return immediately after emitting event to tell decoding failed
	// if memo data cannot be decoded properly
	// this doesn't mean MsgListNft fail. It succeeds anyway.
	if err != nil {
		_ = fmt.Errorf(err.Error())
		_ = ctx.EventManager().EmitTypedEvent(&types.EventFailedParsingMemoInputs{
			ClassId: nftIdentifier.ClassId,
			NftId:   nftIdentifier.NftId,
			Memo:    txMemo,
		})
		return
	}

}

func (h Hooks) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier nftmarkettypes.NftIdentifier, fee sdk.Coin) {
	h.k.AccumulateReward(ctx, nftIdentifier, fee)
}

func (h Hooks) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier nftmarkettypes.NftIdentifier) {
}
