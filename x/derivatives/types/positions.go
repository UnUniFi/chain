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
	position, err := UnpackPerpetualFuturesPosition(positionAny)
	if position != nil || err != nil {
		return position, err
	}

	position, err = UnpackPerpetualOptionsPosition(positionAny)
	if position != nil || err != nil {
		return position, err
	}

	return nil, fmt.Errorf("this Any doesn't have Position value")
}
