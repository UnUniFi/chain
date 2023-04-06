# State

## GenesisState

```protobuf
message GenesisState {
  Params params = 1 [(gogoproto.nullable) = false];
  repeated Position positions = 2 [(gogoproto.nullable) = false];
  PoolMarketCap pool_market_cap = 3 [(gogoproto.nullable) = false];
  repeated PerpetualFuturesNetPositionOfMarket perpetual_futures_net_position_of_market = 4 [(gogoproto.nullable) = false];
}
```

`GenesisState` is the data structure of the genesis state of the derivatives module. It contains the following fields: `params`, `positions`, `pool_market_cap` and `perpetual_futures_net_position_of_market`.

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

## PerpetualFuturesNetPositionOfMarket

`PerpetualFuturesNetPositionOfMarket` is the data which contains the inclusive information regarding x/derivatives's PerpetualFutures market.

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
