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

func (m Position) IsValid(params Params, availableAssetInPool sdk.Coin) error {
	if !m.IsValidMarginAsset() {
		return fmt.Errorf("margin asset is not valid")
	}

	// check the least requirement for the margin
	if !m.RemainingMargin.Amount.IsPositive() {
		return fmt.Errorf("remaining margin must be positive")
	}

	pfPosition, err := NewPerpetualFuturesPositionFromPosition(m)
	if err != nil {
		return err
	}

	if !pfPosition.PositionInstance.IsValidLeverage(params.PerpetualFutures.MaxLeverage) {
		return fmt.Errorf("leverage is not valid")
	}

	if !pfPosition.IsValidPositionSize(params.PoolParams.QuoteTicker) {
		return fmt.Errorf("position size is not valid")
	}

	if availableAssetInPool.Amount.LTE(pfPosition.PositionInstance.SizeInDenomExponent(OneMillionInt)) {
		return ErrInsufficientAssetBalance
	}

	return nil
}

// Margin asset must be one of the market assets.
func (m Position) IsValidMarginAsset() bool {
	return (m.Market.BaseDenom == m.RemainingMargin.Denom || m.Market.QuoteDenom == m.RemainingMargin.Denom)
}

func (m PerpetualFuturesPosition) IsValidPositionSize(quoteTicker string) bool {
	// check position size validity
	baseMetricsRate := NewMetricsRateType(quoteTicker, m.Market.BaseDenom, m.OpenedBaseRate)
	quoteMetricsRate := NewMetricsRateType(quoteTicker, m.Market.QuoteDenom, m.OpenedQuoteRate)
	marginMaintenanceRate := m.MarginMaintenanceRate(baseMetricsRate, quoteMetricsRate)

	return !marginMaintenanceRate.LT(sdk.OneDec())
}

