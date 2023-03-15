package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

var _ PositionInstance = (*PerpetualOptionsPositionInstance)(nil)

func UnpackPerpetualOptionsPosition(positionAny types.Any) PositionInstance {
	if positionAny.TypeUrl == "/"+proto.MessageName(&PerpetualOptionsPositionInstance{}) {
		var position PerpetualOptionsPositionInstance
		err := position.Unmarshal(positionAny.Value)
		if err != nil {
			return nil
		}
		return &position
	}

	return nil
}
