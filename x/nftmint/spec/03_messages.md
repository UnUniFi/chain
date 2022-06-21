# Messages

**NOTE: This is very early draft.**

### CreateClass

```protobuf
message MsgCreateClass {
  string name = 1;
  string base_token_uri = 2;
  string total_supply_cap = 3;
  string sender = 4; // initial owner
  string symbol = 5; // optional
  string description = 6; // optional
  string class_uri = 7; // optional
  string class_uri_hash = 8; // optional
}
```

### TransferClass

```protobuf
message MsgTransferClass {
  string class_id = 1;
  string sender = 2;
  string recipient = 3;
}
```

### UpdateBaseTokenUri

```protobuf
message MsgUpdateBaseTokenUri {
  string class_id = 1; 
  string base_token_uri = 2;
  string sender = 3;
}
```

### MintNFT

```protobuf
message MsgMintNFT {
  string class_id = 1;
  string recipient = 2;
}
```

### BurnNFT

```protobuf
message MsgBurnNFT {
  string class_id = 1;
  string nft_id = 2; 
  string sender = 3;
}
```
