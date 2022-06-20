# Requirement

**NOTES: This is just for development. Once the other legitimate pages are finalized, remove this.**

## Basic

**_The requirements for collective NFT minting mainly._**

- Anyone can create Collective NFT by creating `Class` and its belonging `NFT`s while the owner of `Class` controls the minting right.
- The NFT standard generally follows ERC721 by default.
- The content of `NFT.Uri` (metadata strucure) follows UnUniFi standard.

### Class

1. Anyone can create `Class` to mint NFT.
1. The owner of `Class` is recorded with `Class.Id`.
1. The initial owner of `Class` is the creator of `Class`.
1. The owner of `Class` can transfer the ownership to any recipient by sending a message.

### Mint

1. The owner of `Class` is recorded with `Class.Id`.
1. The owner of `Class` can restrict the right to mint NFT.
1. The total supply of `NFT`s in `Class` increases.
1. The minter address of `NFT` should be recorded with `Class.Id` and `NFT.Id`.

### Burn

1. The NFTs can be burned.
1. The owner of `Class` can choose permittion levels out of the choises that 0 is `Nobody`, 1 is `OnlyClassOwner`, 2 is `OnlyNFTOwner` and 3 is `ClassAndNFTOwner`.
1. In case of 0 level, nobody can burn `NFT`s' data under that `Class`.
1. In case of 1 level, the owner of `Class` of the `NFT` can burn its all `NFT`s' data.
1. In case of 2 level, the owner of `NFT` can burn its `NFT` data.
1. In case of 3 level, both the owner of `NFT` and `Class` of that `NFT` can burn.
1. The total supply of `NFT`s in `Class` deceases.

### Update

#### Class

1. The owner of `Class` of `NFT` can restrict ability of `Class` data to be updated.
1. The owner of `Class` can choose permittion levels out of the choises that 0 is `Nobody`, 1 is `OnlyClassOwner`.
1. In case of 0 level, nobody can update `NFT`s' data under that `Class`.
1. In case of 1 level, the owner of `Class` of the `NFT` can update its all `NFT`s' data.

#### NFT

1. The owner of `Class` of `NFT` can restrict ability of `NFT`s under that `class` to be updated.
1. The owner of `Class` can choose permittion levels out of the choises that 0 is `Nobody`, 1 is `OnlyClassOwner`, 2 is `OnlyNFTOwner` and 3 is `ClassAndNFTOwner`.
1. In case of 0 level, nobody can update `NFT`s' data under that `Class`.
1. In case of 1 level, the owner of `Class` of the `NFT` can update its all `NFT`s' data.
1. In case of 2 level, the owner of `NFT` can update its `NFT` data.
1. In case of 3 level, both the owner of `NFT` and `Class` of that `NFT` can update.

### Transfer

1. The owner of `NFT` can transfer that `NFT` to any recipient.
1. (In the future,) the owner of `NFT` may be able to transfer that `NFT` onto the other blockchain using IBC as long as the `NFT` follows sdk's nft module standard.

### Validation

1. There's no duplicating `Class.Id` on the chain.
1. There's no duplicating `NFT.Id` in a `Class`.
1. TokenURI must be legal by length (idea: len(tokenURI) > 0).
1. The `Class.Id` and `NFT.Id` format must follows sdk's nft module definition.

### Message

- CreateClass(class_id, ..)
- TransferClass(class_id, recipient)
- UpdateClass(class_id, ...)
- MintNFT(class_id, nft_id, ...)
- UpdateNFT(class_id, nft_id, ...)
- BurnNFT(class_id, nft_id)

From cosmosd SDK's x/nft module:

- Send(class_id, nft_id, sender, receiver)

### Query

The cosmos SDK's nft module has many queries.   
Please refer them since I don't write duplicated queries here. ([Query service methods](https://github.com/cosmos/cosmos-sdk/blob/aba9bdc24cb6a7b9a85e6cad617f7b55d6dcdcec/docs/architecture/adr-043-nft-module.md?plain=1#L175))

