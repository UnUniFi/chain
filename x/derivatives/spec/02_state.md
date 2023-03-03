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

### Params

Go to 05_params page.

### Position

### PerpetualFuturesNetPositionOfMarket
