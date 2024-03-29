# DERIVATIVES

The `DERIVATIVES` module provides deriving functions

## Contents

1. **[Concepts](#concepts)**
1. **[State](#state)**
1. **[Keepers](#keepers)**
1. **[Messages And Queries](#messages_and_queries)**
1. **[Params](#params)**
1. **[Events](#events)**

## Concepts

The model of Perpetual Futures feature of this module totally follows GMX perpetual futures model.  
(ref: <https://gmxio.gitbook.io/gmx/>)

Briefly saying, the tradings on the perpetual futures market are supported by a unique multi-asset pool that earns liquidity producers fees from market making, swap fees, and leverage trading.  
Because pool asset will take the counter part of the arbitral trade by a user, users can trade with high leverage and no slippage to the price from oracle with low fee.

## Pool

### Liquidity Provider Token

Our liquidity provider token's ticker is `DLP`.  
In the backend, it has `udlp` denom, which is the micro unit of `DLP`.

DLP consists of an index of assets used for leverage trading. It can be minted using the assets which the protocol supports and burnt to redeem any index asset. The price for minting and redemption is calculated based on the formulas in the WhitePaper in the section "3.1.1".  
WhitePaper: <https://ununifi.io/assets/download/UnUniFi-Whitepaper.pdf>

Fees earned on the platform are directly added to the pool.　 Therefore, DLP holders can benefit from them as a reward through the increasement of the DLP price.

There's dynamic change of the minting and redemption fee rate at this moment. There's the static rate which is defined in the protocol. And, the actual fee rate also consider the difference of asset proportion between target and actual proportion. The static base fee rate can be modified through the governace voting.
The potential range of those fee rate are below:

```text
mintFeeRate is proportion to max(0, (actualAmountInPool[i] - targetAmount[i]) / targetAmount[i])
redeemFeeRate is proportion to max(0, -(actualAmountInPool[i] - targetAmount[i]) / targetAmount[i])
```

One thing to be noted is that the Liquidity Pool will take the counterpart position of a trader’s order, so, if traders get profit, the pool and at the same time liquidity providers get loss.

## Perpetual Futures

### Position (Perpetual Future)

Users can open a perpetual futures position by sending `MsgOpenPosition` tx. There's no fee for opening a position.  
The position can be covered by two types of assets as margin, which are the tokens of the trading pair. If you trade 'BTC/USDC' pair, you can deposit BTC or USDC as a margin. The profit will be distributed in the same token as the margin if there's some.  
The created position cannot be modified except for closing a whole in the current implementation.  
And, the liquidation is triggered against each position. The margin of the position cannot be added afterward now. But, this will be supported in the near future.  
The max leverage rate is defined in the params of the protocol for all trading pairs equally. This can be modified through governance voting.

When a position is created, the corresponding amount of token in the pool will be allocated to the position to assure the distribution of the profit for the position. (This could be considered as lending)  
There's no fee as the default settings for borrowing at this moment. But, it can be modified through governance voting.

### Liquidation

A position will be liquidated if the losses and fees reduce the margin to the point where:  
`EffectiveMargin / (Rate * Price / Leverage) <= MarginMaintenanceRate`
`EffectiveMargin = MArgin + Profit (- Loss) - LevyPeriodFee`

MaintenanceMarginRate is defined as a parameter in the protocol. The default value of it is `0.5`.  
The values are all based on `QuoteTicker` of the protocol, in which the default value is `USD`.

This is achieved through any user sending a `MsgReportLiquidation` tx. If the liquidation is needed, The reporter gets a portion of the commission fee based on the remaining margin. The remaining margin is then returned to the position owner after deducting a commission fee.

`Report Liquidation Fee = PositionSize * CommissionRate`
The default value of `CommissionRate` is `0.001`.

A portion of the fee is paid to the reporter. The default value of `ReportLiquidationRewardRate` is `0.3`.
There's no penalty for reporting a position that is not needed to be liquidated.

### Levy Period

Levy Period is set to reduce the imbalance in the positions of the entire market.
If the Long positions are biased in the entire market, a imaginary funding fee will be deducted from the margin of the Long positions and added to the margin of the Short positions. At the same time, commission fees are also subtracted.

`ImaginaryFunding_fee = PositionSize * ImaginaryFundingRate`
`ImaginaryFundingRate = ImaginaryFundingCoefficient * NetPosition (long - short) / TotalPosition (long + short)`

The default value of `imaginary_funding_coefficient` is `0.0005`.

The calculation method for a commission fee is the same as that for the Liquidation.
The default value of `ReportLevyPeriodRewardRate` is `0.3`.

From the perspective of economics, it can be expressed that this model unifies the conventional funding rate and the time cost of waiting for matchmaking to the imaginary funding rate.

This is achieved through any user sending a `MsgReportLevyPeriod` tx.
If more than 8 hours have passed since the last levy period, this process occurs and the reporter gets a portion of the commission fee based on the remaining margin.

There's no penalty for reporting a position that is not needed to be levied.

## Price Feed

Prices for the accepting token are provided through our pricefeed module.  
Our pricefeed module takes the price data from the restricted addresses which are defined in the protocol in advance.  
The token price is calculated like below:
The price of BTC in the pair of USDC,  
 `price_BTC = price_BTC_USD / price_USDC_USD`  
So, the pricefeed module has the data of BTC price based on USD and USDC price based on USD in this case.  
Price is calculated as the median price of major exchanges for each token. Price will ideally be updated at every block. USDC or other stablecoins are not hard-coded in the protocol.

Price data is treated in this form:

```go
type CurrentPrice struct {
  MarketId string                                 `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty" yaml:"market_id"`
  Price    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
}
```

## State

### GenesisState

```protobuf
message GenesisState {
  Params params = 1 [(gogoproto.nullable) = false];
  repeated Position positions = 2 [(gogoproto.nullable) = false];
  PoolMarketCap pool_market_cap = 3 [(gogoproto.nullable) = false];
  repeated PerpetualFuturesGrossPositionOfMarket perpetual_futures_gross_position_of_market = 4 [(gogoproto.nullable) = false];
}
```

`GenesisState` is the data structure of the genesis state of the derivatives module. It contains the following fields: `params`, `positions`, `pool_market_cap` and `perpetual_futures_gross_position_of_market`.

These fields are made to be able to start the network from an arbitrary genesis state.

### Params

Go to [05_params](https://github.com/UnUniFi/chain/blob/main/proto/derivatives/params.proto) page.

## Position

`Position` field is for the complete data of the all opening position.

```protobuf
enum PositionType {
  POSITION_UNKNOWN = 0;
  LONG = 1;
  SHORT = 2;
}

message Position {
  string id = 1 [
    (gogoproto.moretags) = "yaml:\"id\""
  ];
  Market market = 2 [
    (gogoproto.moretags) = "yaml:\"market\"",
    (gogoproto.nullable) = false
  ];
  string address = 3 [
    (gogoproto.moretags) = "yaml:\"address\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  google.protobuf.Timestamp opened_at = 4 [
    (gogoproto.moretags) = "yaml:\"opened_at\"",
    (gogoproto.nullable) = false,
    (gogoproto.stdtime) = true
  ];
  uint64 opened_height = 5 [
    (gogoproto.moretags) = "yaml:\"opened_height\""
  ];
  string opened_base_rate = 6 [
    (gogoproto.moretags) = "yaml:\"opened_base_rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  string opened_quote_rate = 7 [
    (gogoproto.moretags) = "yaml:\"opened_quote_rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  cosmos.base.v1beta1.Coin remaining_margin = 8 [
    (gogoproto.moretags) = "yaml:\"remaining_margin\"",
    (gogoproto.nullable) = false
  ];
  cosmos.base.v1beta1.Coin remaining_margin = 7 [
    (gogoproto.moretags) = "yaml:\"remaining_margin\"",
    (gogoproto.nullable) = false
  ];
    cosmos.base.v1beta1.Coin levied_amount = 9 [
    (gogoproto.moretags) = "yaml:\"levied_amount\"",
    (gogoproto.nullable) = false
  ];
  bool levied_amount_negative = 10 [
    (gogoproto.moretags) = "yaml:\"levied_amount_negative\"",
    (gogoproto.nullable) = false
  ];
  google.protobuf.Timestamp last_levied_at = 11 [
    (gogoproto.moretags) = "yaml:\"last_levied_at\"",
    (gogoproto.nullable) = false,
    (gogoproto.stdtime) = true
  ];
  google.protobuf.Any position_instance = 12 [
    (gogoproto.moretags) = "yaml:\"position_instance\"",
    (gogoproto.nullable) = false
  ];

}
```

- `id` is the unique identifier of the position.
- `market` defines the trading pair of the position.
- `address` is the address of the position owner.
- `opened_at` is the timestamp of the position opening.
- `opened_height` is the block height of the position opening.
- `opened_base_rate` is the price rate of the base denom at the opening.
- `opened_quote_rate` is the price rate of the quote denom at the opening.
- `remaining_margin` is the remaining margin for the position.
- `levied_amount` is the total of the Levy Period fees, plus the ReportLiquidation fee, if liquidated.
- `levied_amount_negative` is levied_amount is negative or not. Basically true, but false depending on LevyPeriod position bias.
- `last_levied_at` is the timestamp of the last levied time. Initialized to BlockTime when the position is opened
- `position_instance` is the `Any` type which contains the actual position data. If it's about perpetual futures, it contains `PositionType`, `Size_`, and `Leverage` fields.

```protobuf
message PerpetualFuturesPositionInstance {
  PositionType position_type = 1 [(gogoproto.moretags) = "yaml:\"position_type\""];
  string       size          = 2 [
    (gogoproto.moretags)   = "yaml:\"size\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
  uint32 leverage = 3 [(gogoproto.moretags) = "yaml:\"leverage\""];
}
```

## PoolMarketCap

`pool_market_cap` field is to contain `PoolMarketCap` data which contains the comprehensive data regarding pool situation for the derivatives.

```proto
message PoolMarketCap {
  message AssetInfo {
    string denom = 1 [
      (gogoproto.moretags) = "yaml:\"denom\""
    ];
    string amount = 2 [
      (gogoproto.moretags) = "yaml:\"amount\"",
      (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
      (gogoproto.nullable) = false
    ];
    string price = 3 [
      (gogoproto.moretags) = "yaml:\"price\"",
      (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
      (gogoproto.nullable) = false
    ];
    string reserved = 4 [
      (gogoproto.moretags) = "yaml:\"reserved\"",
      (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
      (gogoproto.nullable) = false
    ];
  }

  string quote_ticker = 1 [
    (gogoproto.moretags) = "yaml:\"quote_ticker\""
  ];
  string total = 2 [
    (gogoproto.moretags) = "yaml:\"total\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  repeated AssetInfo asset_info = 3 [
    (gogoproto.moretags) = "yaml:\"asset_info\"",
    (gogoproto.nullable) = false
  ];
}
```

- `quote_ticker` is the ticker of the key currency for the root value.
- `total` is the total value of the all pool assets combined.
- `asset_info` is the list of the information regarding the specific pool asset.

## PerpetualFuturesGrossPositionOfMarket

`PerpetualFuturesGrossPositionOfMarket` is the data which contains the inclusive information regarding x/derivatives's PerpetualFutures market.

```protobuf
message PerpetualFuturesGrossPositionOfMarket {
  Market       market        = 1 [(gogoproto.moretags) = "yaml:\"market\"", (gogoproto.nullable) = false];
  PositionType position_type = 2 [(gogoproto.moretags) = "yaml:\"position_type\""];
  string       position_size_in_denom_exponent = 3 [
    (gogoproto.moretags)   = "yaml:\"position_size_in_denom_exponent\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.nullable)   = false
  ];
}
```

- `market` is the information of the trading pair.
- `position_type` is the position type. LONG or SHORT.
- `position_size_in_denom_exponent` is the gross position size of the market.

NetPosition and TotalPosition are calculated using these gross positions.
`NetPosition = GrossPosition(LONG) - GrossPosition(SHORT)`
`TotalPosition = GrossPosition(LONG) + GrossPosition(SHORT)`

## Reserve

Reserve is to pay the profit in the result of the trade. It works as the counter-position of the trader's. It's separated from the deposit as the collateral (margin) from a trader. It's all from the pool assets.

**The detail of the data structure, name and the related functions are not configured on code basis yet. So, note that the following description is not the final version.**

```protobuf
enum MarketType {
  UNKNOWN = 0;
  FUTURES = 1;
  OPTIONS = 2;
}

message Reserve {
  MarketType market_type = 1 [
    (gogoproto.moretags) = "yaml:\"market_type\""
  ];
  cosmos.base.v1beta1.Coin amount = 2 [
    (gogoproto.moretags) = "yaml:\"amount\"",
    (gogoproto.nullable) = false
  ];
}

// The list of the `Reserve` will be contained in `GenesisState`.
```

And here's the important functions' interfaces relating `Reserve`.

```go
// In keeper package
func (k Keeper) GetReserve(ctx sdk.Context, marketType types.MarketType, denom string) (types.Reserve, error)
func (k Keeper) GetReservesByDenom(ctx sdk.Context, denom string) []types.Reserve
func (k Keeper) GetAllReserves(ctx sdk.Context) []types.Reserve
func (k Keeper) SetReserve(ctx sdk.Context, reserve types.Reserve) error

// In types package
func ReserveKeyPrefix(marketType MarketType, denom string) []byte
func TotalReserveValueInMetrics(reserves []types.Reserve, metrics string) (sdk.Dec, error)
func (m Reserve) ModifyReserveAmount(amount sdk.Int) (Reserve, error)
```

## Keepers

### DerivativesKeeper

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
  GetPerpetualFuturesGrossPositionOfMarket(ctx sdk.Context, market types.Market, positionType types.PositionType) types.PerpetualFuturesGrossPositionOfMarket
  GetPositionSizeOfGrossPositionOfMarket(ctx sdk.Context, market types.Market) sdk.Dec
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

## Messages_And_Queries

### Messages

#### DepositToPool

[DepositToPool](https://github.com/UnUniFi/chain/blob/main/proto/ununifi/derivatives/tx.proto)

MsgDepositToPool deposits tokens into the pool and mint liquidity provider token `DLP`.  
The token's price is determined by the worth of all tokens within the pool and factoring in the profits and losses of all currently opened positions.  
Hence, the `DLP` amount of being minted is determined at the time of minting.  
Fee is charged based on the defined static param.

#### WithdrawFromPool

[MsgWithdrawFromPool](https://github.com/UnUniFi/chain/blob/main/proto/ununifi/derivatives/tx.proto)

MsgWithdrawFromPool withdraws tokens from the pool and burn liquidity provider token `DLP`.  
Fee is charged based on the defined static param.

#### OpenPosition

[MsgOpenPosition](https://github.com/UnUniFi/chain/blob/main/proto/ununifi/derivatives/tx.proto)

Open a perpetual futures position.  
User defines the trading pair, long/short, position size and leverage rate.  
The maximum position size is limited by the amount of the corresponding token in the pool. User cannot take a position that is larger than the pool size.

#### ClosePosition

[MsgClosePosition](https://github.com/UnUniFi/chain/blob/main/proto/ununifi/derivatives/tx.proto)

Close a whole position by defining a unique position id.  
Only the owner of the position can close it. If the position has profit, the profit will be distributed in the same token of the position margin.

#### ReportLiquidation

[MsgReportLiquidation](https://github.com/UnUniFi/chain/blob/main/proto/ununifi/derivatives/tx.proto)

Report a position for which the margin maintenance ratio is below a certain level and need to be liquidated.
If a liquidation is made, the reporter gets a portion of the commission fee as a reward.
This architecture makes the chain avoidable to be aware of liquidation logic in the EndBlock handler to enhance the scalability.

#### ReportLevyPeriod

[ReportLevyPeriod](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L107-L121)

Report a position that have been in place for more than 8 hours since the last levy for correction of position bias.
Adds or subtracts the margin of the reported position depending on the overall position bias. In addition, the commission fee is subtracted from the margin. The commission fee rate is the defined static number in the params.
The reporter gets a portion of the commission fee as a reward.

### Queries

The derivatives module primarily provides the following queries:

- [Params](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L79-L85)
- [Pool](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L88-L106)
- [LiquidityProviderTokenRealAPY](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L108-L122)
- [LiquidityProviderTokenNominalAPY](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L124-L138)
- [PerpetualFutures](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L140-L162)
- [PerpetualFuturesMarket](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L164-L197)
- [AllPositions](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L214-L224)
- [Position](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L226-L249)
- [PerpetualFuturesPositionSize](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L251-L265)
- [AddressPositions](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L267-L280)
- [DLPTokenRates](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L283-L292)
- [EstimateDLPTokenAmount](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L294-L312)
- [EstimateRedeemTokenAmount](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/query.proto#L314-L332)

## Parameters

`Params` is included in `GenesisState`. It has below three properties which will be explaned in each section.

```proto
message Params {
  PoolParams pool_params = 1 [
    (gogoproto.moretags) = "yaml:\"pool_params\"",
    (gogoproto.nullable) = false
  ];
  PerpetualFuturesParams perpetual_futures = 2 [
    (gogoproto.moretags) = "yaml:\"perpetual_futures\"",
    (gogoproto.nullable) = false
  ];
  PerpetualOptionsParams perpetual_options = 3 [
    (gogoproto.moretags) = "yaml:\"perpetual_options\"",
    (gogoproto.nullable) = false
  ];
}
```

### PoolParams

```proto
message PoolAssetConf {
  string denom = 1 [
    (gogoproto.moretags) = "yaml:\"denom\""
  ];
  string target_weight = 2 [
    (gogoproto.moretags) = "yaml:\"target_weight\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
}

message PoolParams {
  string quote_ticker = 1 [
    (gogoproto.moretags) = "yaml:\"quote_ticker\""
  ];
  string base_lpt_mint_fee = 2 [
    (gogoproto.moretags) = "yaml:\"base_lpt_mint_fee\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  string base_lpt_redeem_fee = 3 [
    (gogoproto.moretags) = "yaml:\"base_lpt_redeem_fee\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  string borrowing_fee_rate_per_hour = 4 [
    (gogoproto.moretags) = "yaml:\"borrowing_fee_rate_per_hour\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  string report_liquidation_reward_rate = 5 [
    (gogoproto.moretags) = "yaml:\"report_liquidation_reward_rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  string report_levy_period_reward_rate = 6 [
    (gogoproto.moretags) = "yaml:\"report_levy_period_reward_rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  repeated PoolAssetConf accepted_assets_conf = 7 [
    (gogoproto.moretags) = "yaml:\"accepted_assets_conf\"",
    (gogoproto.nullable) = false
  ];
}
```

- `QuoteTicker` defines the ticker of the currency for the market cap to be calculated. The default value is `usd`.
- `BaseLptMintFee` defines fee ratio in parcentage for the minting DLP token by depositing some token.  
  The default value is `0.001`.
- `BaseLptRedeemFee` defines fee ratio in parcentage for the redeeming DLP token by burning some token.  
  The default value is `0.001`.
- `BorrowingFeeRatePerHour` defines fee ratio for the borrowing token from the pool to the traders.
  The default value is `0.000001`.
- `ReportLiquidationRewardRate` defines reward ratio for the reporting the liquidation of the position for a reporter. The reward is the commission fee multiplied by this rate.
  The default value is `0.3`.
- `ReportLevyPeriodRewardRate` defines reward ratio for the reporting the levy period for a reporter. The reward is the commission fee multiplied by this rate.
  The default value is `0.3`.
- `AcceptedAssetsConf` defines the tokens which can be deposited into a pool to get DLP.  
  The tokens in `AcceptedAssets` have to have `DenomMetadata` in bank module in this current implementation (could be changed).
- `LevyPeriodRequiredSeconds` defines the required time for the next Levy Period. the default value is `28800`(8 hours)

### PerpetualFuturesParams

```proto
message PerpetualFuturesParams {
  string commission_rate = 1 [
    (gogoproto.moretags)   = "yaml:\"commission_rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
  string margin_maintenance_rate = 2 [
    (gogoproto.moretags)   = "yaml:\"margin_maintenance_rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
  string imaginary_funding_rate_proportional_coefficient = 3 [
    (gogoproto.moretags)   = "yaml:\"imaginary_funding_rate_proportonal_coefficient\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable)   = false
  ];
  repeated Market markets      = 4 [(gogoproto.moretags) = "yaml:\"markets\""];
  uint32          max_leverage = 5 [(gogoproto.moretags) = "yaml:\"max_leverage\""];
}
```

- `CommissionRate` is the fee for trading. It's taken when to close a position. The default value is `0.001`.
- `MarginMaintenanceRate` is used for the determination of the liquidation condition. The default value is `0.5`.
- `ImaginaryFundingRateProportionalCoefficient` is the fee ratio for the imaginary funding. The default value is `0.0005`.
- `Markets` defines the available trading pair on the perpetual futures market.
- `MaxLeverage` is the maximum leverage allowed. The default value is `30`.

### PerpetualOptionsParams

nothing is defined yet.

## Events

- [EventPriceIsNotFed](README.md#EventPriceIsNotFed)
- [EventPerpetualFuturesPositionOpened](README.md#EventPerpetualFuturesPositionOpened)
- [EventPerpetualFuturesPositionClosed](README.md#EventPerpetualFuturesPositionClosed)
- [EventPerpetualFuturesPositionLiquidated](README.md#EventPerpetualFuturesPositionLiquidated)
- [EventPerpetualFuturesPositionLevied](README.md#EventPerpetualFuturesPositionLevied)
- [EventPerpetualFuturesLiquidationFee](README.md#EventPerpetualFuturesLiquidationFee)
- [EventPerpetualFuturesLevyFee](README.md#EventPerpetualFuturesLevyFee)
- [EventPerpetualFuturesImaginaryFundingFee](README.md#EventPerpetualFuturesImaginaryFundingFee)

### EventPriceIsNotFed

This event is emitted when at least one necessary price is not fed in the pricefeed module to be referenced in the derivatives module.

## Events for the fee earned by the protocol

The following events are emitted when the protocol earns the fee from the traders.

### EventPerpetualFuturesLevyFee

This event is emitted when the protocol earns the commission fee from the traders in LevyPeriod.

```protobuf
message EventPerpetualFuturesLevyFee {
  cosmos.base.v1beta1.Coin fee = 1 [
    (gogoproto.moretags) = "yaml:\"fee\"",
    (gogoproto.nullable) = false
  ];
  string position_id = 2 [
    (gogoproto.moretags) = "yaml:\"position_id\""
  ];
}
```

### EventPerpetualFuturesImaginaryFundingFee

This event is emitted when the protocol earns the imaginary funding fee from the traders in LevyPeriod.
The imaginary funding fee can be minus if the protocol funds the position.
In that case, the `fee_negative` property has `false`.

```protobuf
message EventPerpetualFuturesImaginaryFundingFee {
  cosmos.base.v1beta1.Coin fee = 1 [
    (gogoproto.moretags) = "yaml:\"fee\"",
    (gogoproto.nullable) = false
  ];
  bool fee_negative = 3 [
    (gogoproto.moretags) = "yaml:\"fee_negative\""
  ];
  string position_id = 2 [
    (gogoproto.moretags) = "yaml:\"position_id\""
  ];
}
```

### EventPerpetualFuturesLiquidationFee

This event is emitted when the protocol earns the liquidation fee from the traders when the position is liquidated.

```protobuf
message EventPerpetualFuturesLiquidationFee {
  cosmos.base.v1beta1.Coin fee = 1 [
    (gogoproto.moretags) = "yaml:\"fee\"",
    (gogoproto.nullable) = false
  ];
  string position_id = 2 [
    (gogoproto.moretags) = "yaml:\"position_id\""
  ];
}
```
