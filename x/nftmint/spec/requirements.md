# Requirement

**NOTES: This is just for development. Once the other legitimate pages are finalized, remove this.**

## Basic

The requirements for collective NFT minting mainly.

### Mint

### Burn


### Validation

1. There's no duplicating `Class.Id` on the chain
1. There's no duplicating `NFT.Id` in a `Class`
1. TokenURI must be legal by length (idea: len(tokenURI) > 0)
1. The `Class.Id` and `NFT.Id` format must follows sdk's nft module definition
## Constant

There're some flexible variable that sdk's nft module has like minimum and maximum `Class.Id` string length.   
Write down those variable which we must define constantly to validate. (var name is not fixed, just momentary)   

`MinClassIdLen`   
`MinClassNameLen`   
`MaxClassNameLen`   
`MinNFTIDLen`   
`MaxNFTIDLen`   
`MaxNameLen`   
`MaxDescriptionLen`   
`MaxURILen`   
`DoNotModify`   
`IDPrefix`   
`DenomPrefix`   

## Non-transferable NFT

The requirements for Non-transferable NFT minting.

### Mint

1. 
### Burn

### Validation
