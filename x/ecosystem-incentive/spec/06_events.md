# Event

## EventRegister

An event to be emitted when to be registered an `incentive_unit` by being called `MsgRegister`.

```proto
message EventRegister {
  string incentive_unit_id = 1 [
    (gogoproto.moretags) = "yaml:\"incentive_unit_id\""
  ];
  repeated SubjectInfo subject_info_list = 2 [
    (gogoproto.moretags) = "yaml:\"subject_info_lists\"",
    (gogoproto.nullable) = false
  ];
}
```

## EventWithdrawAllRewards

An event to be emitted when to be withdrawn all rewards from a subject by being called `MsgWithdrawAllRewards`.

```proto
message EventWithdrawAllRewards {
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

## EventFailedParsingMemoInputs

An event to be emitted to inform that pasing the memo data which is put in, e.g. MsgListNft.

```proto
message EventFailedParsingMemoInputs {
  string class_id = 1 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 2 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
  string memo = 3 [ (gogoproto.moretags) = "yaml:\"memo\"" ];
}
```

## EventRecordedIncentiveUnitId

An event to be emmitted to inform the NFT ID which is listed on UnUniFi is intentionally associated with the Incentive Unit Id which was put in the Memo data.

```proto
message EventRecordedIncentiveUnitId {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  string class_id = 2 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 3 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}
```

## EventDeletedNftIdRecordedForFrontendReward

An event to be emmitted to inform that the NFT ID is deleted, following the NFT is unlisted from UnUniFi NFT market.

```proto
message EventDeletedNftIdRecordedForFrontendReward {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  string class_id = 2 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 3 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}
```

## EventNotRegisteredIncentiveUnitId

An event to be emmitted to inform recording Incentive Unit Id is failed.

```proto
message EventNotRegisteredIncentiveUnitId {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  string class_id = 2 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 3 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}
```

## EventNotRecordedNftId

An event to be emmitted inform the NFT Id is not recorded to be subject to the Ecosystem Incentive reward.

```proto
message EventNotRecordedNftId {
  string class_id = 1 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 2 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}
```

## EventUpdatedReward

An event to be emmitted inform to subject that the reward amount is updated.

```proto
message EventUpdatedReward {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  cosmos.base.v1beta1.Coin reward = 2 [
    (gogoproto.moretags) = "yaml:\"incentive_unit_id\"",
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.nullable) = false
  ];
}
```

## EventVersionUnmatched

An event to be emmitted to inform the version input in Memo data was wrong.

```protobuf
message EventVersionUnmatched {
  string unmatched_version = 1 [ (gogoproto.moretags) = "yaml:\"unmatched_version\"" ];
  string class_id = 2 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 3 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}
```
