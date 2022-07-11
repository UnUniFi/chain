# State

**NOTE: This is very early draft.**

## Class and its Relating Attributes

The parameters in `ClassAttributes` can be updated by implementing messages to achieve it.   
The explanation of each params lies in [here](https://github.com/UnUniFi/chain/blob/design/spec/x/nftmint/spec/02_state.md).

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

message OwningClassIdList {
  string owner = 1 [
    (gogoproto.moretags) = "yaml:\"owner\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated string class_id = 2;
}

message ClassNameIdList {
  string class_name = 1;
  repeated string class_id = 2;
}
```

- ClassAttributes: `format(class_id) -> ClassAttributes`
- OwningClassIdList: `format(owner) -> OwningClassIdList`
- ClassNameIdList: `format(name) -> ClassNameIdList`

## NFT and its Relating Attributes

There aren't types defined in proto for the relating to nft data.
But, in UnUniFi, the minter of each NFT is recorded.
