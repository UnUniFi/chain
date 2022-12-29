# Messages

```protobuf
message MsgSwap {
  string address = 1;
  SwapRoute route = 2;
  Coin amount = 3;
}

message MsgSwapRelay {
  string address = 1;
  repeated SwapRoute routes = 2;
  Coin amount = 3;
}
```
