# Concepts

**NOTE: This is very early draft.**

The `x/nftmint` module implements the feature to mint NFTs on UnUniFi.

### Transrer Mechanism

#### Normal NFT

The normal NFT transfer is achieved by cosmos SDK's x/nft module message.   
It performs changing owner which is connected specific `Class.Id` and `NFT.Id` in the KVStore in that module.   
So, it only requires `Class` and `NFT` objects are in that module and the sender (current owner) signature to transfer it.
