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
		Id:                   positionId,
		Market:               market,
		Address:              sender,
		OpenedAt:             ctx.BlockTime(),
		OpenedHeight:         uint64(ctx.BlockHeight()),
		OpenedBaseRate:       openedBaseRate,
		OpenedQuoteRate:      openedQuoteRate,
		RemainingMargin:      margin,
		LeviedAmount:         sdk.NewInt64Coin(margin.Denom, 0),
		LeviedAmountNegative: true,
		LastLeviedAt:         ctx.BlockTime(),
		PositionInstance:     *any,
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
		if err := k.AddReserveTokensForPosition(ctx, types.MarketType_FUTURES, positionInstance.SizeInDenomExponent(types.OneMillionInt), position.Market.BaseDenom); err != nil {
			return nil, err
		}
	case types.PositionType_SHORT:
		k.AddPerpetualFuturesGrossPositionOfMarket(ctx, market, positionInstance.PositionType, positionInstance.SizeInDenomExponent(types.OneMillionInt))
		// Reserve tokens to pay profit
		if err := k.AddReserveTokensForPosition(ctx, types.MarketType_FUTURES, positionInstance.SizeInDenomExponent(types.OneMillionInt), position.Market.QuoteDenom); err != nil {
			return nil, err
		}
	case types.PositionType_POSITION_UNKNOWN:
		return nil, fmt.Errorf("unknown position type")
	}

	senderAccAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return nil, err
	}

	if err := k.SendMarginToMarginManager(ctx, senderAccAddr, sdk.NewCoins(margin)); err != nil {
		return nil, err
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionOpened{
		Sender:     sender,
		PositionId: positionId,
	})

	return &position, nil
}

// AddReserveTokensForPosition adds the tokens o the amount of the position size to pay the maximum profit
// in reserved property of the PoolMarketCap
func (k Keeper) AddReserveTokensForPosition(ctx sdk.Context, marketType types.MarketType, positionSizeInDenomExponent sdk.Int, denom string) error {
	reserveOld, err := k.GetReservedCoin(ctx, marketType, denom)
	if err != nil {
		return err
	}

	reserveNew := reserveOld.Amount.AddAmount(positionSizeInDenomExponent)

	if err := k.SetReservedCoin(ctx, types.NewReserve(marketType, reserveNew)); err != nil {
		return err
	}
	return nil
}

// SubReserveTokensForPosition subtracts the tokens o the amount of the position size to pay the maximum profit
// in reserved property of the PoolMarketCap
func (k Keeper) SubReserveTokensForPosition(ctx sdk.Context, marketType types.MarketType, positionSizeInDenomExponent sdk.Int, denom string) error {
	reserveOld, err := k.GetReservedCoin(ctx, marketType, denom)
	if err != nil {
		return err
	}

	reserveNew := reserveOld.Amount.SubAmount(positionSizeInDenomExponent)

	if err := k.SetReservedCoin(ctx, types.NewReserve(marketType, reserveNew)); err != nil {
		return err
	}

	return nil
}

func (k Keeper) ClosePerpetualFuturesPosition(ctx sdk.Context, position types.PerpetualFuturesPosition) error {
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

	// profit or loss amount in margin denom
	pnlAmount := position.ProfitAndLoss(baseMetricsRate, quoteMetricsRate)
	if position.LeviedAmountNegative {
		pnlAmount = pnlAmount.Sub(position.LeviedAmount.Amount)
	} else {
		pnlAmount = pnlAmount.Add(position.LeviedAmount.Amount)
	}

	returningAmount, err := k.HandleReturnAmount(ctx, pnlAmount, position)
	if err != nil {
		return err
	}

	switch position.PositionInstance.PositionType {
	// FIXME: Don't use OneMillionInt directly to make it decimal unit. issue #476
	case types.PositionType_LONG:
		k.SubPerpetualFuturesGrossPositionOfMarket(ctx, position.Market, position.PositionInstance.PositionType, position.PositionInstance.SizeInDenomExponent(types.OneMillionInt))
		// Sub reserve tokens of pool
		if err := k.SubReserveTokensForPosition(ctx, types.MarketType_FUTURES, position.PositionInstance.SizeInDenomExponent(types.OneMillionInt), position.Market.BaseDenom); err != nil {
			return err
		}
	case types.PositionType_SHORT:
		k.SubPerpetualFuturesGrossPositionOfMarket(ctx, position.Market, position.PositionInstance.PositionType, position.PositionInstance.SizeInDenomExponent(types.OneMillionInt))
		// Sub reserve tokens of pool
		if err := k.SubReserveTokensForPosition(ctx, types.MarketType_FUTURES, position.PositionInstance.SizeInDenomExponent(types.OneMillionInt), position.Market.QuoteDenom); err != nil {
			return err
		}
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionClosed{
		Sender:          position.Address,
		PositionId:      position.Id,
		PositionSize:    position.PositionInstance.SizeInDenomExponent(types.OneMillionInt).String(),
		PnlAmount:       pnlAmount.String(),
		ReturningAmount: returningAmount.String(),
	})

	return nil
}

