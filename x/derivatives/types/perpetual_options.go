package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

var _ PositionI = (*PerpetualOptionsPosition)(nil)
var _ OpenedPositionI = (*PerpetualOptionsOpenedPosition)(nil)
var _ ClosedPositionI = (*PerpetualOptionsClosedPosition)(nil)

type PerpetualOptionsPositions []PerpetualOptionsPosition

func UnpackPerpetualOptionsPosition(positionAny *types.Any) (PositionI, error) {
	if positionAny == nil {
		return nil, fmt.Errorf("this Any is nil")
	}
	if positionAny.TypeUrl == "/"+proto.MessageName(&PerpetualOptionsPosition{}) {
		var position PerpetualOptionsPosition
		position.Unmarshal(positionAny.Value)
		return &position, nil
	}

	return nil, nil
}

func UnpackPerpetualOptionsOpenedPosition(positionAny *types.Any) (OpenedPositionI, error) {
	if positionAny == nil {
		return nil, fmt.Errorf("this Any is nil")
	}
	if positionAny.TypeUrl == "/"+proto.MessageName(&PerpetualOptionsOpenedPosition{}) {
		var position PerpetualOptionsPosition
		position.Unmarshal(positionAny.Value)
		return &position, nil
	}

	return nil, nil
}

func UnpackPerpetualOptionsClosedPosition(positionAny *types.Any) (ClosedPositionI, error) {
	if positionAny == nil {
		return nil, fmt.Errorf("this Any is nil")
	}
	if positionAny.TypeUrl == "/"+proto.MessageName(&PerpetualOptionsClosedPosition{}) {
		var position PerpetualOptionsPosition
		position.Unmarshal(positionAny.Value)
		return &position, nil
	}

	return nil, nil
}
