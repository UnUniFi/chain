# Messages and Queries

**NOTE: This is early draft.**

## Messages

All messages of `frontend-incentive`.

### Frontend Register

```protobuf
message MsgFrontendRegister {
  string frontend_name = 1;
  repeated string subjects = 2 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated undetermined weights = 3;
}
```

or possibly take json file

### WithdrawAllFrontendReward

A message to withdraw all accumulated rewards across all denoms.

```protobuf
message MsgWithdrawAllFrontendReward {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
}
```

### WithdrawSpecificFrontendReward

A message to withdraw accumulated reward of specified denom.

```protobuf
message MsgWithdrawSpecificFrontendReward {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string denom = 2;
}
```

## Queries

All queries of `frontend-incentive`.

### FrontendIncentive

```protobuf
message QueryFrontendIncentiveRequest {
  string frontend_name = 1;
}
```

### AllFrontendReward

```protobuf
message QueryAllFrontendRewardRequest {
  string subject = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
}
```

### SpecificDenomFrontendReward

```protobuf
message QuerySpecificDenomFrontendRewardRequest {
  string subject = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string denom = 2;
}
