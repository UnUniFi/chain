package keeper

import (
	"fmt"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

// fixme: it has not been tested
// todo:rename GetCurrentPrice to GetCurrentUsdPrice
func (k Keeper) GetCurrentPrice(ctx sdk.Context, denom string) (sdk.Dec, error) {
	ticker, err := k.pricefeedKeeper.GetTicker(ctx, denom)
	if err != nil {
		return sdk.Dec{}, err
	}
	rate, err := k.GetPrice(ctx, ticker, "usd")
	if err != nil {
		return sdk.Dec{}, err
	}
	return rate.Price, nil
}

func (k Keeper) GetPairUsdPrice(ctx sdk.Context, base, quote string) (sdk.Dec, sdk.Dec, error) {
	baseUsdPrice, err := k.GetCurrentPrice(ctx, base)
	if err != nil {
		return sdk.Dec{}, sdk.Dec{}, err
	}
	quoteUsdPrice, err := k.GetCurrentPrice(ctx, quote)
	if err != nil {
		return sdk.Dec{}, sdk.Dec{}, err
	}
	return baseUsdPrice, quoteUsdPrice, nil
}

func (k Keeper) GetPairUsdPriceFromMarket(ctx sdk.Context, market types.Market) (sdk.Dec, sdk.Dec, error) {
	return k.GetPairUsdPrice(ctx, market.BaseDenom, market.QuoteDenom)
}

func (k Keeper) OpenPerpetualFuturesPosition(ctx sdk.Context, positionId string, sender string, margin sdk.Coin, market types.Market, positionInstance types.PerpetualFuturesPositionInstance) (*types.Position, error) {
	// Get base and quote price in quote ticker of the pool, which is "usd"
	openedBaseRate, err := k.GetCurrentPrice(ctx, market.BaseDenom)
	if err != nil {
		return nil, err
	}

	openedQuoteRate, err := k.GetCurrentPrice(ctx, market.QuoteDenom)
	if err != nil {
		return nil, err
	}

	// NOTE: To be consistent with the other numbers, we should use the micro unit for the size
	any, err := codecTypes.NewAnyWithValue(&positionInstance)
	if err != nil {
		return nil, err
	}

	position := types.Position{
		Id:               positionId,
		Market:           market,
		Address:          sender,
		OpenedAt:         ctx.BlockTime(),
		OpenedHeight:     uint64(ctx.BlockHeight()),
		OpenedBaseRate:   openedBaseRate,
		OpenedQuoteRate:  openedQuoteRate,
		PositionInstance: *any,
		RemainingMargin:  margin,
	}

	// General validation for the position creation
	params := k.GetParams(ctx)

	var reserveCoinDenom string
	if positionInstance.PositionType == types.PositionType_LONG {
		reserveCoinDenom = market.BaseDenom
	} else {
		reserveCoinDenom = market.QuoteDenom
	}

	availableAssetInPoolByDenom, err := k.AvailableAssetInPool(ctx, reserveCoinDenom)
	if err != nil {
		return nil, err
	}

	if err := position.IsValid(params, availableAssetInPoolByDenom); err != nil {
		return nil, err
	}

	switch positionInstance.PositionType {
	// FIXME: Don't use OneMillionInt derectly to make it decimal unit. issue #476
	case types.PositionType_LONG:
		k.AddPerpetualFuturesGrossPositionOfMarket(ctx, market, positionInstance.PositionType, positionInstance.SizeInDenomExponent(types.OneMillionInt))
		// Reserve tokens to pay profit
		if err := k.AddReserveTokensForPosition(ctx, positionInstance.SizeInDenomExponent(types.OneMillionInt), position.Market.BaseDenom); err != nil {
			return nil, err
		}
	case types.PositionType_SHORT:
		k.AddPerpetualFuturesGrossPositionOfMarket(ctx, market, positionInstance.PositionType, positionInstance.SizeInDenomExponent(types.OneMillionInt))
		// Reserve tokens to pay profit
		if err := k.AddReserveTokensForPosition(ctx, positionInstance.SizeInDenomExponent(types.OneMillionInt), position.Market.QuoteDenom); err != nil {
			return nil, err
		}
	case types.PositionType_POSITION_UNKNOWN:
		return nil, fmt.Errorf("unknown position type")
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionOpened{
		Sender:     sender,
		PositionId: positionId,
	})

	return &position, nil
}

// AddReserveTokensForPosition adds the tokens o the amount of the popsition size to pay the maximum profit
// in reserved property of the PoolMarketCap
func (k Keeper) AddReserveTokensForPosition(ctx sdk.Context, positionSizeInDenomExponent sdk.Int, denom string) error {
	reserveOld, err := k.GetReservedCoin(ctx, denom)
	if err != nil {
		return err
	}

	reserveNew := reserveOld.AddAmount(positionSizeInDenomExponent)
	if err := k.SetReservedCoin(ctx, reserveNew); err != nil {
		return err
	}
	return nil
}

// SubReserveTokensForPosition subtracts the tokens o the amount of the popsition size to pay the maximum profit
// in reserved property of the PoolMarketCap
func (k Keeper) SubReserveTokensForPosition(ctx sdk.Context, positionSizeInDenomExponent sdk.Int, denom string) error {
	reserveOld, err := k.GetReservedCoin(ctx, denom)
	if err != nil {
		return err
	}

	reserveNew := reserveOld.SubAmount(positionSizeInDenomExponent)
	if err := k.SetReservedCoin(ctx, reserveNew); err != nil {
		return err
	}

	return nil
}

func (k Keeper) ClosePerpetualFuturesPosition(ctx sdk.Context, position types.PerpetualFuturesPosition) error {
	// params := k.GetParams(ctx)
	// commissionRate := params.PerpetualFutures.CommissionRate
	// Set the ClosePosition commission rate to 0. The commission will be deducted by Levy instead.
	commissionRate := sdk.MustNewDecFromStr("0")

	// At closing the position, the trading fee is deducted.
	// fee = positionSize * commissionRate
	positionSizeInDenomUnit := sdk.NewDecFromInt(position.PositionInstance.SizeInDenomExponent(types.OneMillionInt))
	feeAmountDec := positionSizeInDenomUnit.Mul(commissionRate)
	tradeAmount := positionSizeInDenomUnit.Sub(feeAmountDec)
	feeAmount := feeAmountDec.RoundInt()

	baseUsdPrice, err := k.GetCurrentPrice(ctx, position.Market.BaseDenom)
	if err != nil {
		return err
	}
	quoteUsdPrice, err := k.GetCurrentPrice(ctx, position.Market.QuoteDenom)
	if err != nil {
		return err
	}

	quoteTicker := k.GetPoolQuoteTicker(ctx)
	baseMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.BaseDenom, baseUsdPrice)
	quoteMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.QuoteDenom, quoteUsdPrice)
	pnlAmount := position.ProfitAndLoss(baseMetricsRate, quoteMetricsRate)

	returningAmount, err := k.HandleReturnAmount(ctx, pnlAmount, position)
	if err != nil {
		return err
	}

	// TODO: Fix position size in total by removing the closing position
	switch position.PositionInstance.PositionType {
	// FIXME: Don't use OneMillionInt derectly to make it decimal unit. issue #476
	case types.PositionType_LONG:
		k.SubPerpetualFuturesGrossPositionOfMarket(ctx, position.Market, position.PositionInstance.PositionType, position.PositionInstance.SizeInDenomExponent(types.OneMillionInt))
		// break
	case types.PositionType_SHORT:
		k.SubPerpetualFuturesGrossPositionOfMarket(ctx, position.Market, position.PositionInstance.PositionType, position.PositionInstance.SizeInDenomExponent(types.OneMillionInt))
		// break
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionClosed{
		Sender:          position.Address,
		PositionId:      position.Id,
		FeeAmount:       feeAmount.String(),
		TradeAmount:     tradeAmount.String(),
		ReturningAmount: returningAmount.String(),
	})

	return nil
}

