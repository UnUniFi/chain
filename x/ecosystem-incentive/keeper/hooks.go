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
	memoInputs, err := types.ParseMemo([]byte(txMemo))

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

	// guide the execution based on the version in the memo inputs
	// switch by values of AvailableVersions which is defined in ../types/memo.go
	//var AvailableVersions = []string{
	//	"v1",
	//	}
	switch memoInputs.Version {
	// types.AvailableVersions[0] = "v1"
	case types.AvailableVersions[0]:
		// Store the incentive-unit-id in NftIdForFrontend KVStore with nft-id as key
		h.k.RecordNftIdWithIncentiveUnitId(ctx, nftIdentifier, memoInputs.IncentiveUnitId)

	// If the value doesn't match any cases, emit event and don't do anything
	default:
		_ = ctx.EventManager().EmitTypedEvent(&types.EventVersionUnmatched{
			UnmatchedVersion: memoInputs.Version,
			ClassId:          nftIdentifier.ClassId,
			NftId:            nftIdentifier.NftId,
		})
	}
}

func (h Hooks) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier nftmarkettypes.NftIdentifier, fee sdk.Coin) {
	// call AccumulateRewardForFrontend method to update reward information
	// for the subjects defined by incentiveUnitId associated with nftIdentifier
	h.k.AccumulateRewardForFrontend(ctx, nftIdentifier, fee)

	// delete the recorded nft-id with incetive-unit-id
	h.k.DeleteNftId(ctx, nftIdentifier)
}

// AfterNftUnlistedWithoutPayment is called every time nft is unlisted without payment
func (h Hooks) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier nftmarkettypes.NftIdentifier) {
	// delete the recorded nft-id with incetive-unit-id
	h.k.DeleteNftId(ctx, nftIdentifier)
}
