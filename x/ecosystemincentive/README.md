# ecosystem Incentive

## Abstract

The `ecosystem-incentive` module provides the feature to incentivize the parties who bring value to our NFT market place users, especially frontend creator for the UnUniFi's NFTFi features by distributing certain rate of the NFT traded fee to the subjects.
The subjects put the required information in somewhare (current idea is memo field of the target message like MsgPayAuctionFee) and withdraw the accumulated rewards all at once or for one specific denom.

## Contents

TODO: contents

## Concepts

This module aims to provide the incentive for the parties which especially bring value to our ecosystem like frontend service creator.  
Focusing on the case for the frontend service creator, any of them who creates UnUniFi NFT market and NFTFi frontend service are the subjects to receive Ecosystem Incentive reward from the NFT trading fee which are used in NFT market.

### Getting Ecosystem Incentive Reward

This model of distribution reward could be applied to many use-cases. But, we write down only about the case for nftbackedloan Frontend model here for better explanation of the sense of this module.

First, add the following JSON to the TxMemo field when sending the nftbackedloan's MsgListNft to the chain. The address is the address you want to receive the reward.

```json
{ "frontend": { "version": 1, "recipient": "ununifixxxxxxx" } }
```

In this way, the reward recipient is linked to the NFT ID, and a portion of the commission from this NFT transaction is given to the recipient.
This operation is handled by the `AfterNftPaymentWithCommission` hook function.

### Withdrawing Ecosystem Incentive Reward

Rewards are accumulated in the ecosysytemincentive module and can be received in the following ways.

`MsgWithdrawAllRewards`

```protobuf
message MsgWithdrawAllRewards {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\""
  ];
}
```

`MsgWithdrawReward`

```protobuf
message MsgWithdrawReward {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\""
  ];
  string denom = 2 [ (gogoproto.moretags) = "yaml:\"denom\"" ];
}
```

### The Reward Mechanism

All the reward comes from the fees that UnUniFi protocol earned in addition to gas fee which is defined in protocol as global parameter.
There is nothing inflationary effect or depletion by rewarding subjects.

## State

### nftbackedloanFrontendTable

`format(nft_id) -> format(recipient_address)`

This KVStore manages what NFT is linked to which `recipient_address`.

### RewardTable

RewardTable is the record of the `ecosystem-incentive` rewards.

```protobuf
message RewardRecord {
  string address = 1 [
    (gogoproto.moretags) = "yaml:\"address\""
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
}

message RewardParams {
  string module_name = 1 [ (gogoproto.moretags) = "yaml:\"module_name\"" ];
  repeated RewardRate reward_rate = 2 [
    (gogoproto.moretags) = "yaml:\"reward_rate\"",
    (gogoproto.nullable) = false
  ];
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

// STAKERS type reward will be distributed for the stakers of GUU token.
// FRONTEND_DEVELOPERS type reward will be distributed for the creators of frontend of UnUniFi's services.
// COMMUNITY_POOL type reward will be distributed for the community pool.
enum RewardType {
  UNKNOWN = 0;
  STAKERS = 1;
  FRONTEND_DEVELOPERS = 2;
  COMMUNITY_POOL = 3;
}
```

### RewardRate

The factor to multiple the trading fee for the reward of this module.  
e.g. If `reward_rate` is 80% and the trading fee that is made in a target message is 100GUU, the actual reward is `100GUU * 0.80 = 80GUU`.

### RewardType

The reward type manages the types of the reward for the various subject.
At first, we support frontend creator. But, the reward will be able to distributed for the different type of parties in our ecosystem in the future.

## Ante Handler

When an NFT is listed in nftbackedloan, AnteHandlers is used to read the MsgListNft and link the NFT to the recipient address.

AnteHandlers Reference
<https://docs.cosmos.network/v0.45/modules/auth/03_antehandlers.html>

At this time, a memo in the following format should be included in the MsgListNft.

```json
{ "frontend": { "version": 1, "recipient": "ununifixxxxxxx" } }
```

On CLI, send Tx as follows

```bash
ununifid tx nftbackedloan list \
$class_id $token_id \
--note "{\"frontend\":{\"version\": 1, \"recipient\": \"ununifixxxxxxx\"}}"
```

## Hooks

All rewards accumulation are executed when the according hooks function is called.

The example hooks functions interfaces in x/nftbackedloan module:

```go
type nftbackedloanHooks interface {
 AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier NftIdentifier, fee sdk.Coin)
 AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier NftIdentifier)
}
```

### AfterNftPaymentWithCommission

This hook function is called for the accumulation of the reward for the subjects which are connected with the `nftIdentifiler` in the argument.
The calculation of the actual reward amount is executed in methods which this hook function calls in this module.

#### Location to be inserted AfterNftPaymentWithCommission

- `ProcessPaymentWithCommissionFee(ctx sdk.Context, listingOwner sdk.AccAddress, denom string, amount sdk.Int)` from x/nftbackedloan in nft_listing.go

### AfterNftUnlistedWithoutPayment

This hook function is called when a nft is unlisted for some reason like liquidation.  
The purpose is to remove the unlisted nft information from `nftbackedloanFrontendIncentiveIdTable` KVStore to keep the data consystent.

