# State

```protobuf
message Exchange {
  string contract_address = 1;
  repeated string denoms = 2;
  repeated string pairs = 3;
}
```

```protobuf
message AssetPair {
  string base_denom = 1;
  string quote_denom = 2;
}
```

```protobuf
enum SwapDirection {
  BaseToQuote = 1;
  QuoteToBase = 2;
}

message SwapRoute {
  string conract_address = 1;
  AssetPair pair = 2;
  SwapDirection direction = 3;
}
```
