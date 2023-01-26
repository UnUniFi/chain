package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

var _ Position = (*PerpetualFuturesPosition)(nil)

type PerpetualFuturesPositions []PerpetualFuturesPosition

func UnpackPerpetualFuturesPosition(positionAny *types.Any) (Position, error) {
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

func CalculatePrincipal(position PerpetualFuturesPosition) sdk.Dec {
	return position.Size_.Quo(sdk.NewDecFromInt(position.Leverage))
}