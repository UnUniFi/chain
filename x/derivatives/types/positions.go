package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

type PositionI interface {
	proto.Message
}

type OpenedPositionI interface {
	proto.Message
}

type ClosedPositionI interface {
	proto.Message
}

func UnpackPosition(positionAny *types.Any) (PositionI, error) {
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

func UnpackOpenedPosition(positionAny *types.Any) (OpenedPositionI, error) {
	position, err := UnpackPerpetualFuturesOpenedPosition(positionAny)
	if position != nil || err != nil {
		return position, err
	}

	position, err = UnpackPerpetualOptionsOpenedPosition(positionAny)
	if position != nil || err != nil {
		return position, err
	}

	return nil, fmt.Errorf("this Any doesn't have Position value")
}

func UnpackClosedPosition(positionAny *types.Any) (ClosedPositionI, error) {
	position, err := UnpackPerpetualFuturesClosedPosition(positionAny)
	if position != nil || err != nil {
		return position, err
	}

	position, err = UnpackPerpetualOptionsClosedPosition(positionAny)
	if position != nil || err != nil {
		return position, err
	}

	return nil, fmt.Errorf("this Any doesn't have Position value")
}