1. The owner address of `Class` can be queried by `Class.Id`.
1. The minter address of `NFT` can be queried by `Class.Id` and `NFT.Id`.
1. The `update_status_level` of `Class` can be queried by `Class.Id`.
1. The `update_status_level` of `NFT` can be queried by `Class.Id`.
1. The data of `Class` and its belonging `NFT`s can be queried by `Class.Name`. **NOTE: return object has not been decided yet.** Possible choises are just `Class.Id` or []`NFT`.

1. ClassOwner(class_id)
1. NFTMinter(class_id, nft_id)
1. ClassUpdateStatusLevel(class_id)
1. NFTUpdateStatusLevel(class_id)
1. ClassByName(_under consideration_)

From cosmos SDk's x/nft module:

- Balance(class_id, owner) amount
- Owner(class_id, nft_id) owner
- Supply(class_id) amount
- NFTs(class_id, owner, pagination) nfts, pagination
- NFT(class_id, nft_id) nft
- Class(class_id) class
- Classes(pagination) classes, pagination

### What this feature mainly provide and doesn't

#### Positive

- Users can mint NFTs under specific Class like CryptoPunks

#### Negative

- Users cannot mint dynamic NFTs handled by special function in Contract in Ethereum
- The reason why is the module interect level is very low in cosmos SDk, it doesn't provide the ability to add funtions to NFTs other than we provide like smartcontract on Ethereum

## Constant

**NOTE: This sectino might be removed.**

There're some flexible variables that sdk's nft module has like minimum and maximum `Class.Id` string length.   
Write down those variable which we must define constantly to validate. (var name is not fixed, just momentary)   

`MinClassIdLen`   
`MinClassNameLen`   
`MaxClassNameLen`   
`MinNFTIDLen`   
`MaxNFTIDLen`   
`MaxNameLen`   
`MaxDescriptionLen`   
`MaxURILen`    

Or, some of those could be solved by fixing and generating automatically in a protocol.   
The way to handle those elements have not been determined yet.

## Non-transferable NFT (ntNFT)

The requirements for Non-transferable NFT minting.

### Logic Concept

T

### Mint

1. Anyone can create `Class` to mint ntNFT.
1. owner of `Class` can restrict the right to mint ntNFT.
1. The total supply of `NFT`s of `Class` increases.

### Burn

1. The NFTs can be burned.
1. The owner of `Class` can choose restrict levels that 0 is `Nobody` and 1 is `OnlyClassOwner`.
1. In case of 0 level, nobody can burn `NFT`s' data under that `Class`.
1. In case of 1 level, the owner of `Class` of the `NFT` can burn its all `NFT`s' data.
1. The total supply of `NFT`s in `Class` deceases.

### Transfer

1. Any owner of `NFT` can't transfer `NFT`.
1. Transfer is already achieved by sdk's x/nft module.

### Update

1. The owner of `Class` of `NFT` can restrict ability of `NFT`s under that `class` to be updated.
1. The owner of `Class` can choose restrict levels that 0 is `Nobody` and 1 is `OnlyClassOwner`.
1. In case of 0 level, nobody can update `NFT`s' data under that `Class`.
1. In case of 1 level, the owner of `Class` of the `NFT` can update its all `NFT`s' data.

### Validation

1. There's no duplicating `Class.Id` on the chain.
1. There's no duplicating `NFT.Id` in a `Class`.
1. TokenURI must be legal by length (idea: len(tokenURI) > 0).
1. The `Class.Id` and `NFT.Id` format must follows sdk's nft module definition.
1. The `Class.Data` must be something like `"nfNFT"`. (not determined yet).

### Query

1. The owner of `Class` can be queried by `Class.Id`.
1. The minter of `NFT` can be queried by `Class.Id` and `NFT.Id`.
1. The `NFT` total supply in a `Class` can be queried by `Class.Id`.
1. The whole `Class` and `NFT` data can be queried by their `Class.Id` and `Class.Id` and `NFT.Id`.
1. The owner addresses of `NFT`s in a `Class` can all be queried by `Class.Id`.
