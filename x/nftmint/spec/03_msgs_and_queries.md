# Messages and Queries

## Messages

The `nftmint` module provides below messages.

### CreateClass

CreateClass message is used to create `Class` for minting NFTs using cosmos sdk's x/nft module functions.  
In cosmos sdk, the `Class` object is to the contract in EVM. So to mint NFT requires corresponding the `Class`.   

Additionaly on UnUniFi, we create `ClassAttributes` data with some parameters.   
The `MintingPermission` defines who can mint NFTs under this `Class`. The current providing options are `OnlyOwner` (0) and `Anyone` (1).   
If the `Class` have `OnlyOwner` permission, the NFTs can be minted by literaly only owner of the `Class`.   
If the `Class` have `Anyone` permission, the NFTs can be minted by anyone.   

The `Symbol`, `Description` and `ClassUri` are the flag options which can be set blank.

```protobuf
message MsgCreateClass {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string name = 2;
  string base_token_uri = 3;
  uint64 token_supply_cap = 4;
  MintingPermission minting_permission = 5;
  string symbol = 7;
  string description = 8;
  string class_uri = 9;
}

enum MintingPermission {
  OnlyOwner = 0;
  Anyone = 1;
}
```

### SendClass

The SendClass message is used to change the owner of the `Class` on UnUniFi.   
Technically speaking, this message, if accepted, changes the parameter value of the `ClassAttributes.Owner` to the recipient.

```protobuf
message MsgSendClass {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string class_id = 2;
  string recipient = 3 [
    (gogoproto.moretags) = "yaml:\"recipient\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
}
```

### UpdateBaseTokenUri

The UpdateBaseTokenUri message is used to change the `BaseTokenUri` of the `Class` by defining `Class.Id`.
When this message is sended successfully, the all belonging NFT's `NFT.Uri` are changed according to the updating `BaseTokenUri`.

```protobuf
message MsgUpdateBaseTokenUri {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string class_id = 2; 
  string base_token_uri = 3;
}
```

### UpdateTokenSupplyCap

The UpdateTokenSupplyCap message is used to change the `TokenSupplyCap` of the `Class` by defining `Class.Id`.
This message fails if the tokens supply under the `Class` is over the updating `TokenSupplyCap`.

```protobuf
message MsgUpdateTokenSupplyCap {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string class_id = 2; 
  uint64 token_supply_cap = 3;
}
```

### MintNFT

The MintNFT message is used to mint NFT on UnUniFi using sdk's x/nft module function. 
The specifing `NFT.Id` becomes a part of the `NFT.Uri`.

```protobuf
message MsgMintNFT {
    string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string class_id = 2;
  string nft_id = 3;
  string recipient = 4 [
    (gogoproto.moretags) = "yaml:\"recipient\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
}
```

### BurnNFT

The BurnNFT message is used to burn the NFT defined by `Class.Id` and `NFT.Id`.   
Only the owner of the `NFT` can send this message.

```protobuf
message MsgBurnNFT {
  string sender = 1 [
    (gogoproto.moretags) = "yaml:\"sender\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string class_id = 2;
  string nft_id = 3;
}
```

## Queries

The `nftmint` module supports below queries.

### ClassAttributes

The ClassAttributes query is used to get `ClassAttributes` data specified by `Class.Id`.
The `ClassAttributes` data structure is as below.

```protobuf
message QueryClassAttributesRequest {
  string class_id = 1;
}
message QueryClassAttributesResponse {
  ClassAttributes class_attributes = 1;
}

message ClassAttributes {
  string class_id = 1;
  string owner = 2 [
    (gogoproto.moretags) = "yaml:\"owner\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  string base_token_uri = 3;
  MintingPermission minting_permission = 4;
  uint64 token_supply_cap = 5;
}
```

### NFTMinter

The NFTMinter query is used to get the minter of the `NFT` specified by `Class.Id` and `NFT.Id`.

```protobuf
message QueryNFTMinterRequest {
  string class_id = 1;
  string nft_id = 2;
}
message QueryNFTMinterResponse {
  string minter = 1;
}
```

### ClassIdsByName

The ClassIdsByName query is used to get the `ClassNameIdList` data specified by `Class.Name`.

```protobuf
message QueryClassIdsByNameRequest {
  string class_name = 1;
}
message QueryClassIdsByNameResponse {
  ClassNameIdList class_name_id_list = 1;
}

message ClassNameIdList {
  string class_name = 1;
  repeated string class_id = 2;
}
```

### ClassIdsByOwner

The ClassIdsByOwner query is used to get the `OwningClassIdList` data specified by the owner address.

```protobuf
message QueryClassIdsByOwnerRequest {
  string owner = 1;
}
message QueryClassIdsByOwnerResponse {
  OwningClassIdList owning_class_id_list = 1;
}

message OwningClassIdList {
  string owner = 1 [
    (gogoproto.moretags) = "yaml:\"owner\"",
    (gogoproto.customtype) = "github.com/UnUniFi/chain/types.StringAccAddress",
    (gogoproto.nullable) = false
  ];
  repeated string class_id = 2;
}
```
