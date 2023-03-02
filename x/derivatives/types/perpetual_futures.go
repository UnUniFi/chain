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
	return positionInstance.Size_.Quo(sdk.NewDec(int64(positionInstance.Leverage)))
}

func NewPerpetualFuturesNetPositionOfMarket(market Market, position_size sdk.Dec) PerpetualFuturesNetPositionOfMarket {
	return PerpetualFuturesNetPositionOfMarket{
		Market:       market,
		PositionSize: position_size,
	}
}


func (positionInstance PerpetualFuturesPositionInstance) MarginRequirement(currencyRate sdk.Dec) sdk.Dec {
	return positionInstance.Size_.Mul(currencyRate).Quo(sdk.NewDec(int64(positionInstance.Leverage)))
}

func (m PerpetualFuturesPositionInstance) GetOrderSize() sdk.Dec {
	return sdk.NewDec(int64(m.Leverage)).Mul(m.Size_)
}
