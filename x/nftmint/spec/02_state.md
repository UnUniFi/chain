# State

## Class and its Relating Attributes

### ClassAttributes

We use `ClassAttributes` data object to represent the information which sdk's x/nft module's `Class` doesn't have like owner of the `Class`.     
We require to choose the parameter values when to send `MsgCreateClass` message. To change the parameters in `ClassAttributes` can be made by sending messages to achieve it like `MsgUpdateBaseTokenUri`.   
The close explanation of each parameter lies in 01_concept page.


```protobuf
message ClassAttributes {
  string class_id = 1;
  string owner = 2;
  string base_token_uri = 3;
  MintingPermission minting_permission = 4;
  uint64 token_supply_cap = 5;
}

enum MintingPermission {
  OnlyOwner = 0;
  Anyone = 1;
  WhiteList = 2;
}
```

### OwningClassIdList

OwningClassIdList data is to record the class ids which are owned by specific address.   
This is specifically used to query `QueryClassIdsByOwner`.

```protobuf
message OwningClassIdList {
  string owner = 1 [
    (gogoproto.moretags) = "yaml:\"owner\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated string class_id = 2;
}
```

### ClassNameIdList

ClassNameIdList data is to record the class ids which has specific name.   
This is specifically used to query `QueryClassIdsByName`.

```protobuf
message ClassNameIdList {
  string class_name = 1;
  repeated string class_id = 2;
}
```

- ClassAttributes with prefix "0x01": `format(class_id) -> ClassAttributes`
- OwningClassIdList with prefix "0x03": `format(owner) -> OwningClassIdList`
- ClassNameIdList with prefix "0x04": `format(name) -> ClassNameIdList`

## NFT and its Relating Attributes

There aren't types defined in proto for the relating to nft data.
But, in UnUniFi, the minter of each NFT is recorded.

- Minter with prefix "0x02": `format(class_id + nft_id) -> AccAddress.Byte()`