// If the profit exists, the profit always comes from the pool.
// If the loss exists, the loss always goes to the pool from the users' margin.
func (k Keeper) HandleReturnAmount(ctx sdk.Context, pnlAmount sdk.Int, position types.PerpetualFuturesPosition) (returningAmount sdk.Int, err error) {
	if pnlAmount.IsNegative() {
		returningAmount = position.RemainingMargin.Amount.Sub(pnlAmount.Abs())
		// Tell the loss to the LP happened by a trade
		// This has to be restricted by the protocol behavior in the future
		if !(returningAmount.IsNegative()) {
			_ = ctx.EventManager().EmitTypedEvent(&types.EventLossToLP{
				PositionId: position.Id,
				LossAmount: returningAmount.String(),
			})
		} else {
			returningCoin := sdk.NewCoin(position.RemainingMargin.Denom, returningAmount)
			// Send coin including margin
			if err := k.SendBackMargin(ctx, position.Address.AccAddress(), sdk.NewCoins(returningCoin)); err != nil {
				return sdk.ZeroInt(), err
			}

			// Send Loss of the position to the pool
			if err := k.SendCoinFromMarginManagerToPool(ctx, sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, pnlAmount))); err != nil {
				return sdk.ZeroInt(), err
			}
		}
	} else {
		returningAmount = position.RemainingMargin.Amount.Add(pnlAmount)
		fromMarginManagerAmount := position.RemainingMargin
		if err := k.SendBackMargin(ctx, position.Address.AccAddress(), sdk.NewCoins(fromMarginManagerAmount)); err != nil {
			return sdk.ZeroInt(), err
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, position.Address.AccAddress(), sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, pnlAmount))); err != nil {
			return sdk.ZeroInt(), err
		}
	}

	return returningAmount, nil
}

