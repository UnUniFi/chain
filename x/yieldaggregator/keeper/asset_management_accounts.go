package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

// asset management keeper functions
func (k Keeper) AddAssetManagementAccounts(ctx sdk.Context, id string, name string) error {

	return nil
}

func (k Keeper) UpdateAssetManagementAccounts(ctx sdk.Context, id string, obj types.AssetManagementAccount) error {
	return nil
}

func (k Keeper) DeleteAssetManagementAccounts(ctx sdk.Context, id string) error {
	return nil
}

func (k Keeper) GetAssetManagementAccounts(ctx sdk.Context) {
}

// deposit
func (k Keeper) Deposit(ctx sdk.Context, msg *types.MsgDeposit) error {
	return nil
}

// withdraw
func (k Keeper) Withdraw(ctx sdk.Context, msg *types.MsgWithdraw) error {
	return nil
}
