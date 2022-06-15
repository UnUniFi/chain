# SDk nft module specification

Basically saying, the standard type definition and methods follows ERC721.   
There are major two types defined in sdk's nft module.   
Those are `Class` and `NFT`.   
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

## NFT

The NFT struct represents NFT object itself.   
The NFT type's fields are (in x/nft/nft.pb.go):
```go
type NFT struct {
	// class_id associated with the NFT, similar to the contract address of ERC721
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// id is a unique identifier of the NFT
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

### Burn

### Update

### Transfer

### GetNFT

### GetNFTsOfClass

### GetOwner

### GetBalance

### GetTotalSupply
