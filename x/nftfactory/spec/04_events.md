# Events

The `nftmint` module emits the following events.

## EventMsgCreateClass

```protobuf
message EventCreateClass {
  string owner = 1;
  string class_id = 2;
  string base_token_uri = 3;
  string token_supply_cap = 4;
  MintingPermission minting_permission = 5;
}
```

## EventMsgSendClass

```protobuf
message EventSendClass {
  string sender = 1;
  string receiver = 2;
  string class_id = 3;
}
```

## EventMsgUpdateBaseTokenUri

```protobuf
message EventUpdateBaseTokenUri {
  string owner = 1;
  string class_id = 2;
  string base_token_uri = 3;
}
```

## EventMsgUpdateTokenSupplyCap

```protobuf
message EventUpdateTokenSupplyCap {
  string owner = 1;
  string class_id = 2;
  string token_supply_cap = 3;
}
```

## EventMsgMintNFT

```protobuf
message EventMintNFT {
  string class_id = 1;
  string nft_id = 2;
  string owner = 3;
  string minter = 4;
}
```

## EventMsgBurnNFT

```protobuf
message EventBurnNFT {
  string burner = 1;
  string class_id = 2;
  string nft_id = 3;
}
```
