package keeper

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// GetStrategyCount get the total number of Strategy
func (k Keeper) GetStrategyCount(ctx sdk.Context, vaultDenom string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixStrategyCount(vaultDenom)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetStrategyCount set the total number of Strategy
func (k Keeper) SetStrategyCount(ctx sdk.Context, vaultDenom string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixStrategyCount(vaultDenom)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendStrategy appends a Strategy in the store with a new id and update the count
func (k Keeper) AppendStrategy(
	ctx sdk.Context,
	vaultDenom string,
	strategy types.Strategy,
) uint64 {
	// Create the strategy
	count := k.GetStrategyCount(ctx, vaultDenom)

	// Set the ID of the appended value
	strategy.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	bz := k.cdc.MustMarshal(&strategy)
	store.Set(GetStrategyIDBytes(strategy.Id), bz)

	// Update strategy count
	k.SetStrategyCount(ctx, vaultDenom, count+1)

	return count
}

// SetStrategy set a specific Strategy in the store
func (k Keeper) SetStrategy(ctx sdk.Context, vaultDenom string, Strategy types.Strategy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	b := k.cdc.MustMarshal(&Strategy)
	store.Set(GetStrategyIDBytes(Strategy.Id), b)
}

// GetStrategy returns a Strategy from its id
func (k Keeper) GetStrategy(ctx sdk.Context, vaultDenom string, id uint64) (val types.Strategy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	b := store.Get(GetStrategyIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStrategy removes a Strategy from the store
func (k Keeper) RemoveStrategy(ctx sdk.Context, vaultDenom string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	store.Delete(GetStrategyIDBytes(id))
}

func (k Keeper) MigrateAllLegacyStrategies(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(""))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	legacyStrategies := []types.LegacyStrategy{}
	for ; iterator.Valid(); iterator.Next() {
		var val types.LegacyStrategy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		legacyStrategies = append(legacyStrategies, val)
	}

	for _, legacyStrategy := range legacyStrategies {
		strategy := types.Strategy{
			Denom:           legacyStrategy.Denom,
			Id:              legacyStrategy.Id,
			ContractAddress: legacyStrategy.ContractAddress,
			Name:            legacyStrategy.Name,
			Description:     "",
			GitUrl:          legacyStrategy.GitUrl,
		}
		k.SetStrategy(ctx, strategy.Denom, strategy)
	}
}

// GetAllStrategy returns all Strategy
func (k Keeper) GetAllStrategy(ctx sdk.Context, vaultDenom string) (list []types.Strategy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Strategy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetStrategyIDBytes returns the byte representation of the ID
func GetStrategyIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetStrategyIDFromBytes returns ID in uint64 format from a byte array
func GetStrategyIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

// stake into strategy
func (k Keeper) StakeToStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy, amount sdk.Int) error {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	stakeCoin := sdk.NewCoin(vault.Denom, amount)
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		return k.stakeibcKeeper.LiquidStake(
			ctx,
			vaultModAddr,
			stakeCoin,
		)
	default:
		wasmMsg := `{"stake":{}}`
		contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, vaultModAddr, []byte(wasmMsg), sdk.Coins{stakeCoin})
		return err
	}
}

// unstake worth of withdrawal amount from the strategy
func (k Keeper) UnstakeFromStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy, amount sdk.Int) error {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		{
			err := k.stakeibcKeeper.RedeemStake(
				ctx,
				vaultModAddr,
				sdk.NewCoin(vault.Denom, amount),
				vaultModAddr.String(),
			)
			if err != nil {
				return err
			}

			return nil
		}
	default:
		wasmMsg := fmt.Sprintf(`{"unstake":{"amount":"%s"}}`, amount.String())
		contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, vaultModAddr, []byte(wasmMsg), sdk.Coins{})
		return err
	}
}

func (k Keeper) GetAmountFromStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy) (sdk.Coin, error) {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		updatedAmount := k.stakeibcKeeper.GetUpdatedBalance(ctx, vaultModAddr, vault.Denom)
		return sdk.NewCoin(vault.Denom, updatedAmount), nil
	default:
		wasmQuery := fmt.Sprintf(`{"bonded":{"addr": "%s"}}`, vaultModAddr.String())
		contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
		resp, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
		if err != nil {
			return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), err
		}
		amountStr := strings.ReplaceAll(string(resp), "\"", "")
		amount, ok := sdk.NewIntFromString(amountStr)
		if !ok {
			return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), nil
		}
		return sdk.NewCoin(strategy.Denom, amount), err
	}
}

func (k Keeper) GetUnbondingAmountFromStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy) (sdk.Coin, error) {
	vaultModName := types.GetVaultModuleAccountName(vault.Id)
	vaultModAddr := authtypes.NewModuleAddress(vaultModName)
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		zone, err := k.stakeibcKeeper.GetHostZoneFromIBCDenom(ctx, vault.Denom)
		if err != nil {
			return sdk.Coin{}, err
		}
		unbondingAmount := k.recordsKeeper.GetUserRedemptionRecordBySenderAndHostZone(ctx, vaultModAddr, zone.ChainId)
		return sdk.NewCoin(vault.Denom, unbondingAmount), nil
	default:
		wasmQuery := fmt.Sprintf(`{"unbonding":{"addr": "%s"}}`, vaultModAddr.String())
		contractAddr := sdk.MustAccAddressFromBech32(strategy.ContractAddress)
		resp, err := k.wasmReader.QuerySmart(ctx, contractAddr, []byte(wasmQuery))
		if err != nil {
			return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), err
		}
		amountStr := strings.ReplaceAll(string(resp), "\"", "")
		amount, ok := sdk.NewIntFromString(amountStr)
		if !ok {
			return sdk.NewCoin(strategy.Denom, sdk.ZeroInt()), nil
		}
		return sdk.NewCoin(strategy.Denom, amount), err
	}
}
