package keeper

import (
	"errors"
	"fmt"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifiTypes "github.com/UnUniFi/chain/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

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

func (k Keeper) OpenPerpetualFuturesPosition(ctx sdk.Context, positionId string, sender ununifiTypes.StringAccAddress, margin sdk.Coin, market types.Market, positionInstance types.PerpetualFuturesPositionInstance) (*types.Position, error) {
	openedBaseRate, err := k.GetCurrentPrice(ctx, market.BaseDenom)
	if err != nil {
		return nil, err
	}

	openedQuoteRate, err := k.GetCurrentPrice(ctx, market.BaseDenom)
	if err != nil {
		return nil, err
	}
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

	switch positionInstance.PositionType {
	case types.PositionType_LONG:
		k.AddPerpetualFuturesNetPositionOfMarket(ctx, market, positionInstance.Size_)
		break
	case types.PositionType_SHORT:
		k.SubPerpetualFuturesNetPositionOfMarket(ctx, market, positionInstance.Size_)
		break
	case types.PositionType_POSITION_UNKNOWN:
		return nil, fmt.Errorf("unknown position type")
	}

	ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionOpened{
		Sender:     sender.AccAddress().String(),
		PositionId: positionId,
	})

	return &position, nil
}

func (k Keeper) ClosePerpetualFuturesPosition(ctx sdk.Context, position types.Position, positionInstance types.PerpetualFuturesPositionInstance) error {
	params := k.GetParams(ctx)
	commissionRate := params.PerpetualFutures.CommissionRate
	feeAmount := positionInstance.Size_.Mul(commissionRate)
	tradeAmount := positionInstance.Size_.Sub(feeAmount)

	// todo: fixme
	openedRate := position.OpenedBaseRate
	closedRate, err := k.GetPairRate(ctx, position.Market)
	if err != nil {
		return err
	}

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, position.Address.AccAddress(), types.ModuleName, sdk.Coins{sdk.NewCoin(position.Market.BaseDenom, feeAmount.RoundInt())}) // TODO: this is wrong.

	principal := positionInstance.CalculatePrincipal()
	amountToUser := sdk.Dec{}

	switch positionInstance.PositionType {
	case types.PositionType_LONG:
		k.SubPerpetualFuturesNetPositionOfMarket(ctx, position.Market, tradeAmount)

		if closedRate.GTE(openedRate) {
			profit := closedRate.Mul(sdk.NewDec(int64(positionInstance.Leverage))).Sub(positionInstance.Size_)
			profitAmount := profit.Quo(*closedRate)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := positionInstance.Size_.Sub(closedRate.Mul(sdk.NewDec(int64(positionInstance.Leverage))))
			lossAmount := loss.Quo(*closedRate)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_SHORT:
		k.AddPerpetualFuturesNetPositionOfMarket(ctx, position.Market, tradeAmount)

		if closedRate.LTE(openedRate) {
			profit := positionInstance.Size_.Sub(closedRate.Mul(sdk.NewDec(int64(positionInstance.Leverage))))
			profitAmount := profit.Quo(*closedRate)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := closedRate.Mul(sdk.NewDec(int64(positionInstance.Leverage))).Sub(positionInstance.Size_)
			lossAmount := loss.Quo(*closedRate)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, position.Address.AccAddress(), sdk.Coins{sdk.NewCoin(position.Market.BaseDenom, amountToUser.RoundInt())})

	ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionClosed{
		Sender:      position.Address.AccAddress().String(),
		PositionId:  position.Id,
		FeeAmount:   feeAmount.String(),
		TradeAmount: tradeAmount.String(),
	})

	return nil
}

