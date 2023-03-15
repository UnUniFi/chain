# Keepers

## DerivativesKeeper

```go
type Keeper interface {
  DerivativesGetKeeper
  AddAssetToPool(ctx sdk.Context, asset Pool.Asset)
  AddAssetToSubpool(ctx sdk.Context, asset Subpool.Asset)
  SetAssetTargetWeight(ctx sdk.Context, weight )
}

```

## DerivativesGetKeeper

```go
type Keeper interface {
  GetPoolAssets(ctx sdk.Context)
  GetSubpoolAssets(ctx sdk.Context)
  GetPoolAsset(ctx sdk.Context, denom string)
  GetSubpoolAsset(ctx sdk.Context, denom string)
  GetBaseMintFee(ctx sdk.Context)
  GetBaseRedeemFee(ctx sdk.Context)
  GetBorrowingRateFeePerHour(ctx sdk.Context)
  GetMaximumLeverageRatio(ctx sdk.Context)
  GetAllOpenPositions(ctx sdk.Context)
  GetOpenPositionsOfAccount(ctx sdk.Context, account sdk.AccAddress)
}
```