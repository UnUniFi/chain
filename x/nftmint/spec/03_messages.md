# Messages

**NOTE: This is very early draft.**

### CreateClass

```protobuf
message MsgCreateClass {
  string name = 1;
  string base_token_uri = 2;
  string total_supply_cap = 3;
  bool minting_permission = 4; 
  string sender = 5; // initial owner
  string symbol = 6; // flag optional
  string description = 7; // flag optional
  string class_uri = 8; // flag optional
  string class_uri_hash = 9; // flag optional
}
```

### SendClass

```protobuf
message MsgSendClass {
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

### UpdateTokenSupplyCap

```protobuf
message MsgUpdateBaseTokenUri {
  string class_id = 1; 
  string token_supply_cap = 2;
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
