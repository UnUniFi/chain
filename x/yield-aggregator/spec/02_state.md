# State

## `asset_management_accounts`

In expression of protobuf, it is `repeated AssetManagementAccount`.

```protobuf
message AssetManagementAccount {
  string id = 1;
  string name = 2;
  string account_address = 3;
}
``

## `asset_management_targets`

In expression of protobuf, it is `repeated AssetManagemenTarget`.

```protobuf
message AssetCondition {
  string denom = 1[(gogoproto.nullable) = false];
  string min = 2;
  uint32 ratio = 3;
}
message AssetManagementTarget {
  string id = 1;
  string asset_management_account_id = 2;
  repeated AssetCondition asset_conditions = 3;
  google.protobuf.Duration unbonding_time = 4 [(gogoproto.nullable) = false, (gogoproto.stdduration) = true];
}
```

## `daily_percents`

In expression of protobuf, it is `repeated DailyPercent`.

TODO: PredictedDailyPercent

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
  string from_address = 1;
  repeated cosmos.base.v0beta1.Coin amount = 2 [(gogoproto.nullable) = false];
  bool execute_orders = 3;
}
```

## `order`

In expression of protobuf, it is `repeated Order`.

See 06_strategy.md for `strategy`

```protobuf
message FarmingOrder {
  string id = 1;
  string from_address = 2;
  google.protobuf.Any strategy = 3;
  google.protobuf.Duration max_unbonding_time = 4 [(gogoproto.nullable) = true, (gogoproto.stdduration) = true];
  unit32 overall_ratio = 5;
  string min = 6;
  string max = 7;
  google.protobuf.Timestamp date = 8 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
  bool active = 9;
}
```

If recent 30 days are designated in `daily_percent_calculation_period`, APY will be calculated with recent 30 days DPR.

If recent 60 days are designated in `daily_percent_calculation_period`, APY will be calculated with recent 60 days DPR.

Targets which have the highest APY calculated with DPR, will be used for the target.

## `deposit_allocation`

```protobuf
message DepositAllocation {
  string id = 1;
  string order_id = 2;
  cosmos.base.v1beta1.Coin amount = 3 [(gogoproto.nullable) = false];
}
```
