# State

**NOTE: This is early draft.**

## IncentiveUnit

```protobuf
message IncentiveUnit {
  string incentive_id = 1 [
    (gogoproto.moretags) = "yaml:\"incentive_id\""
  ];
  repeated SubjectInfo subject_info_list = 2 [
    (gogoproto.moretags) = "yaml:\"subject_info_lists\"",
    (gogoproto.nullable) = false
  ];
}

message SubjectInfo {
  string address = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string wight = 2 [
    (gogoproto.moretags) = "yaml:\"auction_size\"",
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
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
  repeated RewardType reward_types = 2 [ (gogoproto.moretags) = "yaml:\"reward_types\"" ];
}

message RewardParams {
  string module_name = 1 [
    (gogoproto.moretags) = "yaml:\"module_name\""
  ];
  repeated RewardRate reward_rate = 2 [
    (gogoproto.moretags) = "yaml:\"reward_rate\"",
    (gogoproto.nullable) = false
  ];
}

// RewardRate defines the ratio to take reward for a specific reward_type.
// The total sum of reward_rate in a module cannot be exceed 1
message RewardRate {
  RewardType reward_type = 1;
  string rate = 2[
    (gogoproto.moretags) = "yaml:\"reward_rate\"",
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

`Params` contains `RewardParams` as the configuration of this module parameters.

### RewardRate

The factor to multipy the trading fee for the reward of this module.   
e.g. If `reward_rate` is 80% and the trading fee that is made in a target message is 100GUU, the actual reward for target `incentive_id` subjects is `100GUU * 0.80 = 80GUU`.  

### RewardType

The reward type manages the types of the reward for the various subject.
At first, we support frontend creator. But, the reward will be able to distributed for the different type of parties in our ecosystem.
