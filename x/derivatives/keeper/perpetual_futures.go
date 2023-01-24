package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) OpenPerpetualFuturesPosition(ctx sdk.Context, address sdk.AccAddress, position *types.PerpetualFuturesPosition) error {
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

func (k Keeper) ClosePerpetualFuturesPosition(ctx sdk.Context, address sdk.AccAddress, position *types.PerpetualFuturesPosition) error {
	// TODO: calculate payoffs
	// Didn't consider leverage yet by Alan
	// ^ position.Size_ is already contain the leverage.
	// See x/derivatives/types/perpetual_futures.go CalculatePrincipal.
	// position_size = leverage * principal by Yu
	params := k.GetParams(ctx)
	// decimal is 6
	commissionRate := params.CommissionRate
	feeAmount := position.Size_.Mul(commissionRate).Quo(sdk.NewDecWithPrec(1, 6))
	tradeAmount := position.Size_.Sub(feeAmount)

	// TODO: AddAccumulatedFee may be not needed. Jump to the definition of AddAccumulatedFee to read more comments.
	k.AddAccumulatedFee(ctx, feeAmount)
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, sdk.Coins{sdk.NewCoin(position.Denom, feeAmount.RoundInt())})

	// TODO: transfer the principal (margin = collateral) and profits from Pool to the trader or from the trader to Pool.

	switch position.PositionType {
	case types.PositionType_LONG:
		k.SubPerpetualFuturesNetPositionOfDenom(ctx, position.Denom, tradeAmount)
		break
	case types.PositionType_SHORT:
		k.AddPerpetualFuturesNetPositionOfDenom(ctx, position.Denom, tradeAmount)
		break
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

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
