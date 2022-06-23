# Concepts

**NOTE: This is very early draft.**

The `x/nftmint` module implements the feature to mint NFTs on UnUniFi.

## Class

The `Class` data is defined in sdk's x/nft module.   
We use that type definition and store the generated `Class` data on UnUniFi in x/nft module.   
This is similar to metadata of each belonging `NFT`.   
On UnUniFi, to create `Class` requires to send a `MsgCreateClass` message with `Class.Name`, `ClassAttributes.BaseTokenUri` and `ClassAttributes.TokenSupplyCap` and as options, `ClassAttributes.MintingPermission`, `Class.Symbol`, `Class.Description` and `Class.Uri`.
Some of these data can be updated after the creation by the owner of the `Class`.

### Class Id

`Class.Id` must be unique because this is the identifier of its belonging `NFT`s. It's like contract address of ERC721 contract on evm chains.   
The way to be generated `Class.Id` in this protocol follows: 
(undefined)

#### TokenSupplyCap

`ClassAttributes.TokenSupplyCap` is max token supply number of each `Class`'s `NFT`. This value is set when to be created `Class`.   
There's the limitation by the global oprion.   
`ClassAttributes.TokenSupplyCap` can be updated by the owner of that `Class`.

#### BaseTokenUri

The `ClassAttributes.BaseTokenUri` is the base token uri of each `NFT.Uri` like ERC721's.   
The `NFT.Uri` consists of `ClassAttributes.BaseTokenUri` and `NFT.Id`.
`ClassAttributes.BaseTokenUri` can be updated by the owner of that `Class`.

#### MintingPermission

The `ClassAttributes.MintingPermission` represents the premission level to mint `NFT` under the `Class`.   
There're three status, which are `OnlyOwner`, `Anyone` and `WhiteList`.

#### Owner

The `ClassAttributes.Owner` represents the owner of the `Class`.   
The initial owner is the sender of `MsgCreateClass`. This parameter can be changed by sending `MsgSendClass`.

## NFT

The `NFT` data is defined in sdk's x/nft module.   
We use that type definition and store the generated `NFT` data on UnUnifi in x/nft module.   
This represents the `NFT` content.

## Mint

Minting `NFT` is achieved by sending `MsgMintNFT` message with `Class.Id`. 
But, there's minting permission in some case.   
If the attached `Class` has `False` attribute regarding `ClassAttributes.MintingPermission`, anyone can mint `NFT`.
If the attached `Class` has `True` attribute regarding `ClassAttributes.MintingPermission`, only the `ClassAttributes.Owner` can mint `NFT`.

## Update

### Class related Attributes

At the moment, we plan to support the update of the `ClassAttributes` and `Class` data by managing messages for each data.   
e.g. We set `MsgUpdateBaseTokenUri` fot the `ClassAttributes.Owner` to update the `ClassAttributes.BaseTokenUri`.   
In this case, every belonging `NFT.Uri`s are changed at once.

## Burn



## Transrer

### Class Ownership

The owner of the `Class` is recorded with `Class.Id`.
And we support the transition of the `ClassAttributes.Owner` by the sending `ClassAttributes.Owner` to any recipient.

#### Normal NFT

The normal NFT transfer is achieved by cosmos SDK's x/nft message.   
It performs changing owner which is connected specific `Class.Id` and `NFT.Id` in the KVStore in that module.   
So, it only requires `Class` and `NFT` objects are in that module and the sender (current owner) signature to transfer it.
