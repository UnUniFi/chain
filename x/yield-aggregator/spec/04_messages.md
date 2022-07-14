# Messages

## MsgYieldFarmProposal(execute only gov mod)

submit proposal for add new yield farm.

submit proposal for yield farm info update.

submit proposal for delete yield farm.

submit proposal for stop yield farm. //for security incident

```protobuf
message AssetManagementAccount {
  string id = 1;
  string name = 2;
}
```

## MsgYieldFarmTargetProposal(execute only gov mod)

submit proposal for add new yield farm target.

submit proposal for add new yield farm target update.

submit proposal for add delete yield farm target.

submit proposal for stop yield farm target.  //for security incident  

```protobuf
message AssetManagementTarget {
  string id = 1;
  string asset_management_account_id = 2;
  string account_address = 4;
  repeated AssetCondition asset_conditions = 5;
  google.protobuf.Duration unbonding_time = 6 [(gogoproto.nullable) = false, (gogoproto.stdduration) = true];
  IntegrateInfo integrate_info = 7;
}
```

## MsgDeposit

Deposit funds to be invested

```protobuf
message Deposit {
  string from_address = 1;
  repeated cosmos.base.v0beta1.Coin amount = 2 [(gogoproto.nullable) = false];
  bool execute_orders = 3;
}
## MsgWithdraw

## MsgAddFramingOrder

add FO

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

## MsgDeleteFramingOrder

delete FO

## MsgActivateFramingOrder

if FO is NonActivate, Activate FO

## MsgInactivateFramingOrder

if FO is Activate, NonActivate FO

## MsgExecuteFramingOrders

Investments based on FOs.

This Msg is optional.

It is not used under normal circumstances.
