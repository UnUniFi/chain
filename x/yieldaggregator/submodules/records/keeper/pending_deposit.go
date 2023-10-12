package keeper

import (
	"fmt"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/submodules/records/types"
)

func (k Keeper) GetAllVaultPendingDeposits(ctx sdk.Context) []types.PendingDeposit {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PendingDepositKey))
	deposits := []types.PendingDeposit{}
	iterator := store.Iterator(nil, nil)
	for ; iterator.Valid(); iterator.Next() {
		var amount math.Int
		err := amount.Unmarshal(iterator.Value())
		if err != nil {
			panic(fmt.Errorf("unable to unmarshal supply value %v", err))
		}

		deposits = append(deposits, types.PendingDeposit{
			VaultId: sdk.BigEndianToUint64(iterator.Key()),
			Amount:  amount,
		})
	}
	return deposits
}

func (k Keeper) IncreaseVaultPendingDeposit(ctx sdk.Context, vaultId uint64, amount sdk.Int) {
	deposit := k.GetVaultPendingDeposit(ctx, vaultId)
	k.SetVaultPendingDeposit(ctx, vaultId, deposit.Add(amount))
}

func (k Keeper) DecreaseVaultPendingDeposit(ctx sdk.Context, vaultId uint64, amount sdk.Int) {
	deposit := k.GetVaultPendingDeposit(ctx, vaultId)
	k.SetVaultPendingDeposit(ctx, vaultId, deposit.Sub(amount))
}

func (k Keeper) GetVaultPendingDeposit(ctx sdk.Context, vaultId uint64) sdk.Int {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PendingDepositKey))
	bz := store.Get(sdk.Uint64ToBigEndian(vaultId))
	if bz == nil {
		return sdk.ZeroInt()
	}

	var amount math.Int
	err := amount.Unmarshal(bz)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal supply value %v", err))
	}

	return amount
}

func (k Keeper) SetVaultPendingDeposit(ctx sdk.Context, vaultId uint64, amount sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PendingDepositKey))
	bz, err := amount.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value %v", err))
	}

	store.Set(sdk.Uint64ToBigEndian(vaultId), bz)
}
