# Messages and Queries

**NOTE: This is very early draft.**

## Messages

The `nftmint` module provides below messages.

### CreateClass

```protobuf
message MsgCreateClass {
  string name = 1;
  string base_token_uri = 2;
  string total_supply_cap = 3;
  string sender = 5; // initial owner
  bool minting_permission = 4; // flag option. default: true
  string symbol = 6; // flag option
  string description = 7; // flag option
  string class_uri = 8; // flag option
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
  string class_id = 1;
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

### ClassTransferable

```protobuf
message QueryClassTransferableRequest {
  string class_id = 1;
}

message QueryClassTransferableResponse {
  string transferable = 1;
}
```
