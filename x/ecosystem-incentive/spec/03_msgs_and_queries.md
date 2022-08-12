# Messages and Queries

**NOTE: This is early draft.**

## Messages

All messages of `ecosystem-incentive`.

### Register

```protobuf
message MsgFrontendRegister {
  string incentive_id = 1;
  repeated string subjects = 2 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated undetermined weights = 3;
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

### WithdrawSpecificReward

A message to withdraw accumulated reward of specified denom.

```protobuf
message MsgWithdrawSpecificDenomReward {
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
message QueryIncentiveRequest {
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
