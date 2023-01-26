package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

var _ PositionI = (*PerpetualFuturesPosition)(nil)
var _ OpenedPositionI = (*PerpetualFuturesOpenedPosition)(nil)
var _ ClosedPositionI = (*PerpetualFuturesClosedPosition)(nil)

type PerpetualFuturesPositions []PerpetualFuturesPosition

func UnpackPerpetualFuturesPosition(positionAny *types.Any) (PositionI, error) {
	if positionAny == nil {
		return nil, fmt.Errorf("this Any is nil")
	}
	if positionAny.TypeUrl == "/"+proto.MessageName(&PerpetualFuturesPosition{}) {
		var position PerpetualFuturesPosition
		position.Unmarshal(positionAny.Value)
		return &position, nil
	}

	return nil, nil
}

func UnpackPerpetualFuturesOpenedPosition(positionAny *types.Any) (OpenedPositionI, error) {
	if positionAny == nil {
		return nil, fmt.Errorf("this Any is nil")
	}
	if positionAny.TypeUrl == "/"+proto.MessageName(&PerpetualFuturesOpenedPosition{}) {
		var position PerpetualFuturesPosition
		position.Unmarshal(positionAny.Value)
		return &position, nil
	}

	return nil, nil
}

func UnpackPerpetualFuturesClosedPosition(positionAny *types.Any) (ClosedPositionI, error) {
	if positionAny == nil {
		return nil, fmt.Errorf("this Any is nil")
	}
	if positionAny.TypeUrl == "/"+proto.MessageName(&PerpetualFuturesClosedPosition{}) {
		var position PerpetualFuturesPosition
		position.Unmarshal(positionAny.Value)
		return &position, nil
	}

	return nil, nil
}

func CalculatePrincipal(position PerpetualFuturesPosition) sdk.Dec {
	return position.Size_.Quo(sdk.NewDecFromInt(position.Leverage))
}
