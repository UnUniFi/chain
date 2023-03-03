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

### QuoteTicker

`QuoteTicker` defines the ticker of the currency for the market cap to be calculated. The default value is `usd`.

### BaseLptMintFee

`BaseLptMintFee` defines fee ratio in parcentage for the minting DLP token by depositing some token.   
The default value is `0.001`.

### BaseLptRedeemFee

`BaseLptRedeemFee` defines fee ratio in parcentage for the redeeming DLP token by burning some token.   
The default value is `0.001`.

### BorrowingFeeRatePerHour

### ReportLiquidationRewardRate

### ReportLevyPeriodRewardRate

### AcceptedAssets

`AcceptedAssets` defines the tokens which can be deposited into a pool to get DLP.   
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

### CommissionRate

The default value is `0.001`.

### MarginMaintenanceRate

The default value is `0.5`.

### ImaginaryFundingRateProportionalCoefficient

The default value is `0.0005`.

### Markets

The `Markets` defines the available trading pair on the perpetual futures market.

## PerpetualOptioins

nothing is defined yet.
