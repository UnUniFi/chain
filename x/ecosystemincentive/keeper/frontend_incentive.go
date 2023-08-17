package keeper

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/UnUniFi/chain/x/ecosystemincentive/types"
	nftbackedloantypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) MemoTxHandler(ctx sdk.Context, msgs []sdk.Msg, memo string) {
	if len(memo) == 0 {
		return
	}

	d := make(map[string]json.RawMessage)
	err := json.Unmarshal([]byte(memo), &d)
	if err != nil || d["frontend"] == nil {
		return
	}
	// txMemo := types.FrontendTxMemo{}
	metadata := types.FrontendMetadata{}
	err = json.Unmarshal(d["frontend"], &metadata)
	if err != nil {
		return
	}

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
			k.HandleMemoTxWithMsgListNft(ctx, msg, metadata)
		}
	}
}

func (k Keeper) HandleMemoTxWithMsgListNft(ctx sdk.Context, msg *nftbackedloantypes.MsgListNft, metadata types.FrontendMetadata) {
	// guide the execution based on the version in the memo inputs
	// switch by values of AvailableVersions which is defined in ../types/memo.go
	//var AvailableVersions = []string{
	//	"v1",
	//	}
	switch metadata.Version {
	// types.AvailableVersions[0] = 1
	case types.AvailableVersions[0]:
		// Store the incentive-unit-id in NftIdForFrontend KVStore with nft-id as key
		_ = k.RecordRecipientWithNftId(ctx, msg.NftId, metadata.Recipient)

	// If the value doesn't match any cases, emit event and don't do anything
	default:
		_ = ctx.EventManager().EmitTypedEvent(&types.EventVersionUnmatched{
			UnmatchedVersion: metadata.Version,
			ClassId:          msg.NftId.ClassId,
			TokenId:          msg.NftId.TokenId,
		})
	}
}
