# State

## `asset_management_accounts`

In expression of protobuf, it is `repeated AssetManagementAccount`.

```protobuf
message AssetManagementAccount {
  string id = 1;
  string name = 2;
  string module_account_address = 3;
  google.protobuf.Any pub_key = 4;
}
```

## `asset_management_targets`

In expression of protobuf, it is `repeated AssetManagemenTarget`.

```protobuf
message AssetManagementTarget {
  string id = 1;
  string asset_management_account_id = 2;
  string denom = 3;
  google.protobuf.Duration unbonding_time = 4 [(gogoproto.nullable) = false, (gogoproto.stdduration) = true];
}
```

## `daily_percents`

In expression of protobuf, it is `repeated DailyPercent`.

```protobuf
message DailyPercent {
  string asset_management_target_id = 1;
  string rate = 2 [(gogoproto.customtype) = "Dec", (gogoproto.nullable) = false];
  google.protobuf.Timestamp date = 3 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
}
```

## `deposits`

In expression of protobuf, it is `repeated Deposit`.

```protobuf
message Deposit {
  string id = 1;
  string from_address = 2;
  repeated cosmos.base.v1beta1.Coin amount = 3 [(gogoproto.nullable) = false];
  google.protobuf.Duration daily_percent_calculation_period = 4 [(gogoproto.nullable) = false, (gogoproto.stdduration) = true];
  google.protobuf.Duration max_unbonding_time = 4 [(gogoproto.nullable) = true, (gogoproto.stdduration) = true];
  google.protobuf.Timestamp date = 5 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
}
```

If recent 30 days are designated in `daily_percent_calculation_period`, APY will be calculated with recent 30 days DPR.

If recent 60 days are designated in `daily_percent_calculation_period`, APY will be calculated with recent 60 days DPR.

Targets which have the highest APY calculated with DPR, will be used for the target.

## `deposit_allocation`

```protobuf
message DepositAllocation {
  string id = 1;
  string deposit_id = 2;
  cosmos.base.v1beta1.Coin amount = 3 [(gogoproto.nullable) = false];
}
```