#### Location to be inserted AfterNftUnlistedWithoutPayment

- `CancelNftListing(ctx sdk.Context, msg *types.MsgCancelNftListing)` from x/nftbackedloan in nft_listing.go
- Case which bid's length for the listing is 0 in `SetLiquidation(ctx sdk.Context, msg *types.MsgEndNftListing)` from x/nftbackedloan in nft_listing.go

## Frontends

We use memo field data to know which frontend a lisetd nft used in the case of frontend-incentive model.  
So we have to use the organized data structure of memo field in a listing tx (MsgListNft) to distinguish it as a legitimate entry or not.

Even if you put the wrong formatted data in the memo of tx contains MsgListNft, the MsgListNft itself will still succeed. The registration of the information which nft-id relates to what `reciever_address` will just fail.

## Messages and Queries

### Messages

All messages of `ecosystem-incentive`.

#### WithdrawAllRewards

A message to withdraw all accumulated rewards across all denoms.

```protobuf
message MsgWithdrawAllRewards {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\""
  ];
}
```

#### WithdrawReward

A message to withdraw accumulated reward of specified denom.

```protobuf
message MsgWithdrawReward {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\""
  ];
  string denom = 2 [ (gogoproto.moretags) = "yaml:\"denom\"" ];
}
```

### Queries

All queries of `ecosystem-incentive`.

#### AllRewards

```protobuf
message QueryAllRewardsRequest {
  string address = 1 [(gogoproto.moretags) = "yaml:\"address\""];
}

message QueryAllRewardsResponse {
  RewardRecord reward_record = 1 [(gogoproto.moretags) = "yaml:\"reward_record\"", (gogoproto.nullable) = false];
}
```

#### SpecificDenomReward

```protobuf
message QueryRewardRequest {
  string address = 1 [(gogoproto.moretags) = "yaml:\"address\""];
  string denom   = 2 [(gogoproto.moretags) = "yaml:\"denom\""];
}

message QueryRewardResponse {
  cosmos.base.v1beta1.Coin reward = 1 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.moretags)     = "yaml:\"reward\"",
    (gogoproto.nullable)     = false
  ];
}
```

### RecipientAddressWithNftId

```protobuf
message QueryRecipientAddressWithNftIdRequest {
  string class_id = 1 [(gogoproto.moretags) = "yaml:\"class_id\""];
  string token_id = 2 [(gogoproto.moretags) = "yaml:\"token_id\""];
}

message QueryRecipientAddressWithNftIdResponse {
  string address = 1 [(gogoproto.moretags) = "yaml:\"address\""];
}
```

## Events

```protobuf
message EventWithdrawAllRewards {
  string   sender                                         = 1 [(gogoproto.moretags) = "yaml:\"sender\""];
  repeated cosmos.base.v1beta1.Coin all_withdrawn_rewards = 2 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins",
    (gogoproto.moretags)     = "yaml:\"all_withdrawn_rewards\"",
    (gogoproto.nullable)     = false
  ];
}

message EventWithdrawReward {
  string                   sender           = 1 [(gogoproto.moretags) = "yaml:\"sender\""];
  cosmos.base.v1beta1.Coin withdrawn_reward = 2 [
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.moretags)     = "yaml:\"withdrawn_reward\"",
    (gogoproto.nullable)     = false
  ];
}

message EventRecordedRecipientWithNftId {
  string recipient = 1 [(gogoproto.moretags) = "yaml:\"recipient\""];
  string class_id  = 2 [(gogoproto.moretags) = "yaml:\"class_id\""];
  string token_id  = 3 [(gogoproto.moretags) = "yaml:\"token_id\""];
}

message EventDeletedNftIdRecordedForFrontendReward {
  string recipient = 1 [(gogoproto.moretags) = "yaml:\"recipient\""];
  string class_id  = 2 [(gogoproto.moretags) = "yaml:\"class_id\""];
  string token_id  = 3 [(gogoproto.moretags) = "yaml:\"token_id\""];
}

message EventNotRecordedNftId {
  string class_id = 1 [(gogoproto.moretags) = "yaml:\"class_id\""];
  string token_id   = 2 [(gogoproto.moretags) = "yaml:\"token_id\""];
}

message EventUpdatedReward {
  string                   recipient     = 1 [(gogoproto.moretags) = "yaml:\"recipient\""];
  cosmos.base.v1beta1.Coin earned_reward = 2 [
    (gogoproto.moretags)     = "yaml:\"earned_reward\"",
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.nullable)     = false
  ];
}

message EventDistributionForStakers {
  cosmos.base.v1beta1.Coin distributed_amount = 1 [
    (gogoproto.moretags)     = "yaml:\"distributed_amount\"",
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.nullable)     = false
  ];
  int64 block_height = 2 [(gogoproto.moretags) = "yaml:\"block_height\""];
}

message EventVersionUnmatched {
  uint32 unmatched_version = 1 [(gogoproto.moretags) = "yaml:\"unmatched_version\""];
  string class_id          = 2 [(gogoproto.moretags) = "yaml:\"class_id\""];
  string token_id          = 3 [(gogoproto.moretags) = "yaml:\"token_id\""];
}
```
