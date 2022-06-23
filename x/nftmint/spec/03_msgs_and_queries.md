# Messages and Queries

**NOTE: This is very early draft.**

## Messages

The `nftmint` module provides below messages.

### CreateClass

```protobuf
message MsgCreateClass {
  string sender = 1; // initial owner
  string name = 2;
  string base_token_uri = 3;
  string total_supply_cap = 4;
  MintingPermission minting_permission = 5; // flag option. default: true
  string symbol = 7; // flag option
  string description = 8; // flag option
  string class_uri = 9; // flag option
}
```

### SendClass

```protobuf
message MsgSendClass {
  string sender = 1;
  string class_id = 2;
  string recipient = 3;
}
```

### UpdateBaseTokenUri

```protobuf
message MsgUpdateBaseTokenUri {
  string sender = 1;
  string class_id = 2; 
  string base_token_uri = 3;
}
```

### UpdateTokenSupplyCap

```protobuf
message MsgUpdateTokenSupplyCap {
  string sender = 1;
  string class_id = 2; 
  string token_supply_cap = 3;
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
  string sender = 1;
  string class_id = 2;
  string nft_id = 3;
}
```

## Queries

The `nftmint` module supports below queries.

### ClassOwner

```protobuf
message QueryClassOwnerRequest {
  string class_id = 1;
}

message QueryClassOwnerResponse {
  string owner = 1;
}
```

### ClassNFTMinter

```protobuf
message QueryNFTMinterRequest {
  string class_id = 1;
  string nft_id = 2;
}

message QueryNFTMinterResponse {
  string minter = 1;
}
```

### ClassClassIdByName

```protobuf
message QueryClassIdByNameRequest {
  string class_name = 1;
}

message QueryClassIdByNameResponse {
  repeated string class_id = 1;
}
```

### ClassBaseTokenUri

```protobuf
message QueryClassBaseTokenUriRequest {
  string class_id = 1;
}

message QueryClassBaseTokenUriResponse {
  string base_token_uri = 1;
}
```

### ClassTokenSupplyCap

```protobuf
message QueryClassTokenSupplyCapRequest {
  string class_id = 1;
}

message QueryClassTokenSupplyCapResponse {
  string token_supply_cap = 1;
}
```

### ClassMintingPermission

```protobuf
message QueryClassMintingPermissionRequest {
  string class_id = 1;
}

message QueryClassMintingPermissionResponse {
  string minting_permission = 1;
}
```
