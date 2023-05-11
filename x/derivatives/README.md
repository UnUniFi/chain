# DERIBATIVES

The `DERIBATIVES` module provides deriving functions

## Contents

1. **[Concepts](#concepts)**
1. **[State](#state)**
1. **[Keepers](#keepers)**
1. **[Messages And Queries](#messages_and_queries)**
1. **[Params](#params)**
1. **[Events](#events)**

## Concepts

The model of Perpetual Futures feature of this module totally follows GMX perpetual futures model.  
(ref: https://gmxio.gitbook.io/gmx/)

Briefly saying, the tradings on the perpetual futures market are supported by a unique multi-asset pool that earns liquidity prodicers fees from market making, swap fees and levrage trading.  
Because pool asset will take the counter part of the arbitral trade by a user, users can trade with high leverage and no slippage to the price from oracle with low fee.

## Pool

### Liquidity Provider Token

Our liquidity provider token's ticker is `DLP`.  
In the backend, it has `udlp` denom, which is the micro unit of `DLP`.

DLP consists of an index of assets used for leverage trading. It can be minted using the assets which the protocol supports and burnt to redeem any index asset. The price for minting and redemption is calculated based on the formulas in the WhitePaper in the section "3.1.1".  
WhitePaper: https://ununifi.io/assets/download/UnUniFi-Whitepaper.pdf

Fees earned on the platform are directly added to the pool.　 Therefore, DLP holders can benefit from them as a reward through the increasement of the DLP price.

There's dynamic change of the minting and redemption fee rate at this moment. There's the static rate which is defined in the protocol. And, the actual fee rate also consider the difference of asset proportion between target and actual proportion. The static base fee rate can be modified through the governace voting.
The potential range of those fee rate are below:

```text
mintFeeRate is proportion to max(0, (actualAmountInPool[i] - targetAmount[i]) / targetAmount[i])
redeemFeeRate is proportion to max(0, -(actualAmountInPool[i] - targetAmount[i]) / targetAmount[i])
```

One thing to be noted is that the Liquidity Pool will take the counterpart position of a trader’s order, so, if traders get profit, the pool and at the same time liquidity providers get loss.

## Perpetual Futures

### Position

User can open a perpetual futures position by sending `MsgOpenPosition` tx. There's no fee for opening a position.  
The position can be covered by two types of asset as margin, which are the tokens of the trading pair. If you trade 'BTC/USDC' pair, you can deposit BTC or USDC as margin. The profit will be distributed in the same token as the margin if there's some.  
The created position cannot be modified except for closing a whole in the current implementation.  
And, the liquidation is triggered against each position. The margin of the position cannot be added afterward now. But, this will be supported in the near future.  
The max leverage rate is defined in the params of the protocol for all trading pairs equially. This can be modified through the governance voting.

When a position is created, the corresponding amount of token in the pool will be allocated to the position to assure the distribution of the profit for the position. (This could be considered as lending)  
There's no fee as the default settings for borrowing at this moment. But, it can be modified through the governance voting.

### Liquidation

A position can be liquidated if the losses and fees reduces the margin to the point where:  
`remaining_margin / (position_size / leverage) <= MarginMaintenanceRate`
MaintenanceMarginRate is defined as a parameter in the protocol. The default value of it is `0.5`.  
The values are all based on `QuoteTicker` of the protocol, which the default value is `USD`.

This is achieved through any user seding a `MsgReportLiquidation` tx. The reporter gets the fee based on the remaining margin and ReportLiquidationRewardRate by the protocol. And the remaining amount of token will be sent back the position owner.  
There's no penalty for reporting the position that is not needed to be liquidated.

### Imaginary Funding Rates

To mitigate the effect of the feature of our perpetual futures model which the liquidity provider takes the counterpart of the trader, Imaginary Funding rate exists.  
If the net position of traders lean to one side, the imaginary funding rate work to make the net position of traders neutral. The neutral net position of traders means the neutral position of the pool an at the same time liquidity providers. In the perspective of economics, it can be expressed that this model unifies the conventional funding rate and the time cost of waiting for matchmaking to the imaginary funding rate.

Imaginary Funding are levied at every 8 hours by a reporter who send the `MsgReportLevyPeriod`. The reporter gets the reward based on the imaginary funding and ReportLevyPeriodRewardRate.

## Price Feed

Prices for the accepting token are provided through our pricefeed module.  
Our pricefeed module takes the price data from the restricted addresses which are defined in the protocol in advance.  
The token price is calculated like below:
The price of BTC in the pair of USDC,  
 `price_BTC = price_BTC_USD / price_USDC_USD`  
So, the pricefeed module has the data of BTC price based on USD and USDC price based on USD in this case.  
Price is calculated the meadian price of major exchanges for each token. Price will ideally be updated at every block. USDC or other stablecoins are not hard-coded in the protocol.

Price data is treated in this form:

```go
type CurrentPrice struct {
	  MarketId string                                 `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty" yaml:"market_id"`
	  Price    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
}
```

# State

## GenesisState

```protobuf
message GenesisState {
  Params params = 1 [(gogoproto.nullable) = false];
  repeated Position positions = 2 [(gogoproto.nullable) = false];
  PoolMarketCap pool_market_cap = 3 [(gogoproto.nullable) = false];
  repeated PerpetualFuturesGrossPositionOfMarket perpetual_futures_gross_position_of_market = 4 [(gogoproto.nullable) = false];
}
```

`GenesisState` is the data structure of the genesis state of the derivatives module. It contains the following fields: `params`, `positions`, `pool_market_cap` and `perpetual_futures_gross_position_of_market`.

These fields are made to be able to start the network from aribitrary genesis state.

### Params

Go to [05_params](https://github.com/UnUniFi/chain/blob/newDevelop/proto/derivatives/params.proto) page.

## Position

`Position` field is for the complete data of the all opening position.

```protobuf
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
  string opened_rate = 6 [
    (gogoproto.moretags) = "yaml:\"opened_rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  cosmos.base.v1beta1.Coin remaining_margin = 7 [
    (gogoproto.moretags) = "yaml:\"remaining_margin\"",
    (gogoproto.nullable) = false
  ];
  google.protobuf.Timestamp last_levied_at = 8 [
    (gogoproto.moretags) = "yaml:\"last_levied_at\"",
    (gogoproto.nullable) = false,
    (gogoproto.stdtime) = true
  ];
  google.protobuf.Any position_instance = 9 [
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
- `opened_rate` is the price rate of the trading pair at the opening.
- `remaining_margin` is the remaining margin for the position.
- `last_levied_at` is the timestamp of the last levied time.
- `position_instance` is the `Any` type which contains the actual position data. If it's about perpetual futures, it contains `PositionType`, `Size_`, `SizeInMicro` and `Leverage` fields.

## PoolMarketCap

`pool_market_cap` field is to contain `PoolMarketCap` data which contains the comprehensive data regarding pool situation for the derivatives.

```proto
message PoolMarketCap {
  message Breakdown {
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
  }

  string quote_ticker = 1 [
    (gogoproto.moretags) = "yaml:\"quote_ticker\""
  ];
  string total = 2 [
    (gogoproto.moretags) = "yaml:\"total\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  repeated Breakdown breakdown = 3 [
    (gogoproto.moretags) = "yaml:\"breakdown\"",
    (gogoproto.nullable) = false
  ];
}
```

- `quote_ticker` is the ticker of the key currency for the root value.
- `total` is the total value of the all pool assets combined.
- `breakdown` is the list of the `Breakdown` which contains the information regarding the specific pool asset.

## PerpetualFuturesGrossPositionOfMarket

`PerpetualFuturesGrossPositionOfMarket` is the data which contains the inclusive information regarding x/derivatives's PerpetualFutures market.

```protobuf
message PerpetualFuturesPositionInstance {
  PositionType position_type = 1 [
    (gogoproto.moretags) = "yaml:\"position_type\""
  ];
  string size = 2 [
    (gogoproto.moretags) = "yaml:\"size\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  // Use micro level size in the backend logic to be consistent with the scale of the coin amount
  // and price information.
  string size_in_micro = 3 [
    (gogoproto.moretags) = "yaml:\"size_in_micro\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int"
  ];
  uint32 leverage = 4 [
    (gogoproto.moretags) = "yaml:\"leverage\""
  ];
}
```

- `market` is the `Market` representive. `Market` itself is the information of the trading pair.
- `position_size` is the total net position size of the market. It tries to define the total cap of the opening position of the market.

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

# Messages_And_Queries

## Messages

### MintLiquidityProviderToken

[MintLiquidityProviderToken](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L22-L32)

Mint liquidity provider token `DLP` by providing the acceptable asset into the pool.  
The token's price is determined by the worth of all tokens within the pool and factoring in the profits and losses of all currently opened positions.  
Hence, the `DLP` amount of being minted is determined at the time of minting.  
Fee is charged based on the defined static param.

### BurnLiquidityProviderToken

[BurnLiquidityProviderToken](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L36-L50)

Burn liquidity provider token `DLP` to the arbitrary acceptable token.  
Fee is charged based on the defined static param.

### OpenPosition

[OpenPosition](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L54-L72)

Open a perpetual futures position.  
User defines the trading pair, long/short, position size and levarage rate.  
The maximum position size is limited by the amount of the corresponding token in the pool. User cannot take a position that is larger than the pool size.

### ClosePosition

[ClosePosition](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L76-L85)

Close a whole position by defining a unique position id.  
Only the owner of the position can close it. If the position has profit, the profit will be distributed in the same token of the position margin.

### ReportLiquidation

[ReportLiquidation](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L89-L103)

Report a position for which the margin maintenance ratio is below a certain level and need to be liquidated.
If a liquidation is made, the reporter gets a portion of the commission fee as a reward.
This architecture makes the chain avoidable to be aware of liquidation logic in the EndBlock handler to enhance the scalability.

### ReportLevyPeriod

[ReportLevyPeriod](https://github.com/UnUniFi/chain/blob/caf28770588ef1370f5ca8d58e9b17e2b131064b/proto/derivatives/tx.proto#L107-L121)

Report a position that have been in place for more than 8 hours since the last levy for correction of position bias.
Adds or subtracts the margin of the reported position depending on the overall position bias. In addition, the commission fee is subtracted from the margin. The commission fee rate is the defined static number in the params.
The reporter gets a portion of the commission fee as a reward.

## Queries

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

# Params

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

## PoolParams

```proto
message PoolParams {
  message Asset {
    string denom = 1 [
      (gogoproto.moretags) = "yaml:\"denom\""
    ];
    string target_weight = 2 [
      (gogoproto.moretags) = "yaml:\"target_weight\"",
      (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
      (gogoproto.nullable) = false
    ];
  }

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
  repeated Asset accepted_assets = 7 [
    (gogoproto.moretags) = "yaml:\"accepted_assets\""
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
- `AcceptedAssets` defines the tokens which can be deposited into a pool to get DLP.  
  The tokens in `AcceptedAssets` have to have `DenomMetadata` in bank module in this current implementation (could be changed).

## PerpetualFutures

```proto
message PerpetualFuturesParams {
  string commission_rate = 1 [
    (gogoproto.moretags) = "yaml:\"commission_rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  string margin_maintenance_rate = 2 [
    (gogoproto.moretags) = "yaml:\"margin_maintenance_rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  string imaginary_funding_rate_proportional_coefficient = 3 [
    (gogoproto.moretags) = "yaml:\"imaginary_funding_rate_proportonal_coefficient\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
  repeated Market markets = 4 [
    (gogoproto.moretags) = "yaml:\"markets\""
  ];
}
```

- `CommissionRate` is the fee for trading. It's taken when to close a position. The default value is `0.001`.
- `MarginMaintenanceRate` is used for the determination of the liquidation condition. The default value is `0.5`.
- `ImaginaryFundingRateProportionalCoefficient` is the fee ratio for the imaginary funding. The default value is `0.0005`.
- `Markets` defines the available trading pair on the perpetual futures market.

## PerpetualOptions

nothing is defined yet.

# Events

- [EventPriceIsNotFeeded](https://github.com/UnUniFi/chain/blob/0dc4f717a4ef3e4b32731069d1dba503babe5998/proto/derivatives/derivatives.proto#L175-L179)
- [EventPerpetualFuturesPositionOpened](https://github.com/UnUniFi/chain/blob/0dc4f717a4ef3e4b32731069d1dba503babe5998/proto/derivatives/perpetual_futures.proto#L105-L108)
- [EventPerpetualFuturesPositionClosed](https://github.com/UnUniFi/chain/blob/0dc4f717a4ef3e4b32731069d1dba503babe5998/proto/derivatives/perpetual_futures.proto#L110-L115)
- [EventPerpetualFuturesPositionLiquidated](https://github.com/UnUniFi/chain/blob/0dc4f717a4ef3e4b32731069d1dba503babe5998/proto/derivatives/perpetual_futures.proto#L117-L122)
- [EventPerpetualFuturesPositionLevied](https://github.com/UnUniFi/chain/blob/0dc4f717a4ef3e4b32731069d1dba503babe5998/proto/derivatives/perpetual_futures.proto#L124-L129)

## EventPriceIsNotFeeded

This event is emitted when at least one necessary price is not feeded in the pricefeed module to be referenced in the derivatives module.
