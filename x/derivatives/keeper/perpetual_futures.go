package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) OpenPerpetualFuturesPosition(ctx sdk.Context, address sdk.AccAddress, positionId uint64, margin sdk.Coin, position *types.PerpetualFuturesPosition) error {
	price, err := k.GetAssetPrice(ctx, position.Denom)
	if err != nil {
		return err
	}

	// TODO: levy margin (principal, collateral)
	k.SaveDepositedMargin(ctx, positionId, margin)

	k.SaveOpenPositionPrice(ctx, positionId, price)

	switch position.PositionType {
	case types.PositionType_LONG:
		k.AddPerpetualFuturesNetPositionOfDenom(ctx, position.Denom, position.Size_)
		break
	case types.PositionType_SHORT:
		k.SubPerpetualFuturesNetPositionOfDenom(ctx, position.Denom, position.Size_)
		break
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

	return nil
}

func (k Keeper) ClosePerpetualFuturesPosition(ctx sdk.Context, address sdk.AccAddress, positionId uint64, position *types.PerpetualFuturesPosition) error {
	// TODO: calculate payoffs
	// Didn't consider leverage yet by Alan
	// ^ position.Size_ is already contain the leverage.
	// See x/derivatives/types/perpetual_futures.go CalculatePrincipal.
	// position_size = leverage * principal by Yu
	params := k.GetParams(ctx)
	// decimal is 12
	commissionRate := params.CommissionRate
	feeAmount := position.Size_.Mul(commissionRate).Quo(sdk.NewDecWithPrec(1, 6))
	tradeAmount := position.Size_.Sub(feeAmount)

	price, err := k.GetAssetPrice(ctx, position.Denom)
	if err != nil {
		return err
	}

	k.SaveClosedPositionPrice(ctx, positionId, price)
	openPrice := k.GetOpenPositionPrice(ctx, positionId)

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, sdk.Coins{sdk.NewCoin(position.Denom, feeAmount.RoundInt())})

	principal := types.CalculatePrincipal(*position)
	amountToUser := sdk.Dec{}

	switch position.PositionType {
	case types.PositionType_LONG:
		k.SubPerpetualFuturesNetPositionOfDenom(ctx, position.Denom, tradeAmount)

		if price.Price.GTE(openPrice.Price) {
			profit := price.Price.Mul(sdk.NewDecFromInt(position.Leverage)).Sub(position.Size_)
			profitAmount := profit.Quo(price.Price)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := position.Size_.Sub(price.Price.Mul(sdk.NewDecFromInt(position.Leverage)))
			lossAmount := loss.Quo(price.Price)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_SHORT:
		k.AddPerpetualFuturesNetPositionOfDenom(ctx, position.Denom, tradeAmount)

		if price.Price.LTE(openPrice.Price) {
			profit := position.Size_.Sub(price.Price.Mul(sdk.NewDecFromInt(position.Leverage)))
			profitAmount := profit.Quo(price.Price)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := price.Price.Mul(sdk.NewDecFromInt(position.Leverage)).Sub(position.Size_)
			lossAmount := loss.Quo(price.Price)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.Coins{sdk.NewCoin(position.Denom, amountToUser.RoundInt())})

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
