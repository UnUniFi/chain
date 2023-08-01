package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/cosmos/gogoproto/proto"
)

type NftDataI interface {
	proto.Message
}

func UnpackNftData(dataAny types.Any) NftDataI {
	if dataAny.TypeUrl == "/"+proto.MessageName(&NftData{}) {
		var nftData NftData
		err := nftData.Unmarshal(dataAny.Value)
		if err != nil {
			return nil
		}
		return &nftData
	}

	return nil
}