func (k Keeper) ReportLiquidationNeededPerpetualFuturesPosition(ctx sdk.Context, rewardRecipient string, position types.PerpetualFuturesPosition) error {
	params := k.GetParams(ctx)

	currentBaseUsdRate, currentQuoteUsdRate, err := k.GetPairUsdPriceFromMarket(ctx, position.Market)
	if err != nil {
		panic(err)
	}

	quoteTicker := k.GetPoolQuoteTicker(ctx)
	baseMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.BaseDenom, currentBaseUsdRate)
	quoteMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.QuoteDenom, currentQuoteUsdRate)
	if position.NeedLiquidation(params.PerpetualFutures.MarginMaintenanceRate, baseMetricsRate, quoteMetricsRate) {
		// In case of closing position by Liquidation, a commission fee is charged.
		commissionFee := sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(params.PerpetualFutures.CommissionRate).RoundInt()
		position.RemainingMargin.Amount = position.RemainingMargin.Amount.Sub(commissionFee)
		rewardAmount := sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(params.PoolParams.ReportLiquidationRewardRate).RoundInt()
		reward := sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, rewardAmount))

		if err := k.ClosePerpetualFuturesPosition(ctx, position); err != nil {
			return err
		}
		// Delete Position
		k.DeletePosition(ctx, position.Address, position.Id)
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rewardRecipient.AccAddress(), reward)
		if err != nil {
			return err
		}

		_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionLiquidated{
			RewardRecipient: rewardRecipient,
			PositionId:      position.Id,
			RemainingMargin: position.RemainingMargin.String(),
			RewardAmount:    rewardAmount.String(),
		})
		return nil
	}

	return nil
}