func (m PerpetualFuturesPositionInstance) IsValidLeverage(maxLeverage uint32) bool {
	return m.Leverage > 0 && m.Leverage <= maxLeverage
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

func (m Position) NeedLiquidation(MarginMaintenanceRate sdk.Dec, currentBaseMetricsRate, currentQuoteMetricsRate MetricsRateType) bool {
	ins, err := UnpackPositionInstance(m.PositionInstance)
	if err != nil {
		return false
	}

	switch positionInstance := ins.(type) {
	case *PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := NewPerpetualFuturesPosition(m, *positionInstance)
		return perpetualFuturesPosition.NeedLiquidation(MarginMaintenanceRate, currentBaseMetricsRate, currentQuoteMetricsRate)
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

func (m PerpetualFuturesPosition) NeedLiquidation(minMarginMaintenanceRate sdk.Dec, currentBaseMetricsRate, currentQuoteMetricsRate MetricsRateType) bool {
	marginMaintenanceRate := m.MarginMaintenanceRate(currentBaseMetricsRate, currentQuoteMetricsRate)
	if marginMaintenanceRate.LTE(minMarginMaintenanceRate) {
		return true
	} else {
		return false
	}
}

func (m PerpetualFuturesPosition) OpenedPairRate() sdk.Dec {
	return m.OpenedBaseRate.Quo(m.OpenedQuoteRate)
}

// todo make test
func (m PerpetualFuturesPosition) EvaluatePosition(currentBaseMetricsRate MetricsRateType) sdk.Dec {
	// FIXME: Don't use OneMillionInt derectly. issue #476
	return currentBaseMetricsRate.Amount.Amount.Mul(sdk.NewDecFromInt(m.PositionInstance.SizeInDenomExponent(OneMillionInt)))
}

// TODO: consider to use sdk.DecCoin
func NormalToMicroInt(amount sdk.Dec) sdk.Int {
	return amount.Mul(sdk.MustNewDecFromStr(OneMillionString)).TruncateInt()
}

func NormalToMicroDec(amount sdk.Dec) sdk.Dec {
	return amount.Mul(sdk.MustNewDecFromStr(OneMillionString))
}

// todo make test
func (m Positions) EvaluatePositions(posType PositionType, quoteTicker string, getCurrentPriceF func(denom string) (sdk.Dec, error)) sdk.Dec {
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

			metricsRate := NewMetricsRateType(quoteTicker, position.Market.BaseDenom, usdMap[position.Market.BaseDenom])

			result = result.Add(perpetualFuturesPosition.EvaluatePosition(metricsRate))
			break
		case *PerpetualOptionsPositionInstance:
			panic("not implemented")
		default:
			continue
		}
	}
	return result
}

func (m Positions) EvaluateLongPositions(quoteTicker string, getCurrentPriceF func(denom string) (sdk.Dec, error)) sdk.Dec {
	return m.EvaluatePositions(PositionType_LONG, quoteTicker, getCurrentPriceF)
}

func (m Positions) EvaluateShortPositions(quoteTicker string, getCurrentPriceF func(denom string) (sdk.Dec, error)) sdk.Dec {
	return m.EvaluatePositions(PositionType_SHORT, quoteTicker, getCurrentPriceF)
}

func (positionInstance PerpetualFuturesPositionInstance) MarginRequirement(currencyRate sdk.Dec) sdk.Int {
	// FIXME: Don't use OneMillionInt derectly. issue #476
	return sdk.NewDecFromInt(positionInstance.SizeInDenomExponent(OneMillionInt)).Mul(currencyRate).Quo(sdk.NewDec(int64(positionInstance.Leverage))).TruncateInt()
}

func (m PerpetualFuturesPosition) RequiredMarginInQuote(baseQuoteRate sdk.Dec) sdk.Int {
	// 必要証拠金(quote単位) = 現在のbase/quoteレート * ポジションサイズ(base単位) ÷ レバレッジ
	return m.PositionInstance.MarginRequirement(baseQuoteRate)
}

func (m PerpetualFuturesPosition) RequiredMarginInBase() sdk.Int {
	// 必要証拠金(base単位) = ポジションサイズ(base単位) ÷ レバレッジ // レートでの変動なし
	return m.PositionInstance.MarginRequirement(sdk.MustNewDecFromStr("1"))
}

func (m PerpetualFuturesPosition) RequiredMarginInMetrics(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Int {
	// 必要証拠金(USD単位) = 必要証拠金(quote単位) * 現在のquote/USDレート
	//                    = 必要証拠金(base単位) * 現在のbase/USDレート
	if m.RemainingMargin.Denom == m.Market.QuoteDenom {
		baseQuoteRate := baseMetricsRate.Amount.Amount.Quo(quoteMetricsRate.Amount.Amount)
		return sdk.NewDecFromInt(m.RequiredMarginInQuote(baseQuoteRate)).Mul(quoteMetricsRate.Amount.Amount).TruncateInt()
	} else {
		return sdk.NewDecFromInt(m.RequiredMarginInBase()).Mul(baseMetricsRate.Amount.Amount).TruncateInt()
	}
}

// CalcReturningAmountAtClose calculates the amount of the principal and the profit/loss at the close of the position.
func (m PerpetualFuturesPosition) CalcReturningAmountAtClose(baseMetricsRate, quoteMetricsRate MetricsRateType, tradingFee sdk.Int) (returningAmount math.Int, lossToLP math.Int) {
	principal := m.RemainingMargin.Amount
	// pnlAmountInMetrics represents the profit/loss amount in the metrics asset of the market.
	// In the most cases, it means it's in "usd".
	// AND, MORE IMPORTANTLY,
	// it's not calculated on a micro level. So, it has to be modified to micro level by multiplying
	// one million to represent the returning amount.
	pnlAmount := m.ProfitAndLoss(baseMetricsRate, quoteMetricsRate)

	returningAmount = principal.Add(pnlAmount)

	// If loss is over the margin, it means liquidity provider takes the loss.
	if returningAmount.IsNegative() {
		lossToLP = returningAmount
		returningAmount = sdk.ZeroInt()
	} else {
		lossToLP = sdk.ZeroInt()
	}

	// Subtract the trading fee.
	if returningAmount.LT(tradingFee) {
		// Return 0 returning amount and the trading fee subtracted by the returning amount as LossToLP
		return sdk.ZeroInt(), tradingFee.Sub(returningAmount)
	}

	returningAmount = returningAmount.Sub(tradingFee)

	return returningAmount, lossToLP
}

// ProfitAndLoss returns the profit/loss amount in the margin denom
func (m PerpetualFuturesPosition) ProfitAndLoss(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Int {
	pnlAmountInMetrics := sdk.NewDecFromInt(m.ProfitAndLossInMetrics(baseMetricsRate, quoteMetricsRate))

	// Make it be calculated in the corresponding asset as the principal.
	var pnlAmount sdk.Dec
	if m.RemainingMargin.Denom == m.Market.BaseDenom {
		pnlAmount = pnlAmountInMetrics.Quo(baseMetricsRate.Amount.Amount)
	} else {
		pnlAmount = pnlAmountInMetrics.Quo(quoteMetricsRate.Amount.Amount)
	}

	return pnlAmount.TruncateInt()
}

func (m PerpetualFuturesPosition) ProfitAndLossInQuote(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Int {
	// 損益(quote単位) = (longなら*1,shortなら*-1) * (現在のbase/quoteレート - ポジション開設時base/quoteレート) * ポジションサイズ(base単位)
	baseQuoteRate := baseMetricsRate.Amount.Amount.Quo(quoteMetricsRate.Amount.Amount)
	// FIXME: Don't use OneMillionInt derectly. issue #476
	profitAndLoss := baseQuoteRate.Sub(m.OpenedPairRate()).Mul(sdk.NewDecFromInt(m.PositionInstance.SizeInDenomExponent(OneMillionInt))).TruncateInt()
	if m.PositionInstance.PositionType == PositionType_LONG {
		return profitAndLoss
	} else {
		return profitAndLoss.Neg()
	}
}

func (m PerpetualFuturesPosition) ProfitAndLossInMetrics(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Int {
	// 損益(USD単位) = 損益(quote単位) * 現在のquote/USDレート
	return sdk.NewDecFromInt(m.ProfitAndLossInQuote(baseMetricsRate, quoteMetricsRate)).Mul(quoteMetricsRate.Amount.Amount).TruncateInt()
}

// position size takes 0 decimal although price takes 6 decimal (micro unit)
func (m PerpetualFuturesPosition) MarginMaintenanceRate(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Dec {
	// 証拠金維持率 = 有効証拠金(USD単位) ÷ 必要証拠金(USD単位)
	return sdk.NewDecFromInt(m.EffectiveMarginInMetrics(baseMetricsRate, quoteMetricsRate)).Quo(sdk.NewDecFromInt(m.RequiredMarginInMetrics(baseMetricsRate, quoteMetricsRate)))
}

func (m PerpetualFuturesPosition) RemainingMarginInBase(baseMetricsRate MetricsRateType) sdk.Int {
	// 残存証拠金(USD単位) = 残存証拠金(base単位) * 現在のbase/USDレート
	return sdk.NewDecFromInt(m.RemainingMargin.Amount).Mul(baseMetricsRate.Amount.Amount).TruncateInt()
}

func (m PerpetualFuturesPosition) RemainingMarginInQuote(quoteMetricsRate MetricsRateType) sdk.Int {
	// 残存証拠金(USD単位) = 残存証拠金(quote単位) * 現在のquote/USDレート
	return sdk.NewDecFromInt(m.RemainingMargin.Amount).Mul(quoteMetricsRate.Amount.Amount).TruncateInt()
}

func (m PerpetualFuturesPosition) RemainingMarginInMetrics(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Int {
	// 残存証拠金(USD単位) = 残存証拠金(base単位) * 現在のbase/USDレート
	//                    = 残存証拠金(quote単位) * 現在のquote/USDレート
	if m.RemainingMargin.Denom == m.Market.BaseDenom {
		return m.RemainingMarginInBase(baseMetricsRate)
	} else if m.RemainingMargin.Denom == m.Market.QuoteDenom {
		return m.RemainingMarginInQuote(quoteMetricsRate)
	} else {
		panic("not supported denom")
	}
}

func (m PerpetualFuturesPosition) EffectiveMarginInMetrics(baseMetricsRate, quoteMetricsRate MetricsRateType) sdk.Int {
	// 有効証拠金(USD単位) = 残存証拠金(USD単位) + 損益(USD単位)
	return m.RemainingMarginInMetrics(baseMetricsRate, quoteMetricsRate).Add(m.ProfitAndLossInMetrics(baseMetricsRate, quoteMetricsRate))
}

func NewMetricsRateType(unit string, denom string, amount sdk.Dec) MetricsRateType {
	return MetricsRateType{
		MetricsUnit: unit,
		Amount:      sdk.NewDecCoinFromDec(denom, amount),
	}
}

type MetricsRateType struct {
	MetricsUnit string
	Amount      sdk.DecCoin
}
