# Keepers

## AssetManagementAccountKeeper

```go
type Keeper interface {
  AssetManagementAccountGetKeeper
  AddAssetManagementAccounts(ctx sdk.Context, id string, name string)
  UpdateAssetManagementAccounts(ctx sdk.Context, id string, obj types.AssetManagementAccount)
  DeleteAssetManagementAccounts(ctx sdk.Context, id string)
  AddAssetManagementTargetsOfAccount(ctx sdk.Context, account_id string, obj types.AssetManagementTarget)
  UpdateAssetManagementTargetsOfAccount(ctx sdk.Context, targetId string, obj types.AssetManagementTarget)
  DeleteAssetManagementTargetsOfAccount(ctx sdk.Context, targetId string)

}

```

## AssetManagementAccountGetKeeper

```go
type Keeper interface {
  GetAssetManagementAccounts(ctx sdk.Context)
  GetAssetManagementTargetsOfAccount(ctx sdk.Context, accountId string)
  GetAssetManagementTargetsOfDenom(ctx sdk.Context, accountId string, denom string)
}

```

## AssetManagementAccountBankKeeper

```go
type Keeper interface {
  PayBack(ctx sdk.Context, targetId string, farmingUnit FarmingUnit)
}

```

## AssetManagementKeeper

```go
type Keeper interface {
  Deposit(ctx sdk.Context, sdk.AccAddress, amount sdk.Coins)
  Withdraw(ctx sdk.Context, sender sdk.AccAddress, amount sdk.Coins)
  AddFarmingOrder(ctx sdk.Context, farmingOrder FarmingOrder)
  DeleteFarmingOrder(ctx sdk.Context, sender sdk.AccAddress, farmingOrderId string)
  GetFarmingOrdersOfAddress(ctx sdk.Context, sender sdk.AccAddress)
  ActivateFarmingOrder(ctx sdk.Context, sender sdk.AccAddress, farmingOrderId string)
  InactivateFarmingOrder(ctx sdk.Context, sender sdk.AccAddress, farmingOrderId string)
  ExecuteFarmingOrders(ctx sdk.Context, sender sdk.AccAddress)
}

```
