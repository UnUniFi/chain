package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

type Position interface {
	proto.Message
}

type Positions []Position

func UnpackPosition(positionAny *types.Any) (Position, error) {
	switch positionAny.TypeUrl {
	case "/" + proto.MessageName(&PerpetualFuturesPosition{}):
		var position PerpetualFuturesPosition
		position.Unmarshal(positionAny.Value)
		return &position, nil
	case "/" + proto.MessageName(&PerpetualOptionsPosition{}):
		var position PerpetualOptionsPosition
		position.Unmarshal(positionAny.Value)
		return &position, nil

	default:
		return nil, fmt.Errorf("this Any doesn't have Position value")
	}
}
