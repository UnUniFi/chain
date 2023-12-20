package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/irs/types"
)

// SetTranchePool set a specific TranchePool in the store
func (k Keeper) SetTranchePool(ctx sdk.Context, tranchePool types.TranchePool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	b := k.cdc.MustMarshal(&tranchePool)
	store.Set(sdk.Uint64ToBigEndian(tranchePool.Id), b)

	store = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrancheByStrategyKey))
	store.Set(types.KeyTrancheByStrategy(tranchePool), sdk.Uint64ToBigEndian(tranchePool.Id))
}

// GetTranchePool returns a TranchePool from its identifier
func (k Keeper) GetTranchePool(ctx sdk.Context, id uint64) (val types.TranchePool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	b := store.Get(sdk.Uint64ToBigEndian(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTranchePool removes a TranchePool from the store
func (k Keeper) RemoveTranchePool(ctx sdk.Context, tranchePool types.TranchePool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	store.Delete(sdk.Uint64ToBigEndian(tranchePool.Id))

	store = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrancheByStrategyKey))
	store.Delete(types.KeyTrancheByStrategy(tranchePool))
}

func (k Keeper) GetTranchesByStrategy(ctx sdk.Context, strategyContract string) (list []types.TranchePool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TrancheByStrategyKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte(strategyContract))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		id := sdk.BigEndianToUint64(iterator.Value())
		pool, found := k.GetTranchePool(ctx, id)
		if found {
			list = append(list, pool)
		}
	}

	return
}

// GetAllTranchePool returns all TranchePool
func (k Keeper) GetAllTranchePool(ctx sdk.Context) (list []types.TranchePool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TranchePool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetLastTrancheId(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TranchePoolKey))
	iterator := sdk.KVStoreReversePrefixIterator(store, []byte{})

	defer iterator.Close()
	if iterator.Valid() {
		var val types.TranchePool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		return val.Id
	}

	return 0
}

func (k Keeper) DepositToTranchePool(ctx sdk.Context, sender sdk.AccAddress, trancheId uint64, trancheType types.TrancheType, token sdk.Coin, requiredYt sdk.Int) error {
	tranche, found := k.GetTranchePool(ctx, trancheId)
	if !found {
		return types.ErrTrancheNotFound
	}

	if trancheType == types.TrancheType_NORMAL_YIELD { // Both PT and YT
		_, err := k.MintPtYtPair(ctx, sender, tranche, token)
		if err != nil {
			return err
		}
	} else if trancheType == types.TrancheType_FIXED_YIELD {
		// Buy PT from AMM with msg.TrancheMaturity for msg.SpendAmount
		_, err := k.SwapPoolTokens(ctx, sender, tranche, token)
		if err != nil {
			return err
		}
	} else if trancheType == types.TrancheType_LEVERAGED_VARIABLE_YIELD {
		// Borrow msg.AmountToBuy from AMM pool
		// MintPtYtPair
		// Sell msg.AmountToBuy worth of PT
		// Return borrowed amount
		err := k.SwapUtToYt(ctx, sender, tranche, requiredYt, token)
		if err != nil {
			return err
		}
	} else {
		return types.ErrInvalidTrancheType
	}
	return nil
}

func (k Keeper) WithdrawFromTranchePool(ctx sdk.Context, sender sdk.AccAddress, trancheId uint64, trancheType types.TrancheType, tokens sdk.Coins, requiredUt sdk.Int) error {
	tranche, found := k.GetTranchePool(ctx, trancheId)
	if !found {
		return types.ErrTrancheNotFound
	}

	if trancheType == types.TrancheType_NORMAL_YIELD { // Both PT and YT
		err := k.RedeemPtYtPair(ctx, sender, tranche, requiredUt, tokens)
		if err != nil {
			return err
		}
	} else if trancheType == types.TrancheType_FIXED_YIELD {
		if len(tokens) != 1 {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "FIXED_YIELD, expected 1 coin, got %d", len(tokens))
		}
		// If matured, send required amount from unbonded from the share
		if tranche.StartTime+tranche.Maturity <= uint64(ctx.BlockTime().Unix()) {
			err := k.RedeemPtAtMaturity(ctx, sender, tranche, tokens[0])
			if err != nil {
				return err
			}
		} else {
			// Else, sell PT from AMM with msg.TrancheMaturity for msg.PTAmount
			_, err := k.SwapPoolTokens(ctx, sender, tranche, tokens[0])
			if err != nil {
				return err
			}
		}
	} else if trancheType == types.TrancheType_LEVERAGED_VARIABLE_YIELD {
		// If matured, send required amount from unbonded from the share
		if tranche.StartTime+tranche.Maturity <= uint64(ctx.BlockTime().Unix()) {
			if len(tokens) != 1 {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "LEVERAGED_VARIABLE_YIELD, expected 1 coin, got %d", len(tokens))
			}
			err := k.RedeemYtAtMaturity(ctx, sender, tranche, tokens[0])
			if err != nil {
				return err
			}
		} else {
			// Else
			// Put required amount of msg.PT from user wallet
			// Close position
			// Start redemption for strategy share
			if len(tokens) != 2 {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "not matured, expected 2 coins, got %d", len(tokens))
			}
			err := k.SwapYtToUt(ctx, sender, tranche, requiredUt, tokens)
			if err != nil {
				return err
			}
		}
	} else {
		return types.ErrInvalidTrancheType
	}
	return nil
}
