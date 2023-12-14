# Concepts

The `x/nftmint` module implements the feature to mint NFTs on UnUniFi.  
At first, we describe brief summary about what this module enables and doesn't.

### Possible

- Anyone can create Collective NFT by creating `Class` and its belonging `NFT`s while the owner of `Class` controls the minting permission. (Collective NFT means the NFT project like CryptoPunks.)
- The NFT standard generally follows ERC721 by default.
- The content of `NFT.Uri` (metadata strucure) follows the other major market place standards.
- The owner of `NFT` can Burn its `NFT`.
- The owner of `Class` can update `ClassAttributes` data with each specific message.
- The `Class.Id` can be queried by `Class.Name` if the full name is matched.
- The `Class.Id` can be queried by `ClassAttributes.Owner`.
- ( If the module which performs the data transition between cosmos SDK's x/nft module and wasmd module is implemented, the `NFT` can be extended by using CosmWasm. )
- ( For the sale, listing and selling can be done in x/nftmarket mudule. )

### Impossible

- The NFT data field is limited. So the addition of data field by the creators is impossible.
- The flexible `Class.Id` is not supported.
- The addition of function to `NFT` behavior by the creators and third-party developers is impossible by the features on the chain.
- This module doesn't support simple NFT sale with fixed price at the moment.

## Class

The `Class` data is defined in sdk's x/nft module.  
We use that type definition and store the generated `Class` data on UnUniFi in x/nft module.  
This is similar to metadata of each belonging `NFT`.  
On UnUniFi, to create `Class` requires to send a `MsgCreateClass` message with `Class.Name`, `ClassAttributes.BaseTokenUri` and `ClassAttributes.TokenSupplyCap`, `ClassAttributes.MintingPermission` and as options, `Class.Symbol`, `Class.Description` and `Class.Uri`.
Some of these data can be updated after the creation by the owner of the `Class`.

### Class Id

`Class.Id` must be unique because this is the identifier of its belonging `NFT`s. It's like contract address of ERC721 contract on evm chains.  
The way to be generated `Class.Id` in this protocol follows:
hash(AccAddress.Byte() + accoutn.sequence)

#### TokenSupplyCap

`ClassAttributes.TokenSupplyCap` is max token supply number of each `Class`'s `NFT`. This value is set when to be created `Class`.  
There's the limitation by the global parameter.  
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
The initial owner is the sender of `MsgCreateClass`. This parameter can be changed by sending `MsgSendClass` by the only owner.

## NFT

The `NFT` data structure is defined in sdk's x/nft module.  
We use that type definition and store the generated `NFT` data on UnUnifi in x/nft module, not in the x/nftmint module.  
This represents the `NFT` content.

### NFT Id

The `NFT.Id` is the identifier of the `NFT` in the `Class`.  
The `NFT.Id` will be chose by the minter and become a part of `NFT.Uri`.

### NFT Uri

The `NFT.Uri` is the parameter which `NFT` has. The content which is in `NFT.Uri` location usually represents the core information of the `NFT`.
In UnUniFi, the `NFT.Uri` is defined as following formula:  
`NFT.Uri = ClassAttributes.BaseTokenUri = NFT.id`  
e.g.  
ClassAttributes.BaseTokenUri = "ipfs://sample/",  
NFT.Id = "a00"  
**NFT.Uri = "ipfs://sample/a00"**

#### Content of `NFT.Uri` (metadsata structure)

The standard of `NFT.Uri`(metadata structure) on UnUniFi.  
The thing we have to pay attention is the compotibility to the NFTs on the other network, especially in big market size like Ethereum, solana.  
The possible elements which we should apply from the OpenSea document:

- image - This is the URL to the image of the item.
- external_url - This is the URL that will appear below the asset's image on OpenSea and will allow users to leave OpenSea and view the item on your site.
- description - A human readable description of the item.
- name - Name of the item.
- attributes - These are the attributes for the item, which will show up on the OpenSea page for the item.
- background_color - Background color of the item on OpenSea. Must be a six-character hexadecimal without a pre-pended #.
- animation_url - A URL to a multi-media attachment for the item. The file extensions GLTF, GLB, WEBM, MP4, M4V, OGV, and OGG are supported, along with the audio-only extensions MP3, WAV, and OGA.
- youtube_url - A URL to a YouTube video.

Possible reference: https://docs.opensea.io/docs/metadata-standards

## Mint

Minting `NFT` is achieved by sending `MsgMintNFT` message with `Class.Id`, `NFT.Id` and `Receiver`.
But, there's minting permission in some case.  
If the attached `Class` has `OnlyOwner` attribute regarding `ClassAttributes.MintingPermission`, only owner of `Class` can mint `NFT`.  
If the attached `Class` has `Anyone` attribute regarding `ClassAttributes.MintingPermission`, anyone can mint `NFT`.  
The current minting options are above two cases.

## Update

### Class related Attributes

At the moment, we support the update of the `ClassAttributes` and `Class` data by managing messages for each data.  
e.g. We set `MsgUpdateBaseTokenUri` fot the `ClassAttributes.Owner` to update the `ClassAttributes.BaseTokenUri`.  
In this case, every belonging `NFT.Uri` will be changed at once.
And we provide `MsgUpdateTokenSupplyCap` for the `ClassAttributes.Owner` to update the `ClassAttributes.TokenSupplyCap`. But, you can't make token supply cap lowner that the current token supply number.

## Burn

The NFTs can be burned by the owner of `NFT`.
As we handle NFTs on NFTFi protocol on UnUniFi, the burn permission belongs only the owner of the `NFT`.

## Transrer

### Class Ownership

The owner of the `Class` is recorded in `ClassAttribtues` as `Owner` parameter.  
And we support the transition of the `ClassAttributes.Owner` by the sending `MsgSendClass` message which change `ClassAttributes.Owner` to any recipient.

#### Normal NFT

The normal NFT transfer is achieved by cosmos sdk's x/nft message.  
It performs changing owner which is connected to specific `Class.Id` and `NFT.Id` in the KVStore in that module.  
So, it only requires `Class` and `NFT` objects are in that module and the sender (current owner) signature to transfer it.

## Validation

1. There's no duplicating `Class.Id` on the chain.
1. There's no duplicating `NFT.Id` in a `Class`.
1. TokenURI must be legal by length (idea: len(tokenURI) > 0).
1. The `Class.Id` and `NFT.Id` format must follows sdk's nft module definition.
1. `TokenSupplyCap` must be within the limitation of `NFT.Id` restriction by SDK's x/nft module. (and overflow)

There're many other rules to be set. Please refer to these:  
https://github.com/irisnet/irismod/blob/master/modules/nft/types/validation.go  
https://github.com/cosmos/cosmos-sdk/blob/v0.46.0-rc1/x/nft/validation.go
