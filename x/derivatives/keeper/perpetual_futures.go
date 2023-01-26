package keeper

import (
	"fmt"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func (k Keeper) PerpetualFuturesOpenedPositionFactory(ctx sdk.Context, positionId string, sender sdk.AccAddress, positionInstance *types.PerpetualFuturesPosition) (*types.OpenedPosition, error) {
	openingPrice, err := k.GetPairPrice(ctx, positionInstance.Pair)
	if err != nil {
		return nil, err
	}

	openedPositionInstance := types.PerpetualFuturesOpenedPosition{
		PerpetualFuturesPosition: positionInstance,
		OpeningPrice:             *openingPrice,
	}

	any, err := codecTypes.NewAnyWithValue(&openedPositionInstance)
	if err != nil {
		return nil, err
	}

	return &types.OpenedPosition{
		Id:           positionId,
		Address:      sender,
		OpenedAt:     *timestamppb.New(ctx.BlockTime()),
		OpenedHeight: uint64(ctx.BlockHeight()),
		Position:     *any,
	}, nil
}

func (k Keeper) PerpetualFuturesClosedPositionFactory(ctx sdk.Context, openedPosition types.OpenedPosition, positionInstance *types.PerpetualFuturesOpenedPosition) (*types.ClosedPosition, error) {
	closingPrice, err := k.GetPairPrice(ctx, positionInstance.Pair)
	if err != nil {
		return nil, err
	}

	closedPositionInstance := types.PerpetualFuturesClosedPosition{
		PerpetualFuturesPosition: positionInstance.PerpetualFuturesPosition,
		OpeningPrice:             positionInstance.OpeningPrice,
		ClosingPrice:             *closingPrice,
	}

	any, err := codecTypes.NewAnyWithValue(&closedPositionInstance)
	if err != nil {
		return nil, err
	}

	return &types.ClosedPosition{
		Id:           openedPosition.Id,
		Address:      openedPosition.Address,
		OpenedAt:     openedPosition.OpenedAt,
		OpenedHeight: openedPosition.OpenedHeight,
		ClosedAt:     *timestamppb.New(ctx.BlockTime()),
		ClosedHeight: uint64(ctx.BlockHeight()),
		Position:     *any,
	}, nil
}

func (k Keeper) OpenPerpetualFuturesPosition(ctx sdk.Context, positionId string, sender sdk.AccAddress, positionInstance *types.PerpetualFuturesPosition) error {
	position, err := k.PerpetualFuturesOpenedPositionFactory(ctx, positionId, sender, positionInstance)
	if err != nil {
		return err
	}

	k.CreateOpenedPosition(ctx, *position)

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

func (k Keeper) ClosePerpetualFuturesPosition(ctx sdk.Context, openedPosition types.OpenedPosition, positionInstance *types.PerpetualFuturesOpenedPosition) error {
	params := k.GetParams(ctx)
	commissionRate := params.PerpetualFutures.CommissionRate
	feeAmount := positionInstance.Size_.Mul(commissionRate)
	tradeAmount := positionInstance.Size_.Sub(feeAmount)

	openingPrice := positionInstance.OpeningPrice
	closingPrice, err := k.GetPairPrice(ctx, positionInstance.Pair)
	if err != nil {
		return err
	}

	closedPosition, err := k.PerpetualFuturesClosedPositionFactory(ctx, openedPosition, positionInstance)
	if err != nil {
		return err
	}

	k.CreateClosedPosition(ctx, *closedPosition)

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, openedPosition.Address.AccAddress(), types.ModuleName, sdk.Coins{sdk.NewCoin(positionInstance.Pair.Denom, feeAmount.RoundInt())}) // TODO: this is wrong.

	principal := types.CalculatePrincipal(*positionInstance.PerpetualFuturesPosition)
	amountToUser := sdk.Dec{}

	switch positionInstance.PositionType {
	case types.PositionType_LONG:
		k.SubPerpetualFuturesNetPositionOfDenom(ctx, positionInstance.Pair.Denom, tradeAmount)

		if closingPrice.GTE(openingPrice) {
			profit := closingPrice.Mul(sdk.NewDecFromInt(positionInstance.Leverage)).Sub(positionInstance.Size_)
			profitAmount := profit.Quo(*closingPrice)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := positionInstance.Size_.Sub(closingPrice.Mul(sdk.NewDecFromInt(positionInstance.Leverage)))
			lossAmount := loss.Quo(*closingPrice)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_SHORT:
		k.AddPerpetualFuturesNetPositionOfDenom(ctx, positionInstance.Pair.Denom, tradeAmount)

		if closingPrice.LTE(openingPrice) {
			profit := positionInstance.Size_.Sub(closingPrice.Mul(sdk.NewDecFromInt(positionInstance.Leverage)))
			profitAmount := profit.Quo(*closingPrice)

			amountToUser = principal.Add(profitAmount)
		} else {
			loss := closingPrice.Mul(sdk.NewDecFromInt(positionInstance.Leverage)).Sub(positionInstance.Size_)
			lossAmount := loss.Quo(*closingPrice)

			amountToUser = principal.Sub(lossAmount)
		}
		break
	case types.PositionType_POSITION_UNKNOWN:
		return fmt.Errorf("unknown position type")
	}

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, openedPosition.Address.AccAddress(), sdk.Coins{sdk.NewCoin(positionInstance.Pair.Denom, amountToUser.RoundInt())})

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
