# Event

## EventRegister

An event to be emitted when to be registered an `incentive_unit` by being called `MsgRegister`.

```proto
message EventRegister {
  string incentive_id = 1 [
    (gogoproto.moretags) = "yaml:\"incentive_id\""
  ];
  repeated SubjectInfo subject_info_list = 2 [
    (gogoproto.moretags) = "yaml:\"subject_info_lists\"",
    (gogoproto.nullable) = false
  ];
}
```

An event to be emitted when to be withdrawn all rewards from a subject by being called `MsgWithdrawAllRewards`.

```proto
message EventWithdrawAllReward {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated cosmos.base.v1beta1.Coin rewards = 2 [
    (gogoproto.moretags) = "yaml:\"all_withdrawn_rewards\"",
    (gogoproto.nullable) = false
  ];
}
```

## EventWithdrawReward

An event to be emitted when to be withdrawn a specific reward with `denom` from a subject by being called `MsgWithdrawReward`.


```proto
message EventWithdrawReward {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  cosmos.base.v1beta1.Coin reward = 2 [
    (gogoproto.moretags) = "yaml:\"withdrawn_reward\"",
    (gogoproto.nullable) = false
  ];
}
```
