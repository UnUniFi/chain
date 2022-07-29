package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Deposit(ctx sdk.Context, msg *types.MsgDeposit) error {

	return nil
}

func (k Keeper) Withdraw(ctx sdk.Context, msg *types.MsgWithdraw) error {

	return nil
}

// // AssetManagementAccountKeeper
//   AddAssetManagementAccounts(ctx sdk.Context, id string, name string)
//   UpdateAssetManagementAccounts(ctx sdk.Context, id string, obj types.AssetManagementAccount)
//   DeleteAssetManagementAccounts(ctx sdk.Context, id string)
//   AddAssetManagementTargetsOfAccount(ctx sdk.Context, account_id string, obj types.AssetManagementTarget)
//   UpdateAssetManagementTargetsOfAccount(ctx sdk.Context, targetId string, obj types.AssetManagementTarget)
//   DeleteAssetManagementTargetsOfAccount(ctx sdk.Context, targetId string)

// // AssetManagementAccountBankKeeper
//   PayBack(ctx sdk.Context, targetId string, farmingUnit FarmingUnit)

// 	// AssetManagementAccountGetKeeper
//   GetAssetManagementAccounts(ctx sdk.Context)
//   GetAssetManagementTargetsOfAccount(ctx sdk.Context, accountId string)
//   GetAssetManagementTargetsOfDenom(ctx sdk.Context, accountId string, denom string)

// 	// AssetManagementKeeper
//   AddFarmingOrder(ctx sdk.Context, farmingOrder FarmingOrder)
//   DeleteFarmingOrder(ctx sdk.Context, sender sdk.AccAddress, farmingOrderId string)
//   GetFarmingOrdersOfAddress(ctx sdk.Context, sender sdk.AccAddress)
//   ActivateFarmingOrder(ctx sdk.Context, sender sdk.AccAddress, farmingOrderId string)
//   InactivateFarmingOrder(ctx sdk.Context, sender sdk.AccAddress, farmingOrderId string)
//   ExecuteFarmingOrders(ctx sdk.Context, sender sdk.AccAddress)
