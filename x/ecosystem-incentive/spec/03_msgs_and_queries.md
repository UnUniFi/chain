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
	"incentive-id": "incentive-unit1",
	"subject-addrs": [
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

