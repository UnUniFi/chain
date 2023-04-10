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

// Position Size is considered in denom unit
func NewPerpetualFuturesNetPositionOfMarket(market Market, positionType PositionType, position_size_in_denom_exponent sdk.Int) PerpetualFuturesNetPositionOfMarket {
	return PerpetualFuturesNetPositionOfMarket{
		Market:                      market,
		PositionType:                positionType,
		PositionSizeInDenomExponent: position_size_in_denom_exponent,
	}
}

func (p PerpetualFuturesPositionInstance) SizeInDenomUnit(denomUnit uint32) sdk.Int {
	// return position size in the decimal unit
	return p.Size_.MulInt64(int64(denomUnit)).TruncateInt()
}
