# Keepers

## DerivativesKeeper

The important functions of the `DerivativesKeeper` are described below.

```go
type Keeper interface {

  // Get functions
  GetLPTokenBaseMintFee(ctx sdk.Context) sdk.Dec
  GetLPTokenBaseRedeemFee(ctx sdk.Context) sdk.Dec
  GetLPTokenSupplySnapshot(ctx sdk.Context, height int64) sdk.Int
  GetLPTokenSupply(ctx sdk.Context) sdk.Int
  GetLPTokenPrice(ctx sdk.Context) sdk.Dec
  GetRedeemDenomAmount(ctx sdk.Context, lptAmount sdk.Int, redeemDenom string) (sdk.Coin, sdk.Coin, error)
  GetAssetPrice(ctx sdk.Context, denom string) (*pftypes.CurrentPrice, error)
  GetPrice(ctx sdk.Context, lhsTicker string, rhsTicker string) (pftypes.CurrentPrice, error)
  GetParams(ctx sdk.Context) (params types.Params)
  GetCurrentPrice(ctx sdk.Context, denom string) (sdk.Dec, error)
  GetPairUsdPrice(ctx sdk.Context, base, quote string) (sdk.Dec, sdk.Dec, error)
  GetPairUsdPriceFromMarket(ctx sdk.Context, market types.Market) (sdk.Dec, sdk.Dec, error)
  GetPerpetualFuturesNetPositionOfMarket(ctx sdk.Context, market types.Market) types.PerpetualFuturesNetPositionOfMarket
  GetPositionSizeOfNetPositionOfMarket(ctx sdk.Context, market types.Market) sdk.Dec
  GetAllPerpetualFuturesNetPositionOfMarket(ctx sdk.Context) []type.PerpetualFuturesNetPositionOfMarket
  GetPoolAssets(ctx sdk.Context) []types.PoolParams_Asset
  GetPoolAssetByDenom(ctx sdk.Context, denom string) types.PoolParams_Asset
  GetAssetBalance(ctx sdk.Context, denom string) sdk.Coin
  GetAssetTargetAmount(ctx sdk.Context, denom string) (sdk.Coin, error)
  GetUserDeposits(ctx sdk.Context, depositor sdk.AccAddress) []sdk.Coin
  GetUserDenomDepositAmount(ctx sdk.Context, depositer sdk.AccAddress, denom string) sdk.Int
  GetPoolMarketCapSnapshot(ctx sdk.Context, height int64) types.PoolMarketCap
  GetPoolQuoteTicker(ctx sdk.Context) string
  GetPoolMarketCap(ctx sdk.Context) types.PoolMarketCap
}
```