func (k Keeper) ReportLevyPeriodPerpetualFuturesPosition(ctx sdk.Context, rewardRecipient string, position types.Position, positionInstance types.PerpetualFuturesPositionInstance) error {
	params := k.GetParams(ctx)

	netPosition := k.GetPerpetualFuturesNetPositionOfMarket(ctx, position.Market).PositionSizeInDenomExponent
	totalPosition := k.GetPerpetualFuturesTotalPositionOfMarket(ctx, position.Market).PositionSizeInDenomExponent
	commissionFee := sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(params.PerpetualFutures.CommissionRate).RoundInt()

	// NetPosition / TotalPosition * LevyCoefficient
	imaginaryFundingRate := sdk.NewDecFromInt(netPosition).Quo(sdk.NewDecFromInt(totalPosition)).Mul(params.PerpetualFutures.ImaginaryFundingRateProportionalCoefficient)
	imaginaryFundingBaseFee := sdk.NewDecFromInt(positionInstance.SizeInDenomExponent(types.OneMillionInt)).Mul(imaginaryFundingRate).RoundInt()
	var imaginaryFundingFee sdk.Int
	if position.Market.BaseDenom == position.RemainingMargin.Denom {
		imaginaryFundingFee = imaginaryFundingBaseFee
	} else {
		imaginaryFundingFee = k.ConvertBaseAmountToQuoteAmount(ctx, position.Market, imaginaryFundingBaseFee)
	}
	if positionInstance.PositionType == types.PositionType_LONG {
		position.RemainingMargin.Amount = position.RemainingMargin.Amount.Sub(imaginaryFundingFee).Sub(commissionFee)

	} else {
		position.RemainingMargin.Amount = position.RemainingMargin.Amount.Add(imaginaryFundingFee).Sub(commissionFee)
	}
	// Tranfer the fees from pool to manager or manager to pool appropriately
	// to keep the remaining margin of the position match the actual number to the balance
	if err := k.HandleImaginaryFundingFeeTransfer(ctx, imaginaryFundingFee, commissionFee, positionInstance.PositionType, position.RemainingMargin.Denom); err != nil {
		return err
	}

	position.LastLeviedAt = ctx.BlockTime()

	// Reward is part of the commission fee
	rewardAmount := sdk.NewDecFromInt(commissionFee).Mul(params.PoolParams.ReportLevyPeriodRewardRate).RoundInt()

	reward := sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, rewardAmount))
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rewardRecipient.AccAddress(), reward)
	if err != nil {
		return err
	}

	k.SetPosition(ctx, position)

	ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionLevied{
		RewardRecipient: rewardRecipient.AccAddress().String(),
		PositionId:      position.Id,
		RemainingMargin: position.RemainingMargin.String(),
		RewardAmount:    rewardAmount.String(),
	})

	return nil
}

func (k Keeper) HandleImaginaryFundingFeeTransfer(ctx sdk.Context, imarginaryFundingFee, commissionFee sdk.Int, positionType types.PositionType, denom string) error {
	var totalFee sdk.Int
	if positionType == types.PositionType_LONG {
		totalFee = imarginaryFundingFee.Add(commissionFee)
	} else {
		totalFee = commissionFee.Sub(imarginaryFundingFee)
	}

	if totalFee.IsPositive() {
		if err := k.SendCoinFromMarginManagerToPool(ctx, sdk.NewCoins(sdk.NewCoin(denom, totalFee))); err != nil {
			return err
		}
	} else {
		if err := k.SendCoinFromPoolToMarginManager(ctx, sdk.NewCoins(sdk.NewCoin(denom, totalFee.Abs()))); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) GetPerpetualFuturesGrossPositionOfMarket(ctx sdk.Context, market types.Market, positionType types.PositionType) types.PerpetualFuturesGrossPositionOfMarket {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.DenomGrossPositionPerpetualFuturesKeyPrefix(market, positionType))
	if bz == nil {
		return types.NewPerpetualFuturesGrossPositionOfMarket(
			market,
			positionType,
			sdk.ZeroInt(),
		)
	}

	grossPositionOfMarket := types.PerpetualFuturesGrossPositionOfMarket{}
	k.cdc.MustUnmarshal(bz, &grossPositionOfMarket)
	return grossPositionOfMarket
}

/// GetPositionSizeOfGrossPositionOfMarket is not used anymore.
/// This can be deleted.
// func (k Keeper) GetPositionSizeOfGrossPositionOfMarket(ctx sdk.Context, market types.Market) sdk.Int {
// 	position := k.GetPerpetualFuturesGrossPositionOfMarket(ctx, market, )
// 	if position.PositionSizeInDenomUnit.IsNil() {
// 		return sdk.ZeroInt()
// 	}
// 	return position.PositionSizeInDenomUnit
// }

func (k Keeper) GetAllPerpetualFuturesGrossPositionOfMarket(ctx sdk.Context) []types.PerpetualFuturesGrossPositionOfMarket {
	store := ctx.KVStore(k.storeKey)

	perpetualFuturesGrossPositionOfMarkets := []types.PerpetualFuturesGrossPositionOfMarket{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixPerpetualFutures))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		grossPositionOfMarket := types.PerpetualFuturesGrossPositionOfMarket{}
		k.cdc.MustUnmarshal(it.Value(), &grossPositionOfMarket)

		perpetualFuturesGrossPositionOfMarkets = append(
			perpetualFuturesGrossPositionOfMarkets,
			grossPositionOfMarket,
		)
	}
	return perpetualFuturesGrossPositionOfMarkets
}