// If the profit exists, the profit always comes from the pool.
// If the loss exists, the loss always goes to the pool from the users' margin.
func (k Keeper) HandleReturnAmount(ctx sdk.Context, pnlAmount sdk.Int, position types.PerpetualFuturesPosition) (returningAmount sdk.Int, err error) {
	addr, err := sdk.AccAddressFromBech32(position.Address)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	if pnlAmount.IsNegative() {
		loss := pnlAmount.Abs()
		returningAmount = position.RemainingMargin.Amount.Sub(loss)

		if returningAmount.IsNegative() {
			_ = ctx.EventManager().EmitTypedEvent(&types.EventLossToLP{
				PositionId: position.Id,
				LossAmount: returningAmount.String(),
			})
			// Send margin to the pool from MarginManager
			// The loss is taken by the pool
			if err := k.SendCoinFromMarginManagerToPool(ctx, sdk.NewCoins(position.RemainingMargin)); err != nil {
				return sdk.ZeroInt(), err
			}
		} else {
			returningCoin := sdk.NewCoin(position.RemainingMargin.Denom, returningAmount)
			// Send margin-loss from MarginManager
			if err := k.SendBackMargin(ctx, addr, sdk.NewCoins(returningCoin)); err != nil {
				return sdk.ZeroInt(), err
			}
			// Send loss to the pool
			if err := k.SendCoinFromMarginManagerToPool(ctx, sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, loss))); err != nil {
				return sdk.ZeroInt(), err
			}
		}
	} else {
		returningAmount = position.RemainingMargin.Amount.Add(pnlAmount)
		// Send margin from MarginManager & profit from the pool
		if err := k.SendBackMargin(ctx, addr, sdk.NewCoins(position.RemainingMargin)); err != nil {
			return sdk.ZeroInt(), err
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, pnlAmount))); err != nil {
			return sdk.ZeroInt(), err
		}
	}

	return returningAmount, nil
}

// LiquidateFuturesPosition is called if a position is needed to be liquidated.
// In fact, this function executes the liquidation operation, which closes the position,
// sends the liquidation reward to the rewardRecipient, and sends the rest of the margin back to the position owner.
func (k Keeper) LiquidateFuturesPosition(ctx sdk.Context, rewardRecipient string, position types.PerpetualFuturesPosition, commissionRate, rewardRate sdk.Dec) error {
	// In case of closing position by Liquidation, a commission fee is charged.
	commissionBaseFee := sdk.NewDecFromInt(position.PositionInstance.SizeInDenomExponent(types.OneMillionInt)).Mul(commissionRate).RoundInt()
	var commissionFee sdk.Int
	if position.Market.BaseDenom == position.RemainingMargin.Denom {
		commissionFee = commissionBaseFee
	} else {
		commissionFee = k.ConvertBaseAmountToQuoteAmount(ctx, position.Market, commissionBaseFee)
	}
	if position.LeviedAmountNegative {
		position.LeviedAmount.Amount = position.LeviedAmount.Amount.Add(commissionFee)
	} else {
		rest := position.LeviedAmount.Amount.Sub(commissionFee)
		if rest.IsNegative() {
			position.LeviedAmountNegative = true
			position.LeviedAmount.Amount = rest.Abs()
		} else {
			position.LeviedAmount.Amount = rest
		}
	}

	if err := k.ClosePerpetualFuturesPosition(ctx, position); err != nil {
		return err
	}

	// Delete Position
	positionAddress, err := sdk.AccAddressFromBech32(position.Address)
	if err != nil {
		return err
	}
	k.DeletePosition(ctx, positionAddress, position.Id)

	// Send Reward
	rewardAmount := sdk.NewDecFromInt(commissionFee).Mul(rewardRate).RoundInt()
	err = k.SendRewardFromCommission(ctx, rewardAmount, position.RemainingMargin.Denom, rewardRecipient)
	if err != nil {
		return err
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionLiquidated{
		RewardRecipient: rewardRecipient,
		PositionId:      position.Id,
		RemainingMargin: position.RemainingMargin.String(),
		RewardAmount:    rewardAmount.String(),
		LeviedAmount:    position.LeviedAmount.String(),
	})

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesLiquidationFee{
		Fee:        sdk.NewCoin(position.RemainingMargin.Denom, commissionFee),
		PositionId: position.Id,
	})

	return nil
}

