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

func (h Hooks) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier nftbackedloantypes.NftIdentifier, fee sdk.Coin) {
	// if there's no fee, return
	if !fee.IsZero() {
		// call RewardDistributionOfnftbackedloan method to update reward information
		// for all the subjects of the nftmarke reward
		if err := h.k.RewardDistributionOfnftbackedloan(ctx, nftIdentifier, fee); err != nil {
			panic(err)
		}
	}

	// delete the recorded nft-id with incetive-unit-id
	h.k.DeleteFrontendRecord(ctx, nftIdentifier)
}

// AfterNftUnlistedWithoutPayment is called every time nft is unlisted without payment
func (h Hooks) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier nftbackedloantypes.NftIdentifier) {
	// delete the recorded nft-id with incetive-unit-id
	h.k.DeleteFrontendRecord(ctx, nftIdentifier)
}
