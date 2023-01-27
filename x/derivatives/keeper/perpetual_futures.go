package keeper

import (
	"fmt"

	ununifiTypes "github.com/UnUniFi/chain/types"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) OpenPerpetualFuturesPosition(ctx sdk.Context, positionId string, sender ununifiTypes.StringAccAddress, margin sdk.Coin, market types.Market, positionInstance types.PerpetualFuturesPositionInstance) (*types.Position, error) {
	openedRate, err := k.GetPairRate(ctx, market)
	if err != nil {
		return nil, err
	}
	any, err := codecTypes.NewAnyWithValue(&positionInstance)
	if err != nil {
		return nil, err
	}

	position := types.Position{
		Id:               positionId,
		Address:          sender,
		OpenedAt:         ctx.BlockTime(),
		OpenedHeight:     uint64(ctx.BlockHeight()),
		OpenedRate:       *openedRate,
		PositionInstance: *any,
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

	// TODO: emit event

	return &position, nil
}

func (k Keeper) ClosePerpetualFuturesPosition(ctx sdk.Context, position types.Position, positionInstance types.PerpetualFuturesPositionInstance) error {
	params := k.GetParams(ctx)
	commissionRate := params.PerpetualFutures.CommissionRate
	feeAmount := positionInstance.Size_.Mul(commissionRate)
	tradeAmount := positionInstance.Size_.Sub(feeAmount)

	openedRate := position.OpenedRate
	closedRate, err := k.GetPairRate(ctx, position.Market)
	if err != nil {
		return err
	}

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, position.Address.AccAddress(), types.ModuleName, sdk.Coins{sdk.NewCoin(position.Market.Denom, feeAmount.RoundInt())}) // TODO: this is wrong.

	principal := types.CalculatePrincipal(positionInstance)
	amountToUser := sdk.Dec{}

	switch positionInstance.PositionType {
	case types.PositionType_LONG:
		k.SubPerpetualFuturesNetPositionOfMarket(ctx, position.Market, tradeAmount)

		if closedRate.GTE(openedRate) {
			profit := closedRate.Mul(sdk.NewDecFromInt(positionInstance.Leverage)).Sub(positionInstance.Size_)
			profitAmount := profit.Quo(*closedRate)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := positionInstance.Size_.Sub(closedRate.Mul(sdk.NewDecFromInt(positionInstance.Leverage)))
			lossAmount := loss.Quo(*closedRate)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_SHORT:
		k.AddPerpetualFuturesNetPositionOfMarket(ctx, position.Market, tradeAmount)

		if closedRate.LTE(openedRate) {
			profit := positionInstance.Size_.Sub(closedRate.Mul(sdk.NewDecFromInt(positionInstance.Leverage)))
			profitAmount := profit.Quo(*closedRate)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := closedRate.Mul(sdk.NewDecFromInt(positionInstance.Leverage)).Sub(positionInstance.Size_)
			lossAmount := loss.Quo(*closedRate)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, position.Address.AccAddress(), sdk.Coins{sdk.NewCoin(position.Market.Denom, amountToUser.RoundInt())})

	// TODO: emit event

	return nil
}

func (k Keeper) ReportLiquidationNeededPerpetualFuturesPosition(ctx sdk.Context, rewardRecipient ununifiTypes.StringAccAddress, remainingMargin sdk.Coin, position types.Position, positionInstance types.PerpetualFuturesPositionInstance) error {
	params := k.GetParams(ctx)
	principal := types.CalculatePrincipal(positionInstance)

	if sdk.NewDecFromInt(remainingMargin.Amount).Mul(sdk.NewDecWithPrec(1, 0)).LT(principal.Mul(params.PerpetualFutures.MarginMaintenanceRate)) {
		k.ClosePerpetualFuturesPosition(ctx, position, positionInstance)

		rewardAmount := sdk.NewDecFromInt(remainingMargin.Amount).Mul(params.Pool.LiquidationNeededReportRewardRate).RoundInt()
		reward := sdk.NewCoins(sdk.NewCoin(remainingMargin.Denom, rewardAmount))
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rewardRecipient.AccAddress(), reward)

		// TODO: emit event
	}

	// TODO: return error if report is invalid
	return nil
}

func (k Keeper) GetPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market) sdk.Dec {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.DenomNetPositionPerpetualFuturesKeyPrefix(market.Denom, market.QuoteDenom))
	amount := sdk.MustNewDecFromStr(string(bz))

	return amount
}

func (k Keeper) SetPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market, amount sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := []byte(amount.String())

	store.Set(types.DenomNetPositionPerpetualFuturesKeyPrefix(market.Denom, market.QuoteDenom), bz)
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
