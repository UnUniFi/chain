# State

**NOTE: This is early draft.**

## IncentiveStore

### incentive_id

`incentive_id` is the unique identifier in the `incentive_store` for the subjects. Hence, it can't be duplicated.

### weight

The ratio of the reward distribution in a `incentive_store` unit.   
`incentive_store` can contain several `subject`s and ratio for each.   

```protobuf
message IncentiveStore {
  string incentive_id = 1;
  repeated string subjects = 2;
  repeated undetermined weights = 3;
  RewardType reward_type = 4;
}
```

- Incentive: `"incentive_id" -> format(IncentiveStore)`

## IncentiveIdTable

- incentive_id_table: `format(nft_id) -> format(incentive_id)`

This KVStore manages what NFT is connected to which `incentive_id`.

## RewardTable

RewardTable is the record of the rewards for the subject of the `ecosystem-incentive`.

```protobuf
message Reward {
  string subject = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated cosmos.base.v1beta1.Coin rewards = 2;
}
```

- RewardTable: `format(address) -> format(reward)`

## Params

```protobuf
message Params {
  repeated RewardParam reward_params = 1 [
    (gogoproto.moretags) = "yaml:\"reward_params\"",
    (gogoproto.nullable) = false
  ];
  repeated RewardType reward_types = 2;
}

message RewardParams {
  string module_name = 1;
  repeated RewardRate reward_rate = 2;
}

message RewardRate {
  RewardType reward_type = 1;
  unsure rate = 2;
}

enum RewardType {
  NFTMARKET_FRONTEND = 0; // example
}
```

`Params` contains `RewardParams` as the configuration of this module parameters.

### RewardRate

The factor to multipy the trading fee for the reward of this module.   
e.g. If `reward_rate` is 80% and the trading fee that is made in a target message is 100GUU, the actual reward for target `incentive_id` subjects is `100GUU * 0.80 = 80GUU`.  

### RewardType

The reward type manages the types of the reward for the various subject.
At first, we support frontend creator. But, the reward will be able to distributed for the different type of parties in our ecosystem.

## Data structure for the memo field

We use memo field data to know which frontend a lisetd nft used in the case of frontend-incentive model. So we have to use the organized data structure of memo field in a listing tx to distingush it as a legitimate entry or not.

The v1 archtecture is:

```json
{
  "version": "v1",
  "incentive-type": 0, // for NFTMARKET_FRONTEND in this example
  "incentive-id": "incentive_id"
}
```

There's a lot of chances to be changed this structure with the change of the version. Please note it when to use.
