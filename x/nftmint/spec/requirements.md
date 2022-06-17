# Requirement

**NOTES: This is just for development. Once the other legitimate pages are finalized, remove this.**

## Basic

The requirements for collective NFT minting mainly.

### Mint

1. Anyone can create `Class` to mint NFT.
1. The owner of `Class` is recorded with `Class.Id`.
1. The initial owner of `Class` is the creator of `Class`.
1. The owner of `Class` can restrict the right to mint NFT.
1. The total supply of `NFT`s in `Class` increases.
1. The minter address of `NFT` should be recorded with `Class.Id` and `NFT.Id`.

### Burn

1. The NFTs can be burned.
1. The owner of `Class` can restrict ability of `NFT`s under that `Class` to be burned.
1. The restrict levels have 3 state that 1 is `Nobody`, 2 is `OnlyClassOwner` and 3 is `OnlyNFTOwner`.
1. In case of 1 level, nobody can burn the `NFT`s under the `Class` has this level.
1. In case of 2 level, only owner of `Class` of `NFT` can burn its all belonging `NFT`s.
1. In case of 3 level, only direct owner of `NFT` can burn it.
1. The total supply of `NFT`s in `Class` deceases.

### Update

1. The owner of `Class` of `NFT` can restrict ability of `NFT`s under that `class` to be updated.
1. The owner of `Class` can choose permittion levels out of the choises that 0 is `Nobody`, 1 is `OnlyClassOwner`, 2 is `OnlyNFTOwner` and 3 is `ClassAndNFTOwner`.
1. In case of 0 level, nobody can update `NFT`s' data under that `Class`.1. In case of 1 level, the owner of `Class` of the `NFT` can update its all `NFT`s' data.
1. In case of 2 level, the owner of `NFT` can update its `NFT` data.
1. In case of 3 level, the owner of `NFT` and `Class` of that `NFT` can update.

### Transfer

1. The owner of `NFT` can transfer that `NFT` to any receipient.

### Validation

1. There's no duplicating `Class.Id` on the chain.
1. There's no duplicating `NFT.Id` in a `Class`.
1. TokenURI must be legal by length (idea: len(tokenURI) > 0).
1. The `Class.Id` and `NFT.Id` format must follows sdk's nft module definition.

### Query

The cosmos SDK's nft module has many queries.   
Please refer them since I don't write duplicated queries here. ([Query service methods](https://github.com/cosmos/cosmos-sdk/blob/aba9bdc24cb6a7b9a85e6cad617f7b55d6dcdcec/docs/architecture/adr-043-nft-module.md?plain=1#L175))

1. The owner address of `Class` can be queried by `Class.Id`.
1. The minter address of `NFT` can be queried by `Class.Id` and `NFT.Id`.
1. The `update_status_level` of `NFT` can be queried by `Class.Id` and `NFT.Id`.


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
The way to handle those elements have not determined yet.

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
1. The owner of `Class` can restrict ability of `NFT`s under that `Class` to be burned.
1. The restrict levels have 3 state that 1 is `Nobody`, 2 is `OnlyClassOwner`.
1. In case of 1 level, nobody can burn the `NFT`s under the `Class` has this level.
1. In case of 2 level, only owner of `Class` of `NFT` can burn its all belonging `NFT`s.
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

