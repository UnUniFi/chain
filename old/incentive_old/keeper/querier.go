package keeper

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/lcnem/jpyx/x/incentive/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryGetParams:
			return queryGetParams(ctx, req, k)
		case types.QueryGetJPYXMintingRewards:
			return queryGetJPYXMintingRewards(ctx, req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint", types.ModuleName)
		}
	}
}

// query params in the store
func queryGetParams(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	// Get params
	params := k.GetParams(ctx)

	// Encode results
	bz, err := codec.MarshalJSONIndent(k.cdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryGetJPYXMintingRewards(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryJPYXMintingRewardsParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	owner := len(params.Owner) > 0

	var jpyxMintingClaims types.JPYXMintingClaims
	switch {
	case owner:
		jpyxMintingClaim, foundJpyxMintingClaim := k.GetJPYXMintingClaim(ctx, params.Owner)
		if foundJpyxMintingClaim {
			jpyxMintingClaims = append(jpyxMintingClaims, jpyxMintingClaim)
		}
	default:
		jpyxMintingClaims = k.GetAllJPYXMintingClaims(ctx)
	}

	var paginatedJpyxMintingClaims types.JPYXMintingClaims
	startU, endU := client.Paginate(len(jpyxMintingClaims), params.Page, params.Limit, 100)
	if startU < 0 || endU < 0 {
		paginatedJpyxMintingClaims = types.JPYXMintingClaims{}
	} else {
		paginatedJpyxMintingClaims = jpyxMintingClaims[startU:endU]
	}

	var augmentedJpyxMintingClaims types.JPYXMintingClaims
	for _, claim := range paginatedJpyxMintingClaims {
		augmentedClaim := k.SimulateJPYXMintingSynchronization(ctx, claim)
		augmentedJpyxMintingClaims = append(augmentedJpyxMintingClaims, augmentedClaim)
	}

	// Marshal JPYX minting claims
	bz, err := codec.MarshalJSONIndent(k.cdc, augmentedJpyxMintingClaims)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
