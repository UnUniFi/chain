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

// CalculatePrincipal is not used in the entire codebase.
// TODO: So, maybe we should remove it?
func (positionInstance PerpetualFuturesPositionInstance) CalculatePrincipal() sdk.Dec {
	return positionInstance.Size_.Quo(sdk.NewDec(int64(positionInstance.Leverage)))
}

// Position Size is considered on a micro level in the backend
func NewPerpetualFuturesNetPositionOfMarket(market Market, position_size_in_micro sdk.Int) PerpetualFuturesNetPositionOfMarket {
	return PerpetualFuturesNetPositionOfMarket{
		Market:              market,
		PositionSizeInMicro: position_size_in_micro,
	}
}