func (k Keeper) ReportLiquidationNeededPerpetualFuturesPosition(ctx sdk.Context, rewardRecipient ununifiTypes.StringAccAddress, position types.Position, positionInstance types.PerpetualFuturesPositionInstance) error {
	params := k.GetParams(ctx)
	principal := positionInstance.CalculatePrincipal()

	if sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(sdk.NewDecWithPrec(1, 0)).LT(principal.Mul(params.PerpetualFutures.MarginMaintenanceRate)) {
		k.ClosePerpetualFuturesPosition(ctx, position, positionInstance)

		rewardAmount := sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(params.Pool.ReportLiquidationRewardRate).RoundInt()
		reward := sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, rewardAmount))
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rewardRecipient.AccAddress(), reward)

		ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionLiquidated{
			RewardRecipient: rewardRecipient.AccAddress().String(),
			PositionId:      position.Id,
			RemainingMargin: position.RemainingMargin.String(),
			RewardAmount:    rewardAmount.String(),
		})
		return nil
	}

	return errors.New("no liquidation needed")
}

func (k Keeper) ReportLevyPeriodPerpetualFuturesPosition(ctx sdk.Context, rewardRecipient ununifiTypes.StringAccAddress, position types.Position, positionInstance types.PerpetualFuturesPositionInstance) error {
	params := k.GetParams(ctx)

	netPosition := k.GetPerpetualFuturesNetPositionOfMarket(ctx, position.Market)

	imaginaryFundingRate := netPosition.Mul(params.PerpetualFutures.ImaginaryFundingRateProportionalCoefficient)
	imaginaryFundingFee := sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(imaginaryFundingRate).RoundInt()
	commissionFee := sdk.NewDecFromInt(position.RemainingMargin.Amount).Mul(params.PerpetualFutures.CommissionRate).RoundInt()

	if imaginaryFundingRate.IsNegative() {
		if positionInstance.PositionType == types.PositionType_SHORT {
			position.RemainingMargin.Amount = position.RemainingMargin.Amount.Sub(imaginaryFundingFee)
		} else {
			position.RemainingMargin.Amount = position.RemainingMargin.Amount.Add(imaginaryFundingFee.Sub(commissionFee))
		}
	} else {
		if positionInstance.PositionType == types.PositionType_LONG {
			position.RemainingMargin.Amount = position.RemainingMargin.Amount.Sub(imaginaryFundingFee)
		} else {
			position.RemainingMargin.Amount = position.RemainingMargin.Amount.Add(imaginaryFundingFee.Sub(commissionFee))
		}
	}
	position.LastLeviedAt = ctx.BlockTime()

	rewardAmount := sdk.NewDecFromInt(commissionFee).Mul(params.Pool.ReportLevyPeriodRewardRate).RoundInt()
	reward := sdk.NewCoins(sdk.NewCoin(position.RemainingMargin.Denom, rewardAmount))
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rewardRecipient.AccAddress(), reward)

	k.SetPosition(ctx, position)

	ctx.EventManager().EmitTypedEvent(&types.EventPerpetualFuturesPositionLevied{
		RewardRecipient: rewardRecipient.AccAddress().String(),
		PositionId:      position.Id,
		RemainingMargin: position.RemainingMargin.String(),
		RewardAmount:    rewardAmount.String(),
	})

	return nil
}

func (k Keeper) GetPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market) sdk.Dec {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.DenomNetPositionPerpetualFuturesKeyPrefix(market.BaseDenom, market.QuoteDenom))
	if bz == nil {
		return sdk.ZeroDec()
	}

	amount := sdk.MustNewDecFromStr(string(bz))

	return amount
}

func (k Keeper) SetPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market, amount sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := []byte(amount.String())

	store.Set(types.DenomNetPositionPerpetualFuturesKeyPrefix(market.BaseDenom, market.QuoteDenom), bz)
}

func (k Keeper) AddPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market, rhs sdk.Dec) {
	lhs := k.GetPerpetualFuturesNetPositionOfMarket(ctx, market)
	result := lhs.Add(rhs)

	k.SetPerpetualFuturesNetPositionOfMarket(ctx, market, result)
}

func (k Keeper) SubPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market, rhs sdk.Dec) {
	lhs := k.GetPerpetualFuturesNetPositionOfMarket(ctx, market)
	result := lhs.Sub(rhs)

	k.SetPerpetualFuturesNetPositionOfMarket(ctx, market, result)
}
