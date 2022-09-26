# Messages and Queries

**NOTE: This is early draft.**

## Messages

All messages of `ecosystem-incentive`.

### Register

```protobuf
message MsgRegister {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string incentive_id = 2;
  repeated string subject_addrs = 3 [
    (gogoproto.moretags) = "yaml:\"subject\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated string weights = 4 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.moretags) = "yaml:\"weight\"",
    (gogoproto.nullable) = false
  ];
}
```

or possibly take json file

### WithdrawAllRewards

A message to withdraw all accumulated rewards across all denoms.

```protobuf
message MsgWithdrawAllRewards {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
}
```

### WithdrawReward

A message to withdraw accumulated reward of specified denom.

```protobuf
message MsgWithdrawReward {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string denom = 2;
}
```

## Queries

All queries of `ecosystem-incentive`.

### IncentiveStore

```protobuf
message QueryIncentiveUnitRequest {
  string incentive_id = 1;
}
```

### AllRewards

```protobuf
message QueryAllRewardsRequest {
  string subject = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
}
```

### SpecificDenomReward

```protobuf
message QuerySpecificDenomRewardRequest {
  string subject = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string denom = 2;
}
