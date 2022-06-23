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
```

- ClassAttributes:`format(class_id) -> ClassAttributes`

## NFT and its Relating Attributes

These params aren't updated once they're created at the minting moment.

```protobuf
message NFTAttributes {
  string class_id = 1;
  string nft_id = 2;
  string minter = 3;
}
```

- NFTAttributes: `format(class_id) -> NFTAttributes`
