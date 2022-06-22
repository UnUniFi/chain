# State

**NOTE: This is very early draft.**

## Class and its Relating Attributes

```protobuf

message ClassAttributes {
  string class_id = 1;
  string owner = 2;
  string base_token_uri = 3;
  bool minting_permission = 4;
  (undefined) token_supply_cap = 5;
  // possibly
  // bool transferable = 6;
}
```

## NFT and its Relating Attributes

```protobuf
message NFTAttributes {
  string class_id = 1;
  string nft_id = 2;
  string minter = 3;
}
```
