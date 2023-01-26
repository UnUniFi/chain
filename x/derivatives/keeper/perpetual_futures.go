package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) PerpetualFuturesOpenedPositionFactory(ctx sdk.Context, positionInstance *types.PerpetualFuturesPosition) (*types.PerpetualFuturesOpenedPosition, error) {
	price, err := k.GetPairPrice(ctx, positionInstance.Pair)
	if err != nil {
		return nil, err
	}

	return &types.PerpetualFuturesOpenedPosition{
		PerpetualFuturesPosition: *&positionInstance,
		OpeningPrice:             *price,
	}, nil
}

func (k Keeper) OpenPerpetualFuturesPosition(ctx sdk.Context, openedPosition types.OpenedPosition, positionInstance *types.PerpetualFuturesOpenedPosition) error {
	k.CreateOpenedPosition(ctx, openedPosition)

	switch positionInstance.PositionType {
	case types.PositionType_LONG:
		k.AddPerpetualFuturesNetPositionOfDenom(ctx, positionInstance.Pair.Denom, positionInstance.Size_)
		break
	case types.PositionType_SHORT:
		k.SubPerpetualFuturesNetPositionOfDenom(ctx, positionInstance.Pair.Denom, positionInstance.Size_)
		break
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

	return nil
}

func (k Keeper) ClosePerpetualFuturesPosition(ctx sdk.Context, closedPosition types.ClosedPosition, positionInstance *types.PerpetualFuturesClosedPosition) error {
	params := k.GetParams(ctx)
	commissionRate := params.PerpetualFutures.CommissionRate
	feeAmount := positionInstance.Size_.Mul(commissionRate)
	tradeAmount := positionInstance.Size_.Sub(feeAmount)

	price, err := k.GetAssetPrice(ctx, positionInstance.Pair.Denom)
	if err != nil {
		return err
	}

	openPrice := positionInstance.OpeningPrice

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, closedPosition.Address.AccAddress(), types.ModuleName, sdk.Coins{sdk.NewCoin(positionInstance.Pair.Denom, feeAmount.RoundInt())})

	principal := types.CalculatePrincipal(*positionInstance.PerpetualFuturesPosition)
	amountToUser := sdk.Dec{}

	switch positionInstance.PositionType {
	case types.PositionType_LONG:
		k.SubPerpetualFuturesNetPositionOfDenom(ctx, positionInstance.Pair.Denom, tradeAmount)

		if price.Price.GTE(openPrice) {
			profit := price.Price.Mul(sdk.NewDecFromInt(positionInstance.Leverage)).Sub(positionInstance.Size_)
			profitAmount := profit.Quo(price.Price)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := positionInstance.Size_.Sub(price.Price.Mul(sdk.NewDecFromInt(positionInstance.Leverage)))
			lossAmount := loss.Quo(price.Price)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_SHORT:
		k.AddPerpetualFuturesNetPositionOfDenom(ctx, positionInstance.Pair.Denom, tradeAmount)

		if price.Price.LTE(openPrice) {
			profit := positionInstance.Size_.Sub(price.Price.Mul(sdk.NewDecFromInt(positionInstance.Leverage)))
			profitAmount := profit.Quo(price.Price)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := price.Price.Mul(sdk.NewDecFromInt(positionInstance.Leverage)).Sub(positionInstance.Size_)
			lossAmount := loss.Quo(price.Price)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, closedPosition.Address.AccAddress(), sdk.Coins{sdk.NewCoin(positionInstance.Pair.Denom, amountToUser.RoundInt())})

	return nil
}

func (k Keeper) GetPerpetualFuturesNetPositionOfDenom(ctx sdk.Context, denom string) sdk.Dec {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.DenomNetPositionPerpetualFuturesKeyPrefix(denom))
	amount := sdk.MustNewDecFromStr(string(bz))

	return amount
}

func (k Keeper) SetPerpetualFuturesNetPositionOfDenom(ctx sdk.Context, denom string, amount sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	bz := []byte(amount.String())

	store.Set(types.DenomNetPositionPerpetualFuturesKeyPrefix(denom), bz)
}

func (k Keeper) AddPerpetualFuturesNetPositionOfDenom(ctx sdk.Context, denom string, rhs sdk.Dec) {
	lhs := k.GetPerpetualFuturesNetPositionOfDenom(ctx, denom)
	result := lhs.Add(rhs)

	k.SetPerpetualFuturesNetPositionOfDenom(ctx, denom, result)
}

func (k Keeper) SubPerpetualFuturesNetPositionOfDenom(ctx sdk.Context, denom string, rhs sdk.Dec) {
	lhs := k.GetPerpetualFuturesNetPositionOfDenom(ctx, denom)
	result := lhs.Sub(rhs)

	k.SetPerpetualFuturesNetPositionOfDenom(ctx, denom, result)
}
