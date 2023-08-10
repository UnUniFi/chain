package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/x/nft"

	"github.com/UnUniFi/chain/x/nftfactory/types"
)

// CreateClass creates new class id with `nftfactory/{creatorAddr}/{subdenom}` name.
// Charges creatorAddr fee for creation
func (k Keeper) CreateClass(ctx sdk.Context, creatorAddr, subdenom string) (newTokenDenom string, err error) {
	err = k.chargeFeeForDenomCreation(ctx, creatorAddr)
	if err != nil {
		return "", sdkerrors.Wrapf(types.ErrUnableToCharge, "class fee collection error: %v", err)
	}

	denom, err := k.validateCreateDenom(ctx, creatorAddr, subdenom)
	if err != nil {
		return "", sdkerrors.Wrapf(types.ErrInvalidClassId, "class id validation error: %v", err)
	}

	err = k.createClassAfterValidation(ctx, creatorAddr, denom)
	if err != nil {
		return "", sdkerrors.Wrap(err, "create class after validation error")
	}

	return denom, nil
}

// Runs CreateClass logic after the charge and all denom validation has been handled.
// Made into a second function for genesis initialization.
func (k Keeper) createClassAfterValidation(ctx sdk.Context, creatorAddr, classId string) (err error) {
	k.nftKeeper.SaveClass(ctx, nft.Class{})

	authorityMetadata := types.ClassAuthorityMetadata{
		Admin: creatorAddr,
	}
	err = k.setAuthorityMetadata(ctx, classId, authorityMetadata)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidAuthorityMetadata, "unable to set authority metadata: %v", err)
	}

	k.addDenomFromCreator(ctx, creatorAddr, classId)
	return nil
}

func (k Keeper) validateCreateDenom(ctx sdk.Context, creatorAddr, subclass string) (newClassId string, err error) {

	denom, err := types.GetClassId(creatorAddr, subclass)
	if err != nil {
		return "", sdkerrors.Wrapf(types.ErrInvalidClassId, "invalid class id: %v", err)
	}

	_, found := k.nftKeeper.GetClass(ctx, denom)
	if found {
		return "", types.ErrClassExists
	}

	return denom, nil
}

func (k Keeper) chargeFeeForDenomCreation(ctx sdk.Context, creatorAddr string) (err error) {
	// Send creation fee to community pool
	creationFee := k.GetParams(ctx).ClassCreationFee
	accAddr, err := sdk.AccAddressFromBech32(creatorAddr)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrUnableToCharge, "wrong creator address: %v", err)
	}

	params := k.GetParams(ctx)

	if len(creationFee) > 0 {
		feeCollectorAddr, err := sdk.AccAddressFromBech32(params.FeeCollectorAddress)
		if err != nil {
			return sdkerrors.Wrapf(types.ErrUnableToCharge, "wrong fee collector address: %v", err)
		}

		err = k.bankKeeper.SendCoins(
			ctx,
			accAddr, feeCollectorAddr,
			creationFee,
		)

		if err != nil {
			return sdkerrors.Wrap(err, "unable to send coins to fee collector")
		}
	}

	return nil
}
