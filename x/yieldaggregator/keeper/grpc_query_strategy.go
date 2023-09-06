package keeper

import (
	"context"
	"encoding/json"

	math "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

func (k Keeper) StrategyAll(c context.Context, req *types.QueryAllStrategyRequest) (*types.QueryAllStrategyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var strategies []types.StrategyContainer
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	strategyStore := prefix.NewStore(store, types.KeyPrefixStrategy(req.Denom))

	pageRes, err := query.Paginate(strategyStore, req.Pagination, func(key []byte, value []byte) error {
		var strategy types.Strategy
		if err := k.cdc.Unmarshal(value, &strategy); err != nil {
			return err
		}

		strategyContainer, err := k.GetStrategyContainer(ctx, strategy)
		if err != nil {
			return err
		}

		strategies = append(strategies, strategyContainer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStrategyResponse{Strategies: strategies, Pagination: pageRes}, nil
}

func (k Keeper) GetStrategyContainer(ctx sdk.Context, strategy types.Strategy) (types.StrategyContainer, error) {
	strategyAddr, err := sdk.AccAddressFromBech32(strategy.ContractAddress)
	if err != nil {
		return types.StrategyContainer{
			Strategy:           strategy,
			DepositFeeRate:     math.LegacyZeroDec(),
			WithdrawFeeRate:    math.LegacyZeroDec(),
			PerformanceFeeRate: math.LegacyZeroDec(),
		}, nil
	}

	wasmQuery := `{"fee":{}}`
	result, err := k.wasmReader.QuerySmart(ctx, strategyAddr, []byte(wasmQuery))
	if err != nil {
		return types.StrategyContainer{
			Strategy:           strategy,
			DepositFeeRate:     math.LegacyZeroDec(),
			WithdrawFeeRate:    math.LegacyZeroDec(),
			PerformanceFeeRate: math.LegacyZeroDec(),
		}, nil
	}

	jsonMap := make(map[string]string)
	err = json.Unmarshal(result, &jsonMap)
	if err != nil {
		return types.StrategyContainer{
			Strategy:           strategy,
			DepositFeeRate:     math.LegacyZeroDec(),
			WithdrawFeeRate:    math.LegacyZeroDec(),
			PerformanceFeeRate: math.LegacyZeroDec(),
		}, nil
	}

	version := k.GetStrategyVersion(ctx, strategy)

	switch version {
	case 0:
		depositFeeRate, err := math.LegacyNewDecFromStr(jsonMap["deposit_fee_rate"])
		if err != nil {
			depositFeeRate = math.LegacyZeroDec()
		}
		withdrawFeeRate, err := math.LegacyNewDecFromStr(jsonMap["withdraw_fee_rate"])
		if err != nil {
			withdrawFeeRate = math.LegacyZeroDec()
		}
		performanceFeeRate, err := math.LegacyNewDecFromStr(jsonMap["interest_fee_rate"])
		if err != nil {
			performanceFeeRate = math.LegacyZeroDec()
		}
		return types.StrategyContainer{
			Strategy:           strategy,
			DepositFeeRate:     depositFeeRate,
			WithdrawFeeRate:    withdrawFeeRate,
			PerformanceFeeRate: performanceFeeRate,
		}, nil
	default: // case 1+
		depositFeeRate, err := math.LegacyNewDecFromStr(jsonMap["deposit_fee_rate"])
		if err != nil {
			depositFeeRate = math.LegacyZeroDec()
		}
		withdrawFeeRate, err := math.LegacyNewDecFromStr(jsonMap["withdraw_fee_rate"])
		if err != nil {
			withdrawFeeRate = math.LegacyZeroDec()
		}
		performanceFeeRate, err := math.LegacyNewDecFromStr(jsonMap["performance_fee_rate"])
		if err != nil {
			performanceFeeRate = math.LegacyZeroDec()
		}

		return types.StrategyContainer{
			Strategy:           strategy,
			DepositFeeRate:     depositFeeRate,
			WithdrawFeeRate:    withdrawFeeRate,
			PerformanceFeeRate: performanceFeeRate,
		}, nil
	}
}

func (k Keeper) Strategy(c context.Context, req *types.QueryGetStrategyRequest) (*types.QueryGetStrategyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	strategy, found := k.GetStrategy(ctx, req.Denom, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	strategyContainer, err := k.GetStrategyContainer(ctx, strategy)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetStrategyResponse{
		Strategy: strategyContainer,
	}, nil
}
