# Events

The `nftmint` module emits the following events.

## EventCreateClass

```protobuf
message EventCreateClass {
  string owner = 1;
  string class_id = 2;
}
```

## EventMintNFT

```protobuf
message EventMintNFT {
  string sender    = 1;
  string class_id  = 2;
  string token_id  = 3;
  string recipient = 4;
}
```

## EventBurnNFT

```protobuf
message EventBurnNFT {
  string sender   = 1;
  string class_id = 2;
  string token_id = 3;
}
```