func (k Keeper) SetPerpetualFuturesGrossPositionOfMarket(ctx sdk.Context, grossPositionOfMarket types.PerpetualFuturesGrossPositionOfMarket) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&grossPositionOfMarket)

	store.Set(types.DenomGrossPositionPerpetualFuturesKeyPrefix(grossPositionOfMarket.Market, grossPositionOfMarket.PositionType), bz)
}

// Call AddPerpetualFuturesGrossPositionOfMarket when the position is created.
func (k Keeper) AddPerpetualFuturesGrossPositionOfMarket(ctx sdk.Context, market types.Market, positionType types.PositionType, rhs sdk.Int) {
	perpFutureGrossPositionOfMarket := k.GetPerpetualFuturesGrossPositionOfMarket(ctx, market, positionType)
	perpFutureGrossPositionOfMarket.PositionSizeInDenomExponent = perpFutureGrossPositionOfMarket.PositionSizeInDenomExponent.Add(rhs)

	k.SetPerpetualFuturesGrossPositionOfMarket(ctx, perpFutureGrossPositionOfMarket)
}

// Call AddPerpetualFuturesGrossPositionOfMarket when the position is closed.
func (k Keeper) SubPerpetualFuturesGrossPositionOfMarket(ctx sdk.Context, market types.Market, positionType types.PositionType, rhs sdk.Int) {
	perpFutureGrossPositionOfMarket := k.GetPerpetualFuturesGrossPositionOfMarket(ctx, market, positionType)
	perpFutureGrossPositionOfMarket.PositionSizeInDenomExponent = perpFutureGrossPositionOfMarket.PositionSizeInDenomExponent.Sub(rhs)

	k.SetPerpetualFuturesGrossPositionOfMarket(ctx, perpFutureGrossPositionOfMarket)
}

func (k Keeper) GetPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market) types.PerpetualFuturesGrossPositionOfMarket {
	grossPositionLong := k.GetPerpetualFuturesGrossPositionOfMarket(ctx, market, types.PositionType_LONG).PositionSizeInDenomExponent
	grossPositionShort := k.GetPerpetualFuturesGrossPositionOfMarket(ctx, market, types.PositionType_SHORT).PositionSizeInDenomExponent
	return types.NewPerpetualFuturesGrossPositionOfMarket(
		market,
		types.PositionType_POSITION_UNKNOWN,
		grossPositionLong.Sub(grossPositionShort),
	)
}

func (k Keeper) GetPerpetualFuturesTotalPositionOfMarket(ctx sdk.Context, market types.Market) types.PerpetualFuturesGrossPositionOfMarket {
	grossPositionLong := k.GetPerpetualFuturesGrossPositionOfMarket(ctx, market, types.PositionType_LONG).PositionSizeInDenomExponent
	grossPositionShort := k.GetPerpetualFuturesGrossPositionOfMarket(ctx, market, types.PositionType_SHORT).PositionSizeInDenomExponent
	return types.NewPerpetualFuturesGrossPositionOfMarket(
		market,
		types.PositionType_POSITION_UNKNOWN,
		grossPositionLong.Add(grossPositionShort),
	)
}

func (k Keeper) ConvertBaseAmountToQuoteAmount(ctx sdk.Context, market types.Market, amount sdk.Int) sdk.Int {
	currentBaseUsdRate, currentQuoteUsdRate, _ := k.GetPairUsdPriceFromMarket(ctx, market)
	quoteTicker := k.GetPoolQuoteTicker(ctx)
	baseMetricsRate := types.NewMetricsRateType(quoteTicker, market.BaseDenom, currentBaseUsdRate)
	quoteMetricsRate := types.NewMetricsRateType(quoteTicker, market.QuoteDenom, currentQuoteUsdRate)

	return sdk.NewDecFromInt(amount).Mul(quoteMetricsRate.Amount.Amount).Quo(baseMetricsRate.Amount.Amount).RoundInt()
}
