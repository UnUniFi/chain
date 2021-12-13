package keeper

import (
	"context"

	"github.com/UnUniFi/chain/x/cdp/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CdpAll(c context.Context, req *types.QueryAllCdpRequest) (*types.QueryAllCdpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var augmentedCdps types.AugmentedCdps
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	cdpStore := prefix.NewStore(store, types.KeyPrefix(types.CdpKey))

	pageRes, err := query.Paginate(cdpStore, req.Pagination, func(key []byte, value []byte) error {
		var cdp types.Cdp
		if err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &cdp); err != nil {
			return err
		}
		augmentedCdp := k.LoadAugmentedCdp(ctx, cdp)
		augmentedCdps = append(augmentedCdps, augmentedCdp)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCdpResponse{Cdp: augmentedCdps, Pagination: pageRes}, nil
}

func (k Keeper) Cdp(c context.Context, req *types.QueryGetCdpRequest) (*types.QueryGetCdpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	ownerAddress, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %s", req.Owner)
	}

	_, valid := k.GetCollateralTypePrefix(ctx, req.CollateralType)
	if !valid {
		return nil, status.Errorf(codes.NotFound, "invalid collateral type: %s", req.CollateralType)
	}

	cdp, found := k.GetCdpByOwnerAndCollateralType(ctx, ownerAddress, req.CollateralType)

	if !found {
		return nil, status.Error(codes.NotFound, "cdp not found")
	}
	augmentedCdp := k.LoadAugmentedCdp(ctx, cdp)

	return &types.QueryGetCdpResponse{Cdp: augmentedCdp}, nil
}

func (k Keeper) AccountAll(c context.Context, req *types.QueryAllAccountRequest) (*types.QueryAllAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	cdpAccAccount := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	liquidatorAccAccount := k.accountKeeper.GetModuleAccount(ctx, types.LiquidatorMacc)

	accounts := []authtypes.ModuleAccount{
		*cdpAccAccount.(*authtypes.ModuleAccount),
		*liquidatorAccAccount.(*authtypes.ModuleAccount),
	}

	var accountsAny []*codectypes.Any

	for _, acc := range accounts {
		accAny, _ := codectypes.NewAnyWithValue(&acc)
		accountsAny = append(accountsAny, accAny)
	}

	return &types.QueryAllAccountResponse{Accounts: accountsAny}, nil
}

func (k Keeper) DepositAll(c context.Context, req *types.QueryAllDepositRequest) (*types.QueryAllDepositResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	cdp, found := k.GetCdpByOwnerAndCollateralType(ctx, sdk.AccAddress(req.Owner), req.CollateralType)

	if !found {
		return nil, status.Error(codes.NotFound, "cdp not found")
	}

	deposits := k.GetDeposits(ctx, cdp.Id)

	return &types.QueryAllDepositResponse{Deposits: deposits}, nil
}
