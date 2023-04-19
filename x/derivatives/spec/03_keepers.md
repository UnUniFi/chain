# Keepers

## DerivativesKeeper

The important functions of the `DerivativesKeeper` are described below.

```go
type Keeper interface {
  // Get functions
  GetLPTokenBaseMintFee(ctx sdk.Context) sdk.Dec
  GetLPTokenBaseRedeemFee(ctx sdk.Context) sdk.Dec
  GetLPTokenSupply(ctx sdk.Context) sdk.Int
  GetLPTokenPrice(ctx sdk.Context) sdk.Dec
  GetRedeemDenomAmount(ctx sdk.Context, lptAmount sdk.Int, redeemDenom string) (sdk.Coin, sdk.Coin, error)
  GetPairUsdPriceFromMarket(ctx sdk.Context, market types.Market) (sdk.Dec, sdk.Dec, error)
  GetPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market, positionType types.PositionType) types.PerpetualFuturesNetPositionOfMarket
  GetPositionSizeOfNetPositionOfMarket(ctx sdk.Context, market types.Market) sdk.Dec
  GetAssetBalanceInPoolByDenom(ctx sdk.Context, denom string) sdk.Coin
  GetAssetTargetAmount(ctx sdk.Context, denom string) (sdk.Coin, error)
  GetPoolMarketCapSnapshot(ctx sdk.Context, height int64) types.PoolMarketCap
  GetPoolMarketCap(ctx sdk.Context) types.PoolMarketCap
  GetLastPositionId(ctx sdk.Context) string
  GetAddressPositions(ctx sdk.Context, user sdk.AccAddress) []*types.Position
  GetLPTPriceFromSnapshot(ctx sdk.Context, height int64) sdk.Dec
  GetLPNominalYieldRate(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec
  GetInflationRateOfAssetsInPool(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec
  GetLPRealYieldRate(ctx sdk.Context, beforeHeight int64, afterHeight int64) sdk.Dec
}
```
