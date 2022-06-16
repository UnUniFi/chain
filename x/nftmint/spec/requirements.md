# Requirement

**NOTES: This is just for development. Once the other legitimate pages are finalized, remove this.**

## Basic

The requirements for collective NFT minting mainly.

### Mint

1. Anyone can create `Class` to mint NFT.
1. The creator of `Class` is recorded with `Class.Id`.
1. The creator of `Class` can restrict the right to mint NFT.
1. The total supply of `NFT`s in `Class` increases.

### Burn

1. The NFTs can be burned.
1. The owner of `Class` can restrict ability of `NFT`s under that `Class` to be burned.
1. The restrict levels have 3 state that 1 is `Nobody`, 2 is `OnlyCreator` and 3 is `OnlyOwner`.
1. In case of 1 level, nobody can burn the `NFT`s under the `Class` has this level.
1. In case of 2 level, only creator of `Class` of `NFT` can burn its all belonging `NFT`s.
1. In case of 3 level, only direct owner of `NFT` can burn it.
1. The total supply of `NFT`s in `Class` deceases.

### Update

1. The creator of `Class` of `NFT` can restrict ability of `NFT`s under that `class` to be updated.
1. The creator of `Class` can choose restrict levels that 1 is `OnlyCreator`, 2 is `OnlyOwner` and 3 is `Nobody`.
1. In case of 1 level, the creator of `Class` of the `NFT` can update its all `NFT`s' data.
1. In case of 2 level, the owner of `NFT` can update its `NFT` data.
1. In case of 3 level, nobody can update `NFT`s' data under that `Class`.

### Transfer

1. The owner of `NFT` can transfer that `NFT` to any receipient.

### Validation

1. There's no duplicating `Class.Id` on the chain
1. There's no duplicating `NFT.Id` in a `Class`
1. TokenURI must be legal by length (idea: len(tokenURI) > 0)
1. The `Class.Id` and `NFT.Id` format must follows sdk's nft module definition

## Constant

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

### Mint

1. Anyone can create `Class` to mint ntNFT.
1. Creator of `Class` can restrict the right to mint ntNFT.
1. The total supply of `NFT`s of `Class` increases.

### Burn

1. The NFTs can be burned.
1. The owner of `Class` can restrict ability of `NFT`s under that `Class` to be burned.
1. The restrict levels have 3 state that 1 is `Nobody`, 2 is `OnlyCreator`.
1. In case of 1 level, nobody can burn the `NFT`s under the `Class` has this level.
1. In case of 2 level, only creator of `Class` of `NFT` can burn its all belonging `NFT`s.
1. The total supply of `NFT`s in `Class` deceases.

### Transfer

1. Any owner of `NFT` can't transfer `NFT`.

### Update

1. The creator of `Class` of `NFT` can restrict ability of `NFT`s under that `class` to be updated.
1. The creator of `Class` can choose restrict levels that 1 is `OnlyCreator`and 2 is `Nobody`.
1. In case of 1 level, the creator of `Class` of the `NFT` can update its all `NFT`s' data.
1. In case of 2 level, nobody can update `NFT`s' data under that `Class`.

### Validation

