package types

import (
	"fmt"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type PositionInstance interface {
	proto.Message
}

type Positions []Position

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

func (m Position) NeedLiquidation(MarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate sdk.Dec) bool {
	ins, err := UnpackPositionInstance(m.PositionInstance)
	if err != nil {
		return false
	}

	switch positionInstance := ins.(type) {
	case *PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := NewPerpetualFuturesPosition(m, *positionInstance)
		return perpetualFuturesPosition.NeedLiquidation(MarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate)
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

func NewPerpetualFuturesPositionFromPosition(position Position) (PerpetualFuturesPosition, error) {
	ins, err := UnpackPositionInstance(position.PositionInstance)
	if err != nil {
		return PerpetualFuturesPosition{}, err
	}
	switch positionInstance := ins.(type) {
	case *PerpetualFuturesPositionInstance:
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
			PositionInstance: *positionInstance,
		}, nil
	default:
		return PerpetualFuturesPosition{}, fmt.Errorf("this Any doesn't have PerpetualFuturesPositionInstance value")
		break
	}
	return PerpetualFuturesPosition{}, fmt.Errorf("this Any doesn't have PerpetualFuturesPositionInstance value")
}

func (m PerpetualFuturesPosition) NeedLiquidation(minMarginMaintenanceRate, currentBaseUsdRate, currentQuoteUsdRate sdk.Dec) bool {
	marginMaintenanceRate := m.MarginMaintenanceRate(currentBaseUsdRate, currentQuoteUsdRate)
	if marginMaintenanceRate.LT(minMarginMaintenanceRate) {
		return true
	} else {
		return false
	}
}

func (m PerpetualFuturesPosition) OpenedPairRate() sdk.Dec {
	return m.OpenedBaseRate.Quo(m.OpenedQuoteRate)
}

// todo make test
func (m PerpetualFuturesPosition) EvaluatePosition(currentBaseUsdRate sdk.Dec) sdk.Dec {
	return currentBaseUsdRate.Mul(m.PositionInstance.Size_)
}

func MicroToNormalDenom(amount sdk.Dec) sdk.Int {
	return amount.Mul(sdk.MustNewDecFromStr("1000000")).TruncateInt()
}

func MicroToNormalDec(amount sdk.Dec) sdk.Dec {
	return amount.Mul(sdk.MustNewDecFromStr("1000000"))
}

func (m PerpetualFuturesPosition) CalcReturningAmountAtClose(baseUSDRate, quoteUSDRate sdk.Dec) (returningAmount math.Int, lossToLP math.Int) {
	principal := m.RemainingMargin.Amount
	pnlAmount := m.ProfitAndLossInMetrics(baseUSDRate, quoteUSDRate)

	returningAmount = principal.Add(pnlAmount.TruncateInt())

	// If loss is over the margin, it means liquidity provider takes the loss.
	if returningAmount.IsNegative() {
		lossToLP = returningAmount
		returningAmount = sdk.ZeroInt()
	}

	return returningAmount, lossToLP
}

// todo make test
func (m Positions) EvaluatePositions(posType PositionType, getCurrentPriceF func(denom string) (sdk.Dec, error)) sdk.Dec {
	usdMap := map[string]sdk.Dec{}
	result := sdk.ZeroDec()
	for _, position := range m {
		ins, err := UnpackPositionInstance(position.PositionInstance)
		if err != nil {
			panic(err)
		}

		if _, ok := usdMap[position.Market.BaseDenom]; !ok {
			price, err := getCurrentPriceF(position.Market.BaseDenom)
			if err != nil {
				panic(err)
			}
			usdMap[position.Market.BaseDenom] = price
		}

		switch positionInstance := ins.(type) {
		case *PerpetualFuturesPositionInstance:
			perpetualFuturesPosition := NewPerpetualFuturesPosition(position, *positionInstance)
			if perpetualFuturesPosition.PositionInstance.PositionType != posType {
				continue
			}
			result = result.Add(perpetualFuturesPosition.EvaluatePosition(usdMap[position.Market.BaseDenom]))
			break
		case *PerpetualOptionsPositionInstance:
			panic("not implemented")
		default:
			continue
		}
	}
	return result
}

func (m Positions) EvaluateLongPositions(getCurrentPriceF func(denom string) (sdk.Dec, error)) sdk.Dec {
	return m.EvaluatePositions(PositionType_LONG, getCurrentPriceF)
}

func (m Positions) EvaluateShortPositions(getCurrentPriceF func(denom string) (sdk.Dec, error)) sdk.Dec {
	return m.EvaluatePositions(PositionType_SHORT, getCurrentPriceF)
}

func (m PerpetualFuturesPosition) RequiredMarginInQuote(baseQuoteRate sdk.Dec) sdk.Dec {
	// 必要証拠金(quote単位) = 現在のbase/quoteレート * ポジションサイズ(base単位) ÷ レバレッジ
	return m.PositionInstance.MarginRequirement(baseQuoteRate)
}
func (m PerpetualFuturesPosition) RequiredMarginInBase() sdk.Dec {
	// 必要証拠金(base単位) = ポジションサイズ(base単位) ÷ レバレッジ // レートでの変動なし
	return m.PositionInstance.MarginRequirement(sdk.MustNewDecFromStr("1"))
}

// func (m PerpetualFuturesPosition) RequiredMarginInMetrics(requiredMarginInQuote, quoteUSDRate sdk.Dec) sdk.Dec {
func (m PerpetualFuturesPosition) RequiredMarginInMetrics(baseUSDRate, quoteUSDRate sdk.Dec) sdk.Dec {
	// 必要証拠金(USD単位) = 必要証拠金(quote単位) * 現在のquote/USDレート
	//                    = 必要証拠金(base単位) * 現在のbase/USDレート
	if m.RemainingMargin.Denom == m.Market.QuoteDenom {
		baseQuoteRate := baseUSDRate.Quo(quoteUSDRate)
		return m.RequiredMarginInQuote(baseQuoteRate).Mul(quoteUSDRate)
	} else if m.RemainingMargin.Denom == m.Market.BaseDenom {
		return m.RequiredMarginInBase().Mul(baseUSDRate)
	} else {
		panic("not supported denom")
	}
}
func (m PerpetualFuturesPosition) ProfitAndLossInQuote(baseUSDRate, quoteUSDRate sdk.Dec) sdk.Dec {
	// 損益(quote単位) = (longなら1,shortなら-1) * (現在のbase/quoteレート - ポジション開設時base/quoteレート) * ポジションサイズ(base単位)
	fmt.Println("baseUSDRate")
	fmt.Println(baseUSDRate.String())
	fmt.Println("quoteUSDRate")
	fmt.Println(quoteUSDRate.String())
	baseQuoteRate := baseUSDRate.Quo(quoteUSDRate)
	fmt.Println("baseQuoteRate")
	fmt.Println(baseQuoteRate.String())
	profitOrLoss := baseQuoteRate.Sub(m.OpenedPairRate()).Mul(m.PositionInstance.Size_)
	fmt.Println("profitOrLoss")
	fmt.Println(profitOrLoss.String())
	if m.PositionInstance.PositionType == PositionType_LONG {
		return profitOrLoss
	} else {
		return profitOrLoss.Neg()
	}
}

func (m PerpetualFuturesPosition) ProfitAndLossInMetrics(baseUSDRate, quoteUSDRate sdk.Dec) sdk.Dec {
	// 損益(USD単位) = 損益(quote単位) * 現在のquote/USDレート
	return m.ProfitAndLossInQuote(baseUSDRate, quoteUSDRate).Mul(quoteUSDRate)
}
func (m PerpetualFuturesPosition) MarginMaintenanceRate(baseUSDRate, quoteUSDRate sdk.Dec) sdk.Dec {
	// 証拠金維持率 = 有効証拠金(USD単位) ÷ 必要証拠金(USD単位)
	return m.EffectiveMarginInMetrics(baseUSDRate, quoteUSDRate).Quo(m.RequiredMarginInMetrics(baseUSDRate, quoteUSDRate))
}
func (m PerpetualFuturesPosition) RemainingMarginInBase(baseUSDRate sdk.Dec) sdk.Dec {
	// 残存証拠金(USD単位) = 残存証拠金(base単位) * 現在のbase/USDレート
	return sdk.NewDecFromInt(m.RemainingMargin.Amount).Mul(baseUSDRate)
}
func (m PerpetualFuturesPosition) RemainingMarginInQuote(quoteUSDRate sdk.Dec) sdk.Dec {
	// 残存証拠金(USD単位) = 残存証拠金(quote単位) * 現在のquote/USDレート
	return sdk.NewDecFromInt(m.RemainingMargin.Amount).Mul(quoteUSDRate)
}
func (m PerpetualFuturesPosition) RemainingMarginInMetrics(baseUSDRate, quoteUSDRate sdk.Dec) sdk.Dec {
	// 残存証拠金(USD単位) = 残存証拠金(base単位) * 現在のbase/USDレート
	//                    = 残存証拠金(quote単位) * 現在のquote/USDレート
	if m.RemainingMargin.Denom == m.Market.BaseDenom {
		return m.RemainingMarginInBase(baseUSDRate)
	} else if m.RemainingMargin.Denom == m.Market.QuoteDenom {
		return m.RemainingMarginInQuote(quoteUSDRate)
	} else {
		panic("not supported denom")
	}
}

func (m PerpetualFuturesPosition) EffectiveMarginInMetrics(baseUSDRate, quoteUSDRate sdk.Dec) sdk.Dec {
	// 有効証拠金(USD単位) = 残存証拠金(USD単位) + 損益(USD単位)
	return m.RemainingMarginInMetrics(baseUSDRate, quoteUSDRate).Add(m.ProfitAndLossInMetrics(baseUSDRate, quoteUSDRate))
}
