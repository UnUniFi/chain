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
		position.Unmarshal(positionAny.Value)
		return &position
	}

	return nil
}

func CalculatePrincipal(position PerpetualFuturesPositionInstance) sdk.Dec {
	return position.Size_.Quo(sdk.NewDecFromInt(position.Leverage))
}
