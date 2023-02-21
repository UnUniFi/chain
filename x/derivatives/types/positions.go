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

func (m Position) NeedLiqudation(MarginMaintenanceRate sdk.Dec) bool {
	ins, err := UnpackPositionInstance(m.PositionInstance)
	if err != nil {
		return false
	}

	switch positionInstance := ins.(type) {
	case *PerpetualFuturesPositionInstance:
		return m.NeedLiqudationPerpetualFutures(MarginMaintenanceRate, *positionInstance)
		break
	case *PerpetualOptionsPositionInstance:
		panic("not implemented")
		break
	default:
		panic("not implemented")
	}
	return false
}

func (m Position) NeedLiqudationPerpetualFutures(MarginMaintenanceRate sdk.Dec, positionInstance PerpetualFuturesPositionInstance) bool {
	marginDec := sdk.NewDecFromInt(m.RemainingMargin.Amount).Mul(sdk.NewDecWithPrec(1, 0))
	principal := positionInstance.CalculatePrincipal().Mul(MarginMaintenanceRate)
	if marginDec.LT(principal) {
		return true
	} else {
		return false
	}
}
