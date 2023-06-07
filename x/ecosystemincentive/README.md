# ecosystem Incentive

## Abstract

The `ecosystem-incentive` module provides the feature to incentivize the parties who bring value to our NFT market place users, especially frontend creator for the UnUniFi's NFTFi features by distributing certain rate of the NFT traded fee to the subjects.   
The subjects put the required information in somewhare (current idea is memo field of the target message like MsgPayAuctionFee) and withdraw the accumulated rewards all at once or for one specific denom.   

## Contents

[Concepts](https://github.com/UnUniFi/chain/blob/design/spec/x/ecosystem-incentive/spec/01_concepts.md)   
[State](https://github.com/UnUniFi/chain/blob/design/spec/x/ecosystem-incentive/spec/02_state.md)
[Messages and Queries](https://github.com/UnUniFi/chain/blob/design/spec/x/ecosystem-incentive/spec/03_messages.md)   
[Hooks](https://github.com/UnUniFi/chain/blob/design/spec/x/ecosystem-incentive/spec/04_hooks.md)   
[Memo Structure](https://github.com/UnUniFi/chain/blob/design/spec/x/ecosystem-incentive/spec/05_memo_structure.md)   
[Events](https://github.com/UnUniFi/chain/blob/design/spec/x/ecosystem-incentive/spec/06_events.md)   

### For developers in the core team

[ADR of this module](https://github.com/UnUniFi/chain/blob/design/spec/doc/architecture/adr-ecosystem-incentive.md)   
There's info about the requirement to achieve the purpose of this module.   

# Concepts

**NOTE: This is early draft.**

This module aims to provide the incentive for the parties which especially bring value to our ecosystem like frontend service creator.   
Fucosing on the case for the frontend service creator, any of them who creates UnUniFi NFT market and NFTFi frontend service are the subjects to recieve Ecosystem Incentive reward from the NFT trading fee in many denoms which are used in NFT market.

## Joining Ecosystem Incentive

Any subjects can send a register message `MsgIncentiveRegister` with the `incentive_id` and `subject_weight_map`.   

## Getting Ecosystem Incentive Reward

This model of distribution reward could be applied to many use-cases. But, we write down only about the case for Nftmarket Frontend model here for better explanation of the sense of this module.   
First, the subjects must register to get incentive by sending `MsgIncentiveRegister`.   
Once the `incentive_id` is registered, they insert that `incentive_id` in the target message which is `MsgListNft` memo field precisely to get the reward for the Nftmarket Frontend incentive mode.
Once the `NftIdentifer` on the market is connected with `incentive_id`, `AfterNftPaymentWithCommission` hook function triggers methods to reflect the reward amount for according addresses in `incentive_id`.

## Withdrawing Ecosystem Incentive Reward

Any registered subjects can withdraw thier reward by sending a withdrawal message if they are there.   
They can withdraw all rewards across all denoms by sending `MsgWithdrawAllRewards`.   
In other way, they can withdraw specific denom reward by sending `MsgWithdrawSpecificDenomReward`.

## The Reward Mechanism

All the reward comes from the fees that UnUniFi protocol earned in addition to gas fee which is defined in protocol as glocal parameter.   
There is nothing inflational effect or depletion by rewarding subjects.

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

// NFTMARKET_FRONTEND type reward will be disributed for the creators of frontend of UnUniFi's services.
enum RewardType {
  UNKNOWN = 0;
  STAKERS = 1;
  FRONTEND_DEVELOPERS = 2;
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

# Messages and Queries

**NOTE: This is early draft.**

## Messages

All messages of `ecosystem-incentive`.

### Register

A message to register `incentive_unit` to take reward from `ecosystem-incentive`.

```protobuf
message MsgRegister {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string incentive_unit_id = 2 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  repeated string subject_addrs = 3 [
    (gogoproto.moretags) = "yaml:\"subject_addrs\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated string weights = 4 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.moretags) = "yaml:\"weights\"",
    (gogoproto.nullable) = false
  ];
}
message MsgRegisterResponse {}
```

`weights` must be `1.000000000000000000` (just ok as long as  it represent 1).   
For example,   
ok: [0.5, 0.5], [0.33, 0.33, 0.34]   
not: [0.5, 0.1], [0.33, 0.33, 0.3333]   

And more importantly, don't forget how one `subject_addr` is associated with one `weight`. It's just order for those two lists. For example, in this case   
```shell
subject_addrs = [
"ununifi17gs6kgph4657epky2ctl9sf66ucyua939nexgl",
"ununifi1w9s3wpkh0kfk0t40m4lwjsx6h2v6gktsvfrgux"
]
weights = [
"0.6",
"0.4
]
```

`ununifi17gs6kgph4657epky2ctl9sf66ucyua939nexgl`'s `weight` will be `0.6` and `ununifi1w9s3wpkh0kfk0t40m4lwjsx6h2v6gktsvfrgux`'s will be `0.4`.

#### CLI

We receive a JSON file in CLI command for this message.
Example JSON file for CLI tx command:

```Json
{
	"incentive_id": "incentive-unit1",
	"subject_addrs": [
		"ununifi17gs6kgph4657epky2ctl9sf66ucyua939nexgl",
		"ununifi1w9s3wpkh0kfk0t40m4lwjsx6h2v6gktsvfrgux"
	],
	"weights": [
		"0.50",
		"0.50"
	]
}
```

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
  string denom = 2 [ (gogoproto.moretags) = "yaml:\"denom\"" ];
}
```

## Queries

All queries of `ecosystem-incentive`.

### IncentiveUnit

```protobuf
message QueryIncentiveUnitRequest {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
}

message QueryIncentiveUnitResponse {
  IncentiveUnit incentive_unit = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit\"" ];
}
```

### AllRewards

```protobuf
message QueryAllRewardsRequest {
  string subject_addr = 1 [ (gogoproto.moretags) = "yaml:\"subject_addr\"" ];
}

message QueryAllRewardsResponse {
  Reward rewards = 1 [
    (gogoproto.moretags) = "yaml:\"rewards\"",
    (gogoproto.nullable) = false
  ];
}
```

### SpecificDenomReward

```protobuf
message QueryRewardRequest {
  string subject_addr = 1 [ (gogoproto.moretags) = "yaml:\"subject_addr\"" ];
  string denom = 2 [ (gogoproto.moretags) = "yaml:\"denom\"" ];
}

message QueryRewardResponse {
  cosmos.base.v1beta1.Coin reward = 1 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.moretags) = "yaml:\"reward\"",
    (gogoproto.nullable) = false
  ];
}
```

### IncentiveUnitIdsByAddr

```protobuf
message QueryIncentiveUnitIdsByAddrRequest {
  string address = 1 [
    (gogoproto.moretags) = "yaml:\"address\""
  ];
}

message QueryIncentiveUnitIdsByAddrResponse {
  IncentiveUnitIdsByAddr incentive_unit_ids_by_addr = 1 [
    (gogoproto.moretags) = "yaml:\"incentive_unit_ids_by_addr\"",
    (gogoproto.nullable) = false
  ];
}
```

# Hooks

**NOTE: This is early draft.**

All rewards accumulation are executed when the according hooks function is called.   

The example hooks functions interfaces in x/nftmarket module:

```go
type NftmarketHooks interface {
	AfterNftListed(ctx sdk.Context, nftIdentifier NftIdentifier, txMemo string)
	AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier NftIdentifier, fee sdk.Coin)
	AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier NftIdentifier)
}
```

## AfterNftListed

This hook function is called for the resistration for the `ecosystem-incentive` with the `txMemo` and `nftIdentifiler`.   
To pass the `txMemo` from the memo data of `MsgListNft` requires a method to get memo data in the process of `MsgListNft` in `x/nftmarket` module.

### Location to be inserted

- `ListNft(ctx sdk.Context, msg *types.MsgListNft)` from x/nftmarket in nft_listing.go

## AfterNftPaymentWithCommission

This hook function is called for the accumulation of the reward for the subjects which are connected with the `nftIdentifiler` in the argument.
The calculation of the actual reward amount is executed in methods which this hook function calls in this module.

### Location to be inserted

- `ProcessPaymentWithCommissionFee(ctx sdk.Context, listingOwner sdk.AccAddress, denom string, amount sdk.Int)`  from x/nftmarket in nft_listing.go

## AfterNftUnlistedWituoutPayment

This hook function is called when a nft is unlisted for some reason like liquidation.   
The purpose is to remove the unlisted nft information from `NftmarketFrontendIncentiveIdTable` KVStore to keep the data consystent.

### Location to be inserted

- `CancelNftListing(ctx sdk.Context, msg *types.MsgCancelNftListing)` from x/nftmarket in nft_listing.go
- Case which bid's length for the listing is 0 in `EndNftListing(ctx sdk.Context, msg *types.MsgEndNftListing)` from x/nftmarket in nft_listing.go

# Data structure for the memo field

We use tx memo field data to identify what incentive will be distributed to what `incentive-unit` by putting the correct formatted json data into that.

The v1's formal data archtecture is:

```json
{
  "version": "v1",
  "incentive_unit_id": "incentive_unit-1"
}
```

NOTE: There's a lot of chances to be changed this structure with the change of the version. Please note it when to use.

## Frontends

We use memo field data to know which frontend a lisetd nft used in the case of frontend-incentive model.   
So we have to use the organized data structure of memo field in a listing tx (MsgListNft) to distingush it as a legitimate entry or not.

Even if you put the wrong formatted data in the memo of tx contains MsgListNft, the MsgListNft itself will still succeed. The registration of the information which nft-id relates to what `incentive-unit-id` will just fail.

```protobuf
message EventRegister {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  repeated SubjectInfo subject_info_lists = 2 [
    (gogoproto.moretags) = "yaml:\"subject_info_lists\"",
    (gogoproto.nullable) = false
  ];
}

message EventWithdrawAllRewards {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated cosmos.base.v1beta1.Coin all_withdrawn_rewards = 2 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
    (gogoproto.moretags) = "yaml:\"all_withdrawn_rewards\"",
    (gogoproto.nullable) = false
  ];
}

message EventWithdrawReward {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  cosmos.base.v1beta1.Coin withdrawn_reward = 2 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.moretags) = "yaml:\"withdrawn_reward\"",
    (gogoproto.nullable) = false
  ];
}

message EventFailedParsingMemoInputs {
  string class_id = 1 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 2 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
  string memo = 3 [ (gogoproto.moretags) = "yaml:\"memo\"" ];
}

message EventRecordedIncentiveUnitId {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  string class_id = 2 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 3 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}

message EventDeletedNftIdRecordedForFrontendReward {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  string class_id = 2 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 3 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}

message EventNotRegisteredIncentiveUnitId {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  string class_id = 2 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 3 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}

message EventNotRecordedNftId {
  string class_id = 1 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 2 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}

message EventUpdatedReward {
  string incentive_unit_id = 1 [ (gogoproto.moretags) = "yaml:\"incentive_unit_id\"" ];
  cosmos.base.v1beta1.Coin earned_reward = 2 [
    (gogoproto.moretags) = "yaml:\"earned_reward\"",
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.nullable) = false
  ];
}

message EventVersionUnmatched {
  string unmatched_version = 1 [ (gogoproto.moretags) = "yaml:\"unmatched_version\"" ];
  string class_id = 2 [ (gogoproto.moretags) = "yaml:\"class_id\"" ];
  string nft_id = 3 [ (gogoproto.moretags) = "yaml:\"nft_id\"" ];
}
```
