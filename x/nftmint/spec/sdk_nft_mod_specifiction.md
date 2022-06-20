# Brief SDk nft module specification

## Abstract

The current cosmos sdk's x/nft module can be simply said that it's the data storage for the NFTs of cosmos SDK.   
That module does the minimum data type definition to meets the requirement for the IBC (Inter-Blockchian Communication) and implementation of the methods to store them in the module.   

Basically saying, the standard type definition and methods follows ERC721.   
There are major two types defined in sdk's nft module.   
Those are `Class` and `NFT`. The NFT is identified by using `Class.Id` and `NFT.Id` combined.   
The important NFT module keeper's methods are  `Mint`, `Burn`, `Update`, `Transfer`, `GetNFT`, `GetNFTsOfClass`, `GetOwner`, `GetBalance`, `GetTotalSupply` etc.   
The details are below.

## Major Defined Types

### Class
Class struct is similar to ethereum ERC721 contract itself.   
It has unique `Class.Id` to be distinguished by the collection.   
The fields are (in x/nft/nft.pb.go):
```go
type Class struct {
	// id defines the unique identifier of the NFT classification, similar to the contract address of ERC721
        // [a-zA-Z][a-zA-Z0-9/:-]{2,100}
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name defines the human-readable name of the NFT classification. Optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// symbol is an abbreviated name for nft classification. Optional
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// description is a brief description of nft classification. Optional
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// uri for the class metadata stored off chain. It can define schema for Class and NFT `Data` attributes. Optional
	Uri string `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	// uri_hash is a hash of the document pointed by uri. Optional
	UriHash string `protobuf:"bytes,6,opt,name=uri_hash,json=uriHash,proto3" json:"uri_hash,omitempty"`
	// data is the app specific metadata of the NFT class. Optional
	Data *types.Any `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}
```

### NFT

The NFT struct represents NFT object itself.   
The NFT type's fields are (in x/nft/nft.pb.go):
```go
type NFT struct {
	// class_id associated with the NFT, similar to the contract address of ERC721
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// id is a unique identifier of the NFT
        // [a-zA-Z][a-zA-Z0-9/:-]{2,100}
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// uri for the NFT metadata stored off chain
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// uri_hash is a hash of the document pointed by uri
	UriHash string `protobuf:"bytes,4,opt,name=uri_hash,json=uriHash,proto3" json:"uri_hash,omitempty"`
	// data is an app specific data of the NFT. Optional
	Data *types.Any `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
}
```

## Important Keeper methods

### Mint

Mint(nft NFT，receiver sdk.AccAddress)   // updates totalSupply

### Burn

Burn(classId string, nftId string)    // updates totalSupply

### Update

Update(nft NFT)

### Transfer

Transfer(classId string, nftId string, receiver sdk.AccAddress)

### GetNFT

GetNFT(classId string, nftId string) NFT

### GetNFTsOfClass

GetNFTsOfClass(classId string) []NFT

### GetOwner

GetOwner(classId string, nftId string) sdk.AccAddress

### GetBalance

GetBalance(classId string, owner sdk.AccAddress) uint64

### GetTotalSupply

GetTotalSupply(classId string) uint64

## Message

There's one message in sdk's nft module.   

### MsgSend

This message does the transfer of the NFT that identified in argument from sender.   

```go
// MsgSend represents a message to send a nft from one account to another account.
type MsgSend struct {
	// class_id defines the unique identifier of the nft classification, similar to the contract address of ERC721
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// id defines the unique identification of nft
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// sender is the address of the owner of nft
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// receiver is the receiver address of nft
	Receiver string `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
}
```

## Notes

#### Pros

- community contributions, commit history and decision records
- module functionality does not touch other core modules, allowing for an easy upgrade path for the Cosmos Hub to adopt

#### Cons

- efforts to upgrade to protobuf
- grouping NFTs by denom/classification in a wallet/owners can be a privacy concern
- limited on-chain metadata capabilities
- id string generation logic is not deterministic/reproducible, can cause collisions

### Positive

- NFT identifiers available on Cosmos Hub.
- Ability to build different NFT modules for the Cosmos Hub, e.g., ERC-721.
- NFT module which supports interoperability with IBC and other cross-chain infrastructures like Gravity Bridge

### Negative

- New IBC app is required for x/nft
- CW721 adapter is required

### Neutral

- Other functions need more modules. For example, a custody module is needed for NFT trading function, a -
- collectible module is needed for defining NFT properties.

### What this module doesn't provide

- The metadata structure standard
- The strict rules of the variable which users enter like `Class.Name`, even `Class.Id`
- The extra usage of NFT like non-transferable NFT or dynamic NFT
