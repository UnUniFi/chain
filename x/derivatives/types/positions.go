package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type PositionInstance interface {
	proto.Message
}

func UnpackPositionInstance(positionAny types.Any) (PositionInstance, error) {
	position := UnpackPerpetualFuturesPositionInstance(positionAny)
	if position != nil {
		return position, nil
	}

	position = UnpackPerpetualOptionsPosition(positionAny)
	if position != nil {
		return position, nil
	}

	return nil, fmt.Errorf("this Any doesn't have PositionInstance value")
}

func MustUnpackPositionInstance(positionAny types.Any) PositionInstance {
	position, err := UnpackPositionInstance(positionAny)
	if err != nil {
		panic(err)
	}
	return position
}

func (m Position) NeedLiquidation(MarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate sdk.Dec) bool {
	ins, err := UnpackPositionInstance(m.PositionInstance)
	if err != nil {
		return false
	}

	switch positionInstance := ins.(type) {
	case *PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := NewPerpetualFuturesPosition(m, *positionInstance)
		return perpetualFuturesPosition.NeedLiquidation(MarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate)
		break
	case *PerpetualOptionsPositionInstance:
		panic("not implemented")
		break
	default:
		panic("not implemented")
	}
	return false
}

func NewPerpetualFuturesPosition(position Position, ins PerpetualFuturesPositionInstance) PerpetualFuturesPosition {
	return PerpetualFuturesPosition{
		Id:               position.Id,
		Market:           position.Market,
		Address:          position.Address,
		OpenedAt:         position.OpenedAt,
		OpenedBaseRate:   position.OpenedBaseRate,
		OpenedQuoteRate:  position.OpenedQuoteRate,
		OpenedHeight:     position.OpenedHeight,
		RemainingMargin:  position.RemainingMargin,
		LastLeviedAt:     position.LastLeviedAt,
		PositionInstance: ins,
	}
}

func (m PerpetualFuturesPosition) NeedLiquidation(minMarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate sdk.Dec) bool {
	marginMaintenanceRate := m.MarginMaintenanceRate(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate)
	if marginMaintenanceRate.LT(minMarginMaintenanceRate) {
		return true
	} else {
		return false
	}
}

// todo make test
func (m PerpetualFuturesPosition) EffectiveMargin(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate sdk.Dec) sdk.Dec {
	effectiveMargin := sdk.NewDecFromInt(m.RemainingMargin.Amount).Mul(currentBaseUsdRate.Quo(currentMarginUsdRate))

	revenue := m.CalcProfit(currentBaseUsdRate.Quo(currentQuoteUsdRate))
	if revenue.RevenueType == RevenueType_PROFIT {
		effectiveMargin = effectiveMargin.Add(sdk.NewDecFromInt(revenue.Amount.Amount))
	} else {
		effectiveMargin = effectiveMargin.Sub(sdk.NewDecFromInt(revenue.Amount.Amount))
	}
	return effectiveMargin
}

func (m PerpetualFuturesPosition) MarginMaintenanceRate(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate sdk.Dec) sdk.Dec {
	marginRequirement := m.PositionInstance.MarginRequirement(currentBaseUsdRate.Quo(currentMarginUsdRate))
	effectiveMargin := m.EffectiveMargin(currentBaseUsdRate, currentQuoteUsdRate, currentMarginUsdRate)

	marginMaintenanceRate := effectiveMargin.Quo(marginRequirement)
	return marginMaintenanceRate
}

func (m PerpetualFuturesPosition) OpenedPairRate() sdk.Dec {
	return m.OpenedBaseRate.Quo(m.OpenedQuoteRate)
}

// todo make test
func (m PerpetualFuturesPosition) EvaluatePosition(currentBaseUsdRate sdk.Dec) sdk.Dec {
	return currentBaseUsdRate.Mul(m.PositionInstance.GetOrderSize())
}

func (m PerpetualFuturesPosition) CalcProfit(closedPairRate sdk.Dec) Revenue {
	sub := closedPairRate.Sub(m.OpenedPairRate())
	revenue := m.GetRevenueType(sub)
	if sub.IsNegative() {
		sub = sub.Neg()
	}
	resultDec := sub.Mul(m.PositionInstance.GetOrderSize())
	return Revenue{
		RevenueType: revenue,
		Amount:      sdk.NewCoin(m.Market.BaseDenom, resultDec.RoundInt()),
	}
}

func (m PerpetualFuturesPosition) GetRevenueType(sub sdk.Dec) RevenueType {
	if m.PositionInstance.PositionType == PositionType_LONG {
		if sub.IsPositive() {
			return RevenueType_PROFIT
		} else {
			return RevenueType_LOSS
		}
	} else if m.PositionInstance.PositionType == PositionType_SHORT {
		// todo: think about amount is zero case
		if sub.IsNegative() {
			return RevenueType_PROFIT
		} else {
			return RevenueType_LOSS
		}
	} else {
		panic("not implemented")
	}
}

func (a Revenue) Equal(b Revenue) bool {
	if a.RevenueType != b.RevenueType {
		return false
	}
	if !a.Amount.IsEqual(b.Amount) {
		return false
	}
	return true
}
