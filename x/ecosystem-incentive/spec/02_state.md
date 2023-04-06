# State

**NOTE: This is early draft.**

## IncentiveUnit

```protobuf
message IncentiveUnit {
  string id = 1 [
    (gogoproto.moretags) = "yaml:\"id\""
  ];
  repeated SubjectInfo subject_info_list = 2 [
    (gogoproto.moretags) = "yaml:\"subject_info_lists\"",
    (gogoproto.nullable) = false
  ];
}

message SubjectInfo {
  string address = 1 [
    (gogoproto.moretags) = "yaml:\"subject_addr\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string weight = 2 [
    (gogoproto.moretags) = "yaml:\"weight\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
}
```

- Incentive: `"incentive_id" -> format(IncentiveStore)`

### incentive_id

`incentive_id` is the unique identifier in the `incentive_store` for the subjects. Hence, it can't be duplicated.

## SubjectInfo

### weight

The ratio of the reward distribution in a `incentive_store` unit.   
`incentive_store` can contain several `subject`s and ratio for each.   


## NftmarketFrontendIncentiveIdTable

- nftmarket_frontend_incentive_id_table: `format(nft_id) -> format(incentive_id)`

This KVStore manages what NFT is connected to which `incentive_id`.

## RewardTable

RewardTable is the record of the rewards for the subject of the `ecosystem-incentive`.

```protobuf
message Reward {
  string subject_addr = 1 [
    (gogoproto.moretags) = "yaml:\"subject_addr\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated cosmos.base.v1beta1.Coin rewards = 2 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
    (gogoproto.moretags) = "yaml:\"rewards\"",
    (gogoproto.nullable) = false
  ];
}
```

- RewardTable: `format(address) -> format(reward)`

## Params

```protobuf
message Params {
  repeated RewardParams reward_params = 1 [ (gogoproto.moretags) = "yaml:\"reward_params\"" ];
  uint64 max_incentive_unit_id_len = 2 [ (gogoproto.moretags) = "yaml:\"max_incentive_unit_id_len\"" ];
  uint64 max_subject_info_num_in_unit = 3 [ (gogoproto.moretags) = "yaml:\"max_subject_info_num_in_unit\"" ];
}

message RewardParams {
  string module_name = 1 [(gogoproto.nullable) = false];
  repeated IncentiveUnit incentive_units = 2 [(gogoproto.nullable) = false];
  repeated RewardStore reward_stores = 3 [(gogoproto.nullable) = false];
  repeated IncentiveUnitIdsByAddr incentive_unit_ids_by_addr = 4 [(gogoproto.nullable) = false];
}

// RewardRate defines the ratio to take reward for a specific reward_type.
// The total sum of reward_rate in a module cannot be exceed 1
message RewardRate {
  RewardType reward_type = 1 [ (gogoproto.moretags) = "yaml:\"reward_type\"" ];
  string rate = 2 [
    (gogoproto.moretags) = "yaml:\"rate\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.nullable) = false
  ];
}

// At first, we go with this one type.
// NFTMARKET_FRONTEND type reward will be disributed for the creators of frontend of UnUniFi's services.
enum RewardType {
  NFTMARKET_FRONTEND = 0;
}
```

`Params` contains `RewardParams` as the configuration of this module parameters and `MaxIncentiveUnitIdLen` as to define the max length of the IncentiveUnitId.

### RewardRate

The factor to multipy the trading fee for the reward of this module.   
e.g. If `reward_rate` is 80% and the trading fee that is made in a target message is 100GUU, the actual reward for target `incentive_id` subjects is `100GUU * 0.80 = 80GUU`.  

### RewardType

The reward type manages the types of the reward for the various subject.
At first, we support frontend creator. But, the reward will be able to distributed for the different type of parties in our ecosystem.

### MaxIncentiveUnitIdLen

The length of `IncentiveUnitId` must be between `MaxIncentiveUnitIdLen` and 0.

## IncentiveUnitIdsByAddr

IncentiveUnitIdsByAddr is the collection of the incentive unit ids for each address.

```protobuf
message IncentiveUnitIdsByAddr {
  string address = 1 [
    (gogoproto.moretags) = "yaml:\"address\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated string incentive_unit_ids = 2 [
    (gogoproto.moretags) = "yaml:\"incentive_unit_ids\"",
    (gogoproto.nullable) = false
  ];
}
```
