package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

type Hooks struct {
	k Keeper
}

var _ nftbackedloantypes.NftbackedloanHooks = Hooks{}

// Hooks create new ecosystem-incentive hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// ------------------- nftbackedloan Module Hooks -------------------

func (h Hooks) AfterNftListed(ctx sdk.Context, nftIdentifier nftbackedloantypes.NftIdentifier, txMemo string) {
	if len(txMemo) == 0 {
		return
	}

	memoInputs, err := types.ParseMemo([]byte(txMemo))

	// return immediately after emitting event to tell decoding failed
	// if memo data cannot be decoded properly
	// this doesn't mean MsgListNft fail. It succeeds anyway.
	if err != nil {
		_ = ctx.EventManager().EmitTypedEvent(&types.EventFailedParsingTxMemoData{
			ClassId: nftIdentifier.ClassId,
			NftId:   nftIdentifier.NftId,
			Memo:    txMemo,
		})
		return
	}

	// guide the execution based on the version in the memo inputs
	// switch by values of AvailableVersions which is defined in ../types/memo.go
	//var AvailableVersions = []string{
	//	"v1",
	//	}
	switch memoInputs.Version {
	// types.AvailableVersions[0] = "v1"
	case types.AvailableVersions[0]:
		// Store the incentive-unit-id in NftIdForFrontend KVStore with nft-id as key
		h.k.RecordRecipientContainerIdWithNftId(ctx, nftIdentifier, memoInputs.RecipientContainerId)

	// If the value doesn't match any cases, emit event and don't do anything
	default:
		_ = ctx.EventManager().EmitTypedEvent(&types.EventVersionUnmatched{
			UnmatchedVersion: memoInputs.Version,
			ClassId:          nftIdentifier.ClassId,
			NftId:            nftIdentifier.NftId,
		})
	}
}

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
