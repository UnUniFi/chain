# Term list

**NOTES: This is just for development. Once the other legitimate pages are finalized, remove this.**

`Class`    

NFT **Class** is comparable to an ERC-721 smart contract (provides description of a smart contract), under which a collection of NFTs can be created and managed.

```protobuf
message Class {
  string id          = 1;
  string name        = 2;
  string symbol      = 3;
  string description = 4;
  string uri         = 5;
  string uri_hash    = 6;
  google.protobuf.Any data = 7;
}
```

* `id` is an alphanumeric identifier of the NFT class; it is used as the primary index for storing the class; _required_
* `name` is a descriptive name of the NFT class; _optional_
* `symbol` is the symbol usually shown on exchanges for the NFT class; _optional_
* `description` is a detailed description of the NFT class; _optional_
* `uri` is a URI for the class metadata stored off chain. It should be a JSON file that contains metadata about the NFT class and NFT data schema ([OpenSea example](https://docs.opensea.io/docs/contract-level-metadata)); _optional_
* `uri_hash` is a hash of the document pointed by uri; _optional_
* `data` is app specific metadata of the class; _optional_

`NFT`    

```protobuf
message NFT {
  string class_id           = 1;
  string id                 = 2;
  string uri                = 3;
  string uri_hash           = 4;
  google.protobuf.Any data  = 10;
}
```

* `class_id` is the identifier of the NFT class where the NFT belongs; _required_,`[a-zA-Z][a-zA-Z0-9/:-]{2,100}`
* `id` is an alphanumeric identifier of the NFT, unique within the scope of its class. It is specified by the creator of the NFT and may be expanded to use DID in the future. `class_id` combined with `id` uniquely identifies an NFT and is used as the primary index for storing the NFT; _required_,`[a-zA-Z][a-zA-Z0-9/:-]{2,100}`

  ```text
  {class_id}/{id} --> NFT (bytes)
  ```

* `uri` is a URI for the NFT metadata stored off chain. Should point to a JSON file that contains metadata about this NFT (Ref: [ERC721 standard and OpenSea extension](https://docs.opensea.io/docs/metadata-standards)); _required_
* `uri_hash` is a hash of the document pointed by uri; _optional_
* `data` is an app specific data of the NFT. CAN be used by composing modules to specify additional properties of the NFT; _optional_
