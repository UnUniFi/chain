syntax = "proto3";
package ununifi.ecosystemincentive;

import "ecosystem-incentive/ecosystem_incentive.proto";
import "gogoproto/gogo.proto";
import "cosmos/base/v1beta1/coin.proto";

option go_package = "github.com/UnUniFi/chain/x/ecosystem-incentive/types";

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
