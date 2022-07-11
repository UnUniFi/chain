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
  MintingPermission minting_permission = 5;
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

### ClassAttributes

```protobuf
message QueryClassAttributesRequest {
  string class_id = 1;
}
message QueryClassAttributesResponse {
  ClassAttributes class_attributes = 1;
}
```

### NFTMinter

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
message QueryClassIdsByNameRequest {
  string class_name = 1;
}
message QueryClassIdsByNameResponse {
  ClassNameIdList class_name_id_list = 1;
}
```

### ClassIdsByOwner

```protobuf
message QueryClassIdsByOwnerRequest {
  string owner = 1;
}
message QueryClassIdsByOwnerResponse {
  OwningClassIdList owning_class_id_list = 1;
}
```
