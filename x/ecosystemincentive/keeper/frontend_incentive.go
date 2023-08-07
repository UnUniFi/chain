package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) MemoTxHandler(ctx sdk.Context, msgs []sdk.Msg, memo string) {
	if len(memo) == 0 {
		return
	}

	memoInputs, err := types.ParseMemo([]byte(memo))

	if err != nil {
		// _ = ctx.EventManager().EmitTypedEvent(&types.EventFailedParsingTxMemoData{
		// 	ClassId: nftIdentifier.ClassId,
		// 	NftId:   nftIdentifier.NftId,
		// 	Memo:    memo,
		// })
		return
	}

	for _, msg := range msgs {
		switch msg := msg.(type) {
		case *nftbackedloantypes.MsgListNft:
			k.HandleMemoTxWithMsgListNft(ctx, msg, memoInputs)
		}
	}
}

func (k Keeper) HandleMemoTxWithMsgListNft(ctx sdk.Context, msg *nftbackedloantypes.MsgListNft, memo *types.TxMemoData) {
	// guide the execution based on the version in the memo inputs
	// switch by values of AvailableVersions which is defined in ../types/memo.go
	//var AvailableVersions = []string{
	//	"v1",
	//	}
	switch memo.Version {
	// types.AvailableVersions[0] = "v1"
	case types.AvailableVersions[0]:
		// Store the incentive-unit-id in NftIdForFrontend KVStore with nft-id as key
		k.RecordRecipientContainerIdWithNftId(ctx, msg.NftId, memo.RecipientContainerId)

	// If the value doesn't match any cases, emit event and don't do anything
	default:
		_ = ctx.EventManager().EmitTypedEvent(&types.EventVersionUnmatched{
			UnmatchedVersion: memo.Version,
			ClassId:          msg.NftId.ClassId,
			NftId:            msg.NftId.NftId,
		})
	}
}
