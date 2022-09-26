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
  string incentive_id = 2;
  repeated string subject_addrs = 3 [
    (gogoproto.moretags) = "yaml:\"subject\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated string weights = 4 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec",
    (gogoproto.moretags) = "yaml:\"weight\"",
    (gogoproto.nullable) = false
  ];
}
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
```json
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
  string denom = 2 
    [ (gogoproto.moretags) = "yaml:\"denom\"",
    (gogoproto.nullable) = false
  ];
}
```

## Queries

All queries of `ecosystem-incentive`.

### IncentiveStore

```protobuf
message QueryIncentiveUnitRequest {
  string incentive_id = 1;
}
```

### AllRewards

```protobuf
message QueryAllRewardsRequest {
  string subject = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
}
```

### SpecificDenomReward

```protobuf
message QuerySpecificDenomRewardRequest {
  string subject = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string denom = 2;
}
