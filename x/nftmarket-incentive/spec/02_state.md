# State

**NOTE: This is early draft.**

## FrontendIncentive

```protobuf
message FrontendIncentive {
  string frontend_name = 1;
  repeated string subjects = 2;
  repeated undetermined weights = 3;
}
```

- FrontendIncentive: `"frontend_name" -> format(FrontendIncentive)`

## RewardTable

RewardTable is the record of the rewards for the subject of the `frontend-incentive`.

```protobuf
message FrontendReward {
  string subject = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated undetermined rewards = 2;  
}
```

- RewardTable: `format(address) -> format(rewards)`
