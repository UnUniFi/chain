package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

var _ PositionInstance = (*PerpetualFuturesPositionInstance)(nil)

func UnpackPerpetualFuturesPositionInstance(positionAny types.Any) PositionInstance {
	if positionAny.TypeUrl == "/"+proto.MessageName(&PerpetualFuturesPositionInstance{}) {
		var position PerpetualFuturesPositionInstance
		err := position.Unmarshal(positionAny.Value)
		if err != nil {
			return nil
		}
		return &position
	}

	return nil
}

func (positionInstance PerpetualFuturesPositionInstance) CalculatePrincipal() sdk.Dec {
	return positionInstance.Size_.Quo(sdk.NewDecFromInt(positionInstance.Leverage))
}