func (k Keeper) ReportLevyPeriodPerpetualFuturesPosition(ctx sdk.Context, rewardRecipient string, position types.Position, positionInstance types.PerpetualFuturesPositionInstance) error {
	params := k.GetParams(ctx)

	netPosition := k.GetPerpetualFuturesNetPositionOfMarket(ctx, position.Market).PositionSizeInDenomExponent
	totalPosition := k.GetPerpetualFuturesTotalPositionOfMarket(ctx, position.Market).PositionSizeInDenomExponent
	commissionBaseFee := sdk.NewDecFromInt(positionInstance.SizeInDenomExponent(types.OneMillionInt)).Mul(params.PerpetualFutures.CommissionRate).RoundInt()
	// NetPosition / TotalPosition * LevyCoefficient
	imaginaryFundingRate := sdk.NewDecFromInt(netPosition).Quo(sdk.NewDecFromInt(totalPosition)).Mul(params.PerpetualFutures.ImaginaryFundingRateProportionalCoefficient)
	imaginaryFundingBaseFee := sdk.NewDecFromInt(positionInstance.SizeInDenomExponent(types.OneMillionInt)).Mul(imaginaryFundingRate).RoundInt()
	var commissionFee sdk.Int
	var imaginaryFundingFee sdk.Int
	if position.Market.BaseDenom == position.RemainingMargin.Denom {
		commissionFee = commissionBaseFee
		imaginaryFundingFee = imaginaryFundingBaseFee
	} else {
		commissionFee = k.ConvertBaseAmountToQuoteAmount(ctx, position.Market, commissionBaseFee)
		imaginaryFundingFee = k.ConvertBaseAmountToQuoteAmount(ctx, position.Market, imaginaryFundingBaseFee)
	}
	var totalFee sdk.Int
	if positionInstance.PositionType == types.PositionType_LONG {
		totalFee = commissionFee.Add(imaginaryFundingFee)
	} else {
		totalFee = commissionFee.Sub(imaginaryFundingFee)
	}
	if position.LeviedAmountNegative {
		position.LeviedAmount.Amount = position.LeviedAmount.Amount.Add(totalFee)
	} else {
		rest := position.LeviedAmount.Amount.Sub(totalFee)
		if rest.IsNegative() {
			position.LeviedAmountNegative = true
			position.LeviedAmount.Amount = rest.Abs()
		} else {
			position.LeviedAmount.Amount = rest
		}
	}

	position.LastLeviedAt = ctx.BlockTime()

	err := k.SetPosition(ctx, position)
	if err != nil {
		return err
	}

	// Send Reward
	rewardAmount := sdk.NewDecFromInt(commissionFee).Mul(params.PoolParams.ReportLevyPeriodRewardRate).RoundInt()
	err = k.SendRewardFromCommission(ctx, rewardAmount, position.RemainingMargin.Denom, rewardRecipient)
	if err != nil {
		return err
	}

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionLevied{
		RewardRecipient: rewardRecipient,
		PositionId:      position.Id,
		RemainingMargin: position.RemainingMargin.String(),
		RewardAmount:    rewardAmount.String(),
		LeviedAmount:    position.LeviedAmount.String(),
	})

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesLevyFee{
		Fee:        sdk.NewCoin(position.RemainingMargin.Denom, commissionFee),
		PositionId: position.Id,
	})

	_ = ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesImaginaryFundingFee{
		Fee:         sdk.NewCoin(position.RemainingMargin.Denom, imaginaryFundingFee.Abs()),
		FeeNegative: imaginaryFundingFee.IsNegative(),
		PositionId:  position.Id,
	})

	return nil
}

func (k Keeper) SendRewardFromCommission(ctx sdk.Context, rewardAmount sdk.Int, denom string, recipientAddr string) error {
	recipient, err := sdk.AccAddressFromBech32(recipientAddr)
	if err != nil {
		return err
	}

	reward := sdk.NewCoins(sdk.NewCoin(denom, rewardAmount))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, reward)
	if err != nil {
		return err
	}
	return nil
}

// func (k Keeper) HandleImaginaryFundingFeeTransfer(ctx sdk.Context, imaginaryFundingFee, commissionFee sdk.Int, positionType types.PositionType, denom string) error {
// 	var totalFee sdk.Int
// 	if positionType == types.PositionType_LONG {
// 		totalFee = imaginaryFundingFee.Add(commissionFee)
// 	} else {
// 		totalFee = commissionFee.Sub(imaginaryFundingFee)
// 	}

// 	if totalFee.IsPositive() {
// 		if err := k.SendCoinFromMarginManagerToPool(ctx, sdk.NewCoins(sdk.NewCoin(denom, totalFee))); err != nil {
// 			return err
// 		}
// 	} else {
// 		if err := k.SendCoinFromPoolToMarginManager(ctx, sdk.NewCoins(sdk.NewCoin(denom, totalFee.Abs()))); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

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

func (k Keeper) GetPerpetualFuturesPositionSizeInMetrics(ctx sdk.Context, market types.Market, pType types.PositionType) sdk.Dec {
	perpFuturesLongPositionNum := k.GetPerpetualFuturesGrossPositionOfMarket(ctx, market, pType)
	currentBaseUsdRate, currentQuoteUsdRate, err := k.GetPairUsdPriceFromMarket(ctx, market)
	if err != nil {
		return sdk.ZeroDec()
	}
	baseDenomPrice := currentBaseUsdRate.Quo(currentQuoteUsdRate)
	// baseDenomPrice, err := k.GetPrice(ctx, market.BaseDenom, market.QuoteDenom)
	if err != nil {
		return sdk.ZeroDec()
	}
	return sdk.NewDecFromInt(perpFuturesLongPositionNum.PositionSizeInDenomExponent).Mul(baseDenomPrice)
}
