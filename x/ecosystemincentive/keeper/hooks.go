package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

type Hooks struct {
	k Keeper
}

var _ nftbackedloantypes.NftbackedloanHooks = Hooks{}

// Hooks create new ecosystem-incentive hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// ------------------- nftbackedloan Module Hooks -------------------

func (h Hooks) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier nftbackedloantypes.NftId, fee sdk.Coin) {
	// if there's no fee, return
	if !fee.IsZero() {
		// call RewardDistributionOfnftbackedloan method to update reward information
		// for all the subjects of the nftbackedloan reward
		if err := h.k.RewardDistributionOfNftbackedloan(ctx, nftIdentifier, fee); err != nil {
			panic(err)
		}
	}

	// delete the recorded nft-id
	_ = h.k.DeleteFrontendRecord(ctx, nftIdentifier)
}

// AfterNftUnlistedWithoutPayment is called every time nft is unlisted without payment
func (h Hooks) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier nftbackedloantypes.NftId) {
	// delete the recorded nft-id
	_ = h.k.DeleteFrontendRecord(ctx, nftIdentifier)
}
