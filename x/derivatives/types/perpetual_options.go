package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

var _ Position = (*PerpetualOptionsPosition)(nil)

type PerpetualOptionsPositions []PerpetualOptionsPosition

func UnpackPerpetualOptionsPosition(positionAny *types.Any) (Position, error) {
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
