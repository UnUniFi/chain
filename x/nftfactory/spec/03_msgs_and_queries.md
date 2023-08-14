# Messages and Queries

## Messages

The `nftfactory` module provides below messages.

### CreateClass

```protobuf
message MsgCreateClass {
  string                sender      = 1;
  string                subclass    = 2;
  string                name        = 3;
  string                symbol      = 4;
  string                description = 5;
  string                uri         = 6;
  string                uri_hash    = 7;
}
```

### MintNFT

The MintNFT message is used to mint NFT on UnUniFi using sdk's x/nft module function. 
The specifing `NFT.Id` becomes a part of the `NFT.Uri`.

```protobuf
message MsgMintNFT {
  string sender    = 1;
  string class_id  = 2;
  string token_id  = 3;
  string uri       = 4;
  string uri_hash  = 5;
  string recipient = 6;
}
```

### BurnNFT

The BurnNFT message is used to burn the NFT defined by `Class.Id` and `NFT.Id`.   
Only the owner of the `NFT` can send this message.

```protobuf
message MsgBurnNFT {
  string sender   = 1;
  string class_id = 2;
  string token_id   = 3;
}
```
